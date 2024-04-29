package repositories

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases/repositories"
)

// Adapter represents a repository adapter
type Adapter struct {
	builder repositories.Builder
}

func createAdapter(
	builder repositories.Builder,
) repositories.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins repositories.Repository) ([]byte, error) {
	ptr, err := app.RepositoryToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (repositories.Repository, error) {
	ins := new(Repository)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToRepository(*ins)
}

// RepositoryToStruct converts a repository to struct
func (app *Adapter) RepositoryToStruct(ins repositories.Repository) (*Repository, error) {
	out := Repository{}
	if ins.IsHeight() {
		out.IsHeight = true
	}

	if ins.IsSkeleton() {
		out.IsSkeleton = true
	}

	if ins.IsList() {
		out.List = ins.List()
	}

	if ins.IsRetrieve() {
		out.Retrieve = ins.Retrieve()
	}

	return &out, nil
}

// StructToRepository converts a struct to repository
func (app *Adapter) StructToRepository(str Repository) (repositories.Repository, error) {
	builder := app.builder.Create()
	if str.IsSkeleton {
		builder.IsSkeleton()
	}

	if str.IsHeight {
		builder.IsHeight()
	}

	if str.List != "" {
		builder.WithList(str.List)
	}

	if str.Retrieve != "" {
		builder.WithRetrieve(str.Retrieve)
	}

	return builder.Now()
}
