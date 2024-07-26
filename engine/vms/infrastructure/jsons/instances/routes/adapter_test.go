package routes

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

func TestAdapter_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := routes.NewRouteForTests(
		*pHash,
		tokens.NewTokensForTests([]tokens.Token{
			tokens.NewTokenWithOmissionForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
				omissions.NewOmissionWithPrefixForTests(
					elements.NewElementWithBytesForTests([]byte("prefix")),
				),
			),
			tokens.NewTokenForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
			),
		}),
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

func TestAdapter_withTokenOmission_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := routes.NewRouteWithTokenForTests(
		*pHash,
		tokens.NewTokensForTests([]tokens.Token{
			tokens.NewTokenWithOmissionForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
				omissions.NewOmissionWithPrefixForTests(
					elements.NewElementWithBytesForTests([]byte("prefix")),
				),
			),
			tokens.NewTokenForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
			),
		}),
		omissions.NewOmissionWithPrefixForTests(
			elements.NewElementWithBytesForTests([]byte("prefix")),
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

func TestAdapter_withGlobal_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := routes.NewRouteWithGobalForTests(
		*pHash,
		tokens.NewTokensForTests([]tokens.Token{
			tokens.NewTokenWithOmissionForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
				omissions.NewOmissionWithPrefixForTests(
					elements.NewElementWithBytesForTests([]byte("prefix")),
				),
			),
			tokens.NewTokenForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
			),
		}),
		omissions.NewOmissionWithPrefixForTests(
			elements.NewElementWithBytesForTests([]byte("prefix")),
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

func TestAdapter_withGlobal_withTokenOmission_Success(t *testing.T) {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is a layer hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ins := routes.NewRouteWithGobalAndTokenForTests(
		*pHash,
		tokens.NewTokensForTests([]tokens.Token{
			tokens.NewTokenWithOmissionForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
				omissions.NewOmissionWithPrefixForTests(
					elements.NewElementWithBytesForTests([]byte("prefix")),
				),
			),
			tokens.NewTokenForTests(
				elements.NewElementsForTests([]elements.Element{
					elements.NewElementWithBytesForTests([]byte("this is some bytes")),
					elements.NewElementWithStringForTests("this is a string"),
				}),
				cardinalities.NewCardinalityForTests(45),
			),
		}),
		omissions.NewOmissionWithPrefixForTests(
			elements.NewElementWithBytesForTests([]byte(" ")),
		),
		omissions.NewOmissionWithPrefixForTests(
			elements.NewElementWithBytesForTests([]byte("prefix")),
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
