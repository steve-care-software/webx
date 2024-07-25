package omissions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
)

func TestAdapter_withPrefix_Success(t *testing.T) {
	ins := omissions.NewOmissionWithPrefixForTests(
		elements.NewElementWithBytesForTests([]byte("prefix")),
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

func TestAdapter_withSuffix_Success(t *testing.T) {
	ins := omissions.NewOmissionWithSuffixForTests(
		elements.NewElementWithBytesForTests([]byte("suffix")),
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

func TestAdapter_withPrefix_withSuffix_Success(t *testing.T) {
	ins := omissions.NewOmissionWithPrefixAndSuffixForTests(
		elements.NewElementWithBytesForTests([]byte("prefix")),
		elements.NewElementWithBytesForTests([]byte("suffix")),
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
