package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// IDatabase defines the database interface
type IDatabase interface {
	Initialize() error
	Close() error
	IsInitialized() bool
	CreateTable(sqlTables *string)
	Insert(tableName Table, rows Rows)
	Update(tableName Table, rows Rows, where []Condition)
	Select(tableName Table, clauses SelectClauses, joinClauses []JoinClauses)
	Delete(tableName Table, clauses SelectClauses)
	DeleteDb() error
	Apply() (*Queries, error)
}

// SDatabase implements IDatabase
type SDatabase struct {
	dbAbsPath   string
	queries     *Queries
	sqlDB       IDatabaseSqlDB
	initialized bool
}

// For mocking 🥸.
var (
	db          IDatabaseSqlDB
	dbMutex     sync.Mutex // sync.Locker
	initialized bool
	sqlOpen     = func(driverName, dataSourceName string) (IDatabaseSqlDB, error) {
		return sql.Open(driverName, dataSourceName)
	}
	osRemove = os.Remove
)

// Initialize establishes the database connection
//
// Returns:
// - 🚨 error:
func (s *SDatabase) Initialize() error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if db != nil {
		return nil // Already initialized successfully
	}

	var err error
	db, err = sqlOpen("sqlite3", s.dbAbsPath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Set PRAGMA statements for performance
	//
	// Doc: https://stackoverflow.com/questions/57118674/go-sqlite3-with-journal-mode-wal-gives-database-is-locked-error
	_, err = db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		err := db.Close()
		if err != nil {
			return err
		}
		db = nil
		return fmt.Errorf("error setting journal mode: %v", err)
	}

	// Configure the database connection pool
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	// Initialization successful
	initialized = true

	return nil
}

// Close closes the database connection
//
// Returns:
// - 🚨 error:
func (s *SDatabase) Close() error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if !initialized {
		// Database was never initialized; nothing to close
		return nil
	}

	if db != nil {
		err := db.Close()
		if err != nil {
			return err
		}
		db = nil
		initialized = false
		return nil
	}

	initialized = false
	return nil
}

// IsInitialized returns true if the database has been initialized
func (s *SDatabase) IsInitialized() bool {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	return initialized
}

func (s *SDatabase) query(
	addTo *[]Query,
	rows Rows,
	buildQueryFunc func(rows Rows, columnNames, valuesPlaceholder []string) string,
) {
	for _, row := range rows {
		columnNames, valuesPlaceholder, columnValues := processRow(row)

		// Build the query using the provided function
		query := buildQueryFunc(rows, columnNames, valuesPlaceholder)

		// Add the query to the queries list
		s.addQuery(addTo, query, columnValues)
	}
}

// addQuery adds the queries to the queries struct component
func (s *SDatabase) addQuery(slice *[]Query, query string, values []any) {
	*slice = append(*slice, Query{
		Query:  query,
		Values: values,
	})
}

// CreateTable creates a new table
func (s *SDatabase) CreateTable(sqlTables *string) {
	s.addQuery(&s.queries.CreateTable, *sqlTables, nil)
}

// Insert inserts records into the table
func (s *SDatabase) Insert(tableName Table, rows Rows) {
	s.query(
		&s.queries.Insert,
		rows,
		func(rows Rows, columnNames, valuesPlaceholder []string) string {
			var b strings.Builder
			b.WriteString("INSERT INTO ")
			b.WriteString(string(tableName))
			b.WriteString(" (")
			b.WriteString(strings.Join(columnNames, ", "))
			b.WriteString(") VALUES (")
			b.WriteString(strings.Join(valuesPlaceholder, ", "))
			b.WriteString(");")
			return b.String()
		},
	)
}

// Update updates a record in the table
func (s *SDatabase) Update(tableName Table, rows Rows, where []Condition) {
	s.query(
		&s.queries.Update,
		rows,
		func(rows Rows, columnNames, valuesPlaceholder []string) string {
			var b strings.Builder
			b.WriteString("UPDATE ")
			b.WriteString(string(tableName))
			b.WriteString(" SET ")

			var updates []string
			for _, column := range columnNames {
				updates = append(updates, fmt.Sprintf("%s = ?", column))
			}

			b.WriteString(strings.Join(updates, ", "))
			b.WriteString(buildWhere(where))

			return b.String()
		},
	)
}

// Select retrieves a record from the table
func (s *SDatabase) Select(tableName Table, clauses SelectClauses, joinClauses []JoinClauses) {
	// Build the SELECT query
	var b strings.Builder
	b.WriteString("SELECT * FROM ")
	b.WriteString(string(tableName))

	// Build JOIN clauses, if any
	if joinClauses != nil {
		b.WriteString(" ")
		b.WriteString(buildJoinClauses(joinClauses))
	}

	// Build WHERE, GROUP BY, HAVING, ORDER BY, LIMIT, OFFSET
	b.WriteString(buildWhere(clauses.Where))
	if len(clauses.GroupBy) > 0 {
		// For SELECT, GroupBy is just: GROUP BY col1, col2 ...
		b.WriteString(fmt.Sprintf(" GROUP BY %s", strings.Join(clauses.GroupBy, ", ")))
	}
	if len(clauses.OrderBy) > 0 {
		b.WriteString(buildHaving(clauses.Having))
		b.WriteString(buildOrderBy(clauses.OrderBy))
	}
	if clauses.Limit != nil && *clauses.Limit > 0 {
		b.WriteString(buildLimit(int64(*clauses.Limit)))
	}
	if clauses.Offset != nil && *clauses.Offset > 0 {
		b.WriteString(buildOffset(int64(*clauses.Offset)))
	}
	b.WriteString(";")

	s.addQuery(&s.queries.Select, b.String(), nil)
}

