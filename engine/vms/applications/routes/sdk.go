package routes

import "github.com/steve-care-software/webx/engine/vms/domain/instances/routes"

// NewApplication creates a new application
func NewApplication(
	routeRepository routes.Repository,
) Application {
	return createApplication(
		routeRepository,
	)
}

// Application represents the route application
type Application interface {
	Execute(input []byte, route routes.Route) (bool, []byte, error)
}
