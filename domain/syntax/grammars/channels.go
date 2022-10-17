package grammars

import "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"

type channels struct {
	hash hash.Hash
	list []Channel
}

func createChannels(
	hash hash.Hash,
	list []Channel,
) Channels {
	out := channels{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *channels) Hash() hash.Hash {
	return obj.hash
}

// List returns the channels
func (obj *channels) List() []Channel {
	return obj.list
}
