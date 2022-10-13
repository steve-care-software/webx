package defaults

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	creates_module "github.com/steve-care-software/syntax/applications/engines/creates/modules"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/cardinalities"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/values"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"
)

type module struct {
	builder                        modules.Builder
	moduleBuilder                  modules.ModuleBuilder
	grammarBuilder                 grammars.Builder
	grammarChannelsBuilder         grammars.ChannelsBuilder
	grammarChannelBuilder          grammars.ChannelBuilder
	grammarChannelConditionBuilder grammars.ChannelConditionBuilder
	grammarExternalBuilder         grammars.ExternalBuilder
	grammarInstanceBuilder         grammars.InstanceBuilder
	grammarEverythingBuilder       grammars.EverythingBuilder
	grammarTokenBuilder            grammars.TokenBuilder
	grammarSuitesBuilder           grammars.SuitesBuilder
	grammarSuiteBuilder            grammars.SuiteBuilder
	grammarBlockBuilder            grammars.BlockBuilder
	grammarLineBuilder             grammars.LineBuilder
	grammarElementBuilder          grammars.ElementBuilder
	grammarCardinalityBuilder      cardinalities.Builder
	grammarValueBuilder            values.Builder
}

func createModule(
	builder modules.Builder,
	moduleBuilder modules.ModuleBuilder,
	grammarBuilder grammars.Builder,
	grammarChannelsBuilder grammars.ChannelsBuilder,
	grammarChannelBuilder grammars.ChannelBuilder,
	grammarChannelConditionBuilder grammars.ChannelConditionBuilder,
	grammarExternalBuilder grammars.ExternalBuilder,
	grammarInstanceBuilder grammars.InstanceBuilder,
	grammarEverythingBuilder grammars.EverythingBuilder,
	grammarTokenBuilder grammars.TokenBuilder,
	grammarSuitesBuilder grammars.SuitesBuilder,
	grammarSuiteBuilder grammars.SuiteBuilder,
	grammarBlockBuilder grammars.BlockBuilder,
	grammarLineBuilder grammars.LineBuilder,
	grammarElementBuilder grammars.ElementBuilder,
	grammarCardinalityBuilder cardinalities.Builder,
	grammarValueBuilder values.Builder,
) creates_module.Application {
	out := module{
		builder:                        builder,
		moduleBuilder:                  moduleBuilder,
		grammarBuilder:                 grammarBuilder,
		grammarChannelsBuilder:         grammarChannelsBuilder,
		grammarChannelBuilder:          grammarChannelBuilder,
		grammarChannelConditionBuilder: grammarChannelConditionBuilder,
		grammarExternalBuilder:         grammarExternalBuilder,
		grammarInstanceBuilder:         grammarInstanceBuilder,
		grammarEverythingBuilder:       grammarEverythingBuilder,
		grammarTokenBuilder:            grammarTokenBuilder,
		grammarSuitesBuilder:           grammarSuitesBuilder,
		grammarSuiteBuilder:            grammarSuiteBuilder,
		grammarBlockBuilder:            grammarBlockBuilder,
		grammarLineBuilder:             grammarLineBuilder,
		grammarElementBuilder:          grammarElementBuilder,
		grammarCardinalityBuilder:      grammarCardinalityBuilder,
		grammarValueBuilder:            grammarValueBuilder,
	}

	return &out
}

// Execute executes the application
func (app *module) Execute() (modules.Modules, error) {
	list := []modules.Module{}
	engine, err := app.engine()
	if err != nil {
		return nil, err
	}

	container, err := app.container()
	if err != nil {
		return nil, err
	}

	list = append(list, engine...)
	list = append(list, container...)
	return app.builder.Create().WithList(list).Now()
}

func (app *module) container() ([]modules.Module, error) {
	list, err := app.containerList()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		list,
	}, nil
}

func (app *module) engine() ([]modules.Module, error) {
	list := []modules.Module{}
	grammar, err := app.grammar()
	if err != nil {
		return nil, err
	}

	list = append(list, grammar...)
	return list, nil
}

