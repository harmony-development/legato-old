package core

import (
	v1 "harmony-server/server/api/core/v1"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
}

// Service contains the core service
type Service struct {
	*Dependencies
	V1 *v1.V1
}

// New creates a new core service
func New(deps *Dependencies) *Service {
	core := &Service{
		Dependencies: deps,
	}
	core.V1 = &v1.V1{}
	return core
}
