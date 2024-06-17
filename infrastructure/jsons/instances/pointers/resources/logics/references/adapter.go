package references

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/references"
)

// Adapter represents a references adapter
type Adapter struct {
	builder          references.Builder
	referenceBuilder references.ReferenceBuilder
	instanceAdapter  instances.Adapter
}

func createAdapter(
	builder references.Builder,
	referenceBuilder references.ReferenceBuilder,
	instanceAdapter instances.Adapter,
) references.Adapter {
	out := Adapter{
		builder:          builder,
		referenceBuilder: referenceBuilder,
		instanceAdapter:  instanceAdapter,
	}

	return &out
}

// ToBytes converts references to bytes
func (app *Adapter) ToBytes(ins references.References) ([]byte, error) {
	return nil, nil
}

// ToInstance converts bytes to references
func (app *Adapter) ToInstance(bytes []byte) (references.References, error) {
	return nil, nil
}
