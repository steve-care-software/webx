package attachments

type attachment struct {
	variable    Variable
	application []byte
}

func createAttachment(
	variable Variable,
	application []byte,
) Attachment {
	out := attachment{
		variable:    variable,
		application: application,
	}

	return &out
}

// Variable returns the variable
func (obj *attachment) Variable() Variable {
	return obj.variable
}

// Application returns the application
func (obj *attachment) Application() []byte {
	return obj.application
}
