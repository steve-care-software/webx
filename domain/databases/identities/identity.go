package identities

import "github.com/steve-care-software/webx/domain/databases/entities"

type identity struct {
	identifier    entities.Identifier
	modifications entities.Identifiers
}

func createIdentity(
	identifier entities.Identifier,
	modifications entities.Identifiers,
) Identity {
	out := identity{
		identifier:    identifier,
		modifications: modifications,
	}

	return &out
}

// Identifier returns the identifier
func (obj *identity) Identifier() entities.Identifier {
	return obj.identifier
}

// Modifications retruns the modifications
func (obj *identity) Modifications() entities.Identifiers {
	return obj.modifications
}
