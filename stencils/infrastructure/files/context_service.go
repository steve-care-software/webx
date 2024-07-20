package files

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
	"github.com/steve-care-software/datastencil/stencils/domain/contexts"
)

type contextService struct {
}

func createContextService() contexts.Service {
	out := contextService{}
	return &out
}

// Save saves a context
func (app *contextService) Save(context contexts.Context) error {
	return nil
}

// Delete deletes a context
func (app *contextService) Delete(hash hash.Hash) error {
	return nil
}
