package instances

import (
	"strconv"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/webx/grammars/domain/grammars/values"
)

type grammar struct {
	builder            grammars.Builder
	channelsBuilder    grammars.ChannelsBuilder
	channelBuilder     grammars.ChannelBuilder
	instanceBuilder    grammars.InstanceBuilder
	everythingBuilder  grammars.EverythingBuilder
	tokensBuilder      grammars.TokensBuilder
	tokenBuilder       grammars.TokenBuilder
	suitesBuilder      grammars.SuitesBuilder
	suiteBuilder       grammars.SuiteBuilder
	blockBuilder       grammars.BlockBuilder
	lineBuilder        grammars.LineBuilder
	elementBuilder     grammars.ElementBuilder
	valueBuilder       grammar_values.Builder
	cardinalityBuilder cardinalities.Builder
}

func createGrammar(
	builder grammars.Builder,
	channelsBuilder grammars.ChannelsBuilder,
	channelBuilder grammars.ChannelBuilder,
	instanceBuilder grammars.InstanceBuilder,
	everythingBuilder grammars.EverythingBuilder,
	tokensBuilder grammars.TokensBuilder,
	tokenBuilder grammars.TokenBuilder,
	suitesBuilder grammars.SuitesBuilder,
	suiteBuilder grammars.SuiteBuilder,
	blockBuilder grammars.BlockBuilder,
	lineBuilder grammars.LineBuilder,
	elementBuilder grammars.ElementBuilder,
	valueBuilder grammar_values.Builder,
	cardinalityBuilder cardinalities.Builder,
) *grammar {
	out := grammar{
		builder:            builder,
		channelsBuilder:    channelsBuilder,
		channelBuilder:     channelBuilder,
		instanceBuilder:    instanceBuilder,
		everythingBuilder:  everythingBuilder,
		tokensBuilder:      tokensBuilder,
		tokenBuilder:       tokenBuilder,
		suitesBuilder:      suitesBuilder,
		suiteBuilder:       suiteBuilder,
		blockBuilder:       blockBuilder,
		lineBuilder:        lineBuilder,
		elementBuilder:     elementBuilder,
		valueBuilder:       valueBuilder,
		cardinalityBuilder: cardinalityBuilder,
	}

	return &out
}

// Execute executes a grammar
func (app *grammar) Execute() (grammars.Grammar, error) {
	root := app.instructions()
	channels := app.channels()
	return app.builder.Create().
		WithRoot(root).
		WithChannels(channels).
		Now()
}

func (app grammar) channels() grammars.Channels {
	list := []grammars.Channel{
		app.channelFromValue("space", []byte(" ")[0]),
		app.channelFromValue("tab", []byte("\t")[0]),
		app.channelFromValue("newLine", []byte("\n")[0]),
		app.channelFromValue("retChar", []byte("\r")[0]),
		app.channelFromToken(
			app.tokenFromBlock("singleLineComment",
				app.blockFromlines([]grammars.Line{
					app.lineFromElements([]grammars.Element{
						app.elementFromToken(
							app.allCharacterToken("doubleSlash", "//"),
							app.cardinalityOnce(),
						),
						app.elementFromEverything(
							app.everythingWithoutEscape(
								"everythingExceptEndOfLine",
								app.anyElementToken(
									"endOfLineSpaces",
									[]grammars.Element{
										app.elementFromValue([]byte("\n")[0]),
										app.elementFromValue([]byte("\r")[0]),
									},
								),
							),
						),
					}),
				}),
			),
		),
	}

	ins, err := app.channelsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) instructions() grammars.Token {
	return app.tokenFromBlock(
		"instructions",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.instruction(), app.cardinality(1, nil)),
			}),
		}),
	)
}

