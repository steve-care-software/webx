package resources

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type resources struct {
	hash hash.Hash
	mp   map[string]Resource
	list []Resource
}

func createResources(
	hash hash.Hash,
	mp map[string]Resource,
	list []Resource,
) Resources {
	out := resources{
		hash: hash,
		mp:   mp,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *resources) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *resources) List() []Resource {
	return obj.list
}

// FetchByName fetches a resource by name
func (obj *resources) FetchByName(name string) (Resource, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("there is no Resource named '%s'", name)
	return nil, errors.New(str)
}

// FetchByPath fetches a resource by path
func (obj *resources) FetchByPath(path []string) (Resource, error) {
	if len(path) <= 0 {
		return nil, errors.New("the path must contain at least 1 name, none provided")
	}

	retResource, err := obj.FetchByName(path[0])
	if err != nil {
		return nil, err
	}

	if len(path) <= 1 {
		return retResource, nil
	}

	if !retResource.HasChildren() {
		str := fmt.Sprintf("the path still had names (%s) but the resource (name: %s) had no children", strings.Join(path[1:], "/"), retResource.Name())
		return nil, errors.New(str)
	}

	return retResource.Children().FetchByPath(path[1:])
}
