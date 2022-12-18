package grammars

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type channelCondition struct {
	hash hash.Hash
	prev Token
	next Token
}

func createChannelConditionWithPrevious(
	hash hash.Hash,
	prev Token,
) ChannelCondition {
	return createChannelConditionInternally(hash, prev, nil)
}

func createChannelConditionWithNext(
	hash hash.Hash,
	next Token,
) ChannelCondition {
	return createChannelConditionInternally(hash, nil, next)
}

func createChannelConditionWithPreviousAndNext(
	hash hash.Hash,
	prev Token,
	next Token,
) ChannelCondition {
	return createChannelConditionInternally(hash, prev, next)
}

func createChannelConditionInternally(
	hash hash.Hash,
	prev Token,
	next Token,
) ChannelCondition {
	out := channelCondition{
		hash: hash,
		prev: prev,
		next: next,
	}

	return &out
}

// Hash returns the hash
func (obj *channelCondition) Hash() hash.Hash {
	return obj.hash
}

// Points returns the amount of points a channelCondition contains
func (obj *channelCondition) Points() uint {
	amount := uint(0)
	if obj.HasPrevious() {
		amount += obj.Previous().Block().Points()
	}

	if obj.HasNext() {
		amount += obj.Next().Block().Points()
	}

	return amount
}

// HasPrevious returns true if there is a previous token, false otherwise
func (obj *channelCondition) HasPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous token, if any
func (obj *channelCondition) Previous() Token {
	return obj.prev
}

// HasNext returns true if there is a next token, false otherwise
func (obj *channelCondition) HasNext() bool {
	return obj.next != nil
}

// Next returns the next token, if any
func (obj *channelCondition) Next() Token {
	return obj.next
}
