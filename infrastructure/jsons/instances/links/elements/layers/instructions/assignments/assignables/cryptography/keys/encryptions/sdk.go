package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	json_decrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	json_encrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// NewAdapter creates a new adapter
func NewAdapter() encryptions.Adapter {
	encryptAdapter := json_encrypts.NewAdapter()
	decryptAdapter := json_decrypts.NewAdapter()
	builder := encryptions.NewBuilder()
	return createAdapter(
		encryptAdapter.(*json_encrypts.Adapter),
		decryptAdapter.(*json_decrypts.Adapter),
		builder,
	)
}
