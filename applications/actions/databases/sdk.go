package databases

import "github.com/steve-care-software/syntax/applications/actions/databases/selects"

// Application represents a database application
type Application interface {
	List() ([]string, error)
	Select(name string) (selects.Application, error)
}
