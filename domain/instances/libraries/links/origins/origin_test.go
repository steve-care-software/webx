package origins

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/resources"
)

func TestOrigin_Success(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	resource := resources.NewResourceForTests(*pFirstLayer)
	operator := operators.NewOperatorWithAndForTests()
	next := NewValueWithResourceForTests(
		resources.NewResourceForTests(*pSecondLayer),
	)

	origin := NewOriginForTests(
		resource,
		operator,
		next,
	)

	retResource := origin.Resource()
	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}

	retOperator := origin.Operator()
	if !reflect.DeepEqual(operator, retOperator) {
		t.Errorf("the operator is invalid")
		return
	}

	retNext := origin.Next()
	if !reflect.DeepEqual(next, retNext) {
		t.Errorf("the next is invalid")
		return
	}
}

func TestOrigin__withoutResource_returnsError(t *testing.T) {
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	operator := operators.NewOperatorWithAndForTests()
	next := NewValueWithResourceForTests(
		resources.NewResourceForTests(*pSecondLayer),
	)

	_, err := NewBuilder().Create().WithOperator(operator).WithNext(next).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestOrigin__withoutOperator_returnsError(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	resource := resources.NewResourceForTests(*pFirstLayer)
	next := NewValueWithResourceForTests(
		resources.NewResourceForTests(*pSecondLayer),
	)

	_, err := NewBuilder().Create().WithResource(resource).WithNext(next).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestOrigin__withoutNext_returnsError(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	resource := resources.NewResourceForTests(*pFirstLayer)
	operator := operators.NewOperatorWithAndForTests()
	_, err := NewBuilder().Create().WithResource(resource).WithOperator(operator).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
