package tokens

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

func TestAdapter_list_Success(t *testing.T) {
	ins := tokens.NewTokensForTests([]tokens.Token{
		tokens.NewTokenWithOmissionForTests(
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementWithBytesForTests([]byte("this is some bytes")),
				elements.NewElementWithStringForTests("this is a string"),
			}),
			cardinalities.NewCardinalityForTests(45),
			omissions.NewOmissionWithPrefixForTests(
				elements.NewElementWithBytesForTests([]byte("prefix")),
			),
		),
		tokens.NewTokenForTests(
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementWithBytesForTests([]byte("this is some bytes")),
				elements.NewElementWithStringForTests("this is a string"),
			}),
			cardinalities.NewCardinalityForTests(45),
		),
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

func TestAdapter_single_Success(t *testing.T) {
	ins := tokens.NewTokenForTests(
		elements.NewElementsForTests([]elements.Element{
			elements.NewElementWithBytesForTests([]byte("this is some bytes")),
			elements.NewElementWithStringForTests("this is a string"),
		}),
		cardinalities.NewCardinalityForTests(45),
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

func TestAdapter_single_withOmission_Success(t *testing.T) {
	ins := tokens.NewTokenWithOmissionForTests(
		elements.NewElementsForTests([]elements.Element{
			elements.NewElementWithBytesForTests([]byte("this is some bytes")),
			elements.NewElementWithStringForTests("this is a string"),
		}),
		cardinalities.NewCardinalityForTests(45),
		omissions.NewOmissionWithPrefixForTests(
			elements.NewElementWithBytesForTests([]byte("prefix")),
		),
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
