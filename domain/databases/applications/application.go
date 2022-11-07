package applications

import "github.com/steve-care-software/webx/domain/databases/entities"

type application struct {
	entity  entities.Entity
	content Content
}

func createApplication(
	entity entities.Entity,
	content Content,
) Application {
	out := application{
		entity:  entity,
		content: content,
	}

	return &out
}

// Entity returns the entity
func (obj *application) Entity() entities.Entity {
	return obj.entity
}

// Content returns the content
func (obj *application) Content() Content {
	return obj.content
}
