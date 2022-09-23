package grammars

type external struct {
	name    string
	grammar Grammar
}

func createExternal(
	name string,
	grammar Grammar,
) External {
	out := external{
		name:    name,
		grammar: grammar,
	}

	return &out
}

// Name returns the name
func (obj *external) Name() string {
	return obj.name
}

// Grammar returns the grammar
func (obj *external) Grammar() Grammar {
	return obj.grammar
}
