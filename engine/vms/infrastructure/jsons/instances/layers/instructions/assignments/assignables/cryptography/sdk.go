package cryptography

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography"
	json_decrypts "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	json_encrypts "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	json_keys "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys"
)

// NewAdapter creates a new adapter
func NewAdapter() cryptography.Adapter {
	encryptAdapter := json_encrypts.NewAdapter()
	decryptAdapter := json_decrypts.NewAdapter()
	keyAdapter := json_keys.NewAdapter()
	builder := cryptography.NewBuilder()
	return createAdapter(
		encryptAdapter.(*json_encrypts.Adapter),
		decryptAdapter.(*json_decrypts.Adapter),
		keyAdapter.(*json_keys.Adapter),
		builder,
	)
}
