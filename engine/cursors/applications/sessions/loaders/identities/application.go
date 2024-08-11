package identities

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"time"

	"github.com/steve-care-software/webx/engine/cursors/applications/encryptions"
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/loaders/identities/namespaces"
	storage_pointer_applications "github.com/steve-care-software/webx/engine/cursors/applications/sessions/pointers"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	loaders_identities "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/profiles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/updates"
)

type application struct {
	encryptionApp             encryptions.Application
	storagePointerApplication storage_pointer_applications.Application
	hashAdapter               hash.Adapter
	voteAdapter               signers.VoteAdapter
	builder                   loaders_identities.Builder
	switchersBuilder          switchers.Builder
	switcherBuilder           switchers.SwitcherBuilder
	updateBuilder             updates.Builder
	singlesAdapter            singles.Adapter
	singlesBuilder            singles.Builder
	singleBuilder             singles.SingleBuilder
	profileBuilder            profiles.Builder
	keyBuilder                keys.Builder
	encryptorBuilder          encryptors.Builder
	signerFactory             signers.Factory
	bitsize                   int
}

func createApplication(
	encryptionApp encryptions.Application,
	storagePointerApplication storage_pointer_applications.Application,
	hashAdapter hash.Adapter,
	voteAdapter signers.VoteAdapter,
	builder loaders_identities.Builder,
	switchersBuilder switchers.Builder,
	switcherBuilder switchers.SwitcherBuilder,
	updateBuilder updates.Builder,
	singlesAdapter singles.Adapter,
	singlesBuilder singles.Builder,
	singleBuilder singles.SingleBuilder,
	profileBuilder profiles.Builder,
	keyBuilder keys.Builder,
	encryptorBuilder encryptors.Builder,
	signerFactory signers.Factory,
	bitsize int,
) Application {
	out := application{
		encryptionApp:             encryptionApp,
		storagePointerApplication: storagePointerApplication,
		hashAdapter:               hashAdapter,
		voteAdapter:               voteAdapter,
		builder:                   builder,
		switchersBuilder:          switchersBuilder,
		switcherBuilder:           switcherBuilder,
		updateBuilder:             updateBuilder,
		singlesAdapter:            singlesAdapter,
		singlesBuilder:            singlesBuilder,
		singleBuilder:             singleBuilder,
		profileBuilder:            profileBuilder,
		keyBuilder:                keyBuilder,
		encryptorBuilder:          encryptorBuilder,
		signerFactory:             signerFactory,
		bitsize:                   bitsize,
	}

	return &out
}

// List returns the list of identity names
func (app *application) List(
	input loaders_identities.Identity,
) []string {
	return input.All().Names()
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

	encryptor, err := app.encryptorBuilder.Create().
		WithBitRate(app.bitsize).
		WithPK(*pKey).
		Now()

	if err != nil {
		return nil, err
	}

	createdOn := time.Now().UTC()
	key, err := app.keyBuilder.Create().
		WithSigner(signer).
		WithEncryptor(encryptor).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return nil, err
	}

	profile, err := app.profileBuilder.Create().
		WithName(name).
		WithDescription(description).
		Now()

	if err != nil {
		return nil, err
	}

	single, err := app.singleBuilder.Create().
		WithProfile(profile).
		WithKey(key).
		Now()

	if err != nil {
		return nil, err
	}

	singleBytes, err := app.singlesAdapter.ToBytes(single)
	if err != nil {
		return nil, err
	}

	cipher, err := app.encryptionApp.Encrypt(singleBytes, password)
	if err != nil {
		return nil, err
	}

	update, err := app.updateBuilder.WithSingle(single).WithBytes(cipher).Now()
	if err != nil {
		return nil, err
	}

	return app.updateAuthenticated(input, update)
}

