package modifications

import "github.com/steve-care-software/webx/domain/databases/entities"

type modification struct {
	identifier entities.Identifier
	content    Content
}

func createModification(
	identifier entities.Identifier,
	content Content,
) Modification {
	out := modification{
		identifier: identifier,
		content:    content,
	}

	return &out
}

// Identifier returns the identifier
func (obj *modification) Identifier() entities.Identifier {
	return obj.identifier
}

// Content returns the content
func (obj *modification) Content() Content {
	return obj.content
}
