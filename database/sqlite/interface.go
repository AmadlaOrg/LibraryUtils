package sqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"time"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

// IDatabaseSqlNullString for NullString methods.
// Found: "database/sql".
type IDatabaseSqlNullString interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNullInt64 for NullInt64 methods.
// Found: "database/sql".
type IDatabaseSqlNullInt64 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNullInt32 for NullInt32 methods.
// Found: "database/sql".
type IDatabaseSqlNullInt32 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNullInt16 for NullInt16 methods.
// Found: "database/sql".
type IDatabaseSqlNullInt16 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNullByte for NullByte methods.
// Found: "database/sql".
type IDatabaseSqlNullByte interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNullFloat64 for NullFloat64 methods.
// Found: "database/sql".
type IDatabaseSqlNullFloat64 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNullBool for NullBool methods.
// Found: "database/sql".
type IDatabaseSqlNullBool interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNullTime for NullTime methods.
// Found: "database/sql".
type IDatabaseSqlNullTime interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlNull for Null methods.
// Found: "database/sql".
type IDatabaseSqlNull interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// IDatabaseSqlDriverConn for DriverConn methods.
// Found: "database/sql".
type IDatabaseSqlDriverConn interface {
	Close() error
}

// IDatabaseSqlDriverStmt for DriverStmt methods.
// Found: "database/sql".
type IDatabaseSqlDriverStmt interface {
	Close() error
}

// IDatabaseSqlDsnConnector for dsnConnector methods.
// Found: "database/sql".
type IDatabaseSqlDsnConnector interface {
	Connect(_ context.Context) (driver.Conn, error)
	Driver() driver.Driver
}

// IDatabaseSqlDB for DB methods.
// Found: "database/sql".
type IDatabaseSqlDB interface {
	PingContext(ctx context.Context) error
	Ping() error
	Close() error
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	SetConnMaxIdleTime(d time.Duration)
	Stats() sql.DBStats
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryRow(query string, args ...any) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
	Driver() driver.Driver
	Conn(ctx context.Context) (*sql.Conn, error)
}

// IDatabaseSqlConn for Conn methods.
// Found: "database/sql".
type IDatabaseSqlConn interface {
	PingContext(ctx context.Context) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Raw(f func(driverConn any) error) (err error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
}

// IDatabaseSqlTx for Tx methods.
// Found: "database/sql".
type IDatabaseSqlTx interface {
	Commit() error
	Rollback() error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
	Stmt(stmt *sql.Stmt) *sql.Stmt
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryRow(query string, args ...any) *sql.Row
}

// IDatabaseSqlStmt for Stmt methods.
// Found: "database/sql".
type IDatabaseSqlStmt interface {
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
	Exec(args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	Query(args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, args ...any) *sql.Row
	QueryRow(args ...any) *sql.Row
	Close() error
}

// IDatabaseSqlRows for Rows methods.
// Found: "database/sql".
type IDatabaseSqlRows interface {
	Next() bool
	NextResultSet() bool
	Err() error
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Scan(dest ...any) error
	Close() error
}

// IDatabaseSqlColumnType for ColumnType methods.
// Found: "database/sql".
type IDatabaseSqlColumnType interface {
	Name() string
	Length() (length int64, ok bool)
	DecimalSize() (precision, scale int64, ok bool)
	ScanType() reflect.Type
	Nullable() (nullable, ok bool)
	DatabaseTypeName() string
}

// IDatabaseSqlRow for Row methods.
// Found: "database/sql".
type IDatabaseSqlRow interface {
	Scan(dest ...any) error
	Err() error
}

// IDatabaseSqlConnRequestSet for connRequestSet methods.
// Found: "database/sql".
type IDatabaseSqlConnRequestSet interface {
	CloseAndRemoveAll()
	Len() int
}

type IDatabaseSQLiteTx interface {
	Commit() error
	Rollback() error
}

// ISqlTx
/*type ISqlTx interface {
	Exec(query string, args ...any) (sql.Result, error)
	Rollback() error
	Commit() error
}*/

/*
   ISQLiteTx:
     config:
       dir: "{{.InterfaceDir}}"
       mockName: "MockISQLiteTx"
*/

// IGoSqlite3SQLiteConn for SQLiteConn methods.
// Found: github.com/mattn/go-sqlite3
type IGoSqlite3SQLiteConn interface {
	Serialize(schema string) ([]byte, error)
	Deserialize(b []byte, schema string) error
	LoadExtension(lib string, entry string) error
	RegisterPreUpdateHook(callback func(conn sqlite3.SQLitePreUpdateData))
	//SetTrace(requested *sqlite3.TraceConfig) error
	Authenticate(username, password string) error
	AuthUserAdd(username, password string, admin bool) error
	AuthUserChange(username, password string, admin bool) error
	AuthUserDelete(username string) error
	AuthEnabled() (exists bool)
	Ping(ctx context.Context) error
	QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error)
	ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error)
	PrepareContext(ctx context.Context, query string) (driver.Stmt, error)
	BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error)
	RegisterCollation(name string, cmp func(string, string) int) error
	RegisterCommitHook(callback func() int)
	RegisterRollbackHook(callback func())
	RegisterUpdateHook(callback func(int, string, string, int64))
	RegisterAuthorizer(callback func(int, string, string, string) int)
	RegisterFunc(name string, impl any, pure bool) error
	RegisterAggregator(name string, impl any, pure bool) error
	AutoCommit() bool
	Exec(query string, args []driver.Value) (driver.Result, error)
	Query(query string, args []driver.Value) (driver.Rows, error)
	Begin() (driver.Tx, error)
	Close() error
	Prepare(query string) (driver.Stmt, error)
	GetFilename(schemaName string) string
	GetLimit(id int) int
	SetLimit(id int, newVal int) int
	SetFileControlInt(dbName string, op int, arg int) error
	DeclareVTab(sql string) error
	//CreateModule(moduleName string, module sqlite3.Module) error
	Backup(dest string, srcConn *sqlite3.SQLiteConn, src string) (*sqlite3.SQLiteBackup, error)
}
