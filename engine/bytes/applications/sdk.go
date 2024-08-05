package applications

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/branches"
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/iterations"
	"github.com/steve-care-software/webx/engine/bytes/domain/iterations/developments"
	"github.com/steve-care-software/webx/engine/bytes/domain/iterations/productions"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
	namespace_originals "github.com/steve-care-software/webx/engine/bytes/domain/namespaces/originals"
	"github.com/steve-care-software/webx/engine/bytes/domain/states"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithBasePath(basePath []string) Builder
	Now() (Application, error)
}

// Application represents the database application
type Application interface {
	Begin(name string) (*uint, error)
	Status(context uint) (string, error)

	// namespaces:
	Namespace(context uint, name string) (namespaces.Namespace, error)
	Namespaces(context uint) (namespaces.Namespaces, error)
	DeletedNamespaces(context uint) ([]string, error)
	SetNamespace(context uint, name string) error
	InsertNamespace(context uint, original namespace_originals.Original) error
	UpdateNamespace(context uint, original string, updated namespace_originals.Original) error
	DeleteNamespace(context uint, name string) error
	RecoverNamespace(context uint, name string) error
	PurgeNamespace(context uint, name string) error
	PurgeNamespaces(context uint) error

	// iterations:
	Iteration(context uint, name string) (iterations.Iteration, error)
	Iterations(context uint) (iterations.Iterations, error)
	DeletedIterations(context uint) ([]string, error)
	SetIteration(context uint, name string) error
	InsertIteration(context uint, name string, description string) error
	UpdateIteration(context uint, original string, updated namespace_originals.Original) error
	DeleteIteration(context uint, name string) error
	RecoverIteration(context uint, name string) error
	PurgeIteration(context uint, name string) error
	PurgeIterations(context uint) error

	// iteration development:
	IterationDevelopment(context uint, name string) (developments.Development, error)
	IterationDevelopments(context uint) (developments.Developments, error)
	DeletedIterationDevelopments(context uint) ([]string, error)
	SetIterationDevelopment(context uint, name string) error
	MoveIterationDevelopmentIntoProduction(context uint, name string, devName string, deleteOriginal bool) error
	InsertIterationDevelopment(context uint, name string, description string) error
	UpdateIterationDevelopment(context uint, original string, updated namespace_originals.Original) error
	DeleteIterationDevelopment(context uint, name string) error
	RecoverIterationDevelopment(context uint, name string) error
	PurgeIterationDevelopment(context uint, name string) error
	PurgeIterationDevelopments(context uint) error

	// iteration production:
	IterationProduction(context uint, name string) (productions.Production, error)
	IterationProductions(context uint) (productions.Productions, error)
	DeletedIterationProductions(context uint) ([]string, error)
	SetIterationProduction(context uint, name string) error
	InsertIterationProduction(context uint, name string, description string) error
	UpdateIterationProduction(context uint, original string, updated namespace_originals.Original) error
	DeleteIterationProduction(context uint, name string) error
	RecoverIterationProduction(context uint, name string) error
	PurgeIterationProduction(context uint, name string) error
	PurgeIterationProductions(context uint) error

	// branches:
	Branch(context uint, name string) (branches.Branch, error)
	Branches(context uint) (branches.Branches, error)
	DeletedBranches(context uint) ([]string, error)
	SetBranch(context uint, name string) error
	MergeBranch(context uint, deleteOriginal bool) error
	InsertBranch(context uint, name string, description string) error
	UpdateBranch(context uint, original string, updated namespace_originals.Original) error
	DeleteBranch(context uint, name string) error
	RecoverBranch(context uint, name string) error
	PurgeBranch(context uint, name string) error
	PurgeBranches(context uint) error

	// states:
	State(context uint, name string) (states.State, error)
	States(context uint) (states.States, error)
	DeletedStates(context uint) ([]string, error)
	SetState(context uint, name string) error
	InsertState(context uint, name string, description string) error
	UpdateState(context uint, original string, updated namespace_originals.Original) error
	DeleteState(context uint, name string) error
	RecoverState(context uint, name string) error
	PurgeState(context uint, name string) error
	PurgeStates(context uint) error

	// data:
	Retrieve(context uint, retrival delimiters.Delimiter) ([]byte, error)
	Insert(context uint, data []byte) (delimiters.Delimiter, error)
	Update(context uint, original delimiters.Delimiter, updated []byte) error
	Delete(context uint, delete delimiters.Delimiter) error

	// system:
	Commit(context uint) error
	CommitWithMetaData(context uint, metaData delimiters.Delimiter) error
	Close(context uint) error
	Cleanup(context uint) error
	Purge(context uint) error
}
