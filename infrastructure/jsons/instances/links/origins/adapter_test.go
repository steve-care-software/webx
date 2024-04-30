package origins

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

func TestAdapter_withResource_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pSecondHash, err := hash.NewAdapter().FromBytes([]byte("this is another layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := origins.NewOriginForTests(
		resources.NewResourceForTests(*pHash),
		operators.NewOperatorWithAndForTests(),
		origins.NewValueWithResourceForTests(
			resources.NewResourceWithIsMandatoryForTests(*pSecondHash),
		),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestAdapter_withOrigin_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pSecondHash, err := hash.NewAdapter().FromBytes([]byte("this is another layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := origins.NewOriginForTests(
		resources.NewResourceForTests(*pHash),
		operators.NewOperatorWithAndForTests(),
		origins.NewValueWithOriginForTests(
			origins.NewOriginForTests(
				resources.NewResourceForTests(*pHash),
				operators.NewOperatorWithAndForTests(),
				origins.NewValueWithResourceForTests(
					resources.NewResourceWithIsMandatoryForTests(*pSecondHash),
				),
			),
		),
	)
	adapter := NewAdapter()

	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
