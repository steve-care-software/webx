package signers

type factory struct {
}

func createFactory() Factory {
	out := factory{}
	return &out
}

// Create creates a new Signer instance
func (app *factory) Create() Signer {
	x := curve.Scalar().Pick(curve.RandomStream())
	return createSigner(x)
}
