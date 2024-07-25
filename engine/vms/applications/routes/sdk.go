package routes

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
)

// Application represents a route application
type Application interface {
	Execute(input []byte, route routes.Route) (layers.Layer, error)
}
