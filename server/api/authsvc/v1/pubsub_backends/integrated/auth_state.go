package authstate

import (
	"errors"
	"sync"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/server/api/authsvc/v1/authsteps"
	"github.com/harmony-development/legato/server/logger"
)

type AuthSessionState struct {
	currentStep authsteps.Step
}

// AuthState ...
type AuthState struct {
	authChannels  map[authv1.AuthService_StreamStepsServer]chan struct{}
	authEvents    map[string][]authv1.AuthService_StreamStepsServer
	sessionStates map[string]AuthSessionState
	Logger        logger.ILogger
	sync.Mutex
}

func New(l logger.Logger) *AuthState {
	return &AuthState{
		Logger:        l,
		authChannels:  make(map[authv1.AuthService_StreamStepsServer]chan struct{}),
		authEvents:    make(map[string][]authv1.AuthService_StreamStepsServer),
		sessionStates: make(map[string]AuthSessionState),
	}
}

// NewAuthSession ...
func (h *AuthState) NewAuthSession(authID string, step authsteps.Step) error {
	h.Lock()
	defer h.Unlock()

	if _, exists := h.sessionStates[authID]; exists {
		return errors.New("session already exists")
	}

	h.sessionStates[authID] = AuthSessionState{
		currentStep: step,
	}

	return nil
}

func (h *AuthState) AuthSessionExists(authID string) bool {
	h.Lock()
	defer h.Unlock()

	_, exists := h.sessionStates[authID]

	return exists
}

func (h *AuthState) HasStream(authID string) bool {
	h.Lock()
	defer h.Unlock()

	_, exists := h.authEvents[authID]

	if !exists {
		return false
	}

	return len(h.authEvents[authID]) > 0
}

// Subscribe ...
func (h *AuthState) Subscribe(authID string, s authv1.AuthService_StreamStepsServer) (chan struct{}, error) {
	h.Lock()
	defer h.Unlock()

	if _, ok := h.authEvents[authID]; !ok {
		return nil, errors.New("no auth event")
	}

	go func() {
		<-s.Context().Done()
		h.Unsubscribe(authID, s)
	}()

	h.authChannels[s] = make(chan struct{})
	h.authEvents[authID] = append(h.authEvents[authID], s)

	return h.authChannels[s], nil
}

// Unsubscribe ...
func (h *AuthState) Unsubscribe(authID string, s authv1.AuthService_StreamStepsServer) {
	h.Lock()
	defer h.Unlock()

	val, ok := h.authEvents[authID]
	_ = ok
	for idx, serv := range val {
		if serv == s {
			val[idx] = val[len(val)-1]
			val[len(val)-1] = nil
			val = val[:len(val)-1]
			break
		}
	}
	close(h.authChannels[s])
	delete(h.authChannels, s)
	h.authEvents[authID] = val
}

// Broadcast ...
func (h *AuthState) Broadcast(authID string, e *authv1.AuthStep) {
	h.Lock()
	defer h.Unlock()

	val, ok := h.authEvents[authID]
	_ = ok
	for _, serv := range val {
		if err := serv.Send(e); err != nil {
			println(err)
		}
	}
}

// SetStep ...
func (h *AuthState) SetStep(authID string, newStep authsteps.Step) {
	h.Lock()
	defer h.Unlock()

	val, ok := h.sessionStates[authID]
	if !ok {
		return
	}
	val.currentStep = newStep
}

func (h *AuthState) GetStep(authID string) authsteps.Step {
	h.Lock()
	defer h.Unlock()

	return h.sessionStates[authID].currentStep
}

func (h *AuthState) DeleteAuthSession(authID string) {
	h.Lock()
	defer h.Unlock()

	delete(h.sessionStates, authID)
	if _, exists := h.authEvents[authID]; exists {
		for _, s := range h.authEvents[authID] {
			h.Unsubscribe(authID, s)
		}
	}
}
