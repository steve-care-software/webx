package accounts

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts"
	json_communications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	json_credentials "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/credentials"
	json_encryptions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/encryptions"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/retrieves"
)

// NewAdapter creates a new adapter
func NewAdapter() accounts.Adapter {
	communicationAdapter := json_communications.NewAdapter()
	credentialsAdapter := json_credentials.NewAdapter()
	encryptionAdapter := json_encryptions.NewAdapter()
	retrieveAdapter := json_retrieves.NewAdapter()
	builder := accounts.NewBuilder()
	return createAdapter(
		communicationAdapter.(*json_communications.Adapter),
		credentialsAdapter.(*json_credentials.Adapter),
		encryptionAdapter.(*json_encryptions.Adapter),
		retrieveAdapter.(*json_retrieves.Adapter),
		builder,
	)
}
