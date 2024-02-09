package links

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
)

func TestOrigin_Success(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	resource := NewOriginResourceForTests(*pFirstLayer)
	operator := NewOperatorWithAndForTests()
	next := NewOriginValueWithResourceForTests(
		NewOriginResourceForTests(*pSecondLayer),
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
	operator := NewOperatorWithAndForTests()
	next := NewOriginValueWithResourceForTests(
		NewOriginResourceForTests(*pSecondLayer),
	)

	_, err := NewOriginBuilder().Create().WithOperator(operator).WithNext(next).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestOrigin__withoutOperator_returnsError(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	resource := NewOriginResourceForTests(*pFirstLayer)
	next := NewOriginValueWithResourceForTests(
		NewOriginResourceForTests(*pSecondLayer),
	)

	_, err := NewOriginBuilder().Create().WithResource(resource).WithNext(next).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestOrigin__withoutNext_returnsError(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	resource := NewOriginResourceForTests(*pFirstLayer)
	operator := NewOperatorWithAndForTests()
	_, err := NewOriginBuilder().Create().WithResource(resource).WithOperator(operator).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
