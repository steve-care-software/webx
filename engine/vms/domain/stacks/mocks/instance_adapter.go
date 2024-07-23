package mocks

import (
	"errors"

	"github.com/steve-care-software/webx/engine/vms/domain/instances"
)

type instanceAdapter struct {
	toData     map[string][]byte
	toInstance map[string]instances.Instance
}

func createInstanceAdapter(
	toData map[string][]byte,
	toInstance map[string]instances.Instance,
) instances.Adapter {
	out := instanceAdapter{
		toData:     toData,
		toInstance: toInstance,
	}

	return &out
}

// ToBytes returns the instance to data
func (app *instanceAdapter) ToBytes(ins instances.Instance) ([]byte, error) {
	keyname := ins.Hash().String()
	if bytes, ok := app.toData[keyname]; ok {
		return bytes, nil
	}

	return nil, errors.New("the instance is invalid")
}

// ToInstance returns the data to instance
func (app *instanceAdapter) ToInstance(data []byte) (instances.Instance, error) {
	keyname := string(data)
	if ins, ok := app.toInstance[keyname]; ok {
		return ins, nil
	}

	return nil, errors.New("the data is invalid")
}
