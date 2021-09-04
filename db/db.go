package db

// Auth is a database interface which manages sessions
type Sessions interface {
	GetSession(session string) (int64, error)
	SetSession(session string, userID int64) error
}

type Auth interface {
	GetCurrentStep(authID string) (string, error)
	SetStep(authID string, step string) error
}

type DB interface {
	Sessions
	Auth
}
