package applications

import "github.com/steve-care-software/webx/domain/databases/entities"

type application struct {
	entity      entities.Entity
	module      entities.Identifier
	attachments Attachments
}

func createApplication(
	entity entities.Entity,
	module entities.Identifier,
) Application {
	return createApplicationInternally(entity, module, nil)
}

func createApplicationWithAttachments(
	entity entities.Entity,
	module entities.Identifier,
	attachments Attachments,
) Application {
	return createApplicationInternally(entity, module, attachments)
}

func createApplicationInternally(
	entity entities.Entity,
	module entities.Identifier,
	attachments Attachments,
) Application {
	out := application{
		entity:      entity,
		module:      module,
		attachments: attachments,
	}

	return &out
}

// Entity returns the entity
func (obj *application) Entity() entities.Entity {
	return obj.entity
}

// Module returns the module
func (obj *application) Module() entities.Identifier {
	return obj.module
}

// HasAttachments returns true if there is attachments, false otherwise
func (obj *application) HasAttachments() bool {
	return obj.attachments != nil
}

// Attachments returns the attachments, if any
func (obj *application) Attachments() Attachments {
	return obj.attachments
}
