package files

import (
	"os"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/contexts"
)

type contextService struct {
	adapter     contexts.Adapter
	hashAdapter hash.Adapter
	basePath    []string
	endPath     []string
}

func createContextService(
	adapter contexts.Adapter,
	hashAdapter hash.Adapter,
	basePath []string,
	endPath []string,
) contexts.Service {
	out := contextService{
		adapter:     adapter,
		hashAdapter: hashAdapter,
		basePath:    basePath,
		endPath:     endPath,
	}

	return &out
}

// Save saves a context
func (app *contextService) Save(dbPath []string, context contexts.Context) error {
	filePath, err := prepareFilePath(app.hashAdapter, dbPath, app.basePath, app.endPath)
	if err != nil {
		return err
	}

	bytes, err := app.adapter.ToBytes(context)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, bytes, readWritePermissionBits)
}

// Delete deletes a context
func (app *contextService) Delete(dbPath []string) error {
	filePath, err := prepareFilePath(app.hashAdapter, dbPath, app.basePath, app.endPath)
	if err != nil {
		return err
	}

	return os.Remove(filePath)
}
