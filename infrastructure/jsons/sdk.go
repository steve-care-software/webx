package jsons

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
)

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
	elementsBuilder := links.NewElementsBuilder()
	elementBuilder := links.NewElementBuilder()
	conditionBuilder := links.NewConditionBuilder()
	conditionValueBuilder := links.NewConditionValueBuilder()
	conditionResourceBuilder := links.NewConditionResourceBuilder()
	originBuilder := links.NewOriginBuilder()
	originValueBuilder := links.NewOriginValueBuilder()
	originResourceBuilder := links.NewOriginResourceBuilder()
	operatorBuilder := links.NewOperatorBuilder()
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
