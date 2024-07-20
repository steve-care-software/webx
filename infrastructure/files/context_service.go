package files

import (
	"github.com/steve-care-software/datastencil/domain/contexts"
	"github.com/steve-care-software/historydb/domain/hash"
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
