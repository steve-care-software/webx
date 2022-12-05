package modules

import (
	"crypto/sha512"
	b64 "encoding/base64"
	"errors"
	"fmt"
)

type modules struct {
	list []Module
	mp   map[string]Module
}

func createModules(
	list []Module,
	mp map[string]Module,
) Modules {
	out := modules{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the modules
func (obj *modules) List() []Module {
	return obj.list
}

// Fetch fetches a module by name
func (obj *modules) Fetch(name []byte) (Module, error) {
	hashedData := sha512.New().Sum(name)
	keyname := b64.StdEncoding.EncodeToString(hashedData)

	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the module (name: %v, hash: %s) is undefined", name, keyname)
	return nil, errors.New(str)
}
