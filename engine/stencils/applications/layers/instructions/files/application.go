package files

import (
	"github.com/steve-care-software/webx/engine/states/domain/files"
	instruction_files "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/files"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks/failures"
)

type application struct {
	fileService files.Service
}

func cretaeApplication(
	fileService files.Service,
) Application {
	out := application{
		fileService: fileService,
	}

	return &out
}

// Execute execute the application
func (app *application) Execute(frame stacks.Frame, file instruction_files.File) (*uint, error) {
	if file.IsClose() {
		identifierVar := file.Close()
		retFile, err := frame.FetchFile(identifierVar)
		if err != nil {
			code := failures.CouldNotFetchFileFromFrame
			return &code, err
		}

		err = retFile.Close()
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	pathVar := file.Delete()
	retPath, err := frame.FetchList(pathVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return &code, err
	}

	path := []string{}
	pathDirList := retPath.List()
	for _, oneDir := range pathDirList {
		if !oneDir.IsString() {
			code := failures.CouldNotFetchStringFromList
			return &code, err
		}

		pStr := oneDir.String()
		path = append(path, *pStr)
	}

	err = app.fileService.Delete(path)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