func (app *grammar) instruction() grammars.Token {
	return app.tokenFromBlock(
		"instruction",
		app.blockFromlines([]grammars.Line{
			// module declaration:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.moduleDeclaration(), app.cardinalityOnce()),
				app.elementFromToken(app.endOfLineConstant(), app.cardinalityOnce()),
			}),

			// application declaration:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.applicationDeclaration(), app.cardinalityOnce()),
				app.elementFromToken(app.endOfLineConstant(), app.cardinalityOnce()),
			}),

			// parameter:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.parameter(), app.cardinalityOnce()),
				app.elementFromToken(app.endOfLineConstant(), app.cardinalityOnce()),
			}),

			// assignment:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.assignment(), app.cardinalityOnce()),
				app.elementFromToken(app.endOfLineConstant(), app.cardinalityOnce()),
			}),

			// execute:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.execute(), app.cardinalityOnce()),
				app.elementFromToken(app.endOfLineConstant(), app.cardinalityOnce()),
			}),

			// attach:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.attachment(), app.cardinalityOnce()),
				app.elementFromToken(app.endOfLineConstant(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) attachment() grammars.Token {
	return app.tokenFromBlock(
		"attachment",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.attachConstant(), app.cardinalityOnce()),
				app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
				app.elementFromValue([]byte(":")[0]),
				app.elementFromToken(app.attachmentTarget(), app.cardinalityOnce()),
				app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) attachmentTarget() grammars.Token {
	return app.tokenFromBlock(
		"attachmentTarget",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.number(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) attachConstant() grammars.Token {
	return app.allCharacterToken("attachConstant", "attach")
}

func (app *grammar) assignment() grammars.Token {
	return app.tokenFromBlock(
		"assignment",
		app.blockFromlines([]grammars.Line{
			// variable:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.tokenFromBlock(
						"variableAssignment",
						app.blockFromlines([]grammars.Line{
							app.lineFromElements([]grammars.Element{
								app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
								app.elementFromValue([]byte("=")[0]),
								app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
							}),
						}),
					),
					app.cardinalityOnce(),
				),
			}),

			// execution:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.tokenFromBlock(
						"executionAssignment",
						app.blockFromlines([]grammars.Line{
							app.lineFromElements([]grammars.Element{
								app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
								app.elementFromValue([]byte("=")[0]),
								app.elementFromToken(app.execute(), app.cardinalityOnce()),
							}),
						}),
					),
					app.cardinalityOnce(),
				),
			}),

			// instructions:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.tokenFromBlock(
						"instructionsAssignment",
						app.blockFromlines([]grammars.Line{
							app.lineFromElements([]grammars.Element{
								app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
								app.elementFromValue([]byte("=")[0]),
								app.elementFromValue([]byte("{")[0]),
								app.elementFromRecursiveToken("instructions", app.cardinalityOnce()),
								app.elementFromValue([]byte("}")[0]),
							}),
						}),
					),
					app.cardinalityOnce(),
				),
			}),

			// constant:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.tokenFromBlock(
						"constantAssignment",
						app.blockFromlines([]grammars.Line{
							app.lineFromElements([]grammars.Element{
								app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
								app.elementFromValue([]byte("=")[0]),
								app.elementFromEverything(
									app.everything(
										"everythingExceptEndOfLine",
										app.endOfLineConstant(),
										app.assignmentConstantEscape(),
									),
								),
							}),
						}),
					),
					app.cardinalityOnce(),
				),
			}),
		}),
	)
}

func (app *grammar) execute() grammars.Token {
	return app.tokenFromBlock(
		"execute",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.executeConstant(), app.cardinalityOnce()),
				app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) executeConstant() grammars.Token {
	return app.allCharacterToken("executeConstant", "execute")
}

func (app *grammar) parameter() grammars.Token {
	return app.tokenFromBlock(
		"parameter",
		app.blockFromlines([]grammars.Line{
			// input:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.tokenFromBlock(
						"outputParameter",
						app.blockFromlines([]grammars.Line{
							app.lineFromElements([]grammars.Element{
								app.elementFromToken(app.directionOutputConstant(), app.cardinalityOnce()),
								app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
							}),
						}),
					),
					app.cardinalityOnce(),
				),
			}),

			// output:
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.tokenFromBlock(
						"inputParameter",
						app.blockFromlines([]grammars.Line{
							app.lineFromElements([]grammars.Element{
								app.elementFromToken(app.directionInputConstant(), app.cardinalityOnce()),
								app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
							}),
						}),
					),
					app.cardinalityOnce(),
				),
			}),
		}),
	)
}

