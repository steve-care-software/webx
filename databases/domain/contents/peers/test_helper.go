package peers

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"net/url"
)

// NewPeersForTests creates new peers for tests
func NewPeersForTests(maxPeerAmount uint) []*url.URL {
	peersMap := map[string]*url.URL{}
	for i := 0; i < int(maxPeerAmount); i++ {
		str, err := generateRandomString(200)
		if err != nil {
			panic(err)
		}

		rawURL := fmt.Sprintf("http://%s:8080/myuri/01", str)
		peer, err := url.Parse(rawURL)
		if err != nil {
			panic(err)
		}

		peersMap[rawURL] = peer
	}

	peersList := []*url.URL{}
	for _, onePeer := range peersMap {
		peersList = append(peersList, onePeer)
	}

	return peersList

}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
