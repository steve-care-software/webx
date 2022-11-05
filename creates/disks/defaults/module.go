package defaults

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	creates_module "github.com/steve-care-software/webx/applications/creates/modules"
	grammar_applications "github.com/steve-care-software/webx/applications/grammars"
	identity_applications "github.com/steve-care-software/webx/applications/identities"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/domain/grammars/values"
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
	"github.com/steve-care-software/webx/domain/programs/modules"
)

type module struct {
	identityApplication            identity_applications.Application
	grammarApplication             grammar_applications.Application
	builder                        modules.Builder
	moduleBuilder                  modules.ModuleBuilder
	signaturePrivateKeyFactory     signatures.PrivateKeyFactory
	encryptionPrivateKeyFactory    keys.Factory
	identityBuilder                identities.Builder
	identityModificationsBuilder   modifications.Builder
	identityModificationBuilder    modifications.ModificationBuilder
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
	hashAdapter                    hash.Adapter
}

func createModule(
	identityApplication identity_applications.Application,
	grammarApplication grammar_applications.Application,
	//criteriaApplication criteria_applications.Application,
	builder modules.Builder,
	moduleBuilder modules.ModuleBuilder,
	signaturePrivateKeyFactory signatures.PrivateKeyFactory,
	encryptionPrivateKeyFactory keys.Factory,
	identityBuilder identities.Builder,
	identityModificationsBuilder modifications.Builder,
	identityModificationBuilder modifications.ModificationBuilder,
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
	hashAdapter hash.Adapter,
) creates_module.Application {
	out := module{
		identityApplication:            identityApplication,
		grammarApplication:             grammarApplication,
		builder:                        builder,
		moduleBuilder:                  moduleBuilder,
		signaturePrivateKeyFactory:     signaturePrivateKeyFactory,
		encryptionPrivateKeyFactory:    encryptionPrivateKeyFactory,
		identityBuilder:                identityBuilder,
		identityModificationsBuilder:   identityModificationsBuilder,
		identityModificationBuilder:    identityModificationBuilder,
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
		hashAdapter:                    hashAdapter,
	}

	return &out
}

// Execute executes the application
func (app *module) Execute() (modules.Modules, error) {
	list := []modules.Module{}
	identity, err := app.identity()
	if err != nil {
		return nil, err
	}

	grammar, err := app.grammar()
	if err != nil {
		return nil, err
	}

	/*criteria, err := app.newCriteria()
	if err != nil {
		return nil, err
	}*/

	container, err := app.container()
	if err != nil {
		return nil, err
	}

	castTo, err := app.castTo()
	if err != nil {
		return nil, err
	}

	list = append(list, identity...)
	list = append(list, grammar...)
	//list = append(list, criteria)
	list = append(list, container...)
	list = append(list, castTo...)
	return app.builder.Create().WithList(list).Now()
}

func (app *module) container() ([]modules.Module, error) {
	list, err := app.containerList()
	if err != nil {
		return nil, err
	}

	containerMapWithStringKeynames, err := app.containerMapWithStringKeynames()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		list,
		containerMapWithStringKeynames,
	}, nil
}

func (app *module) identity() ([]modules.Module, error) {
	modifyIdentity, err := app.modifyIdentity()
	if err != nil {
		return nil, err
	}

	retrieveIdentity, err := app.retrieveIdentity()
	if err != nil {
		return nil, err
	}

	insertIdentity, err := app.insertIdentity()
	if err != nil {
		return nil, err
	}

	listIdentityNames, err := app.listIdentityNames()
	if err != nil {
		return nil, err
	}

	newIdentity, err := app.newIdentity()
	if err != nil {
		return nil, err
	}

	newIdentityModification, err := app.newIdentityModification()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		modifyIdentity,
		retrieveIdentity,
		insertIdentity,
		listIdentityNames,
		newIdentity,
		newIdentityModification,
	}, nil
}

