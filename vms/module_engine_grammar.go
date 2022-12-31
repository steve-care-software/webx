package vms

import (
	"errors"
	"fmt"

	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/grammars/domain/grammars/values"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type moduleEngineGrammar struct {
	application             grammar_applications.Application
	builder                 grammars.Builder
	channelsBuilder         grammars.ChannelsBuilder
	channelBuilder          grammars.ChannelBuilder
	channelConditionBuilder grammars.ChannelConditionBuilder
	externalBuilder         grammars.ExternalBuilder
	instanceBuilder         grammars.InstanceBuilder
	everythingBuilder       grammars.EverythingBuilder
	tokenBuilder            grammars.TokenBuilder
	suitesBuilder           grammars.SuitesBuilder
	suiteBuilder            grammars.SuiteBuilder
	blockBuilder            grammars.BlockBuilder
	lineBuilder             grammars.LineBuilder
	elementBuilder          grammars.ElementBuilder
	cardinalityBuilder      cardinalities.Builder
	valueBuilder            values.Builder
}

func createModuleEngineGrammar(
	application grammar_applications.Application,
	builder grammars.Builder,
	channelsBuilder grammars.ChannelsBuilder,
	channelBuilder grammars.ChannelBuilder,
	channelConditionBuilder grammars.ChannelConditionBuilder,
	externalBuilder grammars.ExternalBuilder,
	instanceBuilder grammars.InstanceBuilder,
	everythingBuilder grammars.EverythingBuilder,
	tokenBuilder grammars.TokenBuilder,
	suitesBuilder grammars.SuitesBuilder,
	suiteBuilder grammars.SuiteBuilder,
	blockBuilder grammars.BlockBuilder,
	lineBuilder grammars.LineBuilder,
	elementBuilder grammars.ElementBuilder,
	cardinalityBuilder cardinalities.Builder,
	valueBuilder values.Builder,
) *moduleEngineGrammar {
	out := moduleEngineGrammar{
		application:             application,
		builder:                 builder,
		channelsBuilder:         channelsBuilder,
		channelBuilder:          channelBuilder,
		channelConditionBuilder: channelConditionBuilder,
		externalBuilder:         externalBuilder,
		instanceBuilder:         instanceBuilder,
		everythingBuilder:       everythingBuilder,
		tokenBuilder:            tokenBuilder,
		suitesBuilder:           suitesBuilder,
		suiteBuilder:            suiteBuilder,
		blockBuilder:            blockBuilder,
		lineBuilder:             lineBuilder,
		elementBuilder:          elementBuilder,
		cardinalityBuilder:      cardinalityBuilder,
		valueBuilder:            valueBuilder,
	}

	return &out
}

// Execute executes the application
func (app *moduleEngineGrammar) Execute() map[uint]modules.ExecuteFn {
	return app.grammar()
}

func (app *moduleEngineGrammar) grammar() map[uint]modules.ExecuteFn {
	value := app.newGrammarValue()
	cardinality := app.newGrammarCardinality()
	element := app.newGrammarElement()
	line := app.newGrammarLine()
	block := app.newGrammarBlock()
	suite := app.newGrammarSuite()
	suites := app.newGrammarSuites()
	token := app.newGrammarToken()
	everything := app.newGrammarEverything()
	instance := app.newGrammarInstance()
	external := app.newGrammarExternal()
	channelCondition := app.newGrammarChannelCondition()
	channel := app.newGrammarChannel()
	channels := app.newGrammarChannels()
	grammar := app.newGrammar()
	executeGrammar := app.newGrammar()
	return map[uint]modules.ExecuteFn{
		ModuleEngineGrammarValue:            value,
		ModuleEngineGrammarCardinality:      cardinality,
		ModuleEngineGrammarElement:          element,
		ModuleEngineGrammarLine:             line,
		ModuleEngineGrammarBlock:            block,
		ModuleEngineGrammarSuite:            suite,
		ModuleEngineGrammarSuites:           suites,
		ModuleEngineGrammarToken:            token,
		ModuleEngineGrammarEverything:       everything,
		ModuleEngineGrammarInstance:         instance,
		ModuleEngineGrammarExternal:         external,
		ModuleEngineGrammarChannelCondition: channelCondition,
		ModuleEngineGrammarChannel:          channel,
		ModuleEngineGrammarChannels:         channels,
		ModuleEngineGrammar:                 grammar,
		ModuleEngineExecuteGrammar:          executeGrammar,
	}
}

