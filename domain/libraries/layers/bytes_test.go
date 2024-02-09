package layers

import (
	"reflect"
	"testing"
)

func TestBytes_withJoin_Success(t *testing.T) {
	join := []string{
		"first",
		"second",
	}

	ins := NewBytesWithJoinForTests(join)

	if !ins.IsJoin() {
		t.Errorf("the bytes was expected to contain a join")
		return
	}

	if ins.IsCompare() {
		t.Errorf("the bytes was expected to NOT contain a compare")
		return
	}

	if ins.IsHashBytes() {
		t.Errorf("the bytes was expected to NOT contain a hashBytes")
		return
	}

	retJoin := ins.Join()
	if !reflect.DeepEqual(join, retJoin) {
		t.Errorf("the returned join is invalid")
		return
	}
}

func TestBytes_withCompare_Success(t *testing.T) {
	compare := []string{
		"first",
		"second",
	}

	ins := NewBytesWithCompareForTests(compare)

	if ins.IsJoin() {
		t.Errorf("the bytes was expected to NOT contain a join")
		return
	}

	if !ins.IsCompare() {
		t.Errorf("the bytes was expected to contain a compare")
		return
	}

	if ins.IsHashBytes() {
		t.Errorf("the bytes was expected to NOT contain a hashBytes")
		return
	}

	retCompare := ins.Compare()
	if !reflect.DeepEqual(compare, retCompare) {
		t.Errorf("the returned compare is invalid")
		return
	}
}

func TestBytes_withHashBytes_Success(t *testing.T) {
	hashBytes := "myVariable"
	ins := NewBytesWithHashBytesForTests(hashBytes)

	if ins.IsJoin() {
		t.Errorf("the bytes was expected to NOT contain a join")
		return
	}

	if ins.IsCompare() {
		t.Errorf("the bytes was expected to NOT contain a compare")
		return
	}

	if !ins.IsHashBytes() {
		t.Errorf("the bytes was expected to contain a hashBytes")
		return
	}

	retHashBytes := ins.HashBytes()
	if !reflect.DeepEqual(hashBytes, retHashBytes) {
		t.Errorf("the returned hashBytes is invalid")
		return
	}
}

func TestBytes_withoutParam_returnsError(t *testing.T) {
	_, err := NewBytesBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
