package layers

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/identity/domain/hash"
)

type signerBuilder struct {
	hashAdapter      hash.Adapter
	sign             string
	vote             Vote
	genSignerPubKeys uint
	hashPublicKeys   string
	voteVerify       VoteVerify
	signatureVerify  SignatureVerify
	bytes            string
	isPubKey         bool
}

func createSignerBuilder(
	hashAdapter hash.Adapter,
) SignerBuilder {
	out := signerBuilder{
		hashAdapter:      hashAdapter,
		sign:             "",
		vote:             nil,
		genSignerPubKeys: 0,
		hashPublicKeys:   "",
		voteVerify:       nil,
		signatureVerify:  nil,
		bytes:            "",
		isPubKey:         false,
	}

	return &out
}

// Create initializes the builder
func (app *signerBuilder) Create() SignerBuilder {
	return createSignerBuilder(
		app.hashAdapter,
	)
}

// WithSign adds a sign to the builder
func (app *signerBuilder) WithSign(sign string) SignerBuilder {
	app.sign = sign
	return app
}

// WithVote adds a vote to the builder
func (app *signerBuilder) WithVote(vote Vote) SignerBuilder {
	app.vote = vote
	return app
}

// WithGenerateSignerPublicKey adds a generateSignerPublicKey to the builder
func (app *signerBuilder) WithGenerateSignerPublicKey(genPubKey uint) SignerBuilder {
	app.genSignerPubKeys = genPubKey
	return app
}

// WithHashPublicKeys adds an hash public keys to the builder
func (app *signerBuilder) WithHashPublicKeys(hashPubKeys string) SignerBuilder {
	app.hashPublicKeys = hashPubKeys
	return app
}

// WithVoteVerify adds an hash public keys to the builder
func (app *signerBuilder) WithVoteVerify(voteVerify VoteVerify) SignerBuilder {
	app.voteVerify = voteVerify
	return app
}

// WithSignatureVerify adds a signature verify to the builder
func (app *signerBuilder) WithSignatureVerify(sigVerify SignatureVerify) SignerBuilder {
	app.signatureVerify = sigVerify
	return app
}

// WithBytes add bytes to the builder
func (app *signerBuilder) WithBytes(bytes string) SignerBuilder {
	app.bytes = bytes
	return app
}

// IsPublicKey flags the builder as isPublicKey
func (app *signerBuilder) IsPublicKey() SignerBuilder {
	app.isPubKey = true
	return app
}

// Now builds a new Signer instance
func (app *signerBuilder) Now() (Signer, error) {
	data := [][]byte{}
	if app.sign != "" {
		data = append(data, []byte("signer"))
		data = append(data, []byte(app.sign))
	}

	if app.vote != nil {
		data = append(data, []byte("vote"))
		data = append(data, app.vote.Hash().Bytes())
	}

	if app.genSignerPubKeys > 0 {
		data = append(data, []byte("genSignerPubKeys"))
		data = append(data, []byte(strconv.Itoa(int(app.genSignerPubKeys))))
	}

	if app.hashPublicKeys != "" {
		data = append(data, []byte("hashPublicKeys"))
		data = append(data, []byte(app.hashPublicKeys))
	}

	if app.voteVerify != nil {
		data = append(data, []byte("voteVerify"))
		data = append(data, app.voteVerify.Hash().Bytes())
	}

	if app.signatureVerify != nil {
		data = append(data, []byte("signatureVerify"))
		data = append(data, app.signatureVerify.Hash().Bytes())
	}

	if app.bytes != "" {
		data = append(data, []byte("bytes"))
		data = append(data, []byte(app.bytes))
	}

	if app.isPubKey {
		data = append(data, []byte("isPublicKey"))
	}

	if len(data) <= 0 {
		return nil, errors.New("the Signer is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.sign != "" {
		return createSignerWithSign(*pHash, app.sign), nil
	}

	if app.vote != nil {
		return createSignerWithVote(*pHash, app.vote), nil
	}

	if app.genSignerPubKeys > 0 {
		return createSignerWithGenerateSignerKeys(*pHash, app.genSignerPubKeys), nil
	}

	if app.hashPublicKeys != "" {
		return createSignerWithHashPublicKeys(*pHash, app.hashPublicKeys), nil
	}

	if app.voteVerify != nil {
		return createSignerWithVoteVerify(*pHash, app.voteVerify), nil
	}

	if app.signatureVerify != nil {
		return createSignerWithSignatureVerify(*pHash, app.signatureVerify), nil
	}

	if app.bytes != "" {
		return createSignerWithBytes(*pHash, app.bytes), nil
	}

	return createSignerWithIsPublicKey(*pHash), nil
}
