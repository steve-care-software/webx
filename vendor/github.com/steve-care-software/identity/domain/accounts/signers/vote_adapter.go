package signers

import (
	"bytes"
	"errors"
	"fmt"

	kyber "go.dedis.ch/kyber/v3"
	"github.com/steve-care-software/identity/domain/hash"
)

type voteAdapter struct {
	hashAdapter hash.Adapter
}

func createVoteAdapter(
	hashAdapter hash.Adapter,
) VoteAdapter {
	out := voteAdapter{
		hashAdapter: hashAdapter,
	}

	return &out
}

// ToVote converts bytes to vote
func (app *voteAdapter) ToVote(input []byte) (Vote, error) {
	splitted := bytes.Split(input, []byte(delimiter))
	if len(splitted) != 3 {
		str := fmt.Sprintf("the ring signature string was expected to have %d sections, %d found", 3, len(splitted))
		return nil, errors.New(str)
	}

	ring := []PublicKey{}
	ringPointBytes := bytes.Split(splitted[0], []byte(elementDelimiter))
	for _, oneRingPointBytes := range ringPointBytes {
		if len(oneRingPointBytes) <= 0 {
			continue
		}

		point, err := fromBytesToPoint(oneRingPointBytes)
		if err != nil {
			return nil, err
		}

		pubKey := createPublicKey(point)
		ring = append(ring, pubKey)
	}

	s := []kyber.Scalar{}
	scalarBytes := bytes.Split(splitted[1], []byte(elementDelimiter))
	for _, oneScalarBytes := range scalarBytes {
		if len(oneScalarBytes) <= 0 {
			continue
		}

		scalar, err := fromBytesToScalar(oneScalarBytes)
		if err != nil {
			return nil, err
		}

		s = append(s, scalar)
	}

	e, err := fromBytesToScalar(splitted[2])
	if err != nil {
		return nil, err
	}

	return createVote(ring, s, e), nil
}

// ToVerification executes a verification on the vote
func (app *voteAdapter) ToVerification(
	vote Vote,
	msg []byte,
	pubKeyHashes []hash.Hash,
) (bool, error) {
	if !vote.Verify(msg) {
		return false, errors.New("the signature could not be validated against the message")
	}

	ringPubKeys := vote.Ring()
	if len(pubKeyHashes) != len(ringPubKeys) {
		str := fmt.Sprintf("the length of the given hashes (%d) do not match the length of the signature's []PublicKey (%d)", len(pubKeyHashes), len(ringPubKeys))
		return false, errors.New(str)
	}

	for index, oneRingPubKey := range ringPubKeys {
		bytes, err := oneRingPubKey.Bytes()
		if err != nil {
			return false, err
		}

		ringPubKeyHash, err := app.hashAdapter.FromBytes(bytes)
		if err != nil {
			str := fmt.Sprintf("there was an error while hashing a ring PublicKey: %s", err.Error())
			return false, errors.New(str)
		}

		if !ringPubKeyHash.Compare(pubKeyHashes[index]) {
			str := fmt.Sprintf("the ring PublicKey hash (hash: %s, index: %d) do not match the given PublicKey hash (%s)", ringPubKeyHash.String(), index, pubKeyHashes[index].String())
			return false, errors.New(str)
		}
	}

	return true, nil
}
