package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type application struct {
	hash        hash.Hash
	module      uint
	attachments Attachments
}

func createApplication(
	hash hash.Hash,
	module uint,
) Application {
	return createApplicationInternally(hash, module, nil)
}

func createApplicationWithAttachments(
	hash hash.Hash,
	module uint,
	attachments Attachments,
) Application {
	return createApplicationInternally(hash, module, attachments)
}

func createApplicationInternally(
	hash hash.Hash,
	module uint,
	attachments Attachments,
) Application {
	out := application{
		hash:        hash,
		module:      module,
		attachments: attachments,
	}

	return &out
}

// Hash returns the hash
func (obj *application) Hash() hash.Hash {
	return obj.hash
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
