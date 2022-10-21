package commands

import "github.com/steve-care-software/webx/domain/criterias"

type moduleDeclaration struct {
	name criterias.Criteria
}

func createModuleDeclaration(
	name criterias.Criteria,
) ModuleDeclaration {
	out := moduleDeclaration{
		name: name,
	}

	return &out
}

// Name returns the name
func (obj *moduleDeclaration) Name() criterias.Criteria {
	return obj.name
}
