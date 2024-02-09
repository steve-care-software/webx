package layers

import (
	"reflect"
	"testing"
)

func TestVote_Success(t *testing.T) {
	ring := "myRingVariable"
	message := "myMessage"
	vote := NewVoteForTests(ring, message)
	retRing := vote.Ring()
	if ring != retRing {
		t.Errorf("the ring was expected to be '%s', '%s returned'", ring, retRing)
		return
	}

	retMessage := vote.Message()
	if !reflect.DeepEqual(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestVote_withoutRing_returnsError(t *testing.T) {
	message := "myMessage"
	_, err := NewVoteBuilder().Create().WithMessage(message).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestVote_withoutMessage_returnsError(t *testing.T) {
	ring := "myRingVariable"
	_, err := NewVoteBuilder().Create().WithRing(ring).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
