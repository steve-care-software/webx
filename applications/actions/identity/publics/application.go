package publics

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/connections"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/pendings"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets"
	"github.com/steve-care-software/syntax/domain/identity/units"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

type application struct {
	encryptionPKFactory keys.Factory
	connectionsBuilder  connections.Builder
	connectionBuilder   connections.ConnectionBuilder
	publicBuilder       publics.PublicBuilder
	publicRepository    publics.Repository
	publicService       publics.Service
	depositTo           wallets.Wallet
	unitRingSize        uint
	name                string
}

func createApplication(
	encryptionPKFactory keys.Factory,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	publicBuilder publics.PublicBuilder,
	publicRepository publics.Repository,
	publicService publics.Service,
	depositTo wallets.Wallet,
	unitRingSize uint,
	name string,
) Application {
	out := application{
		encryptionPKFactory: encryptionPKFactory,
		connectionsBuilder:  connectionsBuilder,
		connectionBuilder:   connectionBuilder,
		publicBuilder:       publicBuilder,
		publicRepository:    publicRepository,
		publicService:       publicService,
		depositTo:           depositTo,
		unitRingSize:        unitRingSize,
		name:                name,
	}

	return &out
}

// Retrieve returns the current public identity
func (app *application) Retrieve() (publics.Public, error) {
	return app.publicRepository.RetrieveByName(app.name)
}

// Prepare prepares a pending reception
func (app *application) Prepare() (pendings.Pending, error) {
	return nil, nil
}

// Receive receives a unit
func (app *application) Receive(pending pendings.Pending, unit units.Unit) error {
	return nil
}

// Connect connects with an external user
func (app *application) Connect(connection publics.Public) error {
	currentPublic, err := app.Retrieve()
	if err != nil {
		return err
	}

	err = app.publicService.Save(connection)
	if err != nil {
		return err
	}

	list := []connections.Connection{}
	if currentPublic.HasConnections() {
		list = currentPublic.Connections().List()
	}

	connID := uuid.NewV4()
	pk, err := app.encryptionPKFactory.Create()
	if err != nil {
		return err
	}

	connPublicID := connection.ID()
	conn, err := app.connectionBuilder.Create().WithID(connID).WithPublic(connPublicID).WithEncryption(pk).Now()
	if err != nil {
		return err
	}

	list = append(list, conn)
	connections, err := app.connectionsBuilder.Create().WithList(list).Now()
	if err != nil {
		return err
	}

	id := currentPublic.ID()
	name := currentPublic.Name()
	encPubKey := currentPublic.Encryption()
	sigHash := currentPublic.Signature()
	host := currentPublic.Host()
	port := currentPublic.Port()
	builder := app.publicBuilder.Create().
		WithID(id).
		WithName(name).
		WithEncryption(encPubKey).
		WithSignature(sigHash).
		WithHost(host).
		WithPort(port).
		WithConnections(connections)

	if currentPublic.HasAssets() {
		assets := currentPublic.Assets()
		builder.WithAssets(assets)
	}

	updated, err := builder.Now()
	if err != nil {
		return err
	}

	return app.publicService.Save(updated)
}

// Assets add assets to the given public profile
func (app *application) Assets(publicID uuid.UUID, newAssets assets.Assets) error {
	return nil
}

// Asset add asset to the given public profile
func (app *application) Asset(publicID uuid.UUID, newAsset assets.Asset) error {
	return nil
}
