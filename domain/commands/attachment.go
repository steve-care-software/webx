package commands

import "github.com/steve-care-software/webx/domain/criterias"

type attachment struct {
	current     criterias.Criteria
	target      criterias.Criteria
	application criterias.Criteria
}

func createAttachment(
	current criterias.Criteria,
	target criterias.Criteria,
	application criterias.Criteria,
) Attachment {
	out := attachment{
		current:     current,
		target:      target,
		application: application,
	}

	return &out
}

// Current returns the current
func (obj *attachment) Current() criterias.Criteria {
	return obj.current
}

// Target returns the target
func (obj *attachment) Target() criterias.Criteria {
	return obj.target
}

// Application returns the application
func (obj *attachment) Application() criterias.Criteria {
	return obj.application
}
