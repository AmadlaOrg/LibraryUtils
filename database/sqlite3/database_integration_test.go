package sqlite3

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/tables.sql
var sqlTables string

func Test_Integration_IsInitialized_is_true(t *testing.T) {
	databaseService := NewDatabaseService("/tmp/test-integration-IsInitialized-is-true.db")
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
	databaseService := NewDatabaseService("/tmp/test-integration-create-table.db")
	err := databaseService.Initialize()
	assert.NoError(t, err)

	databaseService.CreateTable(&sqlTables)
	err = databaseService.Apply()
	if err != nil {
		t.Errorf("Failed to apply database %v", err)
	}

	err = databaseService.Close()
	assert.NoError(t, err)

	// TODO: Validate the content

	err = databaseService.DeleteDb()
	assert.NoError(t, err)
}

func Test_Integration_Insert(t *testing.T) {
	databaseService := NewDatabaseService("/tmp/test-integration-insert.db")
	err := databaseService.Initialize()
	assert.NoError(t, err)

	databaseService.CreateTable(&sqlTables)
	err = databaseService.Apply()
	if err != nil {
		t.Errorf("Failed to apply database %v", err)
	}

	entitiesTable := Table{
		Name: "content",
		Rows: []Row{
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
				// TODO: add datetime
			},
		},
	}

	databaseService.Insert(entitiesTable)
	err = databaseService.Apply()
	if err != nil {
		t.Errorf("Failed to apply database %v", err)
	}

	err = databaseService.Close()
	assert.NoError(t, err)

	// TODO: Validate the content

	//err = databaseService.DeleteDb()
	//assert.NoError(t, err)
}

func Test_Integration_Select(t *testing.T) {}

func Test_Integration_Delete(t *testing.T) {}
