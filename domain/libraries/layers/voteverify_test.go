package layers

import (
	"reflect"
	"testing"
)

func TestVoteVerify_Success(t *testing.T) {
	vote := "myVote"
	message := "myMessage"
	hashedRing := "myHashedRingVariable"
	voteVerify := NewVoteVerifyForTests(vote, message, hashedRing)
	retVote := voteVerify.Vote()
	if vote != retVote {
		t.Errorf("the vote was expected to be '%s', '%s returned'", vote, retVote)
		return
	}

	retMessage := voteVerify.Message()
	if !reflect.DeepEqual(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}

	retHashedRing := voteVerify.HashedRing()
	if hashedRing != retHashedRing {
		t.Errorf("the hashed ring was expected to be '%s', '%s returned'", hashedRing, retHashedRing)
		return
	}
}

func TestVoteVerify_withoutVote_returnsError(t *testing.T) {
	message := "myMessage"
	hashedRing := "myHashedRingVariable"
	_, err := NewVoteVerifyBuilder().Create().WithMessage(message).WithHashedRing(hashedRing).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestVoteVerify_withoutMessage_returnsError(t *testing.T) {
	vote := "myVote"
	hashedRing := "myHashedRingVariable"
	_, err := NewVoteVerifyBuilder().Create().WithVote(vote).WithHashedRing(hashedRing).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestVoteVerify_withoutHashedRing_returnsError(t *testing.T) {
	vote := "myVote"
	message := "myMessage"
	_, err := NewVoteVerifyBuilder().Create().WithMessage(message).WithVote(vote).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
