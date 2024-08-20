package blocks

import (
	"errors"
	"fmt"
)

type blocks struct {
	list []Block
	mp   map[string]Block
}

func createBlocks(
	list []Block,
	mp map[string]Block,
) Blocks {
	out := blocks{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of block
func (obj *blocks) List() []Block {
	return obj.list
}

// Fetch fetches a block by name
func (obj *blocks) Fetch(name string) (Block, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the block (name: %s) does not exists", name)
	return nil, errors.New(str)
}
