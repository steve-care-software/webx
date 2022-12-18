package channels

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type channel struct {
	hash  hash.Hash
	token hash.Hash
	prev  *hash.Hash
	next  *hash.Hash
}

func createChannel(
	hash hash.Hash,
	token hash.Hash,
) Channel {
	return createChannelInternally(hash, token, nil, nil)
}

func createChannelWithPrevious(
	hash hash.Hash,
	token hash.Hash,
	prev *hash.Hash,
) Channel {
	return createChannelInternally(hash, token, prev, nil)
}

func createChannelWithNext(
	hash hash.Hash,
	token hash.Hash,
	next *hash.Hash,
) Channel {
	return createChannelInternally(hash, token, nil, next)
}

func createChannelWithPreviousAndNext(
	hash hash.Hash,
	token hash.Hash,
	prev *hash.Hash,
	next *hash.Hash,
) Channel {
	return createChannelInternally(hash, token, prev, next)
}

func createChannelInternally(
	hash hash.Hash,
	token hash.Hash,
	prev *hash.Hash,
	next *hash.Hash,
) Channel {
	out := channel{
		hash:  hash,
		token: token,
		prev:  prev,
		next:  next,
	}

	return &out
}

// Hash returns the hash
func (obj *channel) Hash() hash.Hash {
	return obj.hash
}

// Token returns the token
func (obj *channel) Token() hash.Hash {
	return obj.token
}

// HasPrevious returns true if there is a previous, false otherwise
func (obj *channel) HasPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous, if any
func (obj *channel) Previous() *hash.Hash {
	return obj.prev
}

// HasNext returns true if there is a next, false otherwise
func (obj *channel) HasNext() bool {
	return obj.next != nil
}

// Next returns the next, if any
func (obj *channel) Next() *hash.Hash {
	return obj.next
}
