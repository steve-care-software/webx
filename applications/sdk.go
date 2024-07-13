package applications

import "github.com/steve-care-software/datastencil/domain/instances/executions"

const invalidPatternErr = "the provided context (%d) does not exists"

// Application represents an application
type Application interface {
	Init(dbPath []string, name string, description string) (*uint, error)
	Begin(dbPath []string) (*uint, error)
	Execute(context uint, input []byte) ([]byte, error)
	ExecuteWithPath(context uint, inputPath []string) ([]byte, error)
	ExecuteLayer(context uint, input []byte, layerPath []string) ([]byte, error)
	ExecuteLayerWithPath(context uint, inputPath []string, layerPath []string) ([]byte, error)
	Retrieve(context uint) (executions.Executions, error)
	Commit(context uint) error
	Rollback(context uint) error
	Cancel(context uint) error
}
