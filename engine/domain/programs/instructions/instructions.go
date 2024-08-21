package instructions

import (
	"errors"
	"fmt"
)

type instructions struct {
	list []Instruction
	mp   map[string]Instruction
}

func createInstructions(
	list []Instruction,
	mp map[string]Instruction,
) Instructions {
	out := instructions{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of instruction
func (obj *instructions) List() []Instruction {
	return obj.list
}

// Fetch fetches an instruction by name
func (obj *instructions) Fetch(name string) (Instruction, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the instruction (name: %s) does not exists", name)
	return nil, errors.New(str)
}
