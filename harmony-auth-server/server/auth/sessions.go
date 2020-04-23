package auth

import (
	"github.com/thanhpk/randstr"
	"sync"
	"time"
)

// Sessions is a thread-safe structure for storing user sessions
type Sessions struct {
	SessionMap map[string]Session
	Mut        *sync.RWMutex
}

// Session carries the data for a specific session
type Session struct {
	Expiration time.Time
	UserID     string
}

// MakeSession creates a new session and stores it in memory
func (s Sessions) MakeSession(userID string, expire time.Duration) string {
	s.Mut.Lock()
	defer s.Mut.Unlock()
	session := randstr.Hex(32)
	s.SessionMap[session] = Session{
		Expiration: time.Now().Add(expire),
		UserID:     userID,
	}
	return session
}

// Exists returns if a specific session exists in memory
func (s Sessions) Exists(session string) bool {
	s.Mut.RLock()
	defer s.Mut.RUnlock()
	_, exists := s.SessionMap[session]
	return exists
}

// GetSession returns a session object
func (s Sessions) GetSession(session string) (*Session, bool) {
	s.Mut.RLock()
	defer s.Mut.RUnlock()
	if sess, exists := s.SessionMap[session]; exists {
		return &sess, exists
	}
	return nil, false
}

// ExpireSessions cleans up expired sessions
func (s Sessions) ExpireSessions() {
	for {
		s.Mut.Lock()
		for k := range s.SessionMap {
			if time.Now().After(s.SessionMap[k].Expiration) {
				delete(s.SessionMap, k)
			}
		}
		s.Mut.Unlock()
		time.Sleep(10 * time.Minute)
	}
}
