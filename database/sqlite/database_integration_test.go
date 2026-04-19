package sqlite

import (
	_ "embed"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/tables.sql
var sqlTables string

func Test_Integration_IsInitialized_is_true(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "test.db")
	databaseService := New(dbPath)
	err := databaseService.Initialize()
	assert.NoError(t, err)

	ok := databaseService.IsInitialized()
	assert.True(t, ok)

	err = databaseService.Close()
	assert.NoError(t, err)

	err = databaseService.DeleteDb()
	assert.NoError(t, err)
}

func Test_Integration_CreateTable(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "test.db")
	databaseService := New(dbPath)
	err := databaseService.Initialize()
	assert.NoError(t, err)

	databaseService.CreateTable(&sqlTables)
	_, err = databaseService.Apply()
	if err != nil {
		t.Errorf("Failed to apply database %v", err)
	}

	err = databaseService.Close()
	assert.NoError(t, err)

	err = databaseService.DeleteDb()
	assert.NoError(t, err)
}

func Test_Integration_Insert(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "test.db")
	databaseService := New(dbPath)
	err := databaseService.Initialize()
	assert.NoError(t, err)

	databaseService.CreateTable(&sqlTables)
	_, err = databaseService.Apply()
	if err != nil {
		t.Errorf("Failed to apply database %v", err)
	}

	rows := []DataRow{
		{
			"txt":  "Some text.",
			"num":  1091,
			"bool": true,
			"real": 91.01,
		},
		{
			"txt":  "Some text 2.",
			"num":  1891,
			"bool": false,
			"real": 98.0001,
		},
	}

	databaseService.Insert("content", rows)
	_, err = databaseService.Apply()
	if err != nil {
		t.Errorf("Failed to apply database %v", err)
	}

	err = databaseService.Close()
	assert.NoError(t, err)

	err = databaseService.DeleteDb()
	assert.NoError(t, err)
}

func Test_Integration_Select(t *testing.T) {
	abs, err := filepath.Abs("./testdata/test-integration-insert.db")
	if err != nil {
		t.Errorf("Failed to get absolute path: %v", err)
	}

	databaseService := New(abs)
	err = databaseService.Initialize()
	assert.NoError(t, err)

	databaseService.Select("content", SelectClauses{
		Where: []Condition{
			{
				Column:   "num",
				Operator: OperatorEqual,
				Value:    1091,
			},
		},
	}, []JoinClauses{})
	queryResults, err := databaseService.Apply()
	if err != nil {
		t.Errorf("Failed to apply database %v", err)
	}

	err = databaseService.Close()
	assert.NoError(t, err)

	assert.Equal(t, queryResults, &Queries{
		CreateTable: []Query{},
		DropTable:   []Query{},
		Insert:      []Query{},
		Update:      []Query{},
		Delete:      []Query{},
		Select: []Query{
			{
				Query:  "SELECT * FROM content  WHERE num = '1091';",
				Values: nil,
				Result: []map[string]any{
					{
						"txt":       "Some text.",
						"num":       int64(1091),
						"bool":      true,
						"real":      91.01,
						"date_time": time.Date(2025, 02, 24, 2, 48, 27, 0, time.UTC),
					},
				},
			},
		},
	})
}

func Test_Integration_Delete(t *testing.T) {}
