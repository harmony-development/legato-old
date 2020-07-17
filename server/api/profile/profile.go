package profile

import (
	v1 "github.com/harmony-development/legato/server/api/profile/v1"
	"github.com/harmony-development/legato/server/db"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB db.IHarmonyDB
}

// Service contains the profile service
type Service struct {
	Dependencies
	V1 v1.V1
}

// New creates a new profile service
func New(deps Dependencies) *Service {
	service := &Service{
		Dependencies: deps,
	}
	service.V1 = v1.New(v1.Dependencies{
		DB: service.DB,
	})
	return service
}
