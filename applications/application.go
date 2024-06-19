package applications

import (
	"errors"

	logics_application "github.com/steve-care-software/datastencil/applications/logics"
	resources_application "github.com/steve-care-software/datastencil/applications/resources"
	"github.com/steve-care-software/datastencil/domain/contents"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	execution_links "github.com/steve-care-software/datastencil/domain/instances/executions/links"
	"github.com/steve-care-software/datastencil/domain/resources"
	"github.com/steve-care-software/datastencil/domain/resources/logics"
)

type application struct {
	resourcesApplication resources_application.Application
	logicApplication     logics_application.Application
	contentRepository    contents.Repository
	databaseRepository   databases.Repository
	executionBuiler      executions.ExecutionBuilder
	executionsBuilder    executions.Builder
	dbPath               []string
	resourcesPath        []string
}

func createApplication(
	resourcesApplication resources_application.Application,
	logicApplication logics_application.Application,
	contentRepository contents.Repository,
	databaseRepository databases.Repository,
	executionBuiler executions.ExecutionBuilder,
	executionsBuilder executions.Builder,
	dbPath []string,
	resourcesPath []string,
) Application {
	out := application{
		resourcesApplication: resourcesApplication,
		logicApplication:     logicApplication,
		contentRepository:    contentRepository,
		databaseRepository:   databaseRepository,
		executionBuiler:      executionBuiler,
		executionsBuilder:    executionsBuilder,
		dbPath:               dbPath,
		resourcesPath:        resourcesPath,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte) (executions.Executions, error) {
	return app.execute(input, nil)
}

// ExecuteWithContext executes the application with context
func (app *application) ExecuteWithContext(input []byte, context executions.Executions) (executions.Executions, error) {
	return app.execute(input, context)
}

func (app *application) execute(input []byte, context executions.Executions) (executions.Executions, error) {
	database, err := app.databaseRepository.Retrieve(app.dbPath)
	if err != nil {
		return nil, err
	}

	resources, err := app.retrieveResources(context)
	if err != nil {
		return nil, err
	}

	resourcesList := resources.List()
	for _, oneResource := range resourcesList {
		logicsList := oneResource.Logics().List()
		for _, oneLogic := range logicsList {
			executedLink, err := app.executeLogic(input, oneLogic, context)
			if err != nil {
				continue
			}

			execution, err := app.executionBuiler.Create().
				WithDatabase(database).
				WithLogic(executedLink).
				Now()

			if err != nil {
				return nil, err
			}

			executionsList := context.List()
			executionsList = append(executionsList, execution)
			return app.executionsBuilder.Create().
				WithList(executionsList).
				Now()
		}
	}

	return nil, errors.New("the request could not be executed properly")
}

func (app *application) executeLogic(input []byte, logic logics.Logic, context executions.Executions) (execution_links.Link, error) {
	if context == nil {
		return app.logicApplication.Execute(input, logic)
	}

	return app.logicApplication.ExecuteWithContext(input, logic, context)
}

func (app *application) retrieveResources(context executions.Executions) (resources.Resources, error) {
	if context == nil {
		return app.resourcesApplication.Execute(app.resourcesPath)
	}

	return app.resourcesApplication.ExecuteWithContext(app.resourcesPath, context)
}
