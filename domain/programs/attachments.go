package programs

type attachments struct {
	list []Attachment
}

func createAttachments(
	list []Attachment,
) Attachments {
	out := attachments{
		list: list,
	}

	return &out
}

// List returns the attachments
func (obj *attachments) List() []Attachment {
	return obj.list
}
