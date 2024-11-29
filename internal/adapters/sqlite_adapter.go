package adapters

import (
	"database/sql"
	"fmt"
	"syncway/internal/core"

	_ "github.com/mattn/go-sqlite3"
)

// Compile time check. I am still fuzzy on interfaces in Go.
var _ core.SQLSyncWrapper = (*SQLiteAdapter)(nil)

type SQLiteAdapter struct {
	db *sql.DB
}

func setupChangeTracking(db *sql.DB) error {
	// Create the changeset table
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS changeset (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            table_name TEXT NOT NULL,
            operation TEXT NOT NULL,
            row_id TEXT NOT NULL,
            timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return fmt.Errorf("failed to create changeset table: %w", err)
	}

	fmt.Println("Changeset table created")

	// Create triggers for change tracking on example table 'documents'
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS documents (
            id TEXT PRIMARY KEY,
            name TEXT,
            content TEXT
        );

        CREATE TRIGGER IF NOT EXISTS after_insert_documents
        AFTER INSERT ON documents
        BEGIN
            INSERT INTO changeset (table_name, operation, row_id)
            VALUES ('documents', 'insert', NEW.id);
        END;

        CREATE TRIGGER IF NOT EXISTS after_update_documents
        AFTER UPDATE ON documents
        BEGIN
            INSERT INTO changeset (table_name, operation, row_id)
            VALUES ('documents', 'update', NEW.id);
        END;

        CREATE TRIGGER IF NOT EXISTS after_delete_documents
        AFTER DELETE ON documents
        BEGIN
            INSERT INTO changeset (table_name, operation, row_id)
            VALUES ('documents', 'delete', OLD.id);
        END;
    `)
	if err != nil {
		return fmt.Errorf("failed to set up triggers: %w", err)
	}

	return nil
}

func NewSQLiteAdapter(connString string) (*SQLiteAdapter, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite connection: %w", err)
	}
	err = setupChangeTracking(db)
	if err != nil {
		return nil, fmt.Errorf("failed to setup change tracking: %w", err)
	}

	return &SQLiteAdapter{db: db}, nil
}

// TrackChanges
// Logic for fetching changes from the changeset table
func (s *SQLiteAdapter) TrackChanges() ([]core.ChangeSet, error) {
	// Logic for fetching changes from the changeset table
	return nil, nil
}

// ApplyChanges
// This method applies changes that have come from the Sync Engine.
func (s *SQLiteAdapter) ApplyChanges(changes []core.ChangeSet) error {
	return nil
}
