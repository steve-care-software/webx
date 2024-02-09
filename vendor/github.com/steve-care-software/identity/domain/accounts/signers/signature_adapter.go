package signers

import (
	"bytes"
	"errors"
	"fmt"
)

type signatureAdapter struct {
}

func createSignatureAdapter() SignatureAdapter {
	out := signatureAdapter{}
	return &out
}

// ToSignature converts bytes to a Signature instance
func (app *signatureAdapter) ToSignature(sig []byte) (Signature, error) {
	splitted := bytes.Split(sig, []byte(delimiter))
	if len(splitted) != 2 {
		str := fmt.Sprintf("the signature string was expected to have %d sections, %d found", 2, len(splitted))
		return nil, errors.New(str)
	}

	point, err := fromBytesToPoint(splitted[0])
	if err != nil {
		return nil, err
	}

	scalar, err := fromBytesToScalar(splitted[1])
	if err != nil {
		return nil, err
	}

	pubKey := createPublicKey(point)
	return createSignature(pubKey, scalar)
}
