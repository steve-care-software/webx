package updates

import (
	"errors"
	"math/rand"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	signer      signers.Signer
	ring        []signers.PublicKey
	data        []byte
	addition    []hash.Hash
	removal     []hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		signer:      nil,
		ring:        nil,
		data:        nil,
		addition:    nil,
		removal:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithSigner adds a signer to the builder
func (app *builder) WithSigner(signer signers.Signer) Builder {
	app.signer = signer
	return app
}

// WithRing adds a ring to the builder
func (app *builder) WithRing(ring []signers.PublicKey) Builder {
	app.ring = ring
	return app
}

// WithData adds data to the builder
func (app *builder) WithData(data []byte) Builder {
	app.data = data
	return app
}

// WithWhiteListAddition adds a whitelist addition to the builder
func (app *builder) WithWhiteListAddition(addition []hash.Hash) Builder {
	app.addition = addition
	return app
}

// WithWhiteListRemoval adds a whitelist removal to the builder
func (app *builder) WithWhiteListRemoval(removal []hash.Hash) Builder {
	app.removal = removal
	return app
}

func (app *builder) content() (Content, error) {
	if app.name == "" {

	}

	if app.ring != nil && len(app.ring) <= 0 {
		app.ring = nil
	}

	if app.ring == nil {

	}

	if app.data != nil && len(app.data) <= 0 {
		app.data = nil
	}

	if app.addition != nil && len(app.addition) <= 0 {
		app.addition = nil
	}

	if app.removal != nil && len(app.removal) <= 0 {
		app.removal = nil
	}

	data := [][]byte{
		[]byte(app.name),
	}

	for _, onePubKey := range app.ring {
		data = append(data, []byte(onePubKey.String()))
	}

	if app.addition != nil {
		for _, oneHash := range app.addition {
			data = append(data, oneHash.Bytes())
		}
	}

	if app.removal != nil {
		for _, oneHash := range app.removal {
			data = append(data, oneHash.Bytes())
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.data != nil && app.addition != nil && app.removal != nil {
		return createContentWithDataAndAdditionAndRemoval(*pHash, app.name, app.data, app.addition, app.removal), nil
	}

	if app.data != nil && app.removal != nil {
		return createContentWithDataAndRemoval(*pHash, app.name, app.data, app.removal), nil
	}

	if app.data != nil && app.addition != nil {
		return createContentWithDataAndAddition(*pHash, app.name, app.data, app.removal), nil
	}

	if app.addition != nil && app.removal != nil {
		return createContentWithAdditionAndRemoval(*pHash, app.name, app.addition, app.removal), nil
	}

	if app.addition != nil {
		return createContentWithAddition(*pHash, app.name, app.addition), nil
	}

	if app.removal != nil {
		return createContentWithRemoval(*pHash, app.name, app.removal), nil
	}

	if app.data != nil {
		return createContentWithData(*pHash, app.name, app.data), nil
	}

	return nil, errors.New("the Update is invalid")
}

// Now builds a new Update instance
func (app *builder) Now() (Update, error) {
	content, err := app.content()
	if err != nil {
		return nil, err
	}

	// add our public key and shuffle it:
	ring := append(app.ring, app.signer.PublicKey())
	rand.Shuffle(len(ring), func(i, j int) { ring[i], ring[j] = ring[j], ring[i] })

	// vote on the message:
	message := content.Hash().String()
	vote, err := app.signer.Vote(message, app.ring)
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		content.Hash().Bytes(),
		[]byte(vote.String()),
	})

	if err != nil {
		return nil, err
	}

	return createUpdate(
		*pHash,
		content,
		vote,
	), nil
}
