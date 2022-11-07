package values

import "github.com/steve-care-software/webx/domain/databases/entities"

type value struct {
	entity  entities.Entity
	content Content
}

func createValue(
	entity entities.Entity,
	content Content,
) Value {
	out := value{
		entity:  entity,
		content: content,
	}

	return &out
}

// Entity returns the entity
func (obj *value) Entity() entities.Entity {
	return obj.entity
}

// Content returns the content
func (obj *value) Content() Content {
	return obj.content
}
