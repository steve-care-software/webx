package files

import "github.com/steve-care-software/datastencil/domain/contexts"

type contextRepository struct {
}

func createContextRepository() contexts.Repository {
	out := contextRepository{}
	return &out
}

// Retrieve retrieves a context
func (app *contextRepository) Retrieve(dbPath []string) (contexts.Context, error) {
	return nil, nil
}
