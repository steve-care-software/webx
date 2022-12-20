package peers

import (
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	peers := NewPeersForTests(200)
	adapter := NewAdapter()
	content, err := adapter.ToContent(peers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retPeers, err := adapter.ToPeers(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(peers) != len(retPeers) {
		t.Errorf("%d peers were expected, %d returned", len(peers), len(retPeers))
		return
	}
}
