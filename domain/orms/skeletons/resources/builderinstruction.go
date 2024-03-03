package resources

type builderInstruction struct {
	method        string
	containsParam bool
}

func createBuilderInstruction(
	method string,
	containsParam bool,
) BuilderInstruction {
	out := builderInstruction{
		method:        method,
		containsParam: containsParam,
	}

	return &out
}

// Method returns the method
func (obj *builderInstruction) Method() string {
	return obj.method
}

// ContainsParam returns true if there is params, false otherwise
func (obj *builderInstruction) ContainsParam() bool {
	return obj.containsParam
}
