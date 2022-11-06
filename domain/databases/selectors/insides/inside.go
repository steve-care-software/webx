package insides

import "github.com/steve-care-software/webx/domain/databases/entities"

type inside struct {
	entity  entities.Entity
	content Content
}

func createInside(
	entity entities.Entity,
	content Content,
) Inside {
	out := inside{
		entity:  entity,
		content: content,
	}

	return &out
}

// Entity returns the entity
func (obj *inside) Entity() entities.Entity {
	return obj.entity
}

// Content returns the content
func (obj *inside) Content() Content {
	return obj.content
}
