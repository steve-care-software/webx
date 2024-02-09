package signers

type publicKeyAdapter struct {
}

func createPublicKeyAdapter() PublicKeyAdapter {
	out := publicKeyAdapter{}
	return &out
}

// ToPublicKey converts []byte to a publicKey
func (app *publicKeyAdapter) ToPublicKey(pubKey []byte) (PublicKey, error) {
	point, err := fromBytesToPoint(pubKey)
	if err != nil {
		return nil, err
	}

	return createPublicKey(point), err
}
