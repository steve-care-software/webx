package attachments

type attachment struct {
	variable    Variable
	application string
}

func createAttachment(
	variable Variable,
	application string,
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
func (obj *attachment) Application() string {
	return obj.application
}
