package entities

type entity struct {
	identifier Identifier
	reference  Reference
	signature  Signature
}

func createEntity(
	identifier Identifier,
	reference Reference,
	signature Signature,
) Entity {
	out := entity{
		identifier: identifier,
		reference:  reference,
		signature:  signature,
	}

	return &out
}

// Identifier returns the identifier
func (obj *entity) Identifier() Identifier {
	return obj.identifier
}

// Reference returns the reference
func (obj *entity) Reference() Reference {
	return obj.reference
}

// Signature returns the signature
func (obj *entity) Signature() Signature {
	return obj.signature
}
