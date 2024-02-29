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
		app.concreteResource(),
		app.concreteDashboard(),
	})
}

func (app *skeletonFactory) concreteResource() resources.Resource {
	return app.resourceWithChildren(
		"resource",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				resources.NativeBytes,
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"token",
				[]string{"Token"},
				app.kindWithReference([]string{
					"resource",
					"token",
				}),
				"WithToken",
			),
			app.fieldWithBuilder(
				"signature",
				[]string{"Signature", "Bytes"},
				app.kindWithNative(resources.NativeBytes),
				"WithSignature",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.concreteResourceToken(),
		}),
	)
}

func (app *skeletonFactory) concreteResourceToken() resources.Resource {
	return app.resourceWithChildren(
		"token",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				resources.NativeBytes,
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"content",
				[]string{"Content"},
				app.kindWithReference([]string{
					"resource",
					"token",
					"content",
				}),
				"WithContent",
			),
			app.fieldWithBuilder(
				"created_on",
				[]string{"CreatedOn", "Unix"},
				app.kindWithNative(resources.NativeInteger),
				"CreatedOn",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.resourceWithChildren(
				"content",
				app.field(
					"hash",
					[]string{"Hash", "Bytes"},
					app.kindWithNative(
						resources.NativeBytes,
					),
				),
				app.fields([]resources.Field{
					app.fieldWithBuilderAndCondition(
						"dashboard",
						[]string{"Dashboard"},
						app.kindWithReference([]string{
							"resource",
							"token",
							"content",
							"dashboard",
						}),
						"WithDashboard",
						"IsDashboard",
					),
				}),
				"Create",
				"Now",
				app.resources([]resources.Resource{
					app.concreteResourceTokenDashboard(),
				}),
			),
		}),
	)
}

func (app *skeletonFactory) concreteResourceTokenDashboard() resources.Resource {
	return app.resource(
		"dashboard",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				resources.NativeBytes,
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilderAndCondition(
				"dashboard",
				[]string{"Dashboard"},
				app.kindWithReference([]string{
					"dashboard",
				}),
				"WithDashboard",
				"IsDashboard",
			),
			app.fieldWithBuilderAndCondition(
				"widget",
				[]string{"Widget"},
				app.kindWithReference([]string{
					"dashboard",
					"widget",
				}),
				"WithWidget",
				"IsWidget",
			),
			app.fieldWithBuilderAndCondition(
				"viewport",
				[]string{"Viewport"},
				app.kindWithReference([]string{
					"dashboard",
					"widget",
					"viewport",
				}),
				"WithViewport",
				"IsViewport",
			),
		}),
		"Create",
		"Now",
	)
}

func (app *skeletonFactory) concreteDashboard() resources.Resource {
	return app.resourceWithChildren(
		"dashboard",
		app.field(
			"hash",
			[]string{"Hash", "Bytes"},
			app.kindWithNative(
				resources.NativeBytes,
			),
		),
		app.fields([]resources.Field{
			app.fieldWithBuilder(
				"title",
				[]string{"Title"},
				app.kindWithNative(
					resources.NativeString,
				),
				"WithTitle",
			),
			app.fieldWithBuilder(
				"widgets",
				[]string{"Widgets", "List"},
				app.kindWithConnection("dashboard_widgets"),
				"WithWidgets",
			),
		}),
		"Create",
		"Now",
		app.resources([]resources.Resource{
			app.resourceWithChildren(
				"widget",
				app.field(
					"hash",
					[]string{"Hash", "Bytes"},
					app.kindWithNative(
						resources.NativeBytes,
					),
				),
				app.fields([]resources.Field{
					app.fieldWithBuilder(
						"title",
						[]string{"Title"},
						app.kindWithNative(
							resources.NativeString,
						),
						"WithTitle",
					),
					app.fieldWithBuilder(
						"program",
						[]string{"Program", "Bytes"},
						app.kindWithNative(
							resources.NativeBytes,
						),
						"WithProgram",
					),
					app.fieldWithBuilder(
						"input",
						[]string{"Input"},
						app.kindWithNative(
							resources.NativeBytes,
						),
						"WithInput",
					),
					app.fieldWithBuilderAndCondition(
						"viewport",
						[]string{"Viewport"},
						app.kindWithReference([]string{
							"dashboard",
							"widget",
							"viewport",
						}),
						"WithViewport",
						"HasViewport",
					),
				}),
				"Create",
				"Now",
				app.resources([]resources.Resource{
					app.resource(
						"viewport",
						app.field(
							"hash",
							[]string{"Hash", "Bytes"},
							app.kindWithNative(
								resources.NativeBytes,
							),
						),
						app.fields([]resources.Field{
							app.fieldWithBuilder(
								"row",
								[]string{"Row"},
								app.kindWithNative(
									resources.NativeInteger,
								),
								"WithRow",
							),
							app.fieldWithBuilder(
								"height",
								[]string{"Height"},
								app.kindWithNative(
									resources.NativeInteger,
								),
								"WithHeight",
							),
						}),
						"Create",
						"Now",
					),
				}),
			),
		}),
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
	native uint8,
) resources.Kind {
	ins, err := app.kindBuilder.Create().
		WithNative(native).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
