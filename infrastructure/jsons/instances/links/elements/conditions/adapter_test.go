package conditions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions/resources"
)

func TestAdapter_Success(t *testing.T) {
	ins := conditions.NewConditionForTests(
		resources.NewResourceForTests(uint(45)),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withNext_withResource_Success(t *testing.T) {
	ins := conditions.NewConditionWithNextForTests(
		resources.NewResourceForTests(uint(45)),
		conditions.NewConditionValueWithResourceForTests(
			resources.NewResourceForTests(uint(32)),
		),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withNext_withCondition_Success(t *testing.T) {
	ins := conditions.NewConditionWithNextForTests(
		resources.NewResourceForTests(uint(45)),
		conditions.NewConditionValueWithConditionForTests(
			conditions.NewConditionWithNextForTests(
				resources.NewResourceForTests(uint(45)),
				conditions.NewConditionValueWithResourceForTests(
					resources.NewResourceForTests(uint(32)),
				),
			),
		),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
