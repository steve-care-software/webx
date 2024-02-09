package layers

import "testing"

func TestKind_isPrompt_Success(t *testing.T) {
	kind := NewKindWithPromptForTests()
	if !kind.IsPrompt() {
		t.Errorf("the kind was expected to contain a prompt")
		return
	}

	if kind.IsContinue() {
		t.Errorf("the kind was expected to NOT contain a continue")
		return
	}
}

func TestKind_isContinue_Success(t *testing.T) {
	kind := NewKindWithContinueForTests()
	if kind.IsPrompt() {
		t.Errorf("the kind was expected to NOT contain a prompt")
		return
	}

	if !kind.IsContinue() {
		t.Errorf("the kind was expected to contain a continue")
		return
	}
}

func TestKind_withoutParam_returnsError(t *testing.T) {
	_, err := NewKindBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
