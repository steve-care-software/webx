package sqllites

import (
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons/resources"
)

type skeletonFactory struct {
	builder                skeletons.Builder
	resourcesBuilder       resources.Builder
	resourceBuilder        resources.ResourceBuilder
	fieldsBuilder          resources.FieldsBuilder
	fieldBuilder           resources.FieldBuilder
	kindBuilder            resources.KindBuilder
	nativeBuilder          resources.NativeBuilder
	listBuilder            resources.ListBuilder
	connectionsBuilder     connections.Builder
	connectionBuilder      connections.ConnectionBuilder
	connectionFieldBuilder connections.FieldBuilder
}

func createSkeletonFactory(
	builder skeletons.Builder,
	resourcesBuilder resources.Builder,
	resourceBuilder resources.ResourceBuilder,
	fieldsBuilder resources.FieldsBuilder,
	fieldBuilder resources.FieldBuilder,
	kindBuilder resources.KindBuilder,
	nativeBuilder resources.NativeBuilder,
	listBuilder resources.ListBuilder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	connectionFieldBuilder connections.FieldBuilder,
) skeletons.Factory {
	out := skeletonFactory{
		builder:                builder,
		resourcesBuilder:       resourcesBuilder,
		resourceBuilder:        resourceBuilder,
		fieldsBuilder:          fieldsBuilder,
		fieldBuilder:           fieldBuilder,
		kindBuilder:            kindBuilder,
		nativeBuilder:          nativeBuilder,
		listBuilder:            listBuilder,
		connectionsBuilder:     connectionsBuilder,
		connectionBuilder:      connectionBuilder,
		connectionFieldBuilder: connectionFieldBuilder,
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"layers",
				app.kindWithConnection("library_layers"),
			),
			app.fieldWithCanBeNil(
				"links",
				app.kindWithConnection("library_links"),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"origin",
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
				}),
			),
			app.field(
				"elements",
				app.kindWithConnection("link_elements"),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"layer",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeBytes,
					),
				),
			),
			app.fieldWithCanBeNil(
				"condition",
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
				}),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"resource",
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
					"resource",
				}),
			),
			app.fieldWithCanBeNil(
				"next",
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
					"value",
				}),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithCanBeNil(
				"resource",
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
					"resource",
				}),
			),
			app.fieldWithCanBeNil(
				"condition",
				app.kindWithReference([]string{
					"library",
					"link",
					"element",
					"condition",
				}),
			),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkElementConditionResource() resources.Resource {
	return app.resource(
		"resource",
		app.field(
			"hash",
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"code",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
			app.field(
				"is_raised_in_layer",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkOrigin() resources.Resource {
	return app.resourceWithChildren(
		"origin",
		app.field(
			"hash",
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"resource",
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"resource",
				}),
			),
			app.field(
				"operator",
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"operator",
				}),
			),
			app.field(
				"next",
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"value",
				}),
			),
		}),
		app.resources([]resources.Resource{
			app.concreteLibraryLinkResource(),
			app.concreteLibraryLinkOriginOperator(),
			app.concreteLibraryLinkValue(),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkResource() resources.Resource {
	return app.resource(
		"resource",
		app.field(
			"hash",
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"layer",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeBytes,
					),
				),
			),
			app.field(
				"is_mandatory",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkOriginOperator() resources.Resource {
	return app.resource(
		"operator",
		app.field(
			"hash",
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"is_and",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
			app.field(
				"is_or",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
			app.field(
				"is_xor",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLinkValue() resources.Resource {
	return app.resource(
		"value",
		app.field(
			"hash",
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithCanBeNil(
				"resource",
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
					"resource",
				}),
			),
			app.fieldWithCanBeNil(
				"origin",
				app.kindWithReference([]string{
					"library",
					"link",
					"origin",
				}),
			),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLayer() resources.Resource {
	return app.resourceWithChildren(
		"layer",
		app.field(
			"hash",
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"instructions",
				app.kindWithConnection("layer_instructions"),
			),
			app.field(
				"output",
				app.kindWithReference([]string{
					"library",
					"layer",
					"output",
				}),
			),
			app.field(
				"input",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"variable",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
			),
			app.field(
				"kind",
				app.kindWithReference([]string{
					"library",
					"layer",
					"output",
					"kind",
				}),
			),
			app.fieldWithCanBeNil(
				"execute",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"prompt",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
			app.field(
				"continue",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeInteger,
					),
				),
			),
		}),
	)
}

func (app *skeletonFactory) concreteLibraryLayerInstruction() resources.Resource {
	return app.resourceWithChildren(
		"instruction",
		app.field(
			"hash",
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"assignment",
				app.kindWithReference([]string{
					"library",
					"layer",
					"instruction",
					"assignment",
				}),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.field(
				"name",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeString,
					),
				),
			),
			app.field(
				"assignable",
				app.kindWithReference([]string{
					"library",
					"layer",
					"instruction",
					"assignment",
					"assignable",
				}),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithCanBeNil(
				"bytes",
				app.kindWithReference([]string{
					"library",
					"layer",
					"instruction",
					"assignment",
					"assignable",
					"bytes",
				}),
			),
		}),
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
			app.kindWithNative(
				app.nativeWithSingle(
					resources.NativeBytes,
				),
			),
		),
		app.fields([]resources.Field{
			app.fieldWithCanBeNil(
				"joins",
				app.kindWithNative(
					app.nativeWithList(
						app.list(resources.NativeString, "_"),
					),
				),
			),
			app.fieldWithCanBeNil(
				"compares",
				app.kindWithNative(
					app.nativeWithList(
						app.list(resources.NativeString, "_"),
					),
				),
			),
			app.fieldWithCanBeNil(
				"hash_bytes",
				app.kindWithNative(
					app.nativeWithSingle(
						resources.NativeBytes,
					),
				),
			),
		}),
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
) resources.Resource {
	ins, err := app.resourceBuilder.Create().
		WithName(name).
		WithKey(key).
		WithFields(fields).
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
	children resources.Resources,
) resources.Resource {
	ins, err := app.resourceBuilder.Create().
		WithName(name).
		WithKey(key).
		WithFields(fields).
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
	kind resources.Kind,
) resources.Field {
	ins, err := app.fieldBuilder.Create().
		WithName(name).
		WithKind(kind).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

func (app *skeletonFactory) fieldWithCanBeNil(
	name string,
	kind resources.Kind,
) resources.Field {
	ins, err := app.fieldBuilder.Create().
		WithName(name).
		WithKind(kind).
		CanBeNil().
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
