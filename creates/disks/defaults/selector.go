package defaults

import (
	"errors"
	"fmt"

	creates "github.com/steve-care-software/webx/applications/creates/selectors"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
	"github.com/steve-care-software/webx/domain/selectors"
	"github.com/steve-care-software/webx/domain/trees"
)

type selector struct {
	builder                              selectors.Builder
	selectorFnBuilder                    selectors.SelectorFnBuilder
	tokenBuilder                         selectors.TokenBuilder
	elementBuilder                       selectors.ElementBuilder
	insideBuilder                        selectors.InsideBuilder
	fetchersBuilder                      selectors.FetchersBuilder
	fetcherBuilder                       selectors.FetcherBuilder
	contentFnBuilder                     selectors.ContentFnBuilder
	instructionsBuilder                  instructions.Builder
	instructionBuilder                   instructions.InstructionBuilder
	instructionApplicationBuilder        applications.Builder
	instructionParameterBuilder          parameters.Builder
	instructionAttachmentBuilder         attachments.Builder
	instructionAttachmentVariableBuilder attachments.VariableBuilder
	instructionAssignmentBuilder         instructions.AssignmentBuilder
	instructionValueBuilder              instructions.ValueBuilder
}

func createSelector(
	builder selectors.Builder,
	selectorFnBuilder selectors.SelectorFnBuilder,
	tokenBuilder selectors.TokenBuilder,
	elementBuilder selectors.ElementBuilder,
	insideBuilder selectors.InsideBuilder,
	fetchersBuilder selectors.FetchersBuilder,
	fetcherBuilder selectors.FetcherBuilder,
	contentFnBuilder selectors.ContentFnBuilder,
	instructionsBuilder instructions.Builder,
	instructionBuilder instructions.InstructionBuilder,
	instructionApplicationBuilder applications.Builder,
	instructionParameterBuilder parameters.Builder,
	instructionAttachmentBuilder attachments.Builder,
	instructionAttachmentVariableBuilder attachments.VariableBuilder,
	instructionAssignmentBuilder instructions.AssignmentBuilder,
	instructionValueBuilder instructions.ValueBuilder,
) creates.Application {
	out := selector{
		builder:                              builder,
		selectorFnBuilder:                    selectorFnBuilder,
		tokenBuilder:                         tokenBuilder,
		elementBuilder:                       elementBuilder,
		instructionsBuilder:                  instructionsBuilder,
		insideBuilder:                        insideBuilder,
		fetchersBuilder:                      fetchersBuilder,
		fetcherBuilder:                       fetcherBuilder,
		contentFnBuilder:                     contentFnBuilder,
		instructionBuilder:                   instructionBuilder,
		instructionApplicationBuilder:        instructionApplicationBuilder,
		instructionParameterBuilder:          instructionParameterBuilder,
		instructionAttachmentBuilder:         instructionAttachmentBuilder,
		instructionAttachmentVariableBuilder: instructionAttachmentVariableBuilder,
		instructionAssignmentBuilder:         instructionAssignmentBuilder,
		instructionValueBuilder:              instructionValueBuilder,
	}

	return &out
}

// Execute returns the selector
func (app *selector) Execute() (selectors.Selector, error) {
	return app.instructions(), nil
}

