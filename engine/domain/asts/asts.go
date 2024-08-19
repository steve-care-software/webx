package asts

type asts struct {
	list []AST
	mp   map[string]AST
}

func createASTs(
	list []AST,
) ASTs {
	out := asts{
		list: list,
	}

	return &out
}

// List returns the list of ast
func (obj *asts) List() []AST {
	return obj.list
}
