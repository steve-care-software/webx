package commands

import "github.com/steve-care-software/syntax/domain/bytes/criterias"

type attachment struct {
	global      criterias.Criteria
	local       criterias.Criteria
	application criterias.Criteria
}

func createAttachment(
	global criterias.Criteria,
	local criterias.Criteria,
	application criterias.Criteria,
) Attachment {
	out := attachment{
		global:      global,
		local:       local,
		application: application,
	}

	return &out
}

// Global returns the global
func (obj *attachment) Global() criterias.Criteria {
	return obj.global
}

// Local returns the local
func (obj *attachment) Local() criterias.Criteria {
	return obj.local
}

// Application returns the application
func (obj *attachment) Application() criterias.Criteria {
	return obj.application
}
