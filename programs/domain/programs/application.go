package programs

import "github.com/steve-care-software/webx/programs/domain/programs/modules"

type application struct {
	name        []byte
	module      modules.Module
	attachments Attachments
}

func createApplication(
	name []byte,
	module modules.Module,
) Application {
	return createApplicationInternally(name, module, nil)
}

func createApplicationWithAttachments(
	name []byte,
	module modules.Module,
	attachments Attachments,
) Application {
	return createApplicationInternally(name, module, attachments)
}

func createApplicationInternally(
	name []byte,
	module modules.Module,
	attachments Attachments,
) Application {
	out := application{
		name:        name,
		module:      module,
		attachments: attachments,
	}

	return &out
}

// Name returns the name
func (obj *application) Name() []byte {
	return obj.name
}

// Module returns the module
func (obj *application) Module() modules.Module {
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
