package versions

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions"

// Application represents a version application
type Application interface {
	List(input versions.Version) []string
	Load(input versions.Version, name string) (versions.Version, error)
	Loaded(input versions.Version) ([]string, error)
	Create(input versions.Version, name string, description string) (versions.Version, error)
	Set(input versions.Version, name string) (versions.Version, error)
}
