package signatures

import (
	json_signs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	json_votes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
)

// Signature represents the signature
type Signature struct {
	IsGeneratePrivateKey bool             `json:"generate_pk"`
	FetchPublicKey       string           `json:"fetch_pubkey"`
	Sign                 *json_signs.Sign `json:"sign"`
	Vote                 *json_votes.Vote `json:"vote"`
}
