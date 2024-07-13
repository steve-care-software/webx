package signatures

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter  hash.Adapter
	isGenPrivKey bool
	fetchPubKey  string
	sign         signs.Sign
	vote         votes.Vote
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		isGenPrivKey: false,
		fetchPubKey:  "",
		sign:         nil,
		vote:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithFetchPublicKey adds a fetchPublicKey to the builder
func (app *builder) WithFetchPublicKey(fetchPubKey string) Builder {
	app.fetchPubKey = fetchPubKey
	return app
}

// WithSign adds a sign to the builder
func (app *builder) WithSign(sign signs.Sign) Builder {
	app.sign = sign
	return app
}

// WithVote adds a vote to the builder
func (app *builder) WithVote(vote votes.Vote) Builder {
	app.vote = vote
	return app
}

// IsGeneratePrivateKey flags the builder as generatePrivateKey
func (app *builder) IsGeneratePrivateKey() Builder {
	app.isGenPrivKey = true
	return app
}

// Now builds a new Signature instance
func (app *builder) Now() (Signature, error) {
	data := [][]byte{}
	if app.fetchPubKey != "" {
		data = append(data, []byte("fetchPubKey"))
		data = append(data, []byte(app.fetchPubKey))
	}

	if app.sign != nil {
		data = append(data, []byte("sign"))
		data = append(data, app.sign.Hash().Bytes())
	}

	if app.vote != nil {
		data = append(data, []byte("vote"))
		data = append(data, app.vote.Hash().Bytes())
	}

	if app.isGenPrivKey {
		data = append(data, []byte("generatePK"))
	}

	length := len(data)
	if length != 1 && length != 2 {
		return nil, errors.New("the Signature is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.fetchPubKey != "" {
		return createSignatureWithFetchPublicKey(*pHash, app.fetchPubKey), nil
	}

	if app.sign != nil {
		return createSignatureWithSign(*pHash, app.sign), nil
	}

	if app.vote != nil {
		return createSignatureWithVote(*pHash, app.vote), nil
	}

	return createSignatureWithGeneratePrivateKey(*pHash), nil
}
