package programs

import (
	"errors"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type attachmentsBuilder struct {
	hashAdapter hash.Adapter
	list        []Attachment
}

func createAttachmentsBuilder(
	hashAdapter hash.Adapter,
) AttachmentsBuilder {
	out := attachmentsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *attachmentsBuilder) Create() AttachmentsBuilder {
	return createAttachmentsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *attachmentsBuilder) WithList(list []Attachment) AttachmentsBuilder {
	app.list = list
	return app
}

// Now builds a new Attachments instance
func (app *attachmentsBuilder) Now() (Attachments, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Attachment in order to build a Attachments instance")
	}

	data := [][]byte{}
	for _, oneAttachment := range app.list {
		data = append(data, oneAttachment.Value().Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createAttachments(*pHash, app.list), nil
}
