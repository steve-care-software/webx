package references

import (
	"reflect"
	"testing"
)

func TestAdapter_withoutPeers_Success(t *testing.T) {
	reference := NewReferenceForTests(0)
	adapter := NewAdapter([]byte("0")[0])
	content, err := adapter.ToContent(reference)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retReference, err := adapter.ToReference(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(reference, retReference) {
		t.Errorf("the returned reference is invalid")
		return
	}
}

func TestAdapter_withOnePeer_Success(t *testing.T) {
	reference := NewReferenceForTests(1)
	adapter := NewAdapter([]byte("0")[0])
	content, err := adapter.ToContent(reference)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retReference, err := adapter.ToReference(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(reference.Commits(), retReference.Commits()) {
		t.Errorf("the returned commits is invalid")
		return
	}

	if !reflect.DeepEqual(reference.ContentKeys(), retReference.ContentKeys()) {
		t.Errorf("the returned contentKeys is invalid")
		return
	}

	peers := reference.Peers()
	retPeers := retReference.Peers()
	if len(peers) != len(retPeers) {
		t.Errorf("%d peers were expected, %d returned", len(peers), len(retPeers))
		return
	}
}

func TestAdapter_withMultiplePeers_Success(t *testing.T) {
	reference := NewReferenceForTests(123)
	adapter := NewAdapter([]byte("0")[0])
	content, err := adapter.ToContent(reference)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retReference, err := adapter.ToReference(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(reference.Commits(), retReference.Commits()) {
		t.Errorf("the returned commits is invalid")
		return
	}

	if !reflect.DeepEqual(reference.ContentKeys(), retReference.ContentKeys()) {
		t.Errorf("the returned contentKeys is invalid")
		return
	}

	peers := reference.Peers()
	retPeers := retReference.Peers()
	if len(peers) != len(retPeers) {
		t.Errorf("%d peers were expected, %d returned", len(peers), len(retPeers))
		return
	}
}