func (app *selector) instructions() selectors.Selector {
	return app.selectorWithMultiFn(
		app.token(
			"instructions",
			app.element("instruction", 0),
		),
		app.insideWithSelectors([]selectors.Selector{
			app.moduleDeclaration(),
			app.applicationDeclaration(),
			app.parameter(),
			app.execute("instruction"),
			app.assignment(),
			app.attachment(),
		}),
		func(instances []interface{}) (interface{}, bool, error) {
			list := []instructions.Instruction{}
			for idx, oneIns := range instances {
				if casted, ok := oneIns.(instructions.Instruction); ok {
					list = append(list, casted)
					continue
				}

				str := fmt.Sprintf("the instruction (index: %d) could not be properly casted", idx)
				return nil, false, errors.New(str)
			}

			ins, err := app.instructionsBuilder.Create().
				WithList(list).
				Now()

			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) attachment() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"instruction",
			app.element("attachment", 0),
			0,
		),
		app.insideWithSelectors([]selectors.Selector{
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"attachment",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"attachment",
					app.element("variableReference", 1),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"attachment",
					app.element("variableReference", 2),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
		}),
		func(instances []interface{}) (interface{}, bool, error) {
			if len(instances) != 3 {
				str := fmt.Sprintf("%d elements were expected, %d returned", 3, len(instances))
				return nil, false, errors.New(str)
			}

			variable, err := app.instructionAttachmentVariableBuilder.Create().
				WithCurrent(instances[0].([]byte)).
				WithTarget(instances[1].([]byte)).
				Now()

			if err != nil {
				return nil, false, err
			}

			attachment, err := app.instructionAttachmentBuilder.Create().
				WithVariable(variable).
				WithApplication(instances[2].([]byte)).
				Now()

			if err != nil {
				return nil, false, err
			}

			ins, err := app.instructionBuilder.Create().
				WithAttachment(attachment).
				Now()

			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) assignment() selectors.Selector {
	return app.selectorWithSingleFn(
		app.token(
			"instruction",
			app.element("assignment", 0),
		),
		app.insideWithSelectors([]selectors.Selector{
			app.variableAssignment(),
			app.executionAssignment(),
			app.instructionsAssignment(),
			app.constantAssignment(),
		}),
		func(instance interface{}) (interface{}, bool, error) {
			if casted, ok := instance.(instructions.Assignment); ok {
				ins, err := app.instructionBuilder.Create().
					WithAssignment(casted).
					Now()

				if err != nil {
					return nil, false, err
				}

				return ins, true, nil
			}

			return nil, false, errors.New("the assignment could not be casted properly")
		},
	)
}

func (app *selector) constantAssignment() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"assignment",
			app.element("constantAssignment", 0),
			0,
		),
		app.insideWithSelectors([]selectors.Selector{
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"constantAssignment",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"constantAssignment",
					app.element("everythingExceptEndOfLine", 0),
					0,
				),
				app.fetchAllContentInside(),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
		}),
		func(instances []interface{}) (interface{}, bool, error) {
			if len(instances) != 2 {
				str := fmt.Sprintf("%d elements were expected, %d returned", 2, len(instances))
				return nil, false, errors.New(str)
			}

			value, err := app.instructionValueBuilder.Create().
				WithConstant(instances[1].([]byte)).
				Now()

			if err != nil {
				return nil, false, err
			}

			ins, err := app.instructionAssignmentBuilder.Create().
				WithVariable(instances[0].([]byte)).
				WithValue(value).
				Now()

			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) instructionsAssignment() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"assignment",
			app.element("instructionsAssignment", 0),
			0,
		),
		app.insideWithSelectors([]selectors.Selector{
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"instructionsAssignment",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"instructionsAssignment",
					app.element("instructions", 0),
					0,
				),
				app.insideWithRecursive("instructions"),
				func(instance interface{}) (interface{}, bool, error) {
					if casted, ok := instance.(instructions.Instructions); ok {
						return casted, true, nil
					}

					return nil, false, errors.New("the instance was expected to contain an Instructions instance")
				},
			),
		}),
		func(instances []interface{}) (interface{}, bool, error) {
			if len(instances) != 2 {
				str := fmt.Sprintf("%d elements were expected, %d returned", 2, len(instances))
				return nil, false, errors.New(str)
			}

			builder := app.instructionAssignmentBuilder.Create().
				WithVariable(instances[0].([]byte))

			if casted, ok := instances[1].(instructions.Instructions); ok {
				value, err := app.instructionValueBuilder.Create().
					WithInstructions(casted).
					Now()

				if err != nil {
					return nil, false, err
				}

				builder.WithValue(value)
			}

			ins, err := builder.Now()
			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) executionAssignment() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"assignment",
			app.element("executionAssignment", 0),
			0,
		),
		app.insideWithSelectors([]selectors.Selector{
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"executionAssignment",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
			app.execute("executionAssignment"),
		}),
		func(instances []interface{}) (interface{}, bool, error) {
			if len(instances) != 2 {
				str := fmt.Sprintf("%d elements were expected, %d returned", 2, len(instances))
				return nil, false, errors.New(str)
			}

			builder := app.instructionAssignmentBuilder.Create().
				WithVariable(instances[0].([]byte))

			if casted, ok := instances[1].(instructions.Instruction); ok {
				if !casted.IsExecution() {
					return nil, false, errors.New("the Instruction was expected to contain an Execute name")
				}

				execution := casted.Execution()
				value, err := app.instructionValueBuilder.Create().WithExecution(execution).Now()
				if err != nil {
					return nil, false, err
				}

				builder.WithValue(value)
			}

			ins, err := builder.Now()
			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) variableAssignment() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"assignment",
			app.element("variableAssignment", 0),
			0,
		),
		app.insideWithSelectors([]selectors.Selector{
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"variableAssignment",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"variableAssignment",
					app.element("variableReference", 1),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
		}),
		func(instances []interface{}) (interface{}, bool, error) {
			if len(instances) != 2 {
				str := fmt.Sprintf("%d elements were expected, %d returned", 2, len(instances))
				return nil, false, errors.New(str)
			}

			value, err := app.instructionValueBuilder.Create().WithVariable(instances[1].([]byte)).Now()
			if err != nil {
				return nil, false, err
			}

			ins, err := app.instructionAssignmentBuilder.Create().
				WithVariable(instances[0].([]byte)).
				WithValue(value).
				Now()

			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) execute(tokenName string) selectors.Selector {
	return app.selectorWithSingleFn(
		app.tokenWithContentIndex(
			tokenName,
			app.element("execute", 0),
			0,
		),
		app.insideWithSelector(
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"execute",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					ins, err := app.instructionBuilder.Create().
						WithExecution(instance.([]byte)).
						Now()

					if err != nil {
						return nil, false, err
					}

					return ins, true, nil
				},
			),
		),
		func(instance interface{}) (interface{}, bool, error) {
			if casted, ok := instance.(instructions.Instruction); ok {
				return casted, true, nil
			}

			return nil, false, errors.New("the instruction could not be casted properly")
		},
	)
}