func (app *module) modifyIdentity() (modules.Module, error) {
	name := "modifyIdentity"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if modification, ok := input[app.valueToHashString("modification")].(modifications.Modification); ok {
			if currentPassword, ok := input[app.valueToHashString("currentPassword")].(string); ok {
				selectedApp, err := app.identityApplication.Select(name)
				if err != nil {
					return nil, err
				}

				newPassword := currentPassword
				if newPasswordStr, ok := input[app.valueToHashString("newPassword")].(string); ok {
					newPassword = newPasswordStr
				}

				err = selectedApp.Modify(modification, currentPassword, newPassword)
				if err != nil {
					return nil, err
				}

				return nil, nil
			}

			return nil, errors.New("the currentPassword is mandatory in order to modify an identity instance")
		}

		return nil, errors.New("the modification is mandatory in order to modify an identity instance")
	}

	return app.module(name, fn)
}

func (app *module) retrieveIdentity() (modules.Module, error) {
	name := "retrieveIdentity"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if name, ok := input[app.valueToHashString("name")].(string); ok {
			if password, ok := input[app.valueToHashString("password")].(string); ok {
				selectedApp, err := app.identityApplication.Select(name)
				if err != nil {
					return nil, err
				}

				return selectedApp.Retrieve(password)
			}

			return nil, errors.New("the password is mandatory in order to retrieve an identity instance")
		}

		return nil, errors.New("the name is mandatory in order to retrieve an identity instance")
	}

	return app.module(name, fn)
}

func (app *module) insertIdentity() (modules.Module, error) {
	name := "insertIdentity"
	fn := func(input map[string]interface{}) (interface{}, error) {
		var identity identities.Identity
		if identityIns, ok := input[app.valueToHashString("identity")].(identities.Identity); ok {
			identity = identityIns
		} else {
			return nil, errors.New("the identity is mandatory in order to save a new identity instance")
		}

		password := ""
		if passwordStr, ok := input[app.valueToHashString("password")].(string); ok {
			password = passwordStr
		} else {
			return nil, errors.New("the password is mandatory in order to save a new identity instance")
		}

		err := app.identityApplication.New(identity, password)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	return app.module(name, fn)
}

func (app *module) listIdentityNames() (modules.Module, error) {
	name := "listIdentityNames"
	fn := func(input map[string]interface{}) (interface{}, error) {
		return app.identityApplication.List()
	}

	return app.module(name, fn)
}

func (app *module) newIdentity() (modules.Module, error) {
	name := "newIdentity"
	fn := func(input map[string]interface{}) (interface{}, error) {
		sigPK := app.signaturePrivateKeyFactory.Create()
		encPK, err := app.encryptionPrivateKeyFactory.Create()
		if err != nil {
			return nil, err
		}

		createdOn := time.Now().UTC()
		modificationBuilder := app.identityModificationBuilder.Create().WithSignature(sigPK).WithEncryption(encPK).CreatedOn(createdOn)
		if name, ok := input[app.valueToHashString("name")].(string); ok {
			modificationBuilder.WithName(name)
		}

		modification, err := modificationBuilder.Now()
		if err != nil {
			return nil, err
		}

		modifications, err := app.identityModificationsBuilder.Create().WithList([]modifications.Modification{
			modification,
		}).Now()

		if err != nil {
			return nil, err
		}

		return app.identityBuilder.Create().WithModifications(modifications).Now()
	}

	return app.module(name, fn)
}

func (app *module) newIdentityModification() (modules.Module, error) {
	name := "newIdentityModification"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.identityModificationBuilder.Create()
		if name, ok := input[app.valueToHashString("name")].(string); ok {
			builder.WithName(name)
		}

		if isGenPK, ok := input[app.valueToHashString("genSigPK")].(bool); ok {
			if isGenPK {
				sigPK := app.signaturePrivateKeyFactory.Create()
				builder.WithSignature(sigPK)
			}

		}

		if isGenPK, ok := input[app.valueToHashString("genEncPK")].(bool); ok {
			if isGenPK {
				encPK, err := app.encryptionPrivateKeyFactory.Create()
				if err != nil {
					return nil, err
				}

				builder.WithEncryption(encPK)
			}
		}

		createdOn := time.Now().UTC()
		return builder.CreatedOn(createdOn).Now()
	}

	return app.module(name, fn)
}

