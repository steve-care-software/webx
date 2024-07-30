package bytes

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers/delimiters"
)

func TestStateAdapter_withContainers_single_Success(t *testing.T) {
	state := states.NewStateWithContainersForTests(
		containers.NewContainersForTests([]containers.Container{
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
		}),
		true,
	)

	adapter := NewStateAdapter()
	retBytes, err := adapter.InstanceToBytes(state)
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

	if !reflect.DeepEqual(state, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestStateAdapter_withContainers_single_withRemaining_Success(t *testing.T) {
	state := states.NewStateWithContainersForTests(
		containers.NewContainersForTests([]containers.Container{
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
		}),
		true,
	)

	adapter := NewStateAdapter()
	retBytes, err := adapter.InstanceToBytes(state)
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

	if !reflect.DeepEqual(state, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestStateAdapter_withoutContainers_single_Success(t *testing.T) {
	state := states.NewStateForTests(
		false,
	)

	adapter := NewStateAdapter()
	retBytes, err := adapter.InstanceToBytes(state)
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

	if !reflect.DeepEqual(state, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestStateAdapter_withoutContainers_single_withRemaining_Success(t *testing.T) {
	state := states.NewStateForTests(
		false,
	)

	adapter := NewStateAdapter()
	retBytes, err := adapter.InstanceToBytes(state)
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

	if !reflect.DeepEqual(state, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestStateAdapter_multiple_Success(t *testing.T) {
	states := states.NewStatesForTests([]states.State{
		states.NewStateWithContainersForTests(
			containers.NewContainersForTests([]containers.Container{
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
			}),
			true,
		),
		states.NewStateForTests(
			false,
		),
		states.NewStateForTests(
			true,
		),
	})

	adapter := NewStateAdapter()
	retBytes, err := adapter.InstancesToBytes(states)
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

	if !reflect.DeepEqual(states, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestStateAdapter_multiple_withRemaining_Success(t *testing.T) {
	states := states.NewStatesForTests([]states.State{
		states.NewStateWithContainersForTests(
			containers.NewContainersForTests([]containers.Container{
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
			}),
			true,
		),
		states.NewStateForTests(
			false,
		),
		states.NewStateForTests(
			true,
		),
	})

	adapter := NewStateAdapter()
	retBytes, err := adapter.InstancesToBytes(states)
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

	if !reflect.DeepEqual(states, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
