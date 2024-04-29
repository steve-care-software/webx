package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions"
	json_decrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	json_encrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
)

// NewAdapter creates a new adapter
func NewAdapter() encryptions.Adapter {
	decryptAdapter := json_decrypts.NewAdapter()
	encryptAdapter := json_encrypts.NewAdapter()
	builder := encryptions.NewBuilder()
	return createAdapter(
		decryptAdapter.(*json_decrypts.Adapter),
		encryptAdapter.(*json_encrypts.Adapter),
		builder,
	)
}
