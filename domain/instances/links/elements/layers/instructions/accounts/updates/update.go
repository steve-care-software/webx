package updates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/accounts/updates/criterias"
)

type update struct {
	hash        hash.Hash
	credentials string
	criteria    criterias.Criteria
}

func createUpdate(
	hash hash.Hash,
	credentials string,
	criteria criterias.Criteria,
) Update {
	out := update{
		hash:        hash,
		credentials: credentials,
		criteria:    criteria,
	}

	return &out
}

// Hash returns the hash
func (obj *update) Hash() hash.Hash {
	return obj.hash
}

// Credentials returns the credentials
func (obj *update) Credentials() string {
	return obj.credentials
}

// Criteria returns the criteria
func (obj *update) Criteria() criterias.Criteria {
	return obj.criteria
}
