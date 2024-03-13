package conditions

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	hashAdapter := hash.NewAdapter()
	return createPointerBuilder(
		hashAdapter,
	)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(
		hashAdapter,
	)
}

// NewResourceBuilder creates a new resource builder
func NewResourceBuilder() ResourceBuilder {
	hashAdapter := hash.NewAdapter()
	return createResourceBuilder(
		hashAdapter,
	)
}

// NewOperatorBuilder creates a new operator builder
func NewOperatorBuilder() OperatorBuilder {
	hashAdapter := hash.NewAdapter()
	return createOperatorBuilder(
		hashAdapter,
	)
}

// NewRelationalOperatorBuilder creates a new relational operator builder
func NewRelationalOperatorBuilder() RelationalOperatorBuilder {
	hashAdapter := hash.NewAdapter()
	return createRelationalOperatorBuilder(
		hashAdapter,
	)
}

// NewIntegerOperatorBuilder creates a new integer operator builder
func NewIntegerOperatorBuilder() IntegerOperatorBuilder {
	hashAdapter := hash.NewAdapter()
	return createIntegerOperatorBuilder(
		hashAdapter,
	)
}

// Builder represents a condition builder
type Builder interface {
	Create() Builder
	WithPointer(pointer Pointer) Builder
	WithOperator(operator Operator) Builder
	WithElement(element Element) Builder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Pointer() Pointer
	Operator() Operator
	Element() Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithCondition(condition Condition) ElementBuilder
	WithResource(resource Resource) ElementBuilder
	Now() (Element, error)
}

// Element represents a conditional element
type Element interface {
	Hash() hash.Hash
	IsCondition() bool
	Condition() Condition
	IsResource() bool
	Resource() Resource
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithField(field Pointer) ResourceBuilder
	WithValue(value interface{}) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	IsField() bool
	Field() Pointer
	IsValue() bool
	Value() interface{}
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithEntity(entity string) PointerBuilder
	WithField(field string) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a field pointer
type Pointer interface {
	Hash() hash.Hash
	Entity() string
	Field() string
}

// OperatorBuilder represents an operator builder
type OperatorBuilder interface {
	Create() OperatorBuilder
	WithRelational(relational RelationalOperator) OperatorBuilder
	WithInteger(integer IntegerOperator) OperatorBuilder
	IsEqual() OperatorBuilder
	Now() (Operator, error)
}

// Operator represents an operator
type Operator interface {
	Hash() hash.Hash
	IsEqual() bool
	IsRelational() bool
	Relational() RelationalOperator
	IsInteger() bool
	Integer() IntegerOperator
}

// RelationalOperatorBuilder represents a relational operator builder
type RelationalOperatorBuilder interface {
	Create() RelationalOperatorBuilder
	IsAnd() RelationalOperatorBuilder
	IsOr() RelationalOperatorBuilder
	Now() (RelationalOperator, error)
}

// RelationalOperator represents a relational operator
type RelationalOperator interface {
	Hash() hash.Hash
	IsAnd() bool
	IsOr() bool
}

// IntegerOperatorBuilder represents the integer operator builder
type IntegerOperatorBuilder interface {
	Create() IntegerOperatorBuilder
	IsSmallerThan() IntegerOperatorBuilder
	IsBiggerThan() IntegerOperatorBuilder
	IsEqual() IntegerOperatorBuilder
	Now() (IntegerOperator, error)
}

// IntegerOperator represents an integer operator
type IntegerOperator interface {
	Hash() hash.Hash
	IsSmallerThan() bool
	IsBiggerThan() bool
	IsEqual() bool
}
