package commits

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits"
	json_actions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions"
)

// NewAdapter creates a new adapter
func NewAdapter(
	instanceAdapter instances.Adapter,
) commits.Adapter {
	actionsAdapter := json_actions.NewAdapter(
		instanceAdapter,
	)

	signatureAdapter := signers.NewSignatureAdapter()
	contentBuilder := commits.NewContentBuilder()
	builder := commits.NewBuilder()
	return createAdapter(
		actionsAdapter.(*json_actions.Adapter),
		signatureAdapter,
		contentBuilder,
		builder,
	)
}
