package queries

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/queries/conditions"
)

func TestQuery_Success(t *testing.T) {
	entity := "myEntity"
	condition := conditions.NewConditionForTests(
		conditions.NewPointerForTests("myEntity", "myField"),
		conditions.NewOperatorWithEqualForTests(),
		conditions.NewElementWithResourceForTests(
			conditions.NewResourceWithValueForTests(45),
		),
	)

	ins := NewQueryForTests(entity, condition)

	if ins.HasFields() {
		t.Errorf("the query was expected to NOT contain fields")
		return
	}

	retEntity := ins.Entity()
	if entity != retEntity {
		t.Errorf("the entity was expeected to be '%s', '%s' returned", entity, retEntity)
		return
	}

	retCondition := ins.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the condition is invalid")
		return
	}
}

func TestQuery_withFields_Success(t *testing.T) {
	entity := "myEntity"
	condition := conditions.NewConditionForTests(
		conditions.NewPointerForTests("myEntity", "myField"),
		conditions.NewOperatorWithEqualForTests(),
		conditions.NewElementWithResourceForTests(
			conditions.NewResourceWithValueForTests(45),
		),
	)

	fields := []string{
		"firstField",
		"secondField",
	}

	ins := NewQueryWithFieldsForTests(entity, condition, fields)

	if !ins.HasFields() {
		t.Errorf("the query was expected to contain fields")
		return
	}

	retEntity := ins.Entity()
	if entity != retEntity {
		t.Errorf("the entity was expeected to be '%s', '%s' returned", entity, retEntity)
		return
	}

	retCondition := ins.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the condition is invalid")
		return
	}

	retFields := ins.Fields()
	if !reflect.DeepEqual(fields, retFields) {
		t.Errorf("the fields is invalid")
		return
	}
}

func TestQuery_withoutEntity_returnsError(t *testing.T) {
	condition := conditions.NewConditionForTests(
		conditions.NewPointerForTests("myEntity", "myField"),
		conditions.NewOperatorWithEqualForTests(),
		conditions.NewElementWithResourceForTests(
			conditions.NewResourceWithValueForTests(45),
		),
	)

	_, err := NewBuilder().Create().WithCondition(condition).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}

func TestQuery_withoutCondition_returnsError(t *testing.T) {
	entity := "myEntity"
	_, err := NewBuilder().Create().WithEntity(entity).Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
