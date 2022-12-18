package programs

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/programs/programs/modules"
)

type application struct {
	hash        hash.Hash
	index       uint
	module      modules.Module
	attachments Attachments
}

func createApplication(
	hash hash.Hash,
	index uint,
	module modules.Module,
) Application {
	return createApplicationInternally(hash, index, module, nil)
}

func createApplicationWithAttachments(
	hash hash.Hash,
	index uint,
	module modules.Module,
	attachments Attachments,
) Application {
	return createApplicationInternally(hash, index, module, attachments)
}

func createApplicationInternally(
	hash hash.Hash,
	index uint,
	module modules.Module,
	attachments Attachments,
) Application {
	out := application{
		hash:        hash,
		index:       index,
		module:      module,
		attachments: attachments,
	}

	return &out
}

// Hash returns the hash
func (obj *application) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *application) Index() uint {
	return obj.index
}

// Module returns the module
func (obj *application) Module() modules.Module {
	return obj.module
}

// HasAttachments returns true if there is attachments, false otherwise
func (obj *application) HasAttachments() bool {
	return obj.attachments != nil
}

// Attachments returns the attachments, if any
func (obj *application) Attachments() Attachments {
	return obj.attachments
}
