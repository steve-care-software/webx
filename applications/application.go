package applications

import (
	application_layers "github.com/steve-care-software/datastencil/applications/layers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

type application struct {
	execLayerApp       application_layers.Application
	repository         links.LinkRepository
	layerRepository    layers.LayerRepository
	commandsBuilder    commands.Builder
	commandBuilder     commands.CommandBuilder
	commandLinkBuilder commands.LinkBuilder
	hashAdapter        hash.Adapter
}

func createApplication(
	execLayerApp application_layers.Application,
	repository links.LinkRepository,
	layerRepository layers.LayerRepository,
	commandsBuilder commands.Builder,
	commandBuilder commands.CommandBuilder,
	commandLinkBuilder commands.LinkBuilder,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		execLayerApp:       execLayerApp,
		repository:         repository,
		layerRepository:    layerRepository,
		commandsBuilder:    commandsBuilder,
		commandBuilder:     commandBuilder,
		commandLinkBuilder: commandLinkBuilder,
		hashAdapter:        hashAdapter,
	}

	return &out
}

// Execute executes the link
func (app *application) Execute(input []byte) (commands.Commands, error) {
	link, err := app.repository.Retrieve()
	if err != nil {
		return nil, err
	}

	return app.execute(input, link, nil)
}

// ExecuteWithContext executes the lnk with context
func (app *application) ExecuteWithContext(input []byte, context commands.Commands) (commands.Commands, error) {
	originBytes := context.OriginBytes()
	pPath, err := app.hashAdapter.FromMultiBytes(originBytes)
	if err != nil {
		return nil, err
	}

	link, err := app.repository.RetrieveFromPath(*pPath)
	if err != nil {
		return nil, err
	}

	return app.execute(input, link, context)
}

func (app *application) execute(input []byte, link links.Link, context commands.Commands) (commands.Commands, error) {
	commandLinkBuilder := app.commandLinkBuilder.Create().
		WithInput(input).
		WithLink(link)

	if context != nil {
		last := context.Last()
		commandLinkBuilder.WithCommand(last)
	}

	commandLink, err := commandLinkBuilder.Now()
	if err != nil {
		return nil, err
	}

	currentInput := input
	currentContext := context
	elements := link.Elements().List()
	for _, oneElement := range elements {
		logic := oneElement.Logic()
		layer := logic.Layer()
		retResult, err := app.execLayerApp.Execute(currentInput, layer, currentContext)
		if err != nil {
			return nil, err
		}

		if retResult.IsInterruption() {
			return app.buildResultToContext(
				currentInput,
				layer,
				retResult,
				commandLink,
				currentContext,
			)
		}

		success := retResult.Success()
		kind := success.Kind()
		if kind.IsPrompt() {
			return app.buildResultToContext(
				currentInput,
				layer,
				retResult,
				commandLink,
				currentContext,
			)
		}

		commands, err := app.buildResultToContext(
			currentInput,
			layer,
			retResult,
			commandLink,
			currentContext,
		)

		if err != nil {
			return nil, err
		}

		currentInput = success.Output().Value()
		currentContext = commands
	}

	return currentContext, nil
}

func (app *application) buildResultToContext(
	input []byte,
	layer layers.Layer,
	result results.Result,
	parent commands.Link,
	context commands.Commands,
) (commands.Commands, error) {
	command, err := app.commandBuilder.Create().
		WithInput(input).
		WithLayer(layer).
		WithResult(result).
		WithParent(parent).
		Now()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	commandsList := []commands.Command{}
	if context != nil {
		commandsList = context.List()
	}

	commandsList = append(commandsList, command)
	return app.commandsBuilder.Create().
		WithList(commandsList).
		Now()
}
