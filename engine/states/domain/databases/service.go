package databases

import (
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/pointers"
	"github.com/steve-care-software/webx/engine/states/domain/files"
)

type service struct {
	repository     Repository
	fileService    files.Service
	commitService  commits.Service
	pointerAdapter pointers.Adapter
	pointerBuilder pointers.Builder
}

func createService(
	repository Repository,
	fileService files.Service,
	commitService commits.Service,
	pointerAdapter pointers.Adapter,
	pointerBuilder pointers.Builder,
) Service {
	out := service{
		repository:     repository,
		fileService:    fileService,
		commitService:  commitService,
		pointerAdapter: pointerAdapter,
		pointerBuilder: pointerBuilder,
	}

	return &out
}

// Begin begins a transaction
func (app *service) Begin(path []string) error {
	if !app.repository.Exists(path) {
		err := app.fileService.Init(path)
		if err != nil {
			return err
		}
	}

	err := app.fileService.Lock(path)
	if err != nil {
		return err
	}
	return nil
}

// Save saves a database
func (app *service) Save(ins Database) error {
	err := app.commitService.Save(ins.Head())
	if err != nil {
		return err
	}

	headHash := ins.Head().Hash()
	metaData := ins.MetaData()
	pointer, err := app.pointerBuilder.Create().
		WithHead(headHash).
		WithMetaData(metaData).
		Now()

	if err != nil {
		return err
	}

	bytes, err := app.pointerAdapter.ToBytes(pointer)
	if err != nil {
		return err
	}

	path := ins.MetaData().Path()
	return app.fileService.Save(path, bytes)
}

// SaveAll saves all databases
func (app *service) SaveAll(list []Database) error {
	for _, oneIns := range list {
		err := app.Save(oneIns)
		if err != nil {
			return err
		}
	}

	return nil
}

// End ends a transaction
func (app *service) End(path []string) error {
	return app.fileService.Unlock(path)
}
