package namespaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces"
)

const noLoadedNamespaceErr = "the current user does not have any loaded namespace"

// Application represents a namespace application
type Application interface {
	Loaded(input namespaces.Namespace) ([]string, error)
	Create(
		input namespaces.Namespace,
		name string,
		symbol string,
		description string,
		baseDifficulty uint64,
		increasePerSize uint64,
		sizeBlock uint64,
		totalUnitsAmount uint64,
		purseAmount uint64,
	) (namespaces.Namespace, error)
	Set(input namespaces.Namespace, name string) (namespaces.Namespace, error)
}
