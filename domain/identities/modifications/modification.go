package modifications

import "time"

type modification struct {
	content   Content
	createdOn time.Time
}

func createModification(
	content Content,
	createdOn time.Time,
) Modification {
	out := modification{
		content:   content,
		createdOn: createdOn,
	}

	return &out
}

// Content returns the content
func (obj *modification) Content() Content {
	return obj.content
}

// CreatedOn returns the creation time
func (obj *modification) CreatedOn() time.Time {
	return obj.createdOn
}
