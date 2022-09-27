package composers

import (
	"reflect"

	"github.com/steve-care-software/syntax/domain/syntax/compilers"
)

// Builder represents a composer builder
type Builder interface {
	Create() Builder
	WithValue(value reflect.Value) Builder
	WithCompiler(compiler compilers.Compiler) Builder
	WithElements(elements Elements) Builder
	Now() (Composer, error)
}

// Composer represents a composer
type Composer interface {
	Value() reflect.Value
	Compiler() compilers.Compiler
	Elements() Elements
}

// ElementsBuilder represents an elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents an elements
type Elements interface {
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithName(name []byte) ElementBuilder
	WithMethods(methods []string) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() []byte
	Methods() []string
}
