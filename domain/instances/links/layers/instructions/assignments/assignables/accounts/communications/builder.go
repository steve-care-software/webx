package communications

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/votes"
)

type builder struct {
	hashAdapter  hash.Adapter
	sign         signs.Sign
	vote         votes.Vote
	generateRing string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		sign:         nil,
		vote:         nil,
		generateRing: "",
	}

	return &out
}

// Create intiializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
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

// WithGenerateRing adds a generateRing to the builder
func (app *builder) WithGenerateRing(generateRing string) Builder {
	app.generateRing = generateRing
	return app
}

// WithGenerateRing adds a generateRing to the builder
func (app *builder) Now() (Communication, error) {
	data := [][]byte{}
	if app.sign != nil {
		data = append(data, []byte("sign"))
		data = append(data, app.sign.Hash().Bytes())
	}

	if app.vote != nil {
		data = append(data, []byte("vote"))
		data = append(data, app.vote.Hash().Bytes())
	}

	if app.generateRing != "" {
		data = append(data, []byte("generateRing"))
		data = append(data, []byte(app.generateRing))
	}

	if len(data) != 2 {
		return nil, errors.New("the Communication is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.sign != nil {
		return createCommunicationWithSign(*pHash, app.sign), nil
	}

	if app.vote != nil {
		return createCommunicationWithVote(*pHash, app.vote), nil
	}

	return createCommunicationWithGenerateRing(*pHash, app.generateRing), nil
}
