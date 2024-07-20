package constants

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/constants"
)

func TestAdapter_Success(t *testing.T) {
	ins := constants.NewConstantsForTests([]constants.Constant{
		constants.NewConstantWithBoolForTests(true),
		constants.NewConstantWithStringForTests("this is a string"),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withBool_Success(t *testing.T) {
	ins := constants.NewConstantWithBoolForTests(true)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withString_Success(t *testing.T) {
	ins := constants.NewConstantWithStringForTests("this is a string")

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withInt_Success(t *testing.T) {
	ins := constants.NewConstantWithIntForTests(-22)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withUint_Success(t *testing.T) {
	ins := constants.NewConstantWithUintForTests(22)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withFloat_Success(t *testing.T) {
	ins := constants.NewConstantWithFloatForTests(3.1416)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withList_Success(t *testing.T) {
	ins := constants.NewConstantWithListForTests(
		constants.NewConstantsForTests([]constants.Constant{
			constants.NewConstantWithBoolForTests(false),
			constants.NewConstantWithStringForTests("this is a string"),
			constants.NewConstantWithIntForTests(-22),
			constants.NewConstantWithUintForTests(22),
			constants.NewConstantWithFloatForTests(3.1416),
			constants.NewConstantWithListForTests(
				constants.NewConstantsForTests([]constants.Constant{
					constants.NewConstantWithBoolForTests(true),
					constants.NewConstantWithFloatForTests(3.1416),
				}),
			),
		}),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
