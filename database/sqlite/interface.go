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

// NullString for NullString methods.
// Found: "database/sql".
type NullString interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// NullInt64 for NullInt64 methods.
// Found: "database/sql".
type NullInt64 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// NullInt32 for NullInt32 methods.
// Found: "database/sql".
type NullInt32 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// NullInt16 for NullInt16 methods.
// Found: "database/sql".
type NullInt16 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// NullByte for NullByte methods.
// Found: "database/sql".
type NullByte interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// NullFloat64 for NullFloat64 methods.
// Found: "database/sql".
type NullFloat64 interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// NullBool for NullBool methods.
// Found: "database/sql".
type NullBool interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// NullTime for NullTime methods.
// Found: "database/sql".
type NullTime interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// Null for Null methods.
// Found: "database/sql".
type Null interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// DriverConn for DriverConn methods.
// Found: "database/sql".
type DriverConn interface {
	Close() error
}

// DriverStmt for DriverStmt methods.
// Found: "database/sql".
type DriverStmt interface {
	Close() error
}

// DsnConnector for dsnConnector methods.
// Found: "database/sql".
type DsnConnector interface {
	Connect(_ context.Context) (driver.Conn, error)
	Driver() driver.Driver
}

// DB for DB methods.
// Found: "database/sql".
type DB interface {
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

// Conn for Conn methods.
// Found: "database/sql".
type Conn interface {
	PingContext(ctx context.Context) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Raw(f func(driverConn any) error) (err error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
}

// Tx for Tx methods.
// Found: "database/sql".
type Tx interface {
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

// Stmt for Stmt methods.
// Found: "database/sql".
type Stmt interface {
	ExecContext(ctx context.Context, args ...any) (sql.Result, error)
	Exec(args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, args ...any) (*sql.Rows, error)
	Query(args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, args ...any) *sql.Row
	QueryRow(args ...any) *sql.Row
	Close() error
}

// Rows for Rows methods.
// Found: "database/sql".
type Rows interface {
	Next() bool
	NextResultSet() bool
	Err() error
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Scan(dest ...any) error
	Close() error
}

// ColumnType for ColumnType methods.
// Found: "database/sql".
type ColumnType interface {
	Name() string
	Length() (length int64, ok bool)
	DecimalSize() (precision, scale int64, ok bool)
	ScanType() reflect.Type
	Nullable() (nullable, ok bool)
	DatabaseTypeName() string
}

// Row for Row methods.
// Found: "database/sql".
type Row interface {
	Scan(dest ...any) error
	Err() error
}

// ConnRequestSet for connRequestSet methods.
// Found: "database/sql".
type ConnRequestSet interface {
	CloseAndRemoveAll()
	Len() int
}

type SQLiteTx interface {
	Commit() error
	Rollback() error
}

// SQLiteConn for SQLiteConn methods.
// Found: github.com/mattn/go-sqlite3
type SQLiteConn interface {
	Serialize(schema string) ([]byte, error)
	Deserialize(b []byte, schema string) error
	LoadExtension(lib string, entry string) error
	RegisterPreUpdateHook(callback func(conn sqlite3.SQLitePreUpdateData))
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
	Backup(dest string, srcConn *sqlite3.SQLiteConn, src string) (*sqlite3.SQLiteBackup, error)
}
