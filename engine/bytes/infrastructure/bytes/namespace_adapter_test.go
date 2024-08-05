package bytes

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
)

func TestNamespaceAdapter_multiple_Success(t *testing.T) {
	namespacesIns := namespaces.NewNamespacesForTests([]namespaces.Namespace{
		namespaces.NewNamespaceForTests("first", "This is a description", false),
		namespaces.NewNamespaceWithIterationsForTests(
			"second",
			"",
			true,
			delimiters.NewDelimiterForTests(0, 12),
		),
	})

	adapter := NewNamespaceAdapter()
	retBytes, err := adapter.InstancesToBytes(namespacesIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(namespacesIns, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestNamespaceAdapter_single_Success(t *testing.T) {
	namespaceIns := namespaces.NewNamespaceForTests("myName", "This is a description", false)

	adapter := NewNamespaceAdapter()
	retBytes, err := adapter.InstanceToBytes(namespaceIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(namespaceIns, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestNamespaceAdapter_withIterations_Success(t *testing.T) {
	namespaceIns := namespaces.NewNamespaceWithIterationsForTests(
		"myName",
		"This is a description",
		true,
		delimiters.NewDelimiterForTests(0, 12),
	)

	adapter := NewNamespaceAdapter()
	retBytes, err := adapter.InstanceToBytes(namespaceIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(namespaceIns, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
