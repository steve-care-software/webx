package commands

import "github.com/steve-care-software/syntax/domain/bytes/criterias"

type applicationDeclaration struct {
	module criterias.Criteria
	name   criterias.Criteria
}

func createApplicationDeclaration(
	module criterias.Criteria,
	name criterias.Criteria,
) ApplicationDeclaration {
	out := applicationDeclaration{
		module: module,
		name:   name,
	}

	return &out
}

// Module returns the module
func (obj *applicationDeclaration) Module() criterias.Criteria {
	return obj.module
}

// Name returns the name
func (obj *applicationDeclaration) Name() criterias.Criteria {
	return obj.name
}
