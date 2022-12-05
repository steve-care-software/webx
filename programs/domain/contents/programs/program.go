package programs

import "github.com/steve-care-software/webx/domain/databases/entities"

type program struct {
	entity       entities.Entity
	instructions entities.Identifiers
	outputs      []uint
}

func createProgram(
	entity entities.Entity,
	instructions entities.Identifiers,
) Program {
	return createProgramInternally(entity, instructions, nil)
}

func createProgramWithOutputs(
	entity entities.Entity,
	instructions entities.Identifiers,
	outputs []uint,
) Program {
	return createProgramInternally(entity, instructions, outputs)
}

func createProgramInternally(
	entity entities.Entity,
	instructions entities.Identifiers,
	outputs []uint,
) Program {
	out := program{
		entity:       entity,
		instructions: instructions,
		outputs:      outputs,
	}

	return &out
}

// Entity returns the entity
func (obj *program) Entity() entities.Entity {
	return obj.entity
}

// Instructions returns the instructions
func (obj *program) Instructions() entities.Identifiers {
	return obj.instructions
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *program) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *program) Outputs() []uint {
	return obj.outputs
}
