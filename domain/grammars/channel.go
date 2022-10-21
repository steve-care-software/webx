package grammars

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type channel struct {
	hash      hash.Hash
	name      string
	token     Token
	condition ChannelCondition
}

func createChannel(
	hash hash.Hash,
	name string,
	token Token,
) Channel {
	return createChannelInternally(hash, name, token, nil)
}

func createChannelWithCondition(
	hash hash.Hash,
	name string,
	token Token,
	condition ChannelCondition,
) Channel {
	return createChannelInternally(hash, name, token, condition)
}

func createChannelInternally(
	hash hash.Hash,
	name string,
	token Token,
	condition ChannelCondition,
) Channel {
	out := channel{
		hash:      hash,
		name:      name,
		token:     token,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *channel) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *channel) Name() string {
	return obj.name
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
