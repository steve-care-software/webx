package links

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/identity/domain/hash"
)

func TestOriginValue_withResource_Success(t *testing.T) {
	pLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	originResource := NewOriginResourceForTests(*pLayer)
	originValue := NewOriginValueWithResourceForTests(originResource)

	if !originValue.IsResource() {
		t.Errorf("the originValue was expected to contain a resource")
		return
	}

	if originValue.IsOrigin() {
		t.Errorf("the originValue was expected to NOT contain a resource")
		return
	}

	retResource := originValue.Resource()
	if !reflect.DeepEqual(originResource, retResource) {
		t.Errorf("the resource is invalid")
		return
	}

}

func TestOriginValue_withOrigin_Success(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	origin := NewOriginForTests(
		NewOriginResourceForTests(*pFirstLayer),
		NewOperatorWithAndForTests(),
		NewOriginValueWithResourceForTests(
			NewOriginResourceForTests(*pSecondLayer),
		),
	)
	originValue := NewOriginValueWithOriginForTests(origin)

	if originValue.IsResource() {
		t.Errorf("the originValue was expected to NOT contain a resource")
		return
	}

	if !originValue.IsOrigin() {
		t.Errorf("the originValue was expected to contain a resource")
		return
	}

	retOrigin := originValue.Origin()
	if !reflect.DeepEqual(origin, retOrigin) {
		t.Errorf("the origin is invalid")
		return
	}

}

func TestOriginValue_withoutParam_returnsError(t *testing.T) {
	_, err := NewOriginValueBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
