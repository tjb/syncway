package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"syncway/internal/core"
	"testing"
)

func TestTrackChanges_Success(t *testing.T) {
	// Arrange: Create a mocked database and adapter
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("error closing db: %v", err)
		}
	}(db)

	adapter := &SQLiteAdapter{db: db}

	rows := sqlmock.NewRows([]string{"id", "table_name", "operation", "row_id", "timestamp"}).
		AddRow(1, "documents", "insert", "row1", "2024-11-26 10:00:00").
		AddRow(2, "documents", "update", "row2", "2024-11-26 10:05:00")
	mock.ExpectQuery("SELECT id, table_name, operation, row_id, timestamp FROM changeset").
		WillReturnRows(rows)

	// Act: Call TrackChanges
	changes, err := adapter.TrackChanges()

	// Assert: Verify the results
	assert.NoError(t, err)
	assert.Len(t, changes, 2)
	assert.Equal(t, core.ChangeSet{
		ID:        1,
		TableName: "documents",
		Operation: "insert",
		RowID:     "row1",
		Timestamp: "2024-11-26 10:00:00",
	}, changes[0])
	assert.Equal(t, core.ChangeSet{
		ID:        2,
		TableName: "documents",
		Operation: "update",
		RowID:     "row2",
		Timestamp: "2024-11-26 10:05:00",
	}, changes[1])

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTrackChanges_NoChanges(t *testing.T) {
	// Arrange: Create a mocked database and adapter
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("error closing db: %v", err)
		}
	}(db)

	adapter := &SQLiteAdapter{db: db}

	rows := sqlmock.NewRows([]string{"id", "table_name", "operation", "row_id", "timestamp"})
	mock.ExpectQuery("SELECT id, table_name, operation, row_id, timestamp FROM changeset").
		WillReturnRows(rows)

	// Act: Call TrackChanges
	changes, err := adapter.TrackChanges()

	// Assert: Verify the results
	assert.NoError(t, err)
	assert.Len(t, changes, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTrackChanges_QueryError(t *testing.T) {
	// Arrange: Create a mocked database and adapter
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("error closing db: %v", err)
		}
	}(db)

	adapter := &SQLiteAdapter{db: db}

	mock.ExpectQuery("SELECT id, table_name, operation, row_id, timestamp FROM changeset").
		WillReturnError(errors.New("failed to query changesets: query failed"))

	// Act: Call TrackChanges
	changes, err := adapter.TrackChanges()

	// Assert: Verify the results
	assert.Error(t, err)
	assert.Nil(t, changes)
	assert.NoError(t, mock.ExpectationsWereMet())
}
