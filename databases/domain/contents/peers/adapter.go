package peers

import (
	"encoding/binary"
	"net/url"
)

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToContent converts peer url's to bytes
func (app *adapter) ToContent(peers []*url.URL) ([]byte, error) {
	peerBytes := []byte{}
	for _, onePeer := range peers {
		singlePeerBytes := []byte(onePeer.String())
		peersLength := make([]byte, 8)
		binary.LittleEndian.PutUint64(peersLength, uint64(len(singlePeerBytes)))

		peerBytes = append(peerBytes, []byte(peersLength)...)
		peerBytes = append(peerBytes, []byte(singlePeerBytes)...)
	}

	return peerBytes, nil
}

// ToPeers converts bytes to peer urls
func (app *adapter) ToPeers(content []byte) ([]*url.URL, error) {
	remaining := content
	peersList := []*url.URL{}
	for {
		if len(remaining) <= 0 {
			break
		}

		length := binary.LittleEndian.Uint64(remaining[0:8])

		delimiter := 8 + length
		peerStr := string(remaining[length:delimiter])
		peer, err := url.Parse(peerStr)
		if err != nil {
			return nil, err
		}

		peersList = append(peersList, peer)
		remaining = remaining[delimiter:]
	}

	return peersList, nil
}
