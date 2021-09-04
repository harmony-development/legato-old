package db

// SessionDB is a database interface which manages sessions
type SessionDB interface {
	GetSession(session string) (int64, error)
	SetSession(session string, userID int64) error
}

// AuthDB manages step states
type AuthDB interface {
	GetCurrentStep(authID string) (string, error)
	SetStep(authID string, step string) error
}
