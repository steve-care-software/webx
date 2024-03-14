package origins

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

func TestValue_withResource_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	originResource := resources.NewResourceForTests(*pLayer)
	value := NewValueWithResourceForTests(originResource)

	if !value.IsResource() {
		t.Errorf("the value was expected to contain a resource")
		return
	}

	if value.IsOrigin() {
		t.Errorf("the value was expected to NOT contain a resource")
		return
	}

	retResource := value.Resource()
	if !reflect.DeepEqual(originResource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}

}

func TestValue_withOrigin_Success(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	origin := NewOriginForTests(
		resources.NewResourceForTests(*pFirstLayer),
		operators.NewOperatorWithAndForTests(),
		NewValueWithResourceForTests(
			resources.NewResourceForTests(*pSecondLayer),
		),
	)
	value := NewValueWithOriginForTests(origin)

	if value.IsResource() {
		t.Errorf("the value was expected to NOT contain a resource")
		return
	}

	if !value.IsOrigin() {
		t.Errorf("the value was expected to contain a resource")
		return
	}

	retOrigin := value.Origin()
	if !reflect.DeepEqual(origin, retOrigin) {
		t.Errorf("the origin is invalid")
		return
	}

}

func TestValue_withoutParam_returnsError(t *testing.T) {
	_, err := NewValueBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
