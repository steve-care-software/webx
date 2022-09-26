package authenticates

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity"
	identities "github.com/steve-care-software/syntax/domain/identity"
	"github.com/steve-care-software/syntax/domain/identity/connections"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/units"
	"github.com/steve-care-software/syntax/domain/identity/units/genesis"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
	"github.com/steve-care-software/syntax/domain/identity/wallets/assets"
	"github.com/steve-care-software/syntax/domain/identity/wallets/transactions"
)

type application struct {
	hashAdapter         hash.Adapter
	sigPKFactory        signatures.PrivateKeyFactory
	encryptionPKFactory keys.Factory
	walletsBuilder      wallets.Builder
	walletBuilder       wallets.WalletBuilder
	transactionsBuilder transactions.Builder
	transactionBuilder  transactions.TransactionBuilder
	assetsBuilder       assets.Builder
	assetBuilder        assets.AssetBuilder
	unitsBuilder        units.Builder
	unitBuilder         units.UnitBuilder
	unitContentBuilder  units.ContentBuilder
	genesisBuilder      genesis.Builder
	connectionsBuilder  connections.Builder
	connectionBuilder   connections.ConnectionBuilder
	identityBuilder     identity.Builder
	identityRepository  identity.Repository
	identityService     identity.Service
	publicBuilder       publics.PublicBuilder
	publicRepository    publics.Repository
	publicService       publics.Service
	name                string
	password            []byte
	genesisRingSize     uint
	unitRingSize        uint
}

func createAppliationInternally(
	hashAdapter hash.Adapter,
	sigPKFactory signatures.PrivateKeyFactory,
	encryptionPKFactory keys.Factory,
	assetsBuilder assets.Builder,
	assetBuilder assets.AssetBuilder,
	unitsBuilder units.Builder,
	unitBuilder units.UnitBuilder,
	unitContentBuilder units.ContentBuilder,
	genesisBuilder genesis.Builder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	identityBuilder identity.Builder,
	identityRepository identity.Repository,
	identityService identity.Service,
	publicBuilder publics.PublicBuilder,
	publicRepository publics.Repository,
	publicService publics.Service,
	name string,
	password []byte,
	genesisRingSize uint,
	unitRingSize uint,
) Application {
	out := application{
		hashAdapter:         hashAdapter,
		sigPKFactory:        sigPKFactory,
		encryptionPKFactory: encryptionPKFactory,
		assetsBuilder:       assetsBuilder,
		assetBuilder:        assetBuilder,
		unitsBuilder:        unitsBuilder,
		unitBuilder:         unitBuilder,
		unitContentBuilder:  unitContentBuilder,
		genesisBuilder:      genesisBuilder,
		connectionsBuilder:  connectionsBuilder,
		connectionBuilder:   connectionBuilder,
		identityBuilder:     identityBuilder,
		identityRepository:  identityRepository,
		identityService:     identityService,
		publicBuilder:       publicBuilder,
		publicRepository:    publicRepository,
		publicService:       publicService,
		name:                name,
		password:            password,
		genesisRingSize:     genesisRingSize,
		unitRingSize:        unitRingSize,
	}

	return &out
}

// Retrieve retrieves the current identity
func (app *application) Retrieve() (identities.Identity, error) {
	return app.identityRepository.Retrieve(app.name, app.password)
}

// Delete deletes the current identity
func (app *application) Delete() error {
	identity, err := app.Retrieve()
	if err != nil {
		return err
	}

	return app.identityService.Delete(identity, app.password)
}

