package assignables

import (
	json_accounts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts"
	json_bytes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/databases"
)

// Assignable represents the assignable
type Assignable struct {
	Bytes        *json_bytes.Bytes               `json:"bytes"`
	Constant     *json_constants.Constant        `json:"constant"`
	Account      *json_accounts.Account          `json:"account"`
	Cryptography *json_cryptography.Cryptography `json:"cryptography"`
	Query        string                          `json:"query"`
	Compiler     *json_compiler.Compiler         `json:"compiler"`
	Database     *json_databases.Database        `json:"database"`
}
