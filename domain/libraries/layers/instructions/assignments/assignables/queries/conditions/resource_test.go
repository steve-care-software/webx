package conditions

import (
	"reflect"
	"testing"
)

func TestResource_withField_Success(t *testing.T) {
	field := NewPointerForTests("myEntity", "myField")
	ins := NewResourceWithFieldForTests(field)

	if !ins.IsField() {
		t.Errorf("the resource was expected to contain a field")
		return
	}

	if ins.IsValue() {
		t.Errorf("the resource was expected to NOT contain a value")
		return
	}

	retField := ins.Field()
	if !reflect.DeepEqual(field, retField) {
		t.Errorf("the field was expected to be '%s', '%s' returned", field, retField)
		return
	}
}

func TestResource_withValue_Success(t *testing.T) {
	value := 45
	ins := NewResourceWithValueForTests(value)

	if ins.IsField() {
		t.Errorf("the resource was expected to NOT contain a field")
		return
	}

	if !ins.IsValue() {
		t.Errorf("the resource was expected to contain a value")
		return
	}

	retValue := ins.Value()
	if value != retValue {
		t.Errorf("the value was expected to be '%d', '%d' returned", value, retValue)
		return
	}
}

func TestResource_withoutParam_returnsError(t *testing.T) {
	_, err := NewResourceBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
