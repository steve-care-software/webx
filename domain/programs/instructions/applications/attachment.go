package applications

type attachment struct {
	assignment Assignment
	local      string
}

func createAttachment(
	assignment Assignment,
	local string,
) Attachment {
	out := attachment{
		assignment: assignment,
		local:      local,
	}

	return &out
}

// Assignment returns the assignment
func (obj *attachment) Assignment() Assignment {
	return obj.assignment
}

// Local returns the local
func (obj *attachment) Local() string {
	return obj.local
}
