package profile

import v1 "harmony-server/server/api/profile/v1"

// Dependencies are the backend services this package needs
type Dependencies struct {
}

// Service contains the profile service
type Service struct {
	*Dependencies
	V1 *v1.V1
}

// New creates a new profile service
func New(deps *Dependencies) *Service {
	service := &Service{
		Dependencies: deps,
	}
	service.V1 = &v1.V1{}
	return service
}
