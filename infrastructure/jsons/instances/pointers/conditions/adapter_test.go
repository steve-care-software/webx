package conditions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions/operators"
)

func TestAdapter_Success(t *testing.T) {
	ins := conditions.NewConditionWithComparisonsForTests(
		conditions.NewResourceForTests(
			[]string{"this", "is", "a", "path"},
			false,
		),
		conditions.NewComparisonsForTests([]conditions.Comparison{
			conditions.NewComparisonForTests(
				operators.NewOperatorWithAndForTests(),
				conditions.NewConditionForTests(
					conditions.NewResourceForTests(
						[]string{"this", "is", "another", "path"},
						true,
					),
				),
			),
		}),
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