// Disconnect disconnects from a public connection
func (app *application) Disconnect(id uuid.UUID) error {
	identity, err := app.Retrieve()
	if err != nil {
		return err
	}

	public, err := app.publicRepository.RetrieveByID(id)
	if err != nil {
		return err
	}

	list := []connections.Connection{}
	if identity.Public().HasConnections() {
		list = identity.Public().Connections().ListExcept(id)
	}

	identityPublic := identity.Public()
	identityPublicID := identityPublic.ID()
	identityPublicName := identityPublic.Name()
	identityPublicEncPubKey := identityPublic.Encryption()
	identityPublicSigHash := identityPublic.Signature()
	identityPublicHost := identityPublic.Host()
	identityPublicPort := identityPublic.Port()
	publicBuilder := app.publicBuilder.Create().WithID(identityPublicID).
		WithName(identityPublicName).
		WithEncryption(identityPublicEncPubKey).
		WithSignature(identityPublicSigHash).
		WithHost(identityPublicHost).
		WithPort(identityPublicPort)

	if len(list) > 0 {
		connections, err := app.connectionsBuilder.Create().WithList(list).Now()
		if err != nil {
			return err
		}

		publicBuilder.WithConnections(connections)
	}

	if identityPublic.HasAssets() {
		assets := identityPublic.Assets()
		publicBuilder.WithAssets(assets)
	}

	updatedIdentityPublic, err := publicBuilder.Now()
	if err != nil {
		return err
	}

	identityID := identity.ID()
	sigPK := identity.PrivateKey()
	builder := app.identityBuilder.Create().WithID(identityID).WithPublic(updatedIdentityPublic).WithPrivateKey(sigPK)
	if identity.HasWallets() {
		wallets := identity.Wallets()
		builder.WithWallets(wallets)
	}

	updated, err := builder.Now()
	if err != nil {
		return err
	}

	err = app.publicService.Delete(public)
	if err != nil {
		return err
	}

	return app.identityService.Save(updated, app.password)
}

// Generate generates a new unit
func (app *application) Generate(depositTo wallets.Wallet, supply uint64, ticker string, description string) error {
	genPK := app.sigPKFactory.Create()
	genRing, genPubKeys, err := app.generateRing(app.genesisRingSize, genPK)
	if err != nil {
		return err
	}

	genesis, err := app.genesisBuilder.Create().WithTicker(ticker).WithDescription(description).WithSupply(supply).WithOwner(genRing).Now()
	if err != nil {
		return err
	}

	unitPK := app.sigPKFactory.Create()
	unitRing, unitPubKeys, err := app.generateRing(app.unitRingSize, unitPK)
	if err != nil {
		return err
	}

	content, err := app.unitContentBuilder.Create().WithAmount(supply).WithOwner(unitRing).WithGenesis(genesis).Now()
	if err != nil {
		return err
	}

	msg := content.Hash().String()
	signature, err := genPK.RingSign(msg, genPubKeys)
	if err != nil {
		return err
	}

	unit, err := app.unitBuilder.Create().WithContent(content).WithSignatures([]signatures.RingSignature{
		signature,
	}).Now()

	if err != nil {
		return err
	}

	assetID := uuid.NewV4()
	asset, err := app.assetBuilder.Create().WithID(assetID).WithUnit(unit).WithPrivateKey(unitPK).WithRing(unitPubKeys).Now()
	if err != nil {
		return err
	}

	assetsList := []assets.Asset{}
	if depositTo.HasIncoming() {
		assetsList = depositTo.Incoming().List()
	}

	assetsList = append(assetsList, asset)
	assets, err := app.assetsBuilder.Create().WithList(assetsList).Now()
	if err != nil {
		return err
	}

	walletID := depositTo.ID()
	walletName := depositTo.Name()
	walletDescription := depositTo.Description()
	walletBuilder := app.walletBuilder.Create().WithID(walletID).WithName(walletName).WithDescription(walletDescription).WithIncoming(assets)
	if depositTo.HasOutgoing() {
		outgoing := depositTo.Outgoing()
		walletBuilder.WithOutgoing(outgoing)
	}

	updatedWallet, err := walletBuilder.Now()
	if err != nil {
		return err
	}

	identity, err := app.Retrieve()
	if err != nil {
		return err
	}

	return app.updateWallet(identity, updatedWallet)
}