func (app *grammar) directionOutputConstant() grammars.Token {
	return app.allCharacterToken("directionOutputConstant", "<-")
}

func (app *grammar) directionInputConstant() grammars.Token {
	return app.allCharacterToken("directionInputConstant", "->")
}

func (app *grammar) applicationDeclaration() grammars.Token {
	return app.tokenFromBlock(
		"applicationDeclaration",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.moduleReference(), app.cardinalityOnce()),
				app.elementFromToken(app.variableReference(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) variableReference() grammars.Token {
	return app.tokenFromBlock(
		"variableReference",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte("$")[0]),
				app.elementFromToken(app.name(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) moduleDeclaration() grammars.Token {
	return app.tokenFromBlock(
		"moduleDeclaration",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.moduleConstant(), app.cardinalityOnce()),
				app.elementFromToken(app.moduleReference(), app.cardinalityOnce()),
				app.elementFromToken(app.moduleIndex(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) moduleIndex() grammars.Token {
	return app.tokenFromBlock(
		"moduleIndex",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte(":")[0]),
				app.elementFromToken(app.number(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) number() grammars.Token {
	return app.tokenFromBlock(
		"number",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.anyNumber(),
					app.cardinality(1, nil),
				),
			}),
		}),
	)
}

func (app *grammar) moduleReference() grammars.Token {
	return app.tokenFromBlock(
		"moduleReference",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromValue([]byte("@")[0]),
				app.elementFromToken(app.name(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) name() grammars.Token {
	return app.tokenFromBlock(
		"name",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(
					app.alphaNumeric(),
					app.cardinality(1, nil),
				),
			}),
		}),
	)
}

func (app *grammar) assignmentConstantEscape() grammars.Token {
	return app.allCharacterToken("assignmentConstantEscape", "\\")
}

func (app *grammar) endOfLineConstant() grammars.Token {
	return app.allCharacterToken("endOfLineConstant", ";;")
}

func (app *grammar) moduleConstant() grammars.Token {
	return app.allCharacterToken("moduleConstant", "module")
}

func (app *grammar) alphaNumeric() grammars.Token {
	return app.tokenFromBlock(
		"alphaNumeric",
		app.blockFromlines([]grammars.Line{
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.anyNumber(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.upperCaseLetter(), app.cardinalityOnce()),
			}),
			app.lineFromElements([]grammars.Element{
				app.elementFromToken(app.lowerCaseLetters(), app.cardinalityOnce()),
			}),
		}),
	)
}

func (app *grammar) anyNumber() grammars.Token {
	characters := "0123456789"
	return app.anyCharacterToken("anyNumber", characters)
}

func (app *grammar) upperCaseLetter() grammars.Token {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return app.anyCharacterToken("uppercaseLetter", characters)
}

func (app *grammar) lowerCaseLetters() grammars.Token {
	characters := "abcdefghijklmnopqrstuvwxyz"
	return app.anyCharacterToken("lowerCaseLetter", characters)
}

func (app *grammar) allCharacterToken(tokenName string, values string) grammars.Token {
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.elementFromValue(byte(oneValue))
		elementsList = append(elementsList, element)
	}

	return app.allElementsToken(tokenName, elementsList)
}

func (app *grammar) allElementsToken(tokenName string, elementsList []grammars.Element) grammars.Token {
	return app.tokenFromBlock(
		tokenName,
		app.blockFromlines([]grammars.Line{
			app.lineFromElements(elementsList),
		}),
	)
}

func (app *grammar) anyCharacterToken(tokenName string, values string) grammars.Token {
	elementsList := []grammars.Element{}
	for _, oneValue := range values {
		element := app.elementFromValue(byte(oneValue))
		elementsList = append(elementsList, element)
	}

	return app.anyElementToken(tokenName, elementsList)
}