func (app *module) grammar() ([]modules.Module, error) {
	value, err := app.newGrammarValue()
	if err != nil {
		return nil, err
	}

	cardinality, err := app.newGrammarCardinality()
	if err != nil {
		return nil, err
	}

	element, err := app.newGrammarElement()
	if err != nil {
		return nil, err
	}

	line, err := app.newGrammarLine()
	if err != nil {
		return nil, err
	}

	block, err := app.newGrammarBlock()
	if err != nil {
		return nil, err
	}

	suite, err := app.newGrammarSuite()
	if err != nil {
		return nil, err
	}

	suites, err := app.newGrammarSuites()
	if err != nil {
		return nil, err
	}

	token, err := app.newGrammarToken()
	if err != nil {
		return nil, err
	}

	everything, err := app.newGrammarEverything()
	if err != nil {
		return nil, err
	}

	instance, err := app.newGrammarInstance()
	if err != nil {
		return nil, err
	}

	external, err := app.newGrammarExternal()
	if err != nil {
		return nil, err
	}

	channelCondition, err := app.newGrammarChannelCondition()
	if err != nil {
		return nil, err
	}

	channel, err := app.newGrammarChannel()
	if err != nil {
		return nil, err
	}

	channels, err := app.newGrammarChannels()
	if err != nil {
		return nil, err
	}

	grammar, err := app.newGrammar()
	if err != nil {
		return nil, err
	}

	executeGrammar, err := app.newGrammar()
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
		executeGrammar,
	}, nil
}

