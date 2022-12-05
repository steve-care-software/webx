package instructions

import "github.com/steve-care-software/webx/domain/databases/entities"

type instruction struct {
	entity  entities.Entity
	content Content
}

func createInstruction(
	entity entities.Entity,
	content Content,
) Instruction {
	out := instruction{
		entity:  entity,
		content: content,
	}

	return &out
}

// Entity returns the entity
func (obj *instruction) Entity() entities.Entity {
	return obj.entity
}

// Content returns the content
func (obj *instruction) Content() Content {
	return obj.content
}