func (app *module) grammar() ([]modules.Module, error) {
	value, err := app.engineGrammarValue()
	if err != nil {
		return nil, err
	}

	cardinality, err := app.engineGrammarCardinality()
	if err != nil {
		return nil, err
	}

	element, err := app.engineGrammarElement()
	if err != nil {
		return nil, err
	}

	line, err := app.engineGrammarLine()
	if err != nil {
		return nil, err
	}

	block, err := app.engineGrammarBlock()
	if err != nil {
		return nil, err
	}

	suite, err := app.engineGrammarSuite()
	if err != nil {
		return nil, err
	}

	suites, err := app.engineGrammarSuites()
	if err != nil {
		return nil, err
	}

	token, err := app.engineGrammarToken()
	if err != nil {
		return nil, err
	}

	everything, err := app.engineGrammarEverything()
	if err != nil {
		return nil, err
	}

	instance, err := app.engineGrammarInstance()
	if err != nil {
		return nil, err
	}

	external, err := app.engineGrammarExternal()
	if err != nil {
		return nil, err
	}

	channelCondition, err := app.engineGrammarChannelCondition()
	if err != nil {
		return nil, err
	}

	channel, err := app.engineGrammarChannel()
	if err != nil {
		return nil, err
	}

	channels, err := app.engineGrammarChannels()
	if err != nil {
		return nil, err
	}

	grammar, err := app.engineGrammar()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		value,
		cardinality,
		element,
		line,
		block,
		suite,
		suites,
		token,
		everything,
		instance,
		external,
		channelCondition,
		channel,
		channels,
		grammar,
	}, nil
}

