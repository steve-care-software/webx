package signers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/webx/engine/domain/blockchains/hash"
	kyber "go.dedis.ch/kyber/v3"
)

type voteAdapter struct {
	hashAdapter hash.Adapter
}

func createVoteAdapter(hashAdapter hash.Adapter) VoteAdapter {
	out := voteAdapter{
		hashAdapter: hashAdapter,
	}

	return &out
}

// ToSignature converts a string to a Vote
func (app *voteAdapter) ToSignature(sig string) (Vote, error) {
	decoded, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return nil, err
	}

	splitted := strings.Split(string(decoded), delimiter)
	if len(splitted) != 3 {
		str := fmt.Sprintf("the ring signature string was expected to have %d sections, %d found", 3, len(splitted))
		return nil, errors.New(str)
	}

	ring := []PublicKey{}
	ringPointsStr := strings.Split(splitted[0], elementDelimiter)
	for _, oneRingPointStr := range ringPointsStr {
		if oneRingPointStr == "" {
			continue
		}

		point, err := fromStringToPoint(oneRingPointStr)
		if err != nil {
			return nil, err
		}

		pubKey := createPublicKey(point)
		ring = append(ring, pubKey)
	}

	s := []kyber.Scalar{}
	sScalarsStr := strings.Split(splitted[1], elementDelimiter)
	for _, oneScalarStr := range sScalarsStr {
		if oneScalarStr == "" {
			continue
		}

		scalar, err := fromStringToScalar(oneScalarStr)
		if err != nil {
			return nil, err
		}

		s = append(s, scalar)
	}

	e, err := fromStringToScalar(splitted[2])
	if err != nil {
		return nil, err
	}

	return createVote(ring, s, e), nil
}

// ToVerification verifies that the signature is valid and that it contains exactly the same publicKey hashes
func (app *voteAdapter) ToVerification(sig Vote, msg string, pubKeyHashes []hash.Hash) (bool, error) {
	if !sig.Verify(msg) {
		return false, errors.New("the signature could not be validated against the message")
	}

	ring := sig.Ring()
	filteredHashedKeys := []hash.Hash{}
	for _, onePubKeyHash := range pubKeyHashes {

		isValid := false
		for _, oneRing := range ring {
			pHash, err := app.hashAdapter.FromString(oneRing.String())
			if err != nil {
				return false, err
			}

			if pHash.Compare(onePubKeyHash) {
				isValid = true
				break
			}
		}

		if !isValid {
			continue
		}

		filteredHashedKeys = append(filteredHashedKeys, onePubKeyHash)
	}

	if len(filteredHashedKeys) != len(ring) {
		str := fmt.Sprintf("the length of the given hashes (%d) do not match the length of the signature's []PublicKey (%d)", len(filteredHashedKeys), len(ring))
		return false, errors.New(str)
	}

	for index, oneRingPubKey := range filteredHashedKeys {
		ringPubKeyHash, err := app.hashAdapter.FromString(oneRingPubKey.String())
		if err != nil {
			str := fmt.Sprintf("there was an error while hashing a ring PublicKey: %s", err.Error())
			return false, errors.New(str)
		}

		if !ringPubKeyHash.Compare(filteredHashedKeys[index]) {
			str := fmt.Sprintf("the ring PublicKey hash (hash: %s, index: %d) do not match the given PublicKey hash (%s)", ringPubKeyHash.String(), index, filteredHashedKeys[index].String())
			return false, errors.New(str)
		}
	}

	return true, nil
}