func (app *selector) parameter() selectors.Selector {
	return app.selectorWithSingleFn(
		app.token(
			"instruction",
			app.element("parameter", 0),
		),
		app.insideWithSelectors([]selectors.Selector{
			app.inputParameter(),
			app.outputParameter(),
		}),
		func(instance interface{}) (interface{}, bool, error) {
			if casted, ok := instance.(parameters.Parameter); ok {
				ins, err := app.instructionBuilder.Create().
					WithParameter(casted).
					Now()

				if err != nil {
					return nil, false, err
				}

				return ins, true, nil
			}

			return nil, false, errors.New("the parameter could not be casted properly")
		},
	)
}

func (app *selector) outputParameter() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"parameter",
			app.element("outputParameter", 0),
			0,
		),
		app.insideWithSelector(
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"outputParameter",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
		),
		func(instances []interface{}) (interface{}, bool, error) {
			ins, err := app.instructionParameterBuilder.Create().
				WithName(instances[0].([]byte)).
				Now()

			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) inputParameter() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"parameter",
			app.element("inputParameter", 0),
			0,
		),
		app.insideWithSelector(
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"inputParameter",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
		),
		func(instances []interface{}) (interface{}, bool, error) {
			ins, err := app.instructionParameterBuilder.Create().
				WithName(instances[0].([]byte)).
				IsInput().
				Now()

			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) applicationDeclaration() selectors.Selector {
	return app.selectorWithMultiFn(
		app.tokenWithContentIndex(
			"instruction",
			app.element("applicationDeclaration", 0),
			0,
		),
		app.insideWithSelectors([]selectors.Selector{
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"applicationDeclaration",
					app.element("variableReference", 0),
					0,
				),
				app.insideWithSelector(app.variableReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"applicationDeclaration",
					app.element("moduleReference", 0),
					0,
				),
				app.insideWithSelector(app.moduleReference()),
				func(instance interface{}) (interface{}, bool, error) {
					return instance.([]byte), true, nil
				},
			),
		}),
		func(instances []interface{}) (interface{}, bool, error) {
			if len(instances) != 2 {
				str := fmt.Sprintf("%d elements were expected, %d returned", 2, len(instances))
				return nil, false, errors.New(str)
			}

			appIns, err := app.instructionApplicationBuilder.Create().
				WithName(instances[0].([]byte)).
				WithModule(instances[1].([]byte)).
				Now()

			if err != nil {
				return nil, false, err
			}

			ins, err := app.instructionBuilder.Create().
				WithApplication(appIns).
				Now()

			if err != nil {
				return nil, false, err
			}

			return ins, true, nil
		},
	)
}

func (app *selector) variableReference() selectors.Selector {
	return app.selectorWithSingleFn(
		app.tokenWithContentIndex(
			"variableReference",
			app.element("name", 0),
			0,
		),
		app.fetchAllContentInside(),
		func(instance interface{}) (interface{}, bool, error) {
			return instance.([]byte), true, nil
		},
	)
}

func (app *selector) moduleDeclaration() selectors.Selector {
	return app.selectorWithSingleFn(
		app.tokenWithContentIndex(
			"instruction",
			app.element("moduleDeclaration", 0),
			0,
		),
		app.insideWithSelector(
			app.selectorWithSingleFn(
				app.tokenWithContentIndex(
					"moduleDeclaration",
					app.element("moduleReference", 0),
					0,
				),
				app.insideWithSelector(app.moduleReference()),
				func(instance interface{}) (interface{}, bool, error) {
					ins, err := app.instructionBuilder.Create().
						WithModule(instance.([]byte)).
						Now()

					if err != nil {
						return nil, false, err
					}

					return ins, true, nil
				},
			),
		),
		func(instance interface{}) (interface{}, bool, error) {
			if casted, ok := instance.(instructions.Instruction); ok {
				return casted, true, nil
			}

			return nil, false, errors.New("the instruction could not be casted properly")
		},
	)
}

