package sqllites

import (
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
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
		"assignments",
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
				"WithName",
			),
			app.fieldWithBuilder(
				"assignable",
				[]string{"Assignable"},
				app.kindWithReference([]string{
					"assignments",
					"assignables",
				}),
				"WithAssignable",
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
		"assignables",
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
					"assignments",
					"assignables",
					"bytes",
				}),
				"WithBytes",
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
				"WithJoin",
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
				"WithCompare",
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
				"WithHashBytes",
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
			"dashboard_widgets",
			app.connectionField(
				"dashboard",
				[]string{
					"dashboard",
				},
			),
			app.connectionField(
				"widget",
				[]string{
					"dashboard",
					"widget",
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
	builderMethod string,
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
	builderMethod string,
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
