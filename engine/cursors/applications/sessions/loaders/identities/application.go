package identities

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	loaders_identities "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/profiles"
)

type application struct {
	builder          loaders_identities.Builder
	switchersBuilder switchers.Builder
	switcherBuilder  switchers.SwitcherBuilder
	singlesBuilder   singles.Builder
	singleBuilder    singles.SingleBuilder
	profileBuilder   profiles.Builder
	keyBuilder       keys.Builder
	encryptorBuilder encryptors.Builder
	signerFactory    signers.Factory
	bitsize          int
}

func createApplication(
	builder loaders_identities.Builder,
	switchersBuilder switchers.Builder,
	switcherBuilder switchers.SwitcherBuilder,
	singlesBuilder singles.Builder,
	singleBuilder singles.SingleBuilder,
	profileBuilder profiles.Builder,
	keyBuilder keys.Builder,
	encryptorBuilder encryptors.Builder,
	signerFactory signers.Factory,
	bitsize int,
) Application {
	out := application{
		builder:          builder,
		switchersBuilder: switchersBuilder,
		switcherBuilder:  switcherBuilder,
		singlesBuilder:   singlesBuilder,
		singleBuilder:    singleBuilder,
		profileBuilder:   profileBuilder,
		keyBuilder:       keyBuilder,
		encryptorBuilder: encryptorBuilder,
		signerFactory:    signerFactory,
		bitsize:          bitsize,
	}

	return &out
}

// Create creates an identity
func (app *application) Create(
	input loaders_identities.Identity,
	name string,
	description string,
	password []byte,
) (loaders_identities.Identity, error) {
	signer := app.signerFactory.Create()

	// Generate RSA key.
	pKey, err := rsa.GenerateKey(rand.Reader, app.bitsize)
	if err != nil {
		return nil, err
	}

	encryptor, err := app.encryptorBuilder.Create().WithBitRate(app.bitsize).WithPK(*pKey).Now()
	if err != nil {
		return nil, err
	}

	createdOn := time.Now().UTC()
	key, err := app.keyBuilder.Create().WithSigner(signer).WithEncryptor(encryptor).CreatedOn(createdOn).Now()
	if err != nil {
		return nil, err
	}

	profile, err := app.profileBuilder.Create().WithName(name).WithDescription(description).Now()
	if err != nil {
		return nil, err
	}

	single, err := app.singleBuilder.Create().WithProfile(profile).WithKey(key).Now()
	if err != nil {
		return nil, err
	}

	switcher, err := app.switcherBuilder.Create().WithUpdated(single).Now()
	if err != nil {
		return nil, err
	}

	currentSwitchers := []switchers.Switcher{}
	if input.HasAuthenticated() {
		authenticated := input.Authenticated()
		currentSwitchers = authenticated.List()
	}

	currentSwitchers = append(currentSwitchers, switcher)
	switchers, err := app.switchersBuilder.Create().WithList(currentSwitchers).Now()
	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.builder.Create().WithAll(all).WithAuthenticated(switchers)
	if input.HasCurrent() {
		current := input.Current()
		builder.WithCurrent(current)
	}

	return builder.Now()
}

// Authenticate authenticates
func (app *application) Authenticate(input loaders_identities.Identity, name string, password []byte) (loaders_identities.Identity, error) {
	return nil, nil
}

// SetPassword changes the password of the current authenticated user
func (app *application) SetPassword(input loaders_identities.Identity, newPassword []byte) (loaders_identities.Identity, error) {
	return nil, nil
}

// SetUser sets the authenticated user
func (app *application) SetUser(input loaders_identities.Identity, name string) (loaders_identities.Identity, error) {
	return nil, nil
}

// Follow follows a namespace using the current authenticated user
func (app *application) Follow(input loaders_identities.Identity, namespace string) (loaders_identities.Identity, error) {
	return nil, nil
}

// Encrypt encrypts data using the current authenticated user
func (app *application) Encrypt(input loaders_identities.Identity, message []byte) ([]byte, error) {
	return nil, nil
}

// Decrypt decrypts data using the current authenticated user
func (app *application) Decrypt(input loaders_identities.Identity, cipher []byte) ([]byte, error) {
	return nil, nil
}

// Sign signs data using the current authenticated user
func (app *application) Sign(input loaders_identities.Identity, message []byte) (signers.Signature, error) {
	return nil, nil
}

// ValidateSignature validates a signature
func (app *application) ValidateSignature(input loaders_identities.Identity, message []byte, sig signers.Signature) (bool, error) {
	return false, nil
}

// Vote votes on a message using the current authenticated user
func (app *application) Vote(input loaders_identities.Identity, message []byte, ring []signers.PublicKey) (signers.Vote, error) {
	return nil, nil
}

// ValidateVote validates a vote
func (app *application) ValidateVote(input loaders_identities.Identity, message []byte, vote signers.Vote, ring []hash.Hash) (bool, error) {
	return false, nil
}
