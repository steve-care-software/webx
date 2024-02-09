package signers

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToSigner converts bytes to a Signer instance
func (app *adapter) ToSigner(pk []byte) (Signer, error) {
	scalar, err := fromBytesToScalar(pk)
	if err != nil {
		return nil, err
	}

	return createSigner(scalar), nil
}
