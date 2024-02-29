package resources

const (
	// NativeInteger represents the integer
	NativeInteger (uint8) = iota

	// NativeFloat represents the float
	NativeFloat

	// NativeString represents the string
	NativeString

	// NativeBytes represents the bytes
	NativeBytes
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewResourceBuilder creates a new resource builder
func NewResourceBuilder() ResourceBuilder {
	return createResourceBuilder()
}

// NewFieldsBuilder creates a new fields builder
func NewFieldsBuilder() FieldsBuilder {
	return createFieldsBuilder()
}

// NewFieldBuilder creates a new field builder
func NewFieldBuilder() FieldBuilder {
	return createFieldBuilder()
}

// NewKindBuilder creates a new kind builder
func NewKindBuilder() KindBuilder {
	return createKindBuilder()
}

// Builder represents a resources builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents resources
type Resources interface {
	List() []Resource
	FetchByName(name string) (Resource, error)
	FetchByPath(path []string) (Resource, error)
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithName(name string) ResourceBuilder
	WithKey(key Field) ResourceBuilder
	WithFields(fields Fields) ResourceBuilder
	WithInitialize(initialize string) ResourceBuilder
	WithTrigger(trigger string) ResourceBuilder
	WithChildren(children Resources) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a schema resource
type Resource interface {
	Name() string
	Key() Field
	Fields() Fields
	Initialize() string
	Trigger() string
	HasChildren() bool
	Children() Resources
}

// FieldsBuilder represents a fields builder
type FieldsBuilder interface {
	Create() FieldsBuilder
	WithList(list []Field) FieldsBuilder
	Now() (Fields, error)
}

// Fields represents fields
type Fields interface {
	List() []Field
	Names() []string
}

// FieldBuilder represents a field builder
type FieldBuilder interface {
	Create() FieldBuilder
	WithName(name string) FieldBuilder
	WithRetriever(retriever []string) FieldBuilder
	WithBuilder(builder string) FieldBuilder
	WithCondition(condition string) FieldBuilder
	WithKind(kind Kind) FieldBuilder
	Now() (Field, error)
}

// Field represents a field
type Field interface {
	Name() string
	Retriever() []string
	Kind() Kind
	HasCondition() bool
	Condition() string
	HasBuilder() bool
	Builder() string
}

// KindBuilder represents a kind builder
type KindBuilder interface {
	Create() KindBuilder
	WithNative(native uint8) KindBuilder
	WithReference(reference []string) KindBuilder
	WithConnection(connection string) KindBuilder
	Now() (Kind, error)
}

// Kind represents a kind
type Kind interface {
	IsNative() bool
	Native() *uint8
	IsReference() bool
	Reference() []string
	IsConnection() bool
	Connection() string
}
