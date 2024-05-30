package resources

import (
	"github.com/steve-care-software/datastencil/domain/contents"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/references"
)

type application struct {
	contentRepository contents.Repository
	pointerAdapter    pointers.Adapter
	databaseAdapter   databases.Adapter
	linkAdapter       links.Adapter
	layerAdapter      layers.LayerAdapter
	bridgesBuilder    bridges.Builder
	bridgeBuilder     bridges.BridgeBuilder
	logicBuilder      logics.LogicBuilder
	logicsBuilder     logics.Builder
	instanceAdapter   instances.Adapter
	referenceBuilder  references.ReferenceBuilder
	referencesBuilder references.Builder
	resourcesBuilder  resources.Builder
	resourceBuilder   resources.ResourceBuilder
}

func createApplication(
	contentRepository contents.Repository,
	pointerAdapter pointers.Adapter,
	databaseAdapter databases.Adapter,
	linkAdapter links.Adapter,
	layerAdapter layers.LayerAdapter,
	bridgesBuilder bridges.Builder,
	bridgeBuilder bridges.BridgeBuilder,
	logicBuilder logics.LogicBuilder,
	logicsBuilder logics.Builder,
	instanceAdapter instances.Adapter,
	referenceBuilder references.ReferenceBuilder,
	referencesBuilder references.Builder,
	resourcesBuilder resources.Builder,
	resourceBuilder resources.ResourceBuilder,
) Application {
	out := application{
		contentRepository: contentRepository,
		pointerAdapter:    pointerAdapter,
		databaseAdapter:   databaseAdapter,
		linkAdapter:       linkAdapter,
		layerAdapter:      layerAdapter,
		bridgesBuilder:    bridgesBuilder,
		bridgeBuilder:     bridgeBuilder,
		logicBuilder:      logicBuilder,
		logicsBuilder:     logicsBuilder,
		instanceAdapter:   instanceAdapter,
		referenceBuilder:  referenceBuilder,
		referencesBuilder: referencesBuilder,
		resourcesBuilder:  resourcesBuilder,
		resourceBuilder:   resourceBuilder,
	}

	return &out
}

// Execute executes application
func (app *application) Execute(path []string) (resources.Resources, error) {
	return app.execute(path, nil)
}

// ExecuteWithContext executes application with context
func (app *application) ExecuteWithContext(path []string, context executions.Executions) (resources.Resources, error) {
	return app.execute(path, context)
}

func (app *application) execute(path []string, context executions.Executions) (resources.Resources, error) {
	// load the database pointers from path:
	pointers, err := app.loadPointersFromPath(path)
	if err != nil {
		return nil, err
	}

	// fetch the database paths previously executed from the context:
	databasePaths := [][]string{}
	if context != nil {
		dbPaths, err := context.Databases()
		if err != nil {
			return nil, err
		}

		databasePaths = dbPaths
	}

	// fetch the matched pointers to the current path:
	matchedDatabasePointersList, err := pointers.Match(databasePaths)
	if err != nil {
		return nil, err
	}

	resourcesList := []resources.Resource{}
	for _, oneDatabasePointer := range matchedDatabasePointersList {
		// load the database from the pointer:
		dbPath := append(path, oneDatabasePointer.Path()...)
		database, err := app.loadDatabaseFromPath(dbPath)
		if err != nil {
			return nil, err
		}

		// load the links pointers:
		dbLinksPath := append(dbPath, database.Head().Path()...)
		linksPointers, err := app.loadPointersFromPath(dbLinksPath)
		if err != nil {
			return nil, err
		}

		// fetch the previously executed links path from that database's path:
		linkPointerPaths := [][]string{}
		if context != nil {
			linkPaths, err := context.Links(dbPath)
			if err != nil {
				return nil, err
			}

			linkPointerPaths = linkPaths
		}

		// fetch the matched links pointers:
		matchedLinkPointersList, err := linksPointers.Match(linkPointerPaths)
		if err != nil {
			return nil, err
		}

		logicsList := []logics.Logic{}
		for _, oneLinkPointer := range matchedLinkPointersList {
			// fetch the link:
			linkPath := append(path, oneLinkPointer.Path()...)
			link, err := app.loadLinkFromPath(linkPath)
			if err != nil {
				return nil, err
			}

			referencesList := []references.Reference{}
			linkReferencesList := link.References().List()
			for _, oneReference := range linkReferencesList {
				instancePath := append(linkPath, oneReference.Path()...)
				instance, err := app.loadInstanceFromPath(instancePath)
				if err != nil {
					return nil, err
				}

				variable := oneReference.Variable()
				reference, err := app.referenceBuilder.Create().WithInstance(instance).WithVariable(variable).Now()
				if err != nil {
					return nil, err
				}

				referencesList = append(referencesList, reference)
			}

			references, err := app.referencesBuilder.Create().WithList(referencesList).Now()
			if err != nil {
				return nil, err
			}

			bridgesList := []bridges.Bridge{}
			elementsList := link.Elements().List()
			for _, oneElement := range elementsList {
				layerPath := append(linkPath, oneElement.Layer()...)
				layer, err := app.loadLayerFromPath(layerPath)
				if err != nil {
					return nil, err
				}

				bridge, err := app.bridgeBuilder.Create().WithLayer(layer).WithPath(layerPath).Now()
				if err != nil {
					return nil, err
				}

				bridgesList = append(bridgesList, bridge)
			}

			bridges, err := app.bridgesBuilder.Create().WithList(bridgesList).Now()
			if err != nil {
				return nil, err

			}

			logic, err := app.logicBuilder.Create().WithLink(link).WithBridges(bridges).WithReferences(references).Now()
			if err != nil {
				return nil, err

			}

			logicsList = append(logicsList, logic)
		}

		logics, err := app.logicsBuilder.Create().WithList(logicsList).Now()
		if err != nil {
			return nil, err

		}

		head := database.Head()
		resource, err := app.resourceBuilder.Create().WithDatabase(head).WithLogics(logics).Now()
		if err != nil {
			return nil, err
		}

		resourcesList = append(resourcesList, resource)
	}

	return app.resourcesBuilder.Create().WithList(resourcesList).Now()
}

func (app *application) loadInstanceFromPath(path []string) (instances.Instance, error) {
	bytes, err := app.contentRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	return app.instanceAdapter.ToInstance(bytes)
}

func (app *application) loadLayerFromPath(path []string) (layers.Layer, error) {
	bytes, err := app.contentRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	return app.layerAdapter.ToInstance(bytes)
}

func (app *application) loadLinkFromPath(path []string) (links.Link, error) {
	bytes, err := app.contentRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	return app.linkAdapter.ToInstance(bytes)
}

func (app *application) loadDatabaseFromPath(path []string) (databases.Database, error) {
	bytes, err := app.contentRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	return app.databaseAdapter.ToInstance(bytes)
}

func (app *application) loadPointersFromPath(path []string) (pointers.Pointers, error) {
	bytes, err := app.contentRepository.Retrieve(path)
	if err != nil {
		return nil, err
	}

	return app.pointerAdapter.ToInstance(bytes)
}
