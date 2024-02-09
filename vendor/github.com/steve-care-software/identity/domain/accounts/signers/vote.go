package signers

import (
	kyber "go.dedis.ch/kyber/v3"
)

type vote struct {
	ring []PublicKey
	s    []kyber.Scalar
	e    kyber.Scalar
}

func createVote(
	ring []PublicKey,
	s []kyber.Scalar,
	e kyber.Scalar,
) Vote {
	out := vote{
		ring: ring,
		s:    s,
		e:    e,
	}

	return &out
}

// Ring returns ring pubKeys
func (app *vote) Ring() []PublicKey {
	return app.ring
}

// Verify verifies if the message has been signed by at least 1 shared signature
func (app *vote) Verify(msg []byte) bool {
	// random base:
	g := curve.Point().Base()

	// first e:
	e := app.e

	//e = H(m || s[i] * G + e * P[i]);
	amount := len(app.ring)
	for i := 0; i < amount; i++ {
		sg := curve.Point().Mul(app.s[i], g)
		ep := curve.Point().Mul(e, app.ring[i].Point())
		added, err := curve.Point().Add(sg, ep).MarshalBinary()
		if err != nil {
			return false
		}

		combine := []byte{}
		combine = append(combine, msg...)
		combine = append(combine, added...)
	}

	return app.e.Equal(e)
}

// Bytes returns the string representation of the vote
func (app *vote) Bytes() ([]byte, error) {
	ringBytes := []byte{}
	for _, onePubKey := range app.ring {
		onePubKeyBytes, err := onePubKey.Bytes()
		if err != nil {
			return nil, err
		}

		ringBytes = append(ringBytes, onePubKeyBytes...)
		ringBytes = append(ringBytes, []byte(elementDelimiter)...)
	}

	scalarBytes := []byte{}
	for _, oneScalar := range app.s {
		oneScalarBytes, err := oneScalar.MarshalBinary()
		if err != nil {
			return nil, err
		}

		scalarBytes = append(scalarBytes, oneScalarBytes...)
		scalarBytes = append(scalarBytes, []byte(elementDelimiter)...)
	}

	eBytes, err := app.e.MarshalBinary()
	if err != nil {
		return nil, err
	}

	output := []byte{}
	output = append(output, ringBytes...)
	output = append(output, []byte(delimiter)...)
	output = append(output, scalarBytes...)
	output = append(output, []byte(delimiter)...)
	output = append(output, eBytes...)
	return output, nil
}
