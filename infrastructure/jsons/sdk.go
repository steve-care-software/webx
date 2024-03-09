package jsons

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/operators"
	origins_resources "github.com/steve-care-software/datastencil/domain/libraries/links/origins/resources"
)

// NewLibraryAdapter creates a new library adapter
func NewLibraryAdapter() libraries.Adapter {
	builder := libraries.NewBuilder()
	lnksAdapter := NewLinksAdapter()
	lysAdapter := NewLayersAdapter()
	return createLibraryAdapter(
		builder,
		lnksAdapter.(*linksAdapter),
		lysAdapter.(*layersAdapter),
	)
}

// NewLayersAdapter creates a new layers adapter
func NewLayersAdapter() layers.Adapter {
	builder := layers.NewBuilder()
	pLayerAdapter := NewLayerAdapter()
	return createLayersAdapter(
		builder,
		pLayerAdapter.(*layerAdapter),
	)
}

// NewLayerAdapter creates a new layer adapter
func NewLayerAdapter() layers.LayerAdapter {
	hashAdapter := hash.NewAdapter()
	builder := layers.NewBuilder()
	layerBuilder := layers.NewLayerBuilder()
	outputBuilder := outputs.NewBuilder()
	kindBuilder := kinds.NewBuilder()
	instructionsBuilder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	conditionBuilder := instructions.NewConditionBuilder()
	assignmentBuilder := assignments.NewBuilder()
	assignableBuilder := assignables.NewBuilder()
	constantBuilder := constants.NewBuilder()
	bytesBuilder := bytes.NewBuilder()
	return createLayerAdapter(
		hashAdapter,
		builder,
		layerBuilder,
		outputBuilder,
		kindBuilder,
		instructionsBuilder,
		instructionBuilder,
		conditionBuilder,
		assignmentBuilder,
		assignableBuilder,
		constantBuilder,
		bytesBuilder,
	)
}

// NewLinksAdapter creates a new links adapter
func NewLinksAdapter() links.Adapter {
	builder := links.NewBuilder()
	adapter := NewLinkAdapter()
	return createLinksAdapter(
		builder,
		adapter.(*linkAdapter),
	)
}

// NewLinkAdapter creates a new link adapter
func NewLinkAdapter() links.LinkAdapter {
	hashAdapter := hash.NewAdapter()
	builder := links.NewBuilder()
	linkBuilder := links.NewLinkBuilder()
	elementsBuilder := elements.NewBuilder()
	elementBuilder := elements.NewElementBuilder()
	conditionBuilder := conditions.NewBuilder()
	conditionValueBuilder := conditions.NewConditionValueBuilder()
	conditionResourceBuilder := resources.NewBuilder()
	originBuilder := origins.NewBuilder()
	originValueBuilder := origins.NewValueBuilder()
	originResourceBuilder := origins_resources.NewBuilder()
	operatorBuilder := operators.NewBuilder()
	return createLinkAdapter(
		hashAdapter,
		builder,
		linkBuilder,
		elementsBuilder,
		elementBuilder,
		conditionBuilder,
		conditionValueBuilder,
		conditionResourceBuilder,
		originBuilder,
		originValueBuilder,
		originResourceBuilder,
		operatorBuilder,
	)
}
