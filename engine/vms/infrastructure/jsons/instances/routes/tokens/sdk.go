package tokens

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
	json_elements "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/elements"
	json_omissions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/omissions"
	json_cardinalities "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/tokens/cardinalities"
)

// NewAdapter creates a new adapter
func NewAdapter() tokens.Adapter {
	elementAdapter := json_elements.NewAdapter()
	omissionAdapter := json_omissions.NewAdapter()
	cardinalityAdapter := json_cardinalities.NewAdapter()
	builder := tokens.NewBuilder()
	tokenBuilder := tokens.NewTokenBuilder()
	return createAdapter(
		elementAdapter.(*json_elements.Adapter),
		omissionAdapter.(*json_omissions.Adapter),
		cardinalityAdapter.(*json_cardinalities.Adapter),
		builder,
		tokenBuilder,
	)
}