// Delete deletes records from the table
func (s *SDatabase) Delete(tableName Table, clauses SelectClauses) {
	var b strings.Builder
	b.WriteString("DELETE FROM ")
	b.WriteString(string(tableName))
	b.WriteString(buildWhere(clauses.Where))
	b.WriteString(";")

	s.addQuery(&s.queries.Delete, b.String(), nil)
}

// DeleteDb delete a db file
// - Used when the caching will be reset
//
// Returns:
// - 🚨 error:
func (s *SDatabase) DeleteDb() error {
	if ok, err := ValidateDbAbsPath(s.dbAbsPath); !ok {
		return err
	}

	// Proceed with deleting the database file
	err := osRemove(s.dbAbsPath)
	if err != nil {
		return err
	}

	return nil
}

// Apply executes all queued queries in s.queries
//
// Returns:
// - 🚨 error:
func (s *SDatabase) Apply() (*Queries, error) {
	if !s.IsInitialized() {
		return nil, fmt.Errorf(ErrorDatabaseNotInitialized)
	}

	// Begin a transaction
	sqlTx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	/*
		for _, q := range s.queries.DropTable {
			_, err = s.exec(sqlTx, q.Query, q.Values)
			if err != nil {
				return err
			}
			//allQueries = append(allQueries, q.Query)
		}*/

	// Execute CREATE TABLE queries directly (no values binding)
	for _, q := range s.queries.CreateTable {
		_, err = sqlTx.Exec(q.Query)
		if err != nil {
			err := sqlTx.Rollback()
			if err != nil {
				return nil, err
			} // Rollback transaction if there's an error
			return nil, fmt.Errorf("error creating table: %w", err)
		}
	}

	// Execute other queries with values
	queryTypes := []*[]Query{
		&s.queries.Insert,
		&s.queries.Update,
		&s.queries.Delete,
	}

	for _, querySet := range queryTypes {
		for _, q := range *querySet {
			dbResult, err := s.exec(sqlTx, q.Query, q.Values)
			if err != nil {
				err := sqlTx.Rollback()
				if err != nil {
					return nil, err
				} // Rollback if any query fails
				return nil, err
			}
			lastInsertId, err := dbResult.LastInsertId()
			if err != nil {
				return nil, err
			}
			q.Result = strconv.FormatInt(lastInsertId, 10)
		}
	}

	// Handle SELECT queries separately
	for qi, q := range s.queries.Select {
		s.queries.Select[qi].Result, err = s.querySelect(sqlTx, q.Query, q.Values)
		if err != nil {
			err := sqlTx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	// **Commit transaction once after all queries execute**
	if err = sqlTx.Commit(); err != nil {
		err = sqlTx.Rollback()
		if err != nil {
			return nil, err
		} // Ensure rollback on commit failure
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	// Optionally, clear the queries after applying
	/*s.queries = &Queries{
		CreateTable: []Query{},
		DropTable:   []Query{},
		Insert:      []Query{},
		Update:      []Query{},
		Delete:      []Query{},
		Select:      []Query{},
	}*/

	return s.queries, nil
}

// exec loops through all the queries and executes them.
//
// Params:
// - sqlTx:
// - query:
// - values:
//
// Returns:
// -
// - 🚨 error:
func (s *SDatabase) exec(sqlTx IDatabaseSqlTx, query string, values []any) (sql.Result, error) {
	dbResult, err := sqlTx.Exec(query, values...)
	if err != nil {
		// Roll back the entire transaction on error
		rbErr := sqlTx.Rollback()
		if rbErr != nil {
			return nil, fmt.Errorf("error rolling back transaction: %w (original error: %v)", rbErr, err)
		}
		return nil, fmt.Errorf("error applying query (%s): %w", query, err)
	}
	return dbResult, nil
}

// querySelect
//
// Params:
// - sqlTx:
// - query:
// - values:
//
// Returns:
// -
// - 🚨 error:
func (s *SDatabase) querySelect(sqlTx IDatabaseSqlTx, query string, values []any) (any, error) {
	rows, err := sqlTx.Query(query, values...)
	if err != nil {
		return "", fmt.Errorf("error executing SELECT query (%s): %w", query, err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}(rows)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("error fetching column names: %w", err)
	}

	// Prepare a slice to store the results
	var results []map[string]any

	// Iterate over rows
	for rows.Next() {
		// Create a slice of `any` to hold each column value
		values := make([]any, len(columns))
		valuePtrs := make([]any, len(columns))

		// Assign pointers to values slice
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the row into value pointers
		if err := rows.Scan(valuePtrs...); err != nil {
			return "", fmt.Errorf("error scanning row: %w", err)
		}

		// Convert values to a map
		rowMap := make(map[string]any)
		for i, colName := range columns {
			val := values[i]

			// Convert `[]byte` to string if necessary
			if b, ok := val.([]byte); ok {
				val = string(b)
			}

			rowMap[colName] = val
		}

		results = append(results, rowMap)
	}

	return results, nil

	// Convert results to JSON
	/*jsonData, err := json.Marshal(results)
	if err != nil {
		return "", fmt.Errorf("error converting result to JSON: %w", err)
	}

	return string(jsonData), nil*/
}