func (app *selector) moduleReference() selectors.Selector {
	return app.selectorWithSingleFn(
		app.tokenWithContentIndex(
			"moduleReference",
			app.element("name", 0),
			0,
		),
		app.fetchAllContentInside(),
		func(instance interface{}) (interface{}, bool, error) {
			return instance.([]byte), true, nil
		},
	)
}

func (app *selector) fetchAllContentInside() selectors.Inside {
	return app.insideWithFn(
		app.contentFnWithSingle(
			func(content trees.Content) ([]interface{}, error) {
				return []interface{}{
					content.Bytes(false),
				}, nil
			},
		),
	)
}

func (app *selector) selectorWithSingleFn(
	token selectors.Token,
	inside selectors.Inside,
	fn selectors.SingleSelectorFn,
) selectors.Selector {
	selectorFn, err := app.selectorFnBuilder.Create().
		WithSingle(fn).
		Now()

	if err != nil {
		panic(err)
	}

	return app.selector(token, inside, selectorFn)
}

func (app *selector) selectorWithMultiFn(
	token selectors.Token,
	inside selectors.Inside,
	fn selectors.MultiSelectorFn,
) selectors.Selector {

	selectorFn, err := app.selectorFnBuilder.Create().
		WithMulti(fn).
		Now()

	if err != nil {
		panic(err)
	}

	return app.selector(token, inside, selectorFn)
}

func (app *selector) selector(
	token selectors.Token,
	inside selectors.Inside,
	fn selectors.SelectorFn,
) selectors.Selector {
	selector, err := app.builder.Create().
		WithToken(token).
		WithInside(inside).
		WithFn(fn).
		Now()

	if err != nil {
		panic(err)
	}

	return selector
}

func (app *selector) insideWithRecursive(recursive string) selectors.Inside {
	return app.insideWithRecursives([]string{
		recursive,
	})
}

func (app *selector) insideWithRecursives(recursives []string) selectors.Inside {
	fetchersList := []selectors.Fetcher{}
	for _, oneRecursive := range recursives {
		fetcher, err := app.fetcherBuilder.Create().WithRecursive(oneRecursive).Now()
		if err != nil {
			panic(err)
		}

		fetchersList = append(fetchersList, fetcher)
	}

	fetchers, err := app.fetchersBuilder.Create().WithList(fetchersList).Now()
	if err != nil {
		panic(err)
	}

	inside, err := app.insideBuilder.Create().
		WithFetchers(fetchers).
		Now()

	if err != nil {
		panic(err)
	}

	return inside
}

func (app *selector) insideWithSelector(selector selectors.Selector) selectors.Inside {
	return app.insideWithSelectors([]selectors.Selector{
		selector,
	})
}

func (app *selector) insideWithSelectors(selectorsList []selectors.Selector) selectors.Inside {
	fetchersList := []selectors.Fetcher{}
	for _, oneSelector := range selectorsList {
		fetcher, err := app.fetcherBuilder.Create().WithSelector(oneSelector).Now()
		if err != nil {
			panic(err)
		}

		fetchersList = append(fetchersList, fetcher)
	}

	fetchers, err := app.fetchersBuilder.Create().WithList(fetchersList).Now()
	if err != nil {
		panic(err)
	}

	inside, err := app.insideBuilder.Create().
		WithFetchers(fetchers).
		Now()

	if err != nil {
		panic(err)
	}

	return inside
}

func (app *selector) insideWithFn(fn selectors.ContentFn) selectors.Inside {
	inside, err := app.insideBuilder.Create().
		WithFn(fn).
		Now()

	if err != nil {
		panic(err)
	}

	return inside
}

func (app *selector) contentFnWithSingle(fn selectors.SingleContentFn) selectors.ContentFn {
	contentFn, err := app.contentFnBuilder.Create().
		WithSingle(fn).
		Now()

	if err != nil {
		panic(err)
	}

	return contentFn
}

func (app *selector) token(
	name string,
	element selectors.Element,
) selectors.Token {
	ins, err := app.tokenBuilder.Create().
		WithName(name).
		WithElement(element).
		WithReverseName("reverse").
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *selector) tokenWithContentIndex(
	name string,
	element selectors.Element,
	contentIndex uint,
) selectors.Token {
	ins, err := app.tokenBuilder.Create().
		WithName(name).
		WithElement(element).
		WithReverseName("reverse").
		WithContent(contentIndex).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *selector) element(
	name string,
	index uint,
) selectors.Element {
	ins, err := app.elementBuilder.Create().
		WithName(name).
		WithIndex(index).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
