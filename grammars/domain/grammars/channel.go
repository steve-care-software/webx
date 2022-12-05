package grammars

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type channel struct {
	hash      hash.Hash
	token     Token
	condition ChannelCondition
}

func createChannel(
	hash hash.Hash,
	token Token,
) Channel {
	return createChannelInternally(hash, token, nil)
}

func createChannelWithCondition(
	hash hash.Hash,
	token Token,
	condition ChannelCondition,
) Channel {
	return createChannelInternally(hash, token, condition)
}

func createChannelInternally(
	hash hash.Hash,
	token Token,
	condition ChannelCondition,
) Channel {
	out := channel{
		hash:      hash,
		token:     token,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *channel) Hash() hash.Hash {
	return obj.hash
}

// Token returns the token
func (obj *channel) Token() Token {
	return obj.token
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *channel) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *channel) Condition() ChannelCondition {
	return obj.condition
}