// Authenticate authenticates
func (app *application) Authenticate(input loaders_identities.Identity, name string, password []byte) (loaders_identities.Identity, error) {
	// fetch the identity from the all list:
	storedIdentity, err := input.All().Fetch(name)
	if err != nil {
		return nil, err
	}

	// fetch the data from the pointer in the database:
	pointer, err := app.storagePointerApplication.Retrieve(storedIdentity.Pointer())
	if err != nil {
		return nil, err
	}

	// decrypt the bytes:
	decrypted, err := app.encryptionApp.Decrypt(pointer.Bytes(), password)
	if err != nil {
		return nil, err
	}

	// convert the data to a single instance:
	single, err := app.singlesAdapter.ToInstance(decrypted)
	if err != nil {
		return nil, err
	}

	// create a switcher with the single as its original:
	switcher, err := app.switcherBuilder.Create().WithOriginal(single).Now()
	if err != nil {
		return nil, err
	}

	// add the switcher to the list of authenticated user:
	currentAuthenticated := []switchers.Switcher{}
	if input.HasAuthenticated() {
		currentAuthenticated = input.Authenticated().List()
	}

	currentAuthenticated = append(currentAuthenticated, switcher)
	authenticated, err := app.switchersBuilder.Create().WithList(currentAuthenticated).Now()
	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.builder.Create().
		WithAll(all).
		WithAuthenticated(authenticated)

	if input.HasCurrent() {
		builder.WithCurrent(input.Current())
	}

	return builder.Now()
}

// SetPassword changes the password of the current authenticated user
func (app *application) SetPassword(input loaders_identities.Identity, newPassword []byte) (loaders_identities.Identity, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noCurrentUserErr)
	}

	switcher := input.Current()
	single := switcher.Current()
	singleBytes, err := app.singlesAdapter.ToBytes(single)
	if err != nil {
		return nil, err
	}

	cipher, err := app.encryptionApp.Encrypt(singleBytes, newPassword)
	if err != nil {
		return nil, err
	}

	update, err := app.updateBuilder.Create().
		WithSingle(single).
		WithBytes(cipher).
		Now()

	if err != nil {
		return nil, err
	}

	return app.updateAuthenticated(input, update)
}

// SetUser sets the authenticated user
func (app *application) SetUser(input loaders_identities.Identity, name string) (loaders_identities.Identity, error) {
	if input.HasAuthenticated() {
		return nil, errors.New("there is no authenticated user")
	}

	authenticated := input.Authenticated()
	current, err := authenticated.Fetch(name)
	if err != nil {
		return nil, err
	}

	all := input.All()
	return app.builder.Create().
		WithAll(all).
		WithAuthenticated(authenticated).
		WithCurrent(current).
		Now()
}

// Follow follows a namespace using the current authenticated user
func (app *application) Follow(input loaders_identities.Identity, namespace string, password []byte) (loaders_identities.Identity, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noCurrentUserErr)
	}

	switcher := input.Current()
	if switcher.HasOriginal() {
		return nil, errors.New("the current switcher does not contain an original identity and therefore cannot be updated yet")
	}

	single := switcher.Current()
	profile := single.Profile()
	name := profile.Name()
	description := profile.Description()
	namespaces := []string{}
	if profile.HasNamespaces() {
		namespaces = profile.Namespaces()
	}

	namespaces = append(namespaces, namespace)
	key := single.Key()
	original := switcher.Original()
	updatedSwitcher, err := app.updateProfileInSwitcher(name, description, password, namespaces, key, original)
	if err != nil {
		return nil, err
	}

	return app.updateSwitcher(input, updatedSwitcher)
}

// Encrypt encrypts data using the current authenticated user
func (app *application) Encrypt(input loaders_identities.Identity, message []byte) ([]byte, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noCurrentUserErr)
	}

	encryptor := input.Current().Current().Key().Encryptor()
	return encryptor.Public().Encrypt(message)
}

// Decrypt decrypts data using the current authenticated user
func (app *application) Decrypt(input loaders_identities.Identity, cipher []byte) ([]byte, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noCurrentUserErr)
	}

	encryptor := input.Current().Current().Key().Encryptor()
	return encryptor.Decrypt(cipher)
}

// Sign signs data using the current authenticated user
func (app *application) Sign(input loaders_identities.Identity, message []byte) (signers.Signature, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noCurrentUserErr)
	}

	signer := input.Current().Current().Key().Signer()
	return signer.Sign(string(message))
}

