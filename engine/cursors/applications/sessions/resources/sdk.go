package resources

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
)

const noLoadedResourceErr = "There is no loaded resource"
const noSelectedResourceErr = "There is no selected resource and therefore the requested action is impossible"
const cannotAlterNeverCommittedErr = "The current resource cannot be altered because it has never been comitted"

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDatabase(dbApp databases.Application) Builder
	Now() (Application, error)
}

// Application represents a resource application
type Application interface {
	Insert(input resources.Resource, data []byte, blacklist []hash.Hash, whitelist []hash.Hash) (resources.Resource, error)
	Load(input resources.Resource, delimiterIndex uint64) (resources.Resource, error)
	Select(input resources.Resource, delimiterIndex uint64) (resources.Resource, error)
	Delete(input resources.Resource, vote signers.Vote) (resources.Resource, error)
	Retrieve(input resources.Resource) (singles.Single, error)
	Update(
		input resources.Resource,
		addToBlacklist []hash.Hash,
		removeFromBlacklist []hash.Hash,
		addToWhitelist []hash.Hash,
		removeFromWhitelist []hash.Hash,
	) (resources.Resource, error)
	Commit(input resources.Resource) error
}
