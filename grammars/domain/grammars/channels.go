package grammars

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

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

// Points returns the amount of points a channels contains
func (obj *channels) Points() uint {
	amount := uint(0)
	for _, oneChannel := range obj.list {
		amount += oneChannel.Points()
	}

	return amount
}

// List returns the channels
func (obj *channels) List() []Channel {
	return obj.list
}