func (app *module) executeGrammar() (modules.Module, error) {
	name := "executeGrammar"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if grammar, ok := input[app.valueToHashString("grammar")].(grammars.Grammar); ok {
			if data, ok := input[app.valueToHashString("data")].([]byte); ok {
				return app.grammarApplication.Execute(grammar, data)
			}

			str := fmt.Sprintf("the data was expected to be defined")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the grammar was expected to be defined")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) newGrammar() (modules.Module, error) {
	name := "newGrammar"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarBuilder.Create()
		if root, ok := input[app.valueToHashString("root")].(grammars.Token); ok {
			builder.WithRoot(root)
		}

		if channels, ok := input[app.valueToHashString("channels")].(grammars.Channels); ok {
			builder.WithChannels(channels)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarChannels() (modules.Module, error) {
	name := "newGrammarChannels"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if channelsList, ok := input[app.valueToHashString("channels")].([]interface{}); ok {
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

func (app *module) newGrammarChannel() (modules.Module, error) {
	name := "newGrammarChannel"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarChannelBuilder.Create()
		if token, ok := input[app.valueToHashString("token")].(grammars.Token); ok {
			builder.WithToken(token)
		}

		if condition, ok := input[app.valueToHashString("condition")].(grammars.ChannelCondition); ok {
			builder.WithCondition(condition)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarChannelCondition() (modules.Module, error) {
	name := "newGrammarChannelCondition"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarChannelConditionBuilder.Create()
		if previous, ok := input[app.valueToHashString("previous")].(grammars.Token); ok {
			builder.WithPrevious(previous)
		}

		if next, ok := input[app.valueToHashString("next")].(grammars.Token); ok {
			builder.WithNext(next)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarExternal() (modules.Module, error) {
	name := "newGrammarExternal"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarExternalBuilder.Create()
		if name, ok := input[app.valueToHashString("name")]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if grammar, ok := input[app.valueToHashString("grammar")].(grammars.Grammar); ok {
			builder.WithGrammar(grammar)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarInstance() (modules.Module, error) {
	name := "newGrammarInstance"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarInstanceBuilder.Create()
		if token, ok := input[app.valueToHashString("token")].(grammars.Token); ok {
			builder.WithToken(token)
		}

		if everything, ok := input[app.valueToHashString("everything")].(grammars.Everything); ok {
			builder.WithEverything(everything)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarEverything() (modules.Module, error) {
	name := "newGrammarEverything"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarEverythingBuilder.Create()
		if name, ok := input[app.valueToHashString("name")]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if exception, ok := input[app.valueToHashString("exception")].(grammars.Token); ok {
			builder.WithException(exception)
		}

		if escape, ok := input[app.valueToHashString("escape")].(grammars.Token); ok {
			builder.WithEscape(escape)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarToken() (modules.Module, error) {
	name := "newGrammarToken"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarTokenBuilder.Create()
		if name, ok := input[app.valueToHashString("name")]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if block, ok := input[app.valueToHashString("block")].(grammars.Block); ok {
			builder.WithBlock(block)
		}

		if suites, ok := input[app.valueToHashString("suites")].(grammars.Suites); ok {
			builder.WithSuites(suites)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarSuites() (modules.Module, error) {
	name := "newGrammarSuites"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if suitesList, ok := input[app.valueToHashString("suites")].([]interface{}); ok {
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

func (app *module) newGrammarSuite() (modules.Module, error) {
	name := "newGrammarSuite"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarSuiteBuilder.Create()
		if valid, ok := input[app.valueToHashString("valid")]; ok {
			if casted, ok := valid.(string); ok {
				builder.WithValid([]byte(casted))
			}

			if casted, ok := valid.([]byte); ok {
				builder.WithValid(casted)
			}
		}

		if invalid, ok := input[app.valueToHashString("invalid")]; ok {
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

func (app *module) newGrammarBlock() (modules.Module, error) {
	name := "newGrammarBlock"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if linesList, ok := input[app.valueToHashString("lines")].([]interface{}); ok {
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

func (app *module) newGrammarLine() (modules.Module, error) {
	name := "newGrammarLine"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if elementsList, ok := input[app.valueToHashString("elements")].([]interface{}); ok {
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

func (app *module) newGrammarElement() (modules.Module, error) {
	name := "newGrammarElement"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarElementBuilder.Create()
		if cardinality, ok := input[app.valueToHashString("cardinality")].(cardinalities.Cardinality); ok {
			builder.WithCardinality(cardinality)
		}

		if value, ok := input[app.valueToHashString("value")].(values.Value); ok {
			builder.WithValue(value)
		}

		if external, ok := input[app.valueToHashString("external")].(grammars.External); ok {
			builder.WithExternal(external)
		}

		if instance, ok := input[app.valueToHashString("instance")].(grammars.Instance); ok {
			builder.WithInstance(instance)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) newGrammarCardinality() (modules.Module, error) {
	name := "newGrammarCardinality"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if min, ok := input[app.valueToHashString("min")].(uint); ok {
			if min <= 0 {
				return nil, errors.New("the minimum cannot be smaller or equal than 0")
			}

			builder := app.grammarCardinalityBuilder.Create().WithMin(min)
			if max, ok := input[app.valueToHashString("max")].(uint); ok {
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

	return app.module(name, fn)
}

func (app *module) newGrammarValue() (modules.Module, error) {
	name := "newGrammarValue"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.grammarValueBuilder.Create()
		if name, ok := input[app.valueToHashString("name")]; ok {
			builder.WithName(fmt.Sprintf("%s", name))
		}

		if number, ok := input[app.valueToHashString("number")].(uint); ok {
			if number > 255 {
				return nil, errors.New("the number cannot be bigger than 255")
			}

			builder.WithNumber(byte(number))
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *module) containerMapWithStringKeynames() (modules.Module, error) {
	name := "containerMapWithStringKeynames"
	fn := func(input map[string]interface{}) (interface{}, error) {
		output := map[string]interface{}{}
		if name, ok := input[app.valueToHashString("name")].(string); ok {
			if value, ok := input[app.valueToHashString("value")]; ok {
				name = strings.TrimSpace(name)
				output[name] = value
				return output, nil
			}

			str := fmt.Sprintf("the value was expected to be declared")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the name was expected to be declared and contain a string")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) containerList() (modules.Module, error) {
	name := "containerList"
	fn := func(input map[string]interface{}) (interface{}, error) {
		values := []interface{}{}
		for keyname, element := range input {
			indexKeyname := app.valueToHashString(fmt.Sprintf("%d", len(values)))
			if keyname != indexKeyname {
				continue
			}

			values = append(values, element)
		}

		return values, nil
	}

	return app.module(name, fn)
}

func (app *module) castTo() ([]modules.Module, error) {
	toInt, err := app.castToInt()
	if err != nil {
		return nil, err
	}

	toUint, err := app.castToUint()
	if err != nil {
		return nil, err
	}

	toBool, err := app.castToBool()
	if err != nil {
		return nil, err
	}

	toFloat32, err := app.castToFloat32()
	if err != nil {
		return nil, err
	}

	toFloat64, err := app.castToFloat64()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		toInt,
		toUint,
		toBool,
		toFloat32,
		toFloat64,
	}, nil
}

func (app *module) castToInt() (modules.Module, error) {
	name := "castToInt"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if ins, ok := input[app.valueToHashString("value")]; ok {
			if casted, ok := ins.(string); ok {
				return strconv.Atoi(casted)
			}

			if casted, ok := ins.(uint); ok {
				return int(casted), nil
			}

			str := fmt.Sprintf("the value was expected to contain a string or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) castToUint() (modules.Module, error) {
	name := "castToUint"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if ins, ok := input[app.valueToHashString("value")]; ok {
			str := strings.TrimSpace(fmt.Sprintf("%s", ins))
			intValue, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}

			return uint(intValue), nil
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) castToBool() (modules.Module, error) {
	name := "castToBool"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if ins, ok := input[app.valueToHashString("value")]; ok {
			if casted, ok := ins.(string); ok {
				if strings.TrimSpace(casted) == "true" {
					return true, nil
				}

				if strings.TrimSpace(casted) == "false" {
					return false, nil
				}

				str := fmt.Sprintf("the value was expected to contain true/false when a string is provided")
				return nil, errors.New(str)
			}

			if casted, ok := ins.(int); ok {
				if casted == 0 {
					return false, nil
				}

				return true, nil
			}

			if casted, ok := ins.(uint); ok {
				if casted == 0 {
					return false, nil
				}

				return true, nil
			}

			str := fmt.Sprintf("the value was expected to contain a string, int or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) castToFloat32() (modules.Module, error) {
	name := "castToFloat32"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if ins, ok := input[app.valueToHashString("value")]; ok {
			if casted, ok := ins.(string); ok {
				floatSixtyFour, err := strconv.ParseFloat(casted, 32)
				if err != nil {
					return nil, err
				}

				return float32(floatSixtyFour), nil
			}

			if casted, ok := ins.(int); ok {
				return float32(casted), nil
			}

			if casted, ok := ins.(uint); ok {
				return float32(casted), nil
			}

			str := fmt.Sprintf("the value was expected to contain a string, int or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) castToFloat64() (modules.Module, error) {
	name := "castToFloat64"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if ins, ok := input[app.valueToHashString("value")]; ok {
			if casted, ok := ins.(string); ok {
				return strconv.ParseFloat(casted, 64)
			}

			if casted, ok := ins.(int); ok {
				return float64(casted), nil
			}

			if casted, ok := ins.(uint); ok {
				return float64(casted), nil
			}

			str := fmt.Sprintf("the value was expected to contain a string, int or uint")
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the value was expected to be valid")
		return nil, errors.New(str)
	}

	return app.module(name, fn)
}

func (app *module) valueToHashString(value string) string {
	pHash, err := app.hashAdapter.FromBytes([]byte(value))
	if err != nil {
		panic(err)
	}

	return pHash.String()
}

func (app *module) module(name string, fn modules.ExecuteFn) (modules.Module, error) {
	return app.moduleBuilder.Create().WithName([]byte(name)).WithFunc(fn).Now()
}
