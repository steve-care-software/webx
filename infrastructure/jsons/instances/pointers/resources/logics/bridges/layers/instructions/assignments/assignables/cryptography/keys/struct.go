package keys

import (
	json_encryption "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	json_signatures "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

// Key represents the key
type Key struct {
	Signature  *json_signatures.Signature  `json:"signature"`
	Encryption *json_encryption.Encryption `json:"encryption"`
}
