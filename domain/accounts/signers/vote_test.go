package signers

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
)

func TestVote_Success(t *testing.T) {
	voteAdapter := NewVoteAdapter()
	hashAdapter := hash.NewAdapter()

	// variables:
	msg := []byte("this is a message to sign")
	pk := NewFactory().Create()
	secondPK := NewFactory().Create()
	ringPubKeys := []PublicKey{
		pk.PublicKey(),
		secondPK.PublicKey(),
	}

	ringPubKeyHashes := []hash.Hash{}
	for _, onePubKey := range ringPubKeys {
		pubKeyBytes, err := onePubKey.Bytes()
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		pHash, err := hashAdapter.FromBytes(pubKeyBytes)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		ringPubKeyHashes = append(ringPubKeyHashes, *pHash)
	}

	firstVote, err := pk.Vote(msg, ringPubKeys)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	secondVote, err := secondPK.Vote(msg, ringPubKeys)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	if !firstVote.Verify(msg) {
		t.Errorf("the first ring was expected to be verified")
		return
	}

	firstVoteVerified, err := voteAdapter.ToVerification(firstVote, msg, ringPubKeyHashes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	if !firstVoteVerified {
		t.Errorf("the first ring was expected to be deep verified")
		return
	}

	if !secondVote.Verify(msg) {
		t.Errorf("the second ring was expected to be verified")
		return
	}

	secondVoteVerified, err := voteAdapter.ToVerification(secondVote, msg, ringPubKeyHashes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned; %s", err.Error())
		return
	}

	if !secondVoteVerified {
		t.Errorf("the second ring was expected to be deep verified")
		return
	}

	// encode to string, back and forth:
	firstVoteBytes, err := firstVote.Bytes()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	newVote, err := NewVoteAdapter().ToVote(firstVoteBytes)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
		return
	}

	newVoteBytes, err := newVote.Bytes()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(firstVoteBytes, newVoteBytes) {
		t.Errorf("the rings were expected to be the same.  Expected: %s, Actual: %s", firstVoteBytes, newVoteBytes)
		return
	}
}

func TestVote_PubKeyIsNotInTheRing_returnsError(t *testing.T) {
	// variables:
	msg := []byte("this is a message to sign")
	pk := NewFactory().Create()
	secondPK := NewFactory().Create()
	invalidPK := NewFactory().Create()
	ringPubKeys := []PublicKey{
		pk.PublicKey(),
		secondPK.PublicKey(),
	}

	_, err := invalidPK.Vote(msg, ringPubKeys)
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}
}
