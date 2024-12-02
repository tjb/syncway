package core

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type SyncManager struct {
	adapter SQLSyncWrapper
	socket  *websocket.Conn
}

func NewSyncManager(adapter SQLSyncWrapper, socketUri string) (*SyncManager, error) {
	// TODO: Change socketUri
	conn, _, err := websocket.DefaultDialer.Dial(socketUri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to sync engine: %w", err)
	}

	return &SyncManager{
		adapter: adapter,
		socket:  conn,
	}, nil
}

// InitSync - Start real-time sync process aka listen for changes
func (sm *SyncManager) InitSync() error {
	return nil
}

// sendChangesToEngine - Get changes and send to engine
func (sm *SyncManager) sendChangesToEngine() error {
	return nil
}

// listenForChanges - exactly as the function is named
func (sm *SyncManager) listenForChanges() error {
	return nil
}
