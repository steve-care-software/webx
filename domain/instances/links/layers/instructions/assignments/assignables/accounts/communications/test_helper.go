package communications

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/votes"
)

// NewCommunicationWithSignForTests creates a new communication with sign for tests
func NewCommunicationWithSignForTests(sign signs.Sign) Communication {
	ins, err := NewBuilder().Create().WithSign(sign).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommunicationWithVoteForTests creates a new communication with vote for tests
func NewCommunicationWithVoteForTests(vote votes.Vote) Communication {
	ins, err := NewBuilder().Create().WithVote(vote).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommunicationWithGenerateRingForTests creates a new communication with generate ring for tests
func NewCommunicationWithGenerateRingForTests(generateRing string) Communication {
	ins, err := NewBuilder().Create().WithGenerateRing(generateRing).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
