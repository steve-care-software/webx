package layers

import "testing"

func TestLayers_withList_Success(t *testing.T) {
	list := []Layer{
		NewLayerForTests(
			NewInstructionsForTests([]Instruction{
				NewInstructionWithStopForTests(),
			}),
			NewOutputForTests(
				"myVariable",
				NewKindWithContinueForTests(),
			),
			"myInput",
		),
	}

	ins := NewLayersForTests(list)
	retList := ins.List()
	if len(list) != len(retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestLayers_withEmptyList_returnsError(t *testing.T) {
	list := []Layer{}
	_, err := NewBuilder().Create().WithList(list).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}

func TestLayers_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
