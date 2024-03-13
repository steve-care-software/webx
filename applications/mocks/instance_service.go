package mocks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

type instanceService struct {
	pBeginContext *uint
	fails         bool
}

func createInstanceService(
	pBeginContext *uint,
	fails bool,
) instances.Service {
	out := instanceService{
		pBeginContext: pBeginContext,
		fails:         fails,
	}

	return &out
}

// Init initializes the service
func (app *instanceService) Init() error {
	return nil
}

// Begin begins a transaction
func (app *instanceService) Begin() (*uint, error) {
	if app.pBeginContext == nil {
		return nil, errors.New("the begin was expected to fail")
	}

	return app.pBeginContext, nil
}

// Insert inserts an instance
func (app *instanceService) Insert(context uint, ins instances.Instance, path []string) error {
	return app.serviceFails()
}

// Delete deletes an instance
func (app *instanceService) Delete(context uint, path []string, hash hash.Hash) error {
	return app.serviceFails()
}

// Commit commits actions
func (app *instanceService) Commit(context uint) error {
	return app.serviceFails()
}

// Cancel cancels a context
func (app *instanceService) Cancel(context uint) error {
	return app.serviceFails()
}

// Revert reverts the state of the last commit
func (app *instanceService) Revert() error {
	return app.serviceFails()
}

// Revert reverts the state of the commit to the provided index
func (app *instanceService) RevertToIndex(toIndex uint) error {
	return app.serviceFails()
}

func (app *instanceService) serviceFails() error {
	if app.fails {
		return errors.New("the service was expected to fail")
	}

	return nil
}
