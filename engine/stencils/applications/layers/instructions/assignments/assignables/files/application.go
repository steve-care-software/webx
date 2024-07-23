package files

import (
	"github.com/steve-care-software/webx/engine/states/domain/files"
	assignable_files "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/files"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks/failures"
)

type application struct {
	filesRepository   files.Repository
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	filesRepository files.Repository,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		filesRepository:   filesRepository,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignment assignable_files.File) (stacks.Assignable, *uint, error) {
	if assignment.IsRead() {
		read := assignment.Read()
		identifierVar := read.Identifier()
		retFile, err := frame.FetchFile(identifierVar)
		if err != nil {
			code := failures.CouldNotFetchFileFromFrame
			return nil, &code, err
		}

		hasIndex := read.HasIndex()
		hasLength := read.HasLength()
		builder := app.assignableBuilder.Create()
		index := uint(0)
		if hasIndex {
			indexVar := read.Index()
			pIndex, err := frame.FetchUnsignedInt(indexVar)
			if err != nil {
				code := failures.CouldNotFetchUnsignedIntegerFromFrame
				return nil, &code, err
			}

			index = *pIndex
		}

		length := -1
		if hasLength {
			lengthVar := read.Length()
			pLength, err := frame.FetchUnsignedInt(lengthVar)
			if err != nil {
				code := failures.CouldNotFetchUnsignedIntegerFromFrame
				return nil, &code, err
			}

			length = int(*pLength)
		}

		if hasLength {
			retData, err := app.filesRepository.RetrieveChunk(retFile, index, uint(length))
			if err != nil {
				code := failures.CouldNotRetrieveChunkFromFile
				return nil, &code, err
			}

			builder.WithBytes(retData)
		}

		if hasIndex {
			retData, err := app.filesRepository.RetrieveFrom(retFile, uint(index))
			if err != nil {
				code := failures.CouldNotRetrieveFromFromFile
				return nil, &code, err
			}

			builder.WithBytes(retData)
		}

		if !hasIndex && !hasLength {
			retData, err := app.filesRepository.RetrieveAll(retFile)
			if err != nil {
				code := failures.CouldNotRetrieveAllFromFile
				return nil, &code, err
			}

			builder.WithBytes(retData)
		}

		ins, err := builder.Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, nil, nil
	}

	if assignment.IsLength() {
		lengthVar := assignment.Length()
		retFile, err := frame.FetchFile(lengthVar)
		if err != nil {
			code := failures.CouldNotFetchFileFromFrame
			return nil, &code, err
		}

		stat, err := retFile.Stat()
		if err != nil {
			return nil, nil, err
		}

		length := uint(stat.Size())
		ins, err := app.assignableBuilder.Create().WithUnsignedInt(length).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, nil, nil
	}

	if assignment.IsExists() {
		pathVar := assignment.Exists()
		retPath, err := frame.FetchList(pathVar)
		if err != nil {
			code := failures.CouldNotFetchListFromFrame
			return nil, &code, err
		}

		path := []string{}
		pathDirList := retPath.List()
		for _, oneDir := range pathDirList {
			if !oneDir.IsString() {
				code := failures.CouldNotFetchStringFromList
				return nil, &code, err
			}

			pStr := oneDir.String()
			path = append(path, *pStr)
		}

		exists := app.filesRepository.Exists(path)
		retAssignable, err := app.assignableBuilder.Create().WithBool(exists).Now()
		if err != nil {
			return nil, nil, err
		}

		return retAssignable, nil, nil
	}

	open := assignment.Open()
	pathVar := open.Path()
	retPath, err := frame.FetchList(pathVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, err
	}

	path := []string{}
	pathDirList := retPath.List()
	for _, oneDir := range pathDirList {
		if !oneDir.IsString() {
			code := failures.CouldNotFetchStringFromList
			return nil, &code, err
		}

		pStr := oneDir.String()
		path = append(path, *pStr)
	}

	permissionVar := open.Permission()
	permission, err := frame.FetchUnsignedInt(permissionVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, err
	}

	retFile, err := app.filesRepository.Open(path, *permission)
	if err != nil {
		code := failures.CouldNotOpenFile
		return nil, &code, err
	}

	retAssignable, err := app.assignableBuilder.Create().WithFilePointer(*retFile).Now()
	if err != nil {
		return nil, nil, err
	}

	return retAssignable, nil, nil
}
