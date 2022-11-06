package suites

import "github.com/steve-care-software/webx/domain/databases/entities"

type suite struct {
	entity  entities.Entity
	isValid bool
	content []byte
}

func createSuite(
	entity entities.Entity,
	isValid bool,
	content []byte,
) Suite {
	out := suite{
		entity:  entity,
		isValid: isValid,
		content: content,
	}

	return &out
}

// Entity returns the entity
func (obj *suite) Entity() entities.Entity {
	return obj.entity
}

// IsValid returns true if valid, false otherwise
func (obj *suite) IsValid() bool {
	return obj.isValid
}

// Content returns the content
func (obj *suite) Content() []byte {
	return obj.content
}
