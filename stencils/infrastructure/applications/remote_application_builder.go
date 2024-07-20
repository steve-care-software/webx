package applications

import "github.com/steve-care-software/datastencil/stencils/applications"

type remoteApplicationBuilder struct {
	host string
}

func createRemoteApplicationBuilder() applications.RemoteBuilder {
	out := remoteApplicationBuilder{
		host: "",
	}

	return &out
}

// Create initializes the builder
func (app *remoteApplicationBuilder) Create() applications.RemoteBuilder {
	return createRemoteApplicationBuilder()
}

// WithHost adds a host to the builder
func (app *remoteApplicationBuilder) WithHost(host string) applications.RemoteBuilder {
	app.host = host
	return app
}

// Now builds a new Application instance
func (app *remoteApplicationBuilder) Now() (applications.Application, error) {
	return nil, nil
}
