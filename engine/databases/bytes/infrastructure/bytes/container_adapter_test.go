package bytes

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers/delimiters"
)

func TestContainerAdapter_single_Success(t *testing.T) {
	container := containers.NewContainerForTests(
		"myKeyname",
		pointers.NewPointersForTests([]pointers.Pointer{
			pointers.NewPointerForTests(
				delimiters.NewDelimiterForTests(0, 12),
				true,
			),
			pointers.NewPointerForTests(
				delimiters.NewDelimiterForTests(1, 22),
				false,
			),
		}))

	adapter := NewContainerAdapter()
	retBytes, err := adapter.InstanceToBytes(container)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to contain 0 bytes")
		return
	}

	if !reflect.DeepEqual(container, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestContainerAdapter_single_withRemaining_Success(t *testing.T) {
	container := containers.NewContainerForTests(
		"myKeyname",
		pointers.NewPointersForTests([]pointers.Pointer{
			pointers.NewPointerForTests(
				delimiters.NewDelimiterForTests(0, 12),
				true,
			),
			pointers.NewPointerForTests(
				delimiters.NewDelimiterForTests(1, 22),
				false,
			),
		}))

	adapter := NewContainerAdapter()
	retBytes, err := adapter.InstanceToBytes(container)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some bytes long enough")
	retBytes = append(retBytes, remaining...)
	retIns, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(container, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestContainerAdapter_multiple_Success(t *testing.T) {
	containers := containers.NewContainersForTests([]containers.Container{
		containers.NewContainerForTests(
			"firstContainer",
			pointers.NewPointersForTests([]pointers.Pointer{
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(0, 12),
					true,
				),
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(1, 22),
					false,
				),
			})),
		containers.NewContainerForTests(
			"secondContainer",
			pointers.NewPointersForTests([]pointers.Pointer{
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(0, 12),
					true,
				),
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(1, 22),
					false,
				),
			})),
	})

	adapter := NewContainerAdapter()
	retBytes, err := adapter.InstancesToBytes(containers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to contain 0 bytes")
		return
	}

	if !reflect.DeepEqual(containers, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestContainerAdapter_multiple_withRemaining_Success(t *testing.T) {
	containers := containers.NewContainersForTests([]containers.Container{
		containers.NewContainerForTests(
			"firstContainer",
			pointers.NewPointersForTests([]pointers.Pointer{
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(0, 12),
					true,
				),
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(1, 22),
					false,
				),
			})),
		containers.NewContainerForTests(
			"secondContainer",
			pointers.NewPointersForTests([]pointers.Pointer{
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(0, 12),
					true,
				),
				pointers.NewPointerForTests(
					delimiters.NewDelimiterForTests(1, 22),
					false,
				),
			})),
	})

	adapter := NewContainerAdapter()
	retBytes, err := adapter.InstancesToBytes(containers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some bytes long enough")
	retBytes = append(retBytes, remaining...)
	retIns, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(containers, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
