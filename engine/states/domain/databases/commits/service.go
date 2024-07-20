package commits

import "github.com/steve-care-software/webx/engine/states/domain/files"

type service struct {
	adapter     Adapter
	fileService files.Service
}

func createService(
	adapter Adapter,
	fileService files.Service,
) Service {
	out := service{
		adapter:     adapter,
		fileService: fileService,
	}

	return &out
}

// Save saves a commit
func (app *service) Save(ins Commit) error {
	bytes, err := app.adapter.ToBytes(ins)
	if err != nil {
		return err
	}

	return app.fileService.Transact([]string{
		ins.Hash().String(),
	}, bytes)
}