func (app *module) engineGrammar() (modules.Module, error) {
	name := "engineGrammar"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarBuilder.Create()
		if root, ok := input["root"].(grammars.Token); ok {
			builder.WithRoot(root)
		}

		if channels, ok := input["channels"].(grammars.Channels); ok {
			builder.WithChannels(channels)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarChannels() (modules.Module, error) {
	name := "engineGrammarChannels"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if channelsList, ok := input["channels"].([]interface{}); ok {
			list := []grammars.Channel{}
			for index, oneChannel := range channelsList {
				if casted, ok := oneChannel.(grammars.Channel); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be a Channel instance", index)
				return nil, errors.New(str)
			}

			return app.grammarChannelsBuilder.Create().WithList(list).Now()
		}

		str := fmt.Sprintf("the channels was expected to be valid and contain a list")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarChannel() (modules.Module, error) {
	name := "engineGrammarChannel"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarChannelBuilder.Create()
		if name, ok := input["name"].(string); ok {
			builder.WithName(name)
		}

		if token, ok := input["token"].(grammars.Token); ok {
			builder.WithToken(token)
		}

		if condition, ok := input["condition"].(grammars.ChannelCondition); ok {
			builder.WithCondition(condition)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarChannelCondition() (modules.Module, error) {
	name := "engineGrammarChannelCondition"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarChannelConditionBuilder.Create()
		if previous, ok := input["previous"].(grammars.Token); ok {
			builder.WithPrevious(previous)
		}

		if next, ok := input["next"].(grammars.Token); ok {
			builder.WithNext(next)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarExternal() (modules.Module, error) {
	name := "engineGrammarExternal"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarExternalBuilder.Create()
		if name, ok := input["name"].(string); ok {
			builder.WithName(name)
		}

		if grammar, ok := input["grammar"].(grammars.Grammar); ok {
			builder.WithGrammar(grammar)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarInstance() (modules.Module, error) {
	name := "engineGrammarInstance"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarInstanceBuilder.Create()
		if token, ok := input["token"].(grammars.Token); ok {
			builder.WithToken(token)
		}

		if everything, ok := input["everything"].(grammars.Everything); ok {
			builder.WithEverything(everything)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarEverything() (modules.Module, error) {
	name := "engineGrammarEverything"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarEverythingBuilder.Create()
		if name, ok := input["name"].(string); ok {
			builder.WithName(name)
		}

		if exception, ok := input["exception"].(grammars.Token); ok {
			builder.WithException(exception)
		}

		if escape, ok := input["escape"].(grammars.Token); ok {
			builder.WithEscape(escape)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarToken() (modules.Module, error) {
	name := "engineGrammarToken"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarTokenBuilder.Create()
		if name, ok := input["name"].(string); ok {
			builder.WithName(name)
		}

		if block, ok := input["block"].(grammars.Block); ok {
			builder.WithBlock(block)
		}

		if suites, ok := input["suites"].(grammars.Suites); ok {
			builder.WithSuites(suites)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarSuites() (modules.Module, error) {
	name := "engineGrammarSuites"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if suitesList, ok := input["suites"].([]interface{}); ok {
			list := []grammars.Suite{}
			for index, oneSuite := range suitesList {
				if casted, ok := oneSuite.(grammars.Suite); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be a Suite instance", index)
				return nil, errors.New(str)
			}

			return app.grammarSuitesBuilder.Create().WithList(list).Now()
		}

		str := fmt.Sprintf("the suites was expected to be valid and contain a list")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarSuite() (modules.Module, error) {
	name := "engineGrammarSuite"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarSuiteBuilder.Create()
		if valid, ok := input["valid"]; ok {
			if casted, ok := valid.(string); ok {
				builder.WithValid([]byte(casted))
			}

			if casted, ok := valid.([]byte); ok {
				builder.WithValid(casted)
			}
		}

		if invalid, ok := input["invalid"]; ok {
			if casted, ok := invalid.(string); ok {
				builder.WithInvalid([]byte(casted))
			}

			if casted, ok := invalid.([]byte); ok {
				builder.WithInvalid(casted)
			}
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarBlock() (modules.Module, error) {
	name := "engineGrammarBlock"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if linesList, ok := input["lines"].([]interface{}); ok {
			list := []grammars.Line{}
			for index, oneLine := range linesList {
				if casted, ok := oneLine.(grammars.Line); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be a Line instance", index)
				return nil, errors.New(str)
			}

			return app.grammarBlockBuilder.Create().WithLines(list).Now()
		}

		str := fmt.Sprintf("the lines was expected to be valid and contain a list")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarLine() (modules.Module, error) {
	name := "engineGrammarLine"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if elementsList, ok := input["elements"].([]interface{}); ok {
			list := []grammars.Element{}
			for index, oneElement := range elementsList {
				if casted, ok := oneElement.(grammars.Element); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the value at index: %d was expected to be an Element instance", index)
				return nil, errors.New(str)
			}

			return app.grammarLineBuilder.Create().WithElements(list).Now()
		}

		str := fmt.Sprintf("the elements was expected to be valid and contain a list")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarElement() (modules.Module, error) {
	name := "engineGrammarElement"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarElementBuilder.Create()
		if cardinality, ok := input["cardinality"].(cardinalities.Cardinality); ok {
			builder.WithCardinality(cardinality)
		}

		if value, ok := input["value"].(values.Value); ok {
			builder.WithValue(value)
		}

		if external, ok := input["external"].(grammars.External); ok {
			builder.WithExternal(external)
		}

		if instance, ok := input["instance"].(grammars.Instance); ok {
			builder.WithInstance(instance)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarCardinality() (modules.Module, error) {
	name := "engineGrammarCardinality"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if min, ok := input["min"].(string); ok {
			intMin, err := strconv.Atoi(strings.TrimSpace(min))
			if err != nil {
				return nil, err
			}

			if intMin <= 0 {
				return nil, errors.New("the minimum cannot be smaller or equal than 0")
			}

			builder := app.grammarCardinalityBuilder.Create().WithMin(uint(intMin))
			if max, ok := input["max"].(string); ok {
				intMax, err := strconv.Atoi(strings.TrimSpace(max))
				if err != nil {
					return nil, err
				}

				if intMax < 0 {
					return nil, errors.New("the maximum cannot be smaller or equal than 0")
				}

				builder.WithMax(uint(intMax))
			}

			return builder.Now()
		}

		str := fmt.Sprintf("the name was expected to be valid and contain a string")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) engineGrammarValue() (modules.Module, error) {
	name := "engineGrammarValue"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if name, ok := input["name"].(string); ok {
			if number, ok := input["number"].(string); ok {
				intNumber, err := strconv.Atoi(strings.TrimSpace(number))
				if err != nil {
					return nil, err
				}

				if intNumber < 0 {
					return nil, errors.New("the number cannot be smaller than 0")
				}

				if intNumber > 255 {
					return nil, errors.New("the number cannot be bigger than 255")
				}

				return app.grammarValueBuilder.Create().
					WithName(name).
					WithNumber(byte(intNumber)).
					Now()
			}

			str := fmt.Sprintf("the number was expected to be valid and contain a string")
			return nil, errors.New(str)

		}

		str := fmt.Sprintf("the name was expected to be valid and contain a string")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) containerList() (modules.Module, error) {
	name := "containerList"
	fn := func(input map[string]interface{}) (interface{}, error) {
		values := []interface{}{}
		for keyname, element := range input {
			index := fmt.Sprintf("%d", len(values))
			if strings.TrimSpace(keyname) != index {
				continue
			}

			values = append(values, element)
		}

		return values, nil
	}

	return app.module(name, fn)
}

func (app *module) module(name string, fn modules.ExecuteFn) (modules.Module, error) {
	return app.moduleBuilder.Create().WithName(name).WithFunc(fn).Now()
}