// ValidateSignature validates a signature
func (app *application) ValidateSignature(input loaders_identities.Identity, message []byte, sig signers.Signature) (bool, error) {
	if !input.HasCurrent() {
		return false, errors.New(noCurrentUserErr)
	}

	expectedPubKey := input.Current().Current().Key().Signer().PublicKey()
	return sig.PublicKey(string(message)).Equals(expectedPubKey), nil
}

// Vote votes on a message using the current authenticated user
func (app *application) Vote(input loaders_identities.Identity, message []byte, ring []signers.PublicKey) (signers.Vote, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noCurrentUserErr)
	}

	signer := input.Current().Current().Key().Signer()
	return signer.Vote(string(message), ring)
}

// ValidateVote validates a vote
func (app *application) ValidateVote(input loaders_identities.Identity, message []byte, vote signers.Vote, ring []hash.Hash) (bool, error) {
	if !input.HasCurrent() {
		return false, errors.New(noCurrentUserErr)
	}

	currentPubKeyBytes := []byte(input.Current().Current().Key().Signer().PublicKey().String())
	pHash, err := app.hashAdapter.FromBytes(currentPubKeyBytes)
	if err != nil {
		return false, err
	}

	ring = append(ring, *pHash)
	return app.voteAdapter.ToVerification(vote, string(message), ring)
}

// Namespaces returns the list of namespaces of our current authenticated user
func (app *application) Namespaces(input loaders_identities.Identity) ([]string, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noCurrentUserErr)
	}

	identity := input.Current().Current()
	if !identity.HasNamespaces() {

	}

	return nil, nil
}

// Namespace sets the current namespace of our current authenticated user
func (app *application) Namespace(input loaders_identities.Identity, name string) error {
	return nil
}

// Dive dives into the namespace of our current authenticated user
func (app *application) Dive(input loaders_identities.Identity) (namespaces.Application, error) {
	return nil, nil
}

func (app *application) updateAuthenticated(
	input loaders_identities.Identity,
	update updates.Update,
) (loaders_identities.Identity, error) {
	switcher, err := app.switcherBuilder.Create().
		WithUpdated(update).
		Now()

	if err != nil {
		return nil, err
	}

	return app.updateSwitcher(input, switcher)
}

func (app *application) updateSwitcher(
	input loaders_identities.Identity,
	switcher switchers.Switcher,
) (loaders_identities.Identity, error) {
	currentSwitchers := []switchers.Switcher{}
	if input.HasAuthenticated() {
		authenticated := input.Authenticated()
		currentSwitchers = authenticated.List()
	}

	currentSwitchers = append(currentSwitchers, switcher)
	switchers, err := app.switchersBuilder.Create().
		WithList(currentSwitchers).
		Now()

	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.builder.Create().
		WithAll(all).
		WithAuthenticated(switchers)

	if input.HasCurrent() {
		current := input.Current()
		builder.WithCurrent(current)
	}

	return builder.Now()
}

func (app *application) updateProfileInSwitcher(
	name string,
	description string,
	password []byte,
	namespaces []string,
	key keys.Key,
	original singles.Single,
) (switchers.Switcher, error) {
	profileBuilder := app.profileBuilder.Create().
		WithName(name).
		WithDescription(description)

	if len(namespaces) > 0 {
		profileBuilder.WithNamespaces(namespaces)
	}

	profile, err := profileBuilder.Now()
	if err != nil {
		return nil, err
	}

	single, err := app.singleBuilder.Create().
		WithProfile(profile).
		WithKey(key).
		Now()

	if err != nil {
		return nil, err
	}

	singleBytes, err := app.singlesAdapter.ToBytes(single)
	if err != nil {
		return nil, err
	}

	cipher, err := app.encryptionApp.Encrypt(singleBytes, password)
	if err != nil {
		return nil, err
	}

	update, err := app.updateBuilder.Create().
		WithSingle(single).
		WithBytes(cipher).
		Now()

	if err != nil {
		return nil, err
	}

	return app.switcherBuilder.Create().
		WithOriginal(original).
		WithUpdated(update).
		Now()
}