func (app *grammar) anyElementToken(tokenName string, elementsList []grammars.Element) grammars.Token {
	linesList := []grammars.Line{}
	for _, oneElement := range elementsList {
		linesList = append(
			linesList,
			app.lineFromElements([]grammars.Element{
				oneElement,
			}),
		)
	}

	return app.tokenFromBlock(
		tokenName,
		app.blockFromlines(linesList),
	)
}

func (app *grammar) tokenFromBlock(name string, block grammars.Block) grammars.Token {
	token, err := app.tokenBuilder.Create().
		WithName(name).
		WithBlock(block).
		Now()

	if err != nil {
		panic(err)
	}

	return token
}

func (app *grammar) blockFromlines(lines []grammars.Line) grammars.Block {
	block, err := app.blockBuilder.Create().
		WithLines(lines).
		Now()

	if err != nil {
		panic(err)
	}

	return block
}

func (app *grammar) lineFromElements(elements []grammars.Element) grammars.Line {
	line, err := app.lineBuilder.Create().
		WithElements(elements).
		Now()

	if err != nil {
		panic(err)
	}

	return line
}

func (app *grammar) elementFromEverything(everything grammars.Everything) grammars.Element {
	ins, err := app.instanceBuilder.Create().
		WithEverything(everything).
		Now()

	if err != nil {
		panic(err)
	}

	cardinality := app.cardinalityOnce()
	element, err := app.elementBuilder.Create().
		WithInstance(ins).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

func (app *grammar) everythingWithoutEscape(name string, exception grammars.Token) grammars.Everything {
	return app.everything(name, exception, nil)
}

func (app *grammar) everything(name string, exception grammars.Token, escape grammars.Token) grammars.Everything {
	builder := app.everythingBuilder.Create().
		WithName(name).
		WithException(exception)

	if escape != nil {
		builder.WithEscape(escape)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) elementFromToken(token grammars.Token, cardinality cardinalities.Cardinality) grammars.Element {
	ins, err := app.instanceBuilder.Create().
		WithToken(token).
		Now()

	if err != nil {
		panic(err)
	}

	element, err := app.elementBuilder.Create().
		WithInstance(ins).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

func (app *grammar) elementFromRecursiveToken(tokenName string, cardinality cardinalities.Cardinality) grammars.Element {
	element, err := app.elementBuilder.Create().
		WithRecursive(tokenName).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return element
}

func (app *grammar) tokenFromValue(name string, value byte) grammars.Token {
	return app.tokenFromBlock(name, app.blockFromlines([]grammars.Line{
		app.lineFromElements([]grammars.Element{
			app.elementFromValue(value),
		}),
	}))
}

func (app *grammar) elementFromValue(value byte) grammars.Element {
	valueIns, err := app.valueBuilder.Create().
		WithName(strconv.Itoa(int(value))).
		WithNumber(value).
		Now()

	if err != nil {
		panic(err)
	}

	ins, err := app.elementBuilder.Create().
		WithValue(valueIns).
		WithCardinality(app.cardinalityOnce()).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) cardinalityOnce() cardinalities.Cardinality {
	max := uint(1)
	return app.cardinality(1, &max)
}

func (app *grammar) cardinality(min uint, pMax *uint) cardinalities.Cardinality {
	builder := app.cardinalityBuilder.Create().WithMin(min)
	if pMax != nil {
		builder.WithMax(*pMax)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *grammar) channelFromValue(name string, value byte) grammars.Channel {
	return app.channelFromToken(
		app.tokenFromBlock(
			name,
			app.blockFromlines([]grammars.Line{
				app.lineFromElements([]grammars.Element{
					app.elementFromValue(value),
				}),
			}),
		),
	)
}

func (app grammar) channelFromToken(token grammars.Token) grammars.Channel {
	ins, err := app.channelBuilder.Create().
		WithToken(token).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
