package applications

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type application struct {
	hash        hash.Hash
	index       uint
	module      uint
	attachments Attachments
}

func createApplication(
	hash hash.Hash,
	index uint,
	module uint,
) Application {
	return createApplicationInternally(hash, index, module, nil)
}

func createApplicationWithAttachments(
	hash hash.Hash,
	index uint,
	module uint,
	attachments Attachments,
) Application {
	return createApplicationInternally(hash, index, module, attachments)
}

func createApplicationInternally(
	hash hash.Hash,
	index uint,
	module uint,
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
func (obj *application) Module() uint {
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
