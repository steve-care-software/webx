package sqllites

import (
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

type skeletonFactory struct {
	builder                   skeletons.Builder
	resourcesBuilder          resources.Builder
	resourceBuilder           resources.ResourceBuilder
	fieldsBuilder             resources.FieldsBuilder
	fieldBuilder              resources.FieldBuilder
	builderInstructionBuilder resources.BuilderInstructionBuilder
	kindBuilder               resources.KindBuilder
	nativeBuilder             resources.NativeBuilder
	listBuilder               resources.ListBuilder
	connectionsBuilder        connections.Builder
	connectionBuilder         connections.ConnectionBuilder
	connectionFieldBuilder    connections.FieldBuilder
}

func createSkeletonFactory(
	builder skeletons.Builder,
	resourcesBuilder resources.Builder,
	resourceBuilder resources.ResourceBuilder,
	fieldsBuilder resources.FieldsBuilder,
	fieldBuilder resources.FieldBuilder,
	builderInstructionBuilder resources.BuilderInstructionBuilder,
	kindBuilder resources.KindBuilder,
	nativeBuilder resources.NativeBuilder,
	listBuilder resources.ListBuilder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	connectionFieldBuilder connections.FieldBuilder,
) skeletons.Factory {
	out := skeletonFactory{
		builder:                   builder,
		resourcesBuilder:          resourcesBuilder,
		resourceBuilder:           resourceBuilder,
		fieldsBuilder:             fieldsBuilder,
		fieldBuilder:              fieldBuilder,
		builderInstructionBuilder: builderInstructionBuilder,
		kindBuilder:               kindBuilder,
		nativeBuilder:             nativeBuilder,
		listBuilder:               listBuilder,
		connectionsBuilder:        connectionsBuilder,
		connectionBuilder:         connectionBuilder,
		connectionFieldBuilder:    connectionFieldBuilder,
	}

	return &out
}

// Create creates a new Skeleton instance
func (app *skeletonFactory) Create() (skeletons.Skeleton, error) {
	resources := app.concreteResources()
	connections := app.concreteConnections()
	return app.builder.Create().
		WithResources(resources).
		WithConnections(connections).
		Now()
}

func (app *skeletonFactory) concreteResources() resources.Resources {
	return app.resources([]resources.Resource{
		app.concreteLibrary(),
	})
}

func (app *skeletonFactory) concreteLibrary() resources.Resource {
	return app.resourceWithChildren(
		"library",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"layers",
				[]string{"Layers"},
				app.kindWithConnection("library_layers"),
				app.builderInstructionWithContainsParams(
					"WithLayers",
				),
			),
			app.fieldWithBuilderAndCondition(
				"links",
				[]string{"Links"},
				app.kindWithConnection("library_links"),
				app.builderInstructionWithContainsParams(
					"WithLinks",
				),
				"HasLinks",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLink(),
			app.concreteLibraryLayer(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLink() resources.Resource {
	return app.resourceWithChildren(
		"link",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"origin",
				[]string{"Origin"},
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
				}),
				app.builderInstructionWithContainsParams(
					"WithOrigin",
				),
			),
			app.fieldWithBuilder(
				"elements",
				[]string{"Elements"},
				app.kindWithConnection("link_elements"),
				app.builderInstructionWithContainsParams(
					"WithElements",
				),
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLinkElement(),
			app.concreteLibraryLinkOrigin(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkElement() resources.Resource {
	return app.resourceWithChildren(
		"element",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"layer",
				[]string{"Hash", "Bytes"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeBytes,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithLayerBytes",
				),
			),
			app.fieldWithBuilderAndCondition(
				"condition",
				[]string{"Condition"},
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
				}),
				app.builderInstructionWithContainsParams(
					"WithCondition",
				),
				"HasCondition",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLinkElementCondition(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkElementCondition() resources.Resource {
	return app.resourceWithChildren(
		"condition",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"resource",
				[]string{"Resource"},
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
					"resource",
				}),
				app.builderInstructionWithContainsParams(
					"WithResource",
				),
			),
			app.fieldWithBuilderAndCondition(
				"next",
				[]string{"Next"},
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
					"value",
				}),
				app.builderInstructionWithContainsParams(
					"WithNext",
				),
				"HasNext",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLinkElementConditionValue(),
			app.concreteLibraryLinkElementConditionResource(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkElementConditionValue() resources.Resource {
	return app.resource(
		"value",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilderAndCondition(
				"resource",
				[]string{"Resource"},
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
					"resource",
				}),
				app.builderInstructionWithContainsParams(
					"WithResource",
				),
				"IsResource",
			),
			app.fieldWithBuilderAndCondition(
				"condition",
				[]string{"Condition"},
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
				}),
				app.builderInstructionWithContainsParams(
					"WithCondition",
				),
				"IsCondition",
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteLibraryLinkElementConditionResource() resources.Resource {
	return app.resource(
		"resource",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"code",
				[]string{"Code"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithCode",
				),
			),
			app.fieldWithBuilder(
				"is_raised_in_layer",
				[]string{"IsRaisedInLayer"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstruction(
					"IsRaisedInLayer",
				),
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteLibraryLinkOrigin() resources.Resource {
	return app.resourceWithChildren(
		"origin",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"resource",
				[]string{"Resource"},
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"resource",
				}),
				app.builderInstructionWithContainsParams(
					"WithResource",
				),
			),
			app.fieldWithBuilder(
				"operator",
				[]string{"Operator"},
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"operator",
				}),
				app.builderInstructionWithContainsParams(
					"WithOperator",
				),
			),
			app.fieldWithBuilder(
				"next",
				[]string{"Next"},
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"value",
				}),
				app.builderInstructionWithContainsParams(
					"WithNext",
				),
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLinkOriginResource(),
			app.concreteLibraryLinkOriginOperator(),
			app.concreteLibraryLinkOriginValue(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkOriginResource() resources.Resource {
	return app.resource(
		"resource",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"layer",
				[]string{"Hash", "Bytes"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeBytes,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithLayerBytes",
				),
			),
			app.fieldWithBuilder(
				"is_mandatory",
				[]string{"IsMandatory"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstruction(
					"IsMandatory",
				),
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteLibraryLinkOriginOperator() resources.Resource {
	return app.resource(
		"operator",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"is_and",
				[]string{"IsAnd"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstruction(
					"IsAnd",
				),
			),
			app.fieldWithBuilder(
				"is_or",
				[]string{"IsOr"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstruction(
					"IsOr",
				),
			),
			app.fieldWithBuilder(
				"is_xor",
				[]string{"IsXor"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstruction(
					"IsXor",
				),
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteLibraryLinkOriginValue() resources.Resource {
	return app.resource(
		"value",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilderAndCondition(
				"resource",
				[]string{"Resource"},
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"resource",
				}),
				app.builderInstructionWithContainsParams(
					"WithResource",
				),
				"IsResource",
			),
			app.fieldWithBuilderAndCondition(
				"origin",
				[]string{"Origin"},
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
				}),
				app.builderInstructionWithContainsParams(
					"WithOrigin",
				),
				"IsOrigin",
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteLibraryLayer() resources.Resource {
	return app.resourceWithChildren(
		"layer",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"instructions",
				[]string{"Instructions"},
				app.kindWithConnection("layer_instructions"),
				app.builderInstructionWithContainsParams(
					"WithInstructions",
				),
			),
			app.fieldWithBuilder(
				"output",
				[]string{"Output"},
				app.kindWithReference([]string{
					"library",
					"layer",
					"output",
				}),
				app.builderInstructionWithContainsParams(
					"WithOutput",
				),
			),
			app.fieldWithBuilder(
				"input",
				[]string{"Input"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithInput",
				),
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLayerOutput(),
			app.concreteLibraryLayerInstruction(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLayerOutput() resources.Resource {
	return app.resourceWithChildren(
		"output",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"variable",
				[]string{"Variable"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithVariable",
				),
			),
			app.fieldWithBuilder(
				"kind",
				[]string{"Kind"},
				app.kindWithReference([]string{
					"library",
					"layer",
					"output",
					"kind",
				}),
				app.builderInstructionWithContainsParams(
					"WithKind",
				),
			),
			app.fieldWithBuilderAndCondition(
				"execute",
				[]string{"Execute"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithExecute",
				),
				"HasExecute",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLayerKind(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLayerKind() resources.Resource {
	return app.resource(
		"kind",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"prompt",
				[]string{"IsPrompt"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstruction(
					"IsPrompt",
				),
			),
			app.fieldWithBuilder(
				"continue",
				[]string{"IsContinue"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
				app.builderInstruction(
					"IsContinue",
				),
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteLibraryLayerInstruction() resources.Resource {
	return app.resourceWithChildren(
		"instruction",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"assignment",
				[]string{"Assignment"},
				app.kindWithReference([]string{
					"library",
					"layer",
					"instruction",
					"assignment",
				}),
				app.builderInstructionWithContainsParams(
					"WithAssignment",
				),
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteLibraryLayerAssignment(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLayerAssignment() resources.Resource {
	return app.resourceWithChildren(
		"assignment",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"name",
				[]string{"Name"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithName",
				),
			),
			app.fieldWithBuilder(
				"assignable",
				[]string{"Assignable"},
				app.kindWithReference([]string{
					"library",
					"layer",
					"instruction",
					"assignment",
					"assignable",
				}),
				app.builderInstructionWithContainsParams(
					"WithAssignable",
				),
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.createLibraryLayerAssignable(),
		}),
	)
}

func (app *skeletonFactory) createLibraryLayerAssignable() resources.Resource {
	return app.resourceWithChildren(
		"assignable",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilderAndCondition(
				"bytes",
				[]string{"Bytes"},
				app.kindWithReference([]string{
					"library",
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				}),
				app.builderInstructionWithContainsParams(
					"WithBytes",
				),
				"IsBytes",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.createLibraryLayerBytes(),
		}),
	)
}

func (app *skeletonFactory) createLibraryLayerBytes() resources.Resource {
	return app.resource(
		"bytes",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilderAndCondition(
				"joins",
				[]string{"Join"},
				app.kindWithNative(
					app.nativeWithList(
						app.list(resources.NativeString, "_"),
					),
				),
				app.builderInstructionWithContainsParams(
					"WithJoin",
				),
				"IsJoin",
			),
			app.fieldWithBuilderAndCondition(
				"compares",
				[]string{"Compare"},
				app.kindWithNative(
					app.nativeWithList(
						app.list(resources.NativeString, "_"),
					),
				),
				app.builderInstructionWithContainsParams(
					"WithCompare",
				),
				"IsCompare",
			),
			app.fieldWithBuilderAndCondition(
				"hash_bytes",
				[]string{"HashBytes"},
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeBytes,
					),
				),
				app.builderInstructionWithContainsParams(
					"WithHashBytes",
				),
				"IsHashBytes",
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteConnections() connections.Connections {
	return app.connections([]connections.Connection{
		app.connection(
			"library_layers",
			app.connectionField(
				"library",
				[]string{
					"library",
				},
			),
			app.connectionField(
				"layer",
				[]string{
					"library",
					"layer",
				},
			),
		),
		app.connection(
			"layer_instructions",
			app.connectionField(
				"layer",
				[]string{
					"library",
					"layer",
				},
			),
			app.connectionField(
				"instruction",
				[]string{
					"library",
					"layer",
					"instruction",
				},
			),
		),
		app.connection(
			"library_links",
			app.connectionField(
				"library",
				[]string{
					"library",
				},
			),
			app.connectionField(
				"link",
				[]string{
					"library",
					"link",
				},
			),
		),
		app.connection(
			"link_elements",
			app.connectionField(
				"link",
				[]string{
					"library",
					"link",
				},
			),
			app.connectionField(
				"elements",
				[]string{
					"library",
					"link",
					"element",
				},
			),
		),
	})
}

func (app *skeletonFactory) connections(
	list []connections.Connection,
) connections.Connections {
	ins, err := app.connectionsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) connection(
	name string,
	from connections.Field,
	to connections.Field,
) connections.Connection {
	ins, err := app.connectionBuilder.Create().
		WithName(name).
		From(from).
		To(to).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) connectionField(
	name string,
	path []string,
) connections.Field {
	ins, err := app.connectionFieldBuilder.Create().
		WithName(name).
		WithPath(path).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) resources(
	list []resources.Resource,
) resources.Resources {
	ins, err := app.resourcesBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) resource(
	name string,
	key resources.Field,
	fields resources.Fields,
	initialize string,
	trigger string,
) resources.Resource {
	ins, err := app.resourceBuilder.Create().
		WithName(name).
		WithKey(key).
		WithFields(fields).
		WithInitialize(initialize).
		WithTrigger(trigger).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) resourceWithChildren(
	name string,
	key resources.Field,
	fields resources.Fields,
	initialize string,
	trigger string,
	children resources.Resources,
) resources.Resource {
	ins, err := app.resourceBuilder.Create().
		WithName(name).
		WithKey(key).
		WithFields(fields).
		WithInitialize(initialize).
		WithTrigger(trigger).
		WithChildren(children).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) fields(
	list []resources.Field,
) resources.Fields {
	ins, err := app.fieldsBuilder.Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) field(
	name string,
	retriever []string,
	kind resources.Kind,
) resources.Field {
	ins, err := app.fieldBuilder.Create().
		WithName(name).
		WithRetriever(retriever).
		WithKind(kind).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) fieldWithBuilder(
	name string,
	retriever []string,
	kind resources.Kind,
	builderMethod resources.BuilderInstruction,
) resources.Field {
	ins, err := app.fieldBuilder.Create().
		WithName(name).
		WithRetriever(retriever).
		WithBuilder(builderMethod).
		WithKind(kind).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) fieldWithCondition(
	name string,
	retriever []string,
	kind resources.Kind,
	condition string,
) resources.Field {
	ins, err := app.fieldBuilder.Create().
		WithName(name).
		WithRetriever(retriever).
		WithCondition(condition).
		WithKind(kind).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) fieldWithBuilderAndCondition(
	name string,
	retriever []string,
	kind resources.Kind,
	builderMethod resources.BuilderInstruction,
	condition string,
) resources.Field {
	ins, err := app.fieldBuilder.Create().
		WithName(name).
		WithRetriever(retriever).
		WithBuilder(builderMethod).
		WithCondition(condition).
		WithKind(kind).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) builderInstruction(
	method string,
) resources.BuilderInstruction {
	ins, err := app.builderInstructionBuilder.Create().
		WithMethod(method).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) builderInstructionWithContainsParams(
	method string,
) resources.BuilderInstruction {
	ins, err := app.builderInstructionBuilder.Create().
		WithMethod(method).
		ContainsParam().
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) kindWithConnection(
	name string,
) resources.Kind {
	ins, err := app.kindBuilder.Create().
		WithConnection(name).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) kindWithReference(
	reference []string,
) resources.Kind {
	ins, err := app.kindBuilder.Create().
		WithReference(reference).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) kindWithNative(
	native resources.Native,
) resources.Kind {
	ins, err := app.kindBuilder.Create().
		WithNative(native).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) nativeWithSingle(
	single uint8,
) resources.Native {
	ins, err := app.nativeBuilder.Create().
		WithSingle(single).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) nativeWithList(
	list resources.List,
) resources.Native {
	ins, err := app.nativeBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) list(
	value uint8,
	delimiter string,
) resources.List {
	ins, err := app.listBuilder.Create().
		WithValue(value).
		WithDelimiter(delimiter).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
