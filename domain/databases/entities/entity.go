package entities

type entity struct {
	identifier Identifier
	signature  Signature
}

func createEntity(
	identifier Identifier,
	signature Signature,
) Entity {
	out := entity{
		identifier: identifier,
		signature:  signature,
	}

	return &out
}

// Identifier returns the identifier
func (obj *entity) Identifier() Identifier {
	return obj.identifier
}

// Signature returns the signature
func (obj *entity) Signature() Signature {
	return obj.signature
}
