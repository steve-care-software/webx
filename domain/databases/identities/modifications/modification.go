package modifications

import "github.com/steve-care-software/webx/domain/databases/entities"

type modification struct {
	entity  entities.Entity
	content Content
}

func createModification(
	entity entities.Entity,
	content Content,
) Modification {
	out := modification{
		entity:  entity,
		content: content,
	}

	return &out
}

// Entity returns the entity
func (obj *modification) Entity() entities.Entity {
	return obj.entity
}

// Content returns the content
func (obj *modification) Content() Content {
	return obj.content
}
