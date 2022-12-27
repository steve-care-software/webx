package programs

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type attachments struct {
	hash hash.Hash
	list []Attachment
}

func createAttachments(
	hash hash.Hash,
	list []Attachment,
) Attachments {
	out := attachments{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *attachments) Hash() hash.Hash {
	return obj.hash
}

// List returns the attachments
func (obj *attachments) List() []Attachment {
	return obj.list
}
