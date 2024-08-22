package instructions

import (
	"errors"
	"fmt"
)

type instructions struct {
	list []Instruction
	mp   map[string][]Instruction
}

func createInstructions(
	list []Instruction,
	mp map[string][]Instruction,
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
func (obj *instructions) Fetch(name string, idx uint) (Instruction, error) {
	if ins, ok := obj.mp[name]; ok {
		length := uint(len(ins))
		if idx >= length {
			str := fmt.Sprintf("the instruction (%s) could not be found at index (%d), its length is: %d", name, idx, length)
			return nil, errors.New(str)
		}

		return ins[idx], nil
	}

	str := fmt.Sprintf("the instruction (name: %s) does not exists", name)
	return nil, errors.New(str)
}
