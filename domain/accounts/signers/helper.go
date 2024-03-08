package signers

import (
	"bytes"

	kyber "go.dedis.ch/kyber/v3"
)

func createHash(msg []byte) kyber.Scalar {
	sha256 := curve.Hash()
	sha256.Reset()
	sha256.Write([]byte(msg))

	return curve.Scalar().SetBytes(sha256.Sum(nil))
}

func genK(x kyber.Scalar, msg []byte) (kyber.Scalar, error) {
	xBytes, err := x.MarshalBinary()
	if err != nil {
		return nil, err
	}

	combine := []byte{}
	combine = append(combine, msg...)
	combine = append(combine, xBytes...)
	return createHash(combine), nil
}

func fromBytesToScalar(input []byte) (kyber.Scalar, error) {
	x := curve.Scalar()
	reader := bytes.NewReader(input)
	_, err := x.UnmarshalFrom(reader)
	if err != nil {
		return nil, err
	}

	return x, nil
}

func fromBytesToPoint(input []byte) (kyber.Point, error) {
	p := curve.Point()
	err := p.UnmarshalBinary(input)
	if err != nil {
		return nil, err
	}

	return p, nil
}