func (app *moduleEngineGrammar) executeGrammar() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if grammar, ok := input[0].(grammars.Grammar); ok {
			if data, ok := input[1].([]byte); ok {
				return app.application.Execute(grammar, data)
			}

			str := fmt.Sprintf("the data was expected to be defined")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the grammar was expected to be defined")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineGrammar) newGrammar() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.builder.Create()
		if root, ok := input[0].(grammars.Token); ok {
			builder.WithRoot(root)
		}

		if channels, ok := input[1].(grammars.Channels); ok {
			builder.WithChannels(channels)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarChannels() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if channelsList, ok := input[0].([]interface{}); ok {
			list := []grammars.Channel{}
			for index, oneChannel := range channelsList {
				if casted, ok := oneChannel.(grammars.Channel); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be a Channel instance", index)
				return nil, errors.New(str)
			}

			return app.channelsBuilder.Create().WithList(list).Now()
		}

		str := fmt.Sprintf("the channels was expected to be valid and contain a list")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineGrammar) newGrammarChannel() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.channelBuilder.Create()
		if token, ok := input[0].(grammars.Token); ok {
			builder.WithToken(token)
		}

		if condition, ok := input[1].(grammars.ChannelCondition); ok {
			builder.WithCondition(condition)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarChannelCondition() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.channelConditionBuilder.Create()
		if previous, ok := input[0].(grammars.Token); ok {
			builder.WithPrevious(previous)
		}

		if next, ok := input[1].(grammars.Token); ok {
			builder.WithNext(next)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarExternal() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.externalBuilder.Create()
		if name, ok := input[0]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if grammar, ok := input[1].(grammars.Grammar); ok {
			builder.WithGrammar(grammar)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarInstance() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.instanceBuilder.Create()
		if token, ok := input[0].(grammars.Token); ok {
			builder.WithToken(token)
		}

		if everything, ok := input[1].(grammars.Everything); ok {
			builder.WithEverything(everything)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarEverything() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.everythingBuilder.Create()
		if name, ok := input[0]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if exception, ok := input[1].(grammars.Token); ok {
			builder.WithException(exception)
		}

		if escape, ok := input[2].(grammars.Token); ok {
			builder.WithEscape(escape)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarToken() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.tokenBuilder.Create()
		if name, ok := input[0]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if block, ok := input[1].(grammars.Block); ok {
			builder.WithBlock(block)
		}

		if suites, ok := input[2].(grammars.Suites); ok {
			builder.WithSuites(suites)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarSuites() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if suitesList, ok := input[0].([]interface{}); ok {
			list := []grammars.Suite{}
			for index, oneSuite := range suitesList {
				if casted, ok := oneSuite.(grammars.Suite); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be a Suite instance", index)
				return nil, errors.New(str)
			}

			return app.suitesBuilder.Create().WithList(list).Now()
		}

		str := fmt.Sprintf("the suites was expected to be valid and contain a list")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineGrammar) newGrammarSuite() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.suiteBuilder.Create()
		if valid, ok := input[0]; ok {
			if casted, ok := valid.(string); ok {
				builder.WithValid([]byte(casted))
			}

			if casted, ok := valid.([]byte); ok {
				builder.WithValid(casted)
			}
		}

		if invalid, ok := input[1]; ok {
			if casted, ok := invalid.(string); ok {
				builder.WithInvalid([]byte(casted))
			}

			if casted, ok := invalid.([]byte); ok {
				builder.WithInvalid(casted)
			}
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarBlock() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if linesList, ok := input[0].([]interface{}); ok {
			list := []grammars.Line{}
			for index, oneLine := range linesList {
				if casted, ok := oneLine.(grammars.Line); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be a Line instance", index)
				return nil, errors.New(str)
			}

			return app.blockBuilder.Create().WithLines(list).Now()
		}

		str := fmt.Sprintf("the lines was expected to be valid and contain a list")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineGrammar) newGrammarLine() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if elementsList, ok := input[0].([]interface{}); ok {
			list := []grammars.Element{}
			for index, oneElement := range elementsList {
				if casted, ok := oneElement.(grammars.Element); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be an Element instance", index)
				return nil, errors.New(str)
			}

			return app.lineBuilder.Create().WithElements(list).Now()
		}

		str := fmt.Sprintf("the elements was expected to be valid and contain a list")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineGrammar) newGrammarElement() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.elementBuilder.Create()
		if cardinality, ok := input[0].(cardinalities.Cardinality); ok {
			builder.WithCardinality(cardinality)
		}

		if value, ok := input[1].(values.Value); ok {
			builder.WithValue(value)
		}

		if external, ok := input[2].(grammars.External); ok {
			builder.WithExternal(external)
		}

		if instance, ok := input[3].(grammars.Instance); ok {
			builder.WithInstance(instance)
		}

		return builder.Now()
	}
}

func (app *moduleEngineGrammar) newGrammarCardinality() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		if min, ok := input[0].(uint); ok {
			if min <= 0 {
				return nil, errors.New("the minimum cannot be smaller or equal than 0")
			}

			builder := app.cardinalityBuilder.Create().WithMin(min)
			if max, ok := input[1].(uint); ok {
				if max < 0 {
					return nil, errors.New("the maximum cannot be smaller or equal than 0")
				}

				builder.WithMax(max)
			}

			return builder.Now()
		}

		str := fmt.Sprintf("the name was expected to be valid and contain a string")
		return nil, errors.New(str)
	}
}

func (app *moduleEngineGrammar) newGrammarValue() modules.ExecuteFn {
	return func(input map[uint]interface{}) (interface{}, error) {
		builder := app.valueBuilder.Create()
		if name, ok := input[1]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if number, ok := input[0].(uint); ok {
			if number > 255 {
				return nil, errors.New("the number cannot be bigger than 255")
			}

			builder.WithNumber(byte(number))
		}

		return builder.Now()
	}
}
