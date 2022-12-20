package clients

import (
	"net/url"

	"github.com/steve-care-software/webx/databases/applications"
)

const patternURI = "%s/%s"
const rootURI = "/"
const existsURI = "/exists"
const connectionsURI = "/connections"
const contextURI = "/context"
const contentKeysByKindURI = "/content-keys/%d/%d"
const commitsURI = "/commits/%d"
const commitByHashURI = "/commits/%d/%s"

// Builder represents the client database application builder
type Builder interface {
	Create() Builder
	WithServer(server *url.URL) Builder
	Now() (applications.Application, error)
}
