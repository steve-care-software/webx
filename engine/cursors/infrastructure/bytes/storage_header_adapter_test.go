package bytes

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers"
)

func TestStorageHeaderAdapter_withIdentities_Success(t *testing.T) {
	header := headers.NewHeaderWithIdentitiesForTests(
		storages.NewStorageForTests(
			delimiters.NewDelimiterForTests(0, 12),
			true,
		),
	)

	adapter := NewStorageHeaderAdapter()
	retBytes, err := adapter.ToBytes(header)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(header, retIns) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
