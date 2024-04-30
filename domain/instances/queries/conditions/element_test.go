package conditions

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/resources"
)

func TestElement_withResource_Success(t *testing.T) {
	resource := resources.NewResourceWithFieldForTests(
		pointers.NewPointerForTests("myEntity", "myField"),
	)

	ins := NewElementWithResourceForTests(resource)

	if !ins.IsResource() {
		t.Errorf("the element was expected to contain a resource")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the element was expected to NOT contain a condition")
		return
	}

	retResource := ins.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}
}

func TestElement_withCondition_Success(t *testing.T) {
	condition := NewConditionForTests(
		pointers.NewPointerForTests("myEntity", "myField"),
		operators.NewOperatorWithEqualForTests(),
		NewElementWithResourceForTests(
			resources.NewResourceWithValueForTests(45),
		),
	)

	ins := NewElementWithConditionForTests(condition)

	if ins.IsResource() {
		t.Errorf("the element was expected to NOT contain a resource")
		return
	}

	if !ins.IsCondition() {
		t.Errorf("the element was expected to contain a condition")
		return
	}

	retCondition := ins.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the condition is invalid")
		return
	}
}

func TestElement_withoutParam_returnsError(t *testing.T) {
	_, err := NewElementBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be nil, error returned")
		return
	}
}
