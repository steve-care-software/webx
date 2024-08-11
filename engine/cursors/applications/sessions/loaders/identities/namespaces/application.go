package namespaces

import (
	"crypto/rand"
	"errors"
	"math"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers"
	namespace_switchers "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers"
	namespace_singles "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/roots"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/roots/units"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/roots/units/purses"
	namespace_switchers_updates "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/updates"
)

type application struct {
	builder                   namespaces.Builder
	namespaceBuilder          namespaces.NamespaceBuilder
	namespaceSwitchersBuilder namespace_switchers.Builder
	namespaceSwitcherBuilder  namespace_switchers.SwitcherBuilder
	namespaceUpdateBuilder    namespace_switchers_updates.Builder
	namespaceSingleAdapter    namespace_singles.Adapter
	namespaceSingleBuilder    namespace_singles.Builder
	storageBuilder            storages.Builder
	blockchainBuilder         blockchains.Builder
	rootBuilder               roots.Builder
	unitBuilder               units.Builder
	pursesBuilder             purses.Builder
	purseBuilder              purses.PurseBuilder
}

func createApplication(
	builder namespaces.Builder,
	namespaceBuilder namespaces.NamespaceBuilder,
	namespaceSwitchersBuilder namespace_switchers.Builder,
	namespaceSwitcherBuilder namespace_switchers.SwitcherBuilder,
	namespaceUpdateBuilder namespace_switchers_updates.Builder,
	namespaceSingleAdapter namespace_singles.Adapter,
	namespaceSingleBuilder namespace_singles.Builder,
	storageBuilder storages.Builder,
	blockchainBuilder blockchains.Builder,
	rootBuilder roots.Builder,
	unitBuilder units.Builder,
	pursesBuilder purses.Builder,
	purseBuilder purses.PurseBuilder,
) Application {
	out := application{
		builder:                   builder,
		namespaceBuilder:          namespaceBuilder,
		namespaceSwitchersBuilder: namespaceSwitchersBuilder,
		namespaceSwitcherBuilder:  namespaceSwitcherBuilder,
		namespaceUpdateBuilder:    namespaceUpdateBuilder,
		namespaceSingleAdapter:    namespaceSingleAdapter,
		namespaceSingleBuilder:    namespaceSingleBuilder,
		storageBuilder:            storageBuilder,
		blockchainBuilder:         blockchainBuilder,
		rootBuilder:               rootBuilder,
		unitBuilder:               unitBuilder,
		pursesBuilder:             pursesBuilder,
		purseBuilder:              purseBuilder,
	}

	return &out
}

// Loaded returns the loaded namespaces
func (app *application) Loaded(input namespaces.Namespace) ([]string, error) {
	if !input.HasLoaded() {
		return nil, errors.New(noLoadedNamespaceErr)
	}

	return input.Loaded().Names(), nil
}

// Create creates a new namespace
func (app *application) Create(
	input namespaces.Namespace,
	name string,
	symbol string,
	description string,
	baseDifficulty uint64,
	increasePerSize uint64,
	sizeBlock uint64,
	totalUnitsAmount uint64,
	purseAmount uint64,
) (namespaces.Namespace, error) {
	blockchain, err := app.createBlockchain(
		symbol,
		baseDifficulty,
		increasePerSize,
		sizeBlock,
		totalUnitsAmount,
		purseAmount,
	)

	if err != nil {
		return nil, err
	}

	single, err := app.namespaceSingleBuilder.Create().
		WithBlockchain(blockchain).
		WithName(name).
		WithDescription(description).
		Now()

	if err != nil {
		return nil, err
	}

	singleBytes, err := app.namespaceSingleAdapter.ToBytes(single)
	if err != nil {
		return nil, err
	}

	update, err := app.namespaceUpdateBuilder.Create().WithSingle(single).WithBytes(singleBytes).Now()
	if err != nil {
		return nil, err
	}

	switcher, err := app.namespaceSwitcherBuilder.Create().WithUpdated(update).Now()
	if err != nil {
		return nil, err
	}

	loaded := []switchers.Switcher{}
	if input.HasLoaded() {
		loaded = append(loaded, input.Loaded().List()...)
	}

	loaded = append(loaded, switcher)
	loadedSwitchers, err := app.namespaceSwitchersBuilder.Create().WithList(loaded).Now()
	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.namespaceBuilder.Create().WithAll(all).WithLoaded(loadedSwitchers)
	if input.HasCurrent() {
		current := input.Current()
		builder.WithCurrent(current)
	}

	return builder.Now()
}

// Set sets the current namespace
func (app *application) Set(input namespaces.Namespace, name string) (namespaces.Namespace, error) {
	return nil, nil
}

func (app *application) createBlockchain(
	symbol string,
	baseDifficulty uint64,
	increasePerSize uint64,
	sizeBlock uint64,
	totalUnitsAmount uint64,
	purseAmount uint64,
) (blockchains.Blockchain, error) {
	amountPerPurse := uint64(totalUnitsAmount / purseAmount)
	remaining := uint64(math.Mod(float64(totalUnitsAmount), float64(purseAmount)))

	pursesList := []purses.Purse{}
	castedPurseAmount := int(purseAmount)
	for i := 0; i < castedPurseAmount; i++ {
		amountInPurse := amountPerPurse
		isLastIteration := i+1 >= castedPurseAmount
		if isLastIteration {
			amountInPurse = remaining
		}

		purse, err := app.purseBuilder.Create().WithIndex(uint64(i)).WithAmount(amountInPurse).Now()
		if err != nil {
			return nil, err
		}

		pursesList = append(pursesList, purse)
	}

	purses, err := app.pursesBuilder.Create().WithList(pursesList).Now()
	if err != nil {
		return nil, err
	}

	lock := make([]byte, 512)
	_, err = rand.Read(lock)
	if err != nil {
		return nil, err
	}

	units, err := app.unitBuilder.Create().
		WithSymbol(symbol).
		WithPurses(purses).
		WithLock(lock).
		Now()

	if err != nil {
		return nil, err
	}

	root, err := app.rootBuilder.Create().
		WithUnits(units).
		WithBaseDifficulty(baseDifficulty).
		WithIncreasePerSize(increasePerSize).
		WithSizeBlock(sizeBlock).
		Now()

	if err != nil {
		return nil, err
	}

	return app.blockchainBuilder.Create().
		WithRoot(root).
		Now()
}
