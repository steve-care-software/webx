package tmp

import (
	"github.com/steve-care-software/webx/engine/bytes/applications"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

const contextIdentifierUndefinedPattern = "the context identifier (%d) does not exists"

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder(
	namespaceAdapter namespaces.Adapter,
) applications.Builder {
	namespacesBuilder := namespaces.NewBuilder()
	namespaceBuilder := namespaces.NewNamespaceBuilder()
	hashAdapter := hash.NewAdapter()
	return createApplicationBuilder(
		namespaceAdapter,
		namespacesBuilder,
		namespaceBuilder,
		hashAdapter,
	)
}
