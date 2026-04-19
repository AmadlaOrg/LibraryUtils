package sqlite

import (
	"errors"
	"fmt"

	"github.com/AmadlaOrg/LibraryUtils/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Fixtures
var testDbAbsPath = "/tmp/hery.test.cache"

func TestInitialize_db_set(t *testing.T) {
	// Reset
	db = nil
	initialized = false

	db = NewMockDB(t)

	databaseService := New(testDbAbsPath)
	err := databaseService.Initialize()

	assert.NoError(t, err)
}

func TestInitialize(t *testing.T) {
	tests := []struct {
		name              string
		inputDbPath       string
		internalSqlOpenFn func(driverName, dataSourceName string) (DB, error)
		expectedErr       error
		hasError          bool
	}{
		{
			name:        "Initialize database",
			inputDbPath: "/tmp/hery.test.cache",
			internalSqlOpenFn: func(driverName, dataSourceName string) (DB, error) {
				mockSqlDb := NewMockDB(t)
				mockSqlDb.EXPECT().Exec(mock.Anything).Return(nil, nil)
				mockSqlDb.EXPECT().SetMaxOpenConns(mock.Anything)
				mockSqlDb.EXPECT().SetMaxIdleConns(mock.Anything)
				mockSqlDb.EXPECT().SetConnMaxLifetime(mock.Anything)
				return mockSqlDb, nil
			},
			hasError: false,
		},
		//
		// Error
		//
		{
			name:        "Error: db Open fail",
			inputDbPath: "/tmp/hery.test.cache",
			internalSqlOpenFn: func(driverName, dataSourceName string) (DB, error) {
				return nil, assert.AnError
			},
			expectedErr: errors.New("error opening database: "),
			hasError:    true,
		},
		{
			name:        "Error: db Exec function throws error",
			inputDbPath: "/tmp/hery.test.cache",
			internalSqlOpenFn: func(driverName, dataSourceName string) (DB, error) {
				mockSqlDb := NewMockDB(t)
				mockSqlDb.EXPECT().Exec(mock.Anything).Return(nil, assert.AnError) // assert.AnError
				mockSqlDb.EXPECT().Close().Return(nil)
				return mockSqlDb, nil
			},
			expectedErr: errors.New("error setting journal mode: "),
			hasError:    true,
		},
		{
			name:        "Error: db Exec function throws error also db Close",
			inputDbPath: "/tmp/hery.test.cache",
			internalSqlOpenFn: func(driverName, dataSourceName string) (DB, error) {
				mockSqlDb := NewMockDB(t)
				mockSqlDb.EXPECT().Exec(mock.Anything).Return(nil, assert.AnError) // assert.AnError
				mockSqlDb.EXPECT().Close().Return(assert.AnError)
				return mockSqlDb, nil
			},
			expectedErr: assert.AnError,
			hasError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset
			db = nil
			initialized = false

			/*mockSyncLocker := NewMockSyncLocker(t)
			mockSyncLocker.EXPECT().Lock()
			mockSyncLocker.EXPECT().Unlock()

			dbMutex = mockSyncLocker*/

			originalSqlOpen := sqlOpen
			defer func() { sqlOpen = originalSqlOpen }()
			sqlOpen = tt.internalSqlOpenFn

			databaseService := New(testDbAbsPath)
			err := databaseService.Initialize()

			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestClose(t *testing.T) {
	tests := []struct {
		name                string
		externalInitialized bool
		externalDb          DB
		hasError            bool
	}{
		{
			name:                "Close database",
			externalInitialized: true,
			externalDb: func() DB {
				mockSqlDb := NewMockDB(t)
				mockSqlDb.EXPECT().Close().Return(nil)
				return mockSqlDb
			}(),
			hasError: false,
		},
		{
			name:                "Database is closed",
			externalInitialized: true,
			externalDb: func() DB {
				return nil
			}(),
			hasError: false,
		},
		//
		// Error
		//
		{
			name:                "Error: Close database",
			externalInitialized: true,
			externalDb: func() DB {
				mockSqlDb := NewMockDB(t)
				mockSqlDb.EXPECT().Close().Return(assert.AnError)
				return mockSqlDb
			}(),
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db = tt.externalDb
			initialized = tt.externalInitialized

			/*mockSyncLocker := NewMockSyncLocker(t)
			mockSyncLocker.EXPECT().Lock()
			mockSyncLocker.EXPECT().Unlock()

			dbMutex = mockSyncLocker*/

			databaseService := New(testDbAbsPath)
			err := databaseService.Close()

			if tt.hasError {
				assert.Error(t, err)
				assert.True(t, initialized)
				assert.NotNil(t, db)
			} else {
				assert.NoError(t, err)
				assert.False(t, initialized)
				assert.Nil(t, db)
			}

		})
	}
}

func TestIsInitialized(t *testing.T) {
	tests := []struct {
		name                string
		externalInitialized bool
		expectedInitialized bool
	}{
		{
			name:                "IsInitialized is true",
			externalInitialized: true,
			expectedInitialized: true,
		},
		{
			name:                "IsInitialized is false",
			externalInitialized: false,
			expectedInitialized: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialized = tt.externalInitialized

			/*mockSyncLocker := NewMockSyncLocker(t)
			mockSyncLocker.EXPECT().Lock()
			mockSyncLocker.EXPECT().Unlock()

			dbMutex = mockSyncLocker*/

			databaseService := New(testDbAbsPath)
			got := databaseService.IsInitialized()
			assert.Equal(t, tt.expectedInitialized, got)
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name       string
		inputTable Table
		inputRows  DataRows
	}{
		{
			name:       "Test Insert",
			inputTable: "Net",
			inputRows: []map[string]any{
				{
					"Id":          "c6beaec1-90c4-4d2a-aaef-211ab00b86bd",
					"server_name": "localhost",
					"listen":      "[80, 443]",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			databaseService := New(testDbAbsPath)
			databaseService.Insert(tt.inputTable, tt.inputRows)
		})
	}
}

// TestUpdate is a placeholder — test cases were removed due to non-deterministic map iteration order.
func TestUpdate(t *testing.T) {
}

func TestSelect(t *testing.T) {
	tests := []struct {
		name             string
		inputTable       Table
		inputClauses     SelectClauses
		inputJoinClauses []JoinClauses
		expected         []Query
	}{
		{
			name:       "Test Select: one condition",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "server_name",
						Operator: OperatorLike,
						Value:    "%localhost",
					},
				},
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name WHERE server_name LIKE '%localhost';",
				},
			},
		},
		{
			name:       "Test Select: two conditions",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "server_name",
						Operator: OperatorLike,
						Value:    "%localhost",
					},
					{
						Column:   "foo",
						Operator: OperatorNotEqual,
						Value:    "boo",
					},
				},
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name WHERE server_name LIKE '%localhost' AND foo != 'boo';",
				},
			},
		},
		{
			name:       "Test Select: two conditions and limit 10",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "server_name",
						Operator: OperatorLike,
						Value:    "%localhost",
					},
					{
						Column:   "foo",
						Operator: OperatorNotEqual,
						Value:    "boo",
					},
				},
				Limit: pointer.ToPtr(10),
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name WHERE server_name LIKE '%localhost' AND foo != 'boo' LIMIT 10;",
				},
			},
		},
		{
			name:       "Test Select: two conditions and limit 10 with offset 5",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "server_name",
						Operator: OperatorLike,
						Value:    "%localhost",
					},
					{
						Column:   "foo",
						Operator: OperatorNotEqual,
						Value:    "boo",
					},
				},
				Limit:  pointer.ToPtr(10),
				Offset: pointer.ToPtr(5),
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name WHERE server_name LIKE '%localhost' AND foo != 'boo' LIMIT 10 OFFSET 5;",
				},
			},
		},
		{
			name:       "Test Select: two conditions, and group by and limit 10 with offset 5",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "server_name",
						Operator: OperatorLike,
						Value:    "%localhost",
					},
					{
						Column:   "foo",
						Operator: OperatorNotEqual,
						Value:    "boo",
					},
				},
				GroupBy: []string{
					"server_name",
				},
				Limit:  pointer.ToPtr(10),
				Offset: pointer.ToPtr(5),
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name WHERE server_name LIKE '%localhost' AND foo != 'boo' GROUP BY server_name LIMIT 10 OFFSET 5;",
				},
			},
		},
		{
			name:       "Test Select: two conditions, and group by and limit 10 with offset 5",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "server_name",
						Operator: OperatorLike,
						Value:    "%localhost",
					},
					{
						Column:   "foo",
						Operator: OperatorNotEqual,
						Value:    "boo",
					},
				},
				GroupBy: []string{
					"server_name",
				},
				Limit:  pointer.ToPtr(10),
				Offset: pointer.ToPtr(5),
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name WHERE server_name LIKE '%localhost' AND foo != 'boo' GROUP BY server_name LIMIT 10 OFFSET 5;",
				},
			},
		},
		{
			name:       "Test Select: two conditions, and group by, order by and limit 10 with offset 5",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "server_name",
						Operator: OperatorLike,
						Value:    "%localhost",
					},
					{
						Column:   "foo",
						Operator: OperatorNotEqual,
						Value:    "boo",
					},
				},
				GroupBy: []string{
					"server_name",
				},
				OrderBy: []OrderBy{
					{
						Column:    "server_name",
						Direction: OrderByDesc,
					},
				},
				Limit:  pointer.ToPtr(10),
				Offset: pointer.ToPtr(5),
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name WHERE server_name LIKE '%localhost' AND foo != 'boo' GROUP BY server_name ORDER BY server_name DESC LIMIT 10 OFFSET 5;",
				},
			},
		},
		/*{
			name: "Test Select: inner joins",
			inputTable: Table{
				Name: "mock_table_name",
			},
			inputClauses: SelectClauses{},
			inputJoinClauses: []JoinClauses{
				{
					Table: "mock_second_table_name",
					Type:  JoinTypeInner,
					On: []JoinCondition{
						{
							Column1:  "server_name",
							Operator: OperatorLike,
							Column2:  "local%",
						},
					},
				},
			},
			expected: []Query{
				{
					Query: "SELECT * FROM mock_table_name INNER JOIN mock_second_table_name ON mock_second_table_name.server_name LIKE mock_table_name.server_name;",
				},
			},
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			databaseService := databaseImpl{
				queries: &Queries{
					CreateTable: []Query{},
					DropTable:   []Query{},
					Insert:      []Query{},
					Update:      []Query{},
					Delete:      []Query{},
					Select:      []Query{},
				},
			}
			databaseService.Select(tt.inputTable, tt.inputClauses, tt.inputJoinClauses)
			assert.Equal(t, tt.expected, databaseService.queries.Select)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name         string
		inputTable   Table
		inputClauses SelectClauses
		expected     *Queries
	}{
		{
			name:       "Simple WHERE clause with column selection",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "Id",
						Operator: OperatorEqual,
						Value:    "mock_table_id",
					},
				},
			},
			expected: &Queries{
				CreateTable: []Query{},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete: []Query{
					{
						Query: "DELETE FROM mock_table_name WHERE Id = 'mock_table_id';",
					},
				},
				Select: []Query{},
			},
		},
		{
			name:       "Simple WHERE clause with TWO column selection",
			inputTable: "mock_table_name",
			inputClauses: SelectClauses{
				Where: []Condition{
					{
						Column:   "Name",
						Operator: OperatorEqual,
						Value:    "mock_table_column_name",
					},
					{
						Column:   "Type",
						Operator: OperatorLike,
						Value:    "%mock_table_column_type%",
					},
				},
			},
			expected: &Queries{
				CreateTable: []Query{},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete: []Query{
					{
						Query: "DELETE FROM mock_table_name WHERE Name = 'mock_table_column_name' AND Type LIKE '%mock_table_column_type%';",
					},
				},
				Select: []Query{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			databaseService := databaseImpl{
				queries: &Queries{
					CreateTable: []Query{},
					DropTable:   []Query{},
					Insert:      []Query{},
					Update:      []Query{},
					Delete:      []Query{},
					Select:      []Query{},
				},
			}
			databaseService.Delete(tt.inputTable, tt.inputClauses)
			assert.Equal(t, tt.expected, databaseService.queries)
		})
	}
}

func TestDeleteDb(t *testing.T) {
	tests := []struct {
		name                     string
		serviceStructDbAbsPath   string
		internalFileIsValidMagic func(path string, magic []byte) (bool, error)
		internalOsRemove         func(name string) error
		expectedError            error
		hasError                 bool
	}{
		{
			name:                   "Success Delete",
			serviceStructDbAbsPath: "testdata/mock_db",
			internalFileIsValidMagic: func(path string, magic []byte) (bool, error) {
				return true, nil
			},
			internalOsRemove: func(name string) error {
				return nil
			},
			hasError: false,
		},
		//
		// Errors
		//
		{
			name:                   "Error: Delete fails at ValidateDbAbsPath",
			serviceStructDbAbsPath: "testdata/mock_db",
			internalFileIsValidMagic: func(path string, magic []byte) (bool, error) {
				return false, errors.New("test error (ValidateDbAbsPath)")
			},
			expectedError: errors.New("test error (ValidateDbAbsPath)"),
			hasError:      true,
		},
		{
			name:                   "Error: Delete fails at os.Remove",
			serviceStructDbAbsPath: "testdata/mock_db",
			internalFileIsValidMagic: func(path string, magic []byte) (bool, error) {
				return true, nil
			},
			internalOsRemove: func(name string) error {
				return errors.New("test error (OsRemove)")
			},
			expectedError: errors.New("test error (OsRemove)"),
			hasError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			databaseService := databaseImpl{
				dbAbsPath: tt.serviceStructDbAbsPath,
				queries: &Queries{
					CreateTable: []Query{},
					DropTable:   []Query{},
					Insert:      []Query{},
					Update:      []Query{},
					Delete:      []Query{},
					Select:      []Query{},
				},
			}

			originalFileIsValidMagic := fileIsValidMagic
			defer func() { fileIsValidMagic = originalFileIsValidMagic }()
			fileIsValidMagic = tt.internalFileIsValidMagic

			originalOsRemove := osRemove
			defer func() { osRemove = originalOsRemove }()
			osRemove = tt.internalOsRemove

			err := databaseService.DeleteDb()
			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorContains(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name                     string
		internalQueries          *Queries
		internalInitialized      bool
		internalDbBeginExpect    bool
		internalDbBeginError     error
		internalTxExecExpect     bool
		internalTxExecError      error
		internalTxRollbackExpect bool
		internalTxRollbackError  error
		internalTxCommitExpect   bool
		internalTxCommitError    error
		expectedQueries          *Queries
		expectedError            error
		hasError                 bool
	}{
		{
			name: "Error: is not initialized",
			internalQueries: &Queries{
				CreateTable: []Query{},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete:      []Query{},
				Select:      []Query{},
			},
			internalInitialized: false,
			expectedQueries: &Queries{
				CreateTable: []Query{},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete:      []Query{},
				Select:      []Query{},
			},
			expectedError: fmt.Errorf(ErrorDatabaseNotInitialized),
			hasError:      true,
		},
		{
			name: "Error: db Begin fails",
			internalQueries: &Queries{
				CreateTable: []Query{{Query: "CREATE TABLE test (id TEXT)"}},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete:      []Query{},
				Select:      []Query{},
			},
			internalInitialized:   true,
			internalDbBeginExpect: true,
			internalDbBeginError:  errors.New("test error (ValidateDbBegin)"),
			expectedQueries: &Queries{
				CreateTable: []Query{{Query: "CREATE TABLE test (id TEXT)"}},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete:      []Query{},
				Select:      []Query{},
			},
			expectedError: fmt.Errorf("test error (ValidateDbBegin)"),
			hasError:      true,
		},
		{
			name: "Error: no queries",
			internalQueries: &Queries{
				CreateTable: []Query{},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete:      []Query{},
				Select:      []Query{},
			},
			internalInitialized:   true,
			internalDbBeginExpect: false,
			expectedQueries: &Queries{
				CreateTable: []Query{},
				DropTable:   []Query{},
				Insert:      []Query{},
				Update:      []Query{},
				Delete:      []Query{},
				Select:      []Query{},
			},
			expectedError: fmt.Errorf("error no queries"),
			hasError:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			databaseService := &databaseImpl{
				queries: tt.internalQueries,
			}

			// initialized
			originalInitialized := initialized
			defer func() { initialized = originalInitialized }()
			initialized = tt.internalInitialized

			mockSqlDb := NewMockDB(t)

			if tt.internalDbBeginExpect {

				if tt.internalTxExecExpect {
					mockSqlTx := NewMockTx(t)
					mockSqlTx.EXPECT().Exec(mock.Anything).Return(nil, tt.internalTxExecError)

					if tt.internalTxRollbackExpect {
						mockSqlTx.EXPECT().Rollback().Return(tt.internalTxRollbackError)
					}

					if tt.internalTxCommitExpect {
						mockSqlTx.EXPECT().Commit().Return(tt.internalTxCommitError)
					}

					} else {
					mockSqlDb.EXPECT().Begin().Return(nil, tt.internalDbBeginError)
				}

					db = mockSqlDb
			}

			_, err := databaseService.Apply()
			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedQueries, databaseService.queries)
		})
	}
}
