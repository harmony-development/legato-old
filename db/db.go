package db

// Auth is a database interface which manages sessions
type Auth interface {
	GetSession(session string) (int64, error)
	SetSession(session string, userID int64) error
}
