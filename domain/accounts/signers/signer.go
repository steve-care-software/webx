package signers

import (
	"errors"
	"strconv"

	kyber "go.dedis.ch/kyber/v3"
)

type signer struct {
	x kyber.Scalar
}

func createSigner(x kyber.Scalar) Signer {
	out := signer{
		x: x,
	}

	return &out
}

// PublicKey returns the public key
func (app *signer) PublicKey() PublicKey {
	g := curve.Point().Base()
	p := curve.Point().Mul(app.x, g)
	return createPublicKey(p)
}

// Vote signs a vote on the given message, using the provided ring
func (app *signer) Vote(msg []byte, ring []PublicKey) (Vote, error) {
	retrieveSignerIndexFn := func(ring []PublicKey, pk Signer) int {
		pubKey := pk.PublicKey()
		for index, oneRingPubKey := range ring {
			if oneRingPubKey.Equals(pubKey) {
				return index
			}
		}

		return -1
	}

	// retrieve our signerIndex:
	signerIndex := retrieveSignerIndexFn(ring, app)
	if signerIndex == -1 {
		return nil, errors.New("the signer PublicKey is not in the ring")
	}

	// generate k:
	k, err := genK(app.x, msg)
	if err != nil {
		return nil, err
	}

	// random base:
	g := curve.Point().Base()

	// length:
	r := len(ring)

	// initialize:
	es := make([]kyber.Scalar, r)
	ss := make([]kyber.Scalar, r)
	beginIndex := (signerIndex + 1) % r

	// ei = H(m || k * G)
	mulBytes, err := curve.Point().Mul(k, g).MarshalBinary()
	if err != nil {
		return nil, err
	}

	combine := []byte{}
	combine = append(combine, msg...)
	combine = append(combine, mulBytes...)
	es[beginIndex] = createHash(combine)

	// loop:
	for i := beginIndex; i != signerIndex; i = (i + 1) % r {
		// si = random value
		siCombine := []byte{}
		siCombine = append(siCombine, msg...)
		siCombine = append(siCombine, []byte(strconv.Itoa(i))...)
		ss[i], err = genK(app.x, siCombine)
		if err != nil {
			return nil, err
		}

		//eiPlus1ModR = H(m || si * G + ei * Pi)
		sig := curve.Point().Mul(ss[i], g)
		eipi := curve.Point().Mul(es[i], ring[i].Point())

		addBytes, err := curve.Point().Add(sig, eipi).MarshalBinary()
		if err != nil {
			return nil, err
		}

		combine := []byte{}
		combine = append(combine, msg...)
		combine = append(combine, addBytes...)
		es[(i+1)%r] = createHash(combine)

	}

	// close the ring:
	esx := curve.Scalar().Mul(es[signerIndex], app.x)
	ss[signerIndex] = curve.Scalar().Sub(k, esx)
	out := createVote(ring, ss, es[0])
	return out, nil
}

// Sign signs a message
func (app *signer) Sign(msg []byte) (Signature, error) {
	// generate k:
	k, err := genK(app.x, msg)
	if err != nil {
		return nil, err
	}

	// random base:
	g := curve.Point().Base()

	// r = k * G (a.k.a the same operation as r = g^k)
	r := curve.Point().Mul(k, g)

	// hash(m || r)
	rBytes, err := r.MarshalBinary()
	if err != nil {
		return nil, err
	}

	combine := []byte{}
	combine = append(combine, msg...)
	combine = append(combine, rBytes...)
	e := createHash(combine)

	// s = k - e * x
	s := curve.Scalar().Sub(k, curve.Scalar().Mul(e, app.x))

	// create signature:
	pubKey := createPublicKey(r)
	return createSignature(pubKey, s)
}

// Bytes returns the bytes representation of the Signer
func (app *signer) Bytes() ([]byte, error) {
	return app.x.MarshalBinary()
}
