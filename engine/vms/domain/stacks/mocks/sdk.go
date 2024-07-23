package mocks

import (
	"github.com/steve-care-software/webx/engine/vms/domain/encryptors"
	"github.com/steve-care-software/webx/engine/vms/domain/instances"
)

// NewEncryptor creates a new encryptor
func NewEncryptor(
	encrypt map[string]map[string][]byte,
	decrypt map[string]map[string][]byte,
) encryptors.Encryptor {
	return createEncryptor(
		encrypt,
		decrypt,
	)
}

// NewInstanceAdapter creates a new instance adapter
func NewInstanceAdapter(
	toData map[string][]byte,
	toInstance map[string]instances.Instance,
) instances.Adapter {
	return createInstanceAdapter(
		toData,
		toInstance,
	)
}
