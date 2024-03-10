package conditions

import (
	"testing"
)

func TestPointer_Success(t *testing.T) {
	entity := "myEntity"
	field := "myField"
	ins := NewPointerForTests(entity, field)

	retEntity := ins.Entity()
	if entity != retEntity {
		t.Errorf("the entity was expected to be '%s', '%s' returned", entity, retEntity)
		return
	}

	retField := ins.Field()
	if field != retField {
		t.Errorf("the field was expected to be '%s', '%s' returned", field, retField)
		return
	}
}

func TestPointer_withoutEntity_returnsError(t *testing.T) {
	field := "myField"
	_, err := NewPointerBuilder().Create().WithField(field).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestPointer_withoutField_returnsError(t *testing.T) {
	entity := "myEntity"
	_, err := NewPointerBuilder().Create().WithEntity(entity).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
