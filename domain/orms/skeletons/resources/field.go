package resources

type field struct {
	name      string
	retriever []string
	kind      Kind
	canBeNil  bool
	condition string
	builder   BuilderInstruction
}

func createField(
	name string,
	retriever []string,
	kind Kind,
	canBeNil bool,
) Field {
	return createFieldInternally(name, retriever, kind, canBeNil, "", nil)
}

func createFieldWithCondition(
	name string,
	retriever []string,
	kind Kind,
	canBeNil bool,
	condition string,
) Field {
	return createFieldInternally(name, retriever, kind, canBeNil, condition, nil)
}

func createFieldWithBuilder(
	name string,
	retriever []string,
	kind Kind,
	canBeNil bool,
	builder BuilderInstruction,
) Field {
	return createFieldInternally(name, retriever, kind, canBeNil, "", builder)
}

func createFieldWithConditionAndBuilder(
	name string,
	retriever []string,
	kind Kind,
	canBeNil bool,
	condition string,
	builder BuilderInstruction,
) Field {
	return createFieldInternally(name, retriever, kind, canBeNil, condition, builder)
}

func createFieldInternally(
	name string,
	retriever []string,
	kind Kind,
	canBeNil bool,
	condition string,
	builder BuilderInstruction,
) Field {
	out := field{
		name:      name,
		retriever: retriever,
		kind:      kind,
		canBeNil:  canBeNil,
		condition: condition,
		builder:   builder,
	}

	return &out
}

// Name returns the name
func (obj *field) Name() string {
	return obj.name
}

// Retriever returns the retriever
func (obj *field) Retriever() []string {
	return obj.retriever
}

// Kind returns the kind
func (obj *field) Kind() Kind {
	return obj.kind
}

// CanBeNil returns true if canBeNil, false otherwise
func (obj *field) CanBeNil() bool {
	return obj.canBeNil
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *field) HasCondition() bool {
	return obj.condition != ""
}

// Condition returns the condition, if any
func (obj *field) Condition() string {
	return obj.condition
}

// HasBuilder returns true if there is a builder, false otherwise
func (obj *field) HasBuilder() bool {
	return obj.builder != nil
}

// Builder returns the builder
func (obj *field) Builder() BuilderInstruction {
	return obj.builder
}
