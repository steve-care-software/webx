package compilers

import (
	"github.com/steve-care-software/webx/roots/domain/programs/programs"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewExecutionsBuilder creates a new executions builder
func NewExecutionsBuilder() ExecutionsBuilder {
	return createExecutionsBuilder()
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	return createExecutionBuilder()
}

// NewParametersBuilder creates a new parameters builder
func NewParametersBuilder() ParametersBuilder {
	return createParametersBuilder()
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	return createParameterBuilder()
}

// Builder represents a compiler builder
type Builder interface {
	Create() Builder
	WithExecutions(executions Executions) Builder
	WithOutputs(outputs []uint) Builder
	Now() (Compiler, error)
}

// Compiler represents a compiler
type Compiler interface {
	Executions() Executions
	HasOutputs() bool
	Outputs() []uint
}

// ExecutionsBuilder represents an executions builder
type ExecutionsBuilder interface {
	Create() ExecutionsBuilder
	WithList(list []Execution) ExecutionsBuilder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	List() []Execution
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithParameters(parameters Parameters) ExecutionBuilder
	WithProgram(program programs.Program) ExecutionBuilder
	WithExecuteProgramModule(execProgramModule uint) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an instructions execution
type Execution interface {
	Parameters() Parameters
	Program() programs.Program
	ExecuteProgramModule() uint
}

// ParametersBuilder represents a parameters builder
type ParametersBuilder interface {
	Create() ParametersBuilder
	WithList(list []Parameter) ParametersBuilder
	Now() (Parameters, error)
}

// Parameters represents parameters
type Parameters interface {
	List() []Parameter
}

// ParameterBuilder represents a parameter value
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithIndex(index uint) ParameterBuilder
	WithSelector(selector selectors.Selector) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a parameter
type Parameter interface {
	Index() uint
	Selector() selectors.Selector
}
