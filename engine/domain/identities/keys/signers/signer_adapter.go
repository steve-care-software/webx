package signers

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToSigner converts a string to a signer instance
func (app *adapter) ToSigner(pk string) (Signer, error) {
	scalar, err := fromStringToScalar(pk)
	if err != nil {
		return nil, err
	}

	return createSigner(scalar), nil
}
