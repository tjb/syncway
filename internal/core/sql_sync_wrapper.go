package core

// ChangeSet represents a single change to be synced
type ChangeSet struct {
	ID        int64                  // Unique identifier for the change
	TableName string                 // Name of the table affected
	Operation string                 // Type of operation: insert, update, delete
	RowID     string                 // Primary key of the affected row
	Timestamp string                 // When the change occurred
	Data      map[string]interface{} // Key-value representation of the row's data
}

type SQLSyncWrapper interface {
	// TrackChanges fetches all pending changes from the local database
	TrackChanges() ([]ChangeSet, error)

	// ApplyChanges applies a set of changes to the local database
	ApplyChanges([]ChangeSet) error

	// SyncWithEngine sends local changes to the sync engine and applies remote changes
	SyncWithEngine(serverURL string) error
}
