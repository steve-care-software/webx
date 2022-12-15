package applications

import (
	blockchain_applications "github.com/steve-care-software/webx/blockchains/applications"
	compiler_applications "github.com/steve-care-software/webx/compilers/applications"
	grammar_applications "github.com/steve-care-software/webx/roots/applications/grammars"
	"github.com/steve-care-software/webx/roots/domain/grammars/grammars"
	program_applications "github.com/steve-care-software/webx/programs/applications"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	"github.com/steve-care-software/webx/roots/domain/roots/compilers"
	roots_grammar "github.com/steve-care-software/webx/roots/domain/roots/grammars"
	"github.com/steve-care-software/webx/roots/domain/roots/programs"
	"github.com/steve-care-software/webx/roots/domain/roots/selectors"
	selector_applications "github.com/steve-care-software/webx/roots/applications/selectors"
)

const (
	// GrammarDatabaseKind represents the grammar database kind
	GrammarDatabaseKind = iota
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithBlockchain(blockchain blockchain_applications.Application) Builder
	WithModules(modules modules.Modules) Builder
	Now() (Application, error)
}

// Application represents a root application
type Application interface {
	New(context uint, name string) error
	Database
}

// Database represents the root database application
type Database interface {
	Grammar() Grammar
	Program() Program
	Selector() Selector
	Compiler() Compiler
}

// Grammar represents the grammar database application
type Grammar interface {
	Retrieve(context uint) (roots_grammar.Grammar, error)
	Application(context uint) (grammar_applications.Application, error)
}

// Program represents the program database application
type Program interface {
	List() ([]string, error)
	New(name string, modules []uint) error
	Retrieve(name string) (programs.Program, error)
	Application(name string) (program_applications.Application, error)
	Delete(name string) error
}

// Selector represebrs the root selector database application
type Selector interface {
	List() ([]string, error)
	New(name string, grammar grammars.Grammar) error
	Retrieve(name string) (selectors.Selector, error)
	Application(name string) (selector_applications.Application, error)
	Delete(name string) error
}

// Compiler represents the root compiler database application
type Compiler interface {
	List() ([]string, error)
	New(name string, modules []uint, selectors []string, programs []string) error
	Retrieve(name string) (compilers.Compiler, error)
	Application(name string) (compiler_applications.Application, error)
	Delete(name string) error
}
