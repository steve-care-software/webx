package attachments

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewVariableBuilder creates a new variable builder
func NewVariableBuilder() VariableBuilder {
	return createVariableBuilder()
}

// Builder represents an attachment builder
type Builder interface {
	Create() Builder
	WithVariable(variable Variable) Builder
	WithApplication(application []byte) Builder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Variable() Variable
	Application() []byte
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithCurrent(current []byte) VariableBuilder
	WithTarget(target uint) VariableBuilder
	Now() (Variable, error)
}

// Variable represents an attachment variable
type Variable interface {
	Current() []byte
	Target() uint
}
