package keys

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys"
	json_encryption "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	json_signatures "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

// NewAdapter creates a new adapter
func NewAdapter() keys.Adapter {
	encryptionAdapter := json_encryption.NewAdapter()
	signatureAdapter := json_signatures.NewAdapter()
	builder := keys.NewBuilder()
	return createAdapter(
		encryptionAdapter.(*json_encryption.Adapter),
		signatureAdapter.(*json_signatures.Adapter),
		builder,
	)
}