// Transfer transfer a unit
func (app *application) Transfer(fromUnitHashes []hash.Hash, toOwner []hash.Hash, amount uint64, details string) error {
	identity, err := app.Retrieve()
	if err != nil {
		return err
	}

	if !identity.HasWallets() {
		str := fmt.Sprintf("the current identity (name: %s) does not own any wallet", identity.Public().Name())
		return errors.New(str)
	}

	wallet, err := identity.Wallets().FetchByUnits(fromUnitHashes)
	if err != nil {
		return err
	}

	if !wallet.HasIncoming() {
		str := fmt.Sprintf("the wallet (ID: %s, name: %s) owned by the current identity (name: %s) does not contain any asset", wallet.ID().String(), wallet.Name(), identity.Public().Name())
		return errors.New(str)
	}

	assetsList, err := wallet.Incoming().FetchByUnits(fromUnitHashes)
	if err != nil {
		return err
	}

	unitsList := []units.Unit{}
	for _, oneAsset := range assetsList {
		unitsList = append(unitsList, oneAsset.Unit())
	}

	fromUnits, err := app.unitsBuilder.Create().WithList(unitsList).Now()
	if err != nil {
		return err
	}

	content, err := app.unitContentBuilder.Create().WithAmount(amount).WithOwner(toOwner).WithUnits(fromUnits).Now()
	if err != nil {
		return err
	}

	msg := content.Hash().String()
	signatures := []signatures.RingSignature{}
	for _, oneAsset := range assetsList {
		pk := oneAsset.PrivateKey()
		ring := oneAsset.Ring()
		signature, err := pk.RingSign(msg, ring)
		if err != nil {
			return err
		}

		signatures = append(signatures, signature)
	}

	unit, err := app.unitBuilder.Create().WithContent(content).WithSignatures(signatures).Now()
	if err != nil {
		return err
	}

	trxID := uuid.NewV4()
	createdOn := time.Now().UTC()
	trxBuilder := app.transactionBuilder.Create().WithID(trxID).WithUnit(unit).CreatedOn(createdOn)
	if details != "" {
		trxBuilder.WithDetails(details)
	}

	trx, err := trxBuilder.Now()
	if err != nil {
		return err
	}

	trxList := []transactions.Transaction{}
	if wallet.HasOutgoing() {
		trxList = wallet.Outgoing().List()
	}

	trxList = append(trxList, trx)
	transactionsIns, err := app.transactionsBuilder.Create().WithList(trxList).Now()
	if err != nil {
		return err
	}

	walletID := wallet.ID()
	walletName := wallet.Name()
	walletDescription := wallet.Description()
	walletBuilder := app.walletBuilder.Create().WithID(walletID).WithName(walletName).WithDescription(walletDescription).WithOutgoing(transactionsIns)
	if wallet.HasIncoming() {
		incoming := wallet.Incoming()
		walletBuilder.WithIncoming(incoming)
	}

	updatedWallet, err := walletBuilder.Now()
	if err != nil {
		return err
	}

	return app.updateWallet(identity, updatedWallet)
}

func (app *application) updateWallet(identity identities.Identity, updatedWallet wallets.Wallet) error {
	walletsList := []wallets.Wallet{}
	if identity.HasWallets() {
		walletID := updatedWallet.ID()
		retWallets, err := identity.Wallets().FetchListExceptID(walletID)
		if err != nil {
			return err
		}

		walletsList = retWallets
	}

	walletsList = append(walletsList, updatedWallet)
	updatedWallets, err := app.walletsBuilder.Create().WithList(walletsList).Now()
	if err != nil {
		return err
	}

	identityID := identity.ID()
	public := identity.Public()
	sigPK := identity.PrivateKey()

	updated, err := app.identityBuilder.Create().
		WithID(identityID).
		WithPublic(public).
		WithPrivateKey(sigPK).
		WithWallets(updatedWallets).
		Now()

	if err != nil {
		return err
	}

	return app.identityService.Save(updated, app.password)
}

func (app *application) generateRing(size uint, pk signatures.PrivateKey) ([]hash.Hash, []signatures.PublicKey, error) {
	pubKey := pk.PublicKey()
	pubKeyHash, err := app.hashAdapter.FromBytes([]byte(pubKey.String()))
	if err != nil {
		return nil, nil, err
	}

	ring := []hash.Hash{}
	pubKeys := []signatures.PublicKey{}
	for i := uint(0); i < size; i++ {
		pubKey := app.sigPKFactory.Create().PublicKey()
		hash, err := app.hashAdapter.FromBytes([]byte(pubKey.String()))
		if err != nil {
			return nil, nil, err
		}

		ring = append(ring, *hash)
		pubKeys = append(pubKeys, pubKey)
	}

	ring = append(ring, *pubKeyHash)
	pubKeys = append(pubKeys, pubKey)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ring), func(i, j int) {
		ring[i], ring[j] = ring[j], ring[i]
		pubKeys[i], pubKeys[j] = pubKeys[j], pubKeys[i]
	})

	return ring, pubKeys, nil
}
