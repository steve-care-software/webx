package references

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net/url"
)

type adapter struct {
	contentKeysAdapter ContentKeysAdapter
	commitsAdapter     CommitsAdapter
	builder            Builder
}

func createAdapter(
	contentKeysAdapter ContentKeysAdapter,
	commitsAdapter CommitsAdapter,
	builder Builder,
) Adapter {
	out := adapter{
		contentKeysAdapter: contentKeysAdapter,
		commitsAdapter:     commitsAdapter,
		builder:            builder,
	}
	return &out
}

// ToContent converts reference to bytes
func (app *adapter) ToContent(ins Reference) ([]byte, error) {
	contentKeyBytes, err := app.contentKeysAdapter.ToContent(ins.ContentKeys())
	if err != nil {
		return nil, err
	}

	contentKeysLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(contentKeysLengthBytes, uint64(len(contentKeyBytes)))

	commitsBytes, err := app.commitsAdapter.ToContent(ins.Commits())
	if err != nil {
		return nil, err
	}

	commitLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(commitLengthBytes, uint64(len(commitsBytes)))

	output := []byte{}
	output = append(output, contentKeysLengthBytes...)
	output = append(output, contentKeyBytes...)
	output = append(output, commitLengthBytes...)
	output = append(output, commitsBytes...)

	if ins.HasPeers() {
		peersList := ins.Peers()
		peerBytes := []byte{}
		for _, onePeer := range peersList {
			singlePeerBytes := []byte(onePeer.String())
			peersLength := make([]byte, 8)
			binary.LittleEndian.PutUint64(peersLength, uint64(len(singlePeerBytes)))

			peerBytes = append(peerBytes, []byte(peersLength)...)
			peerBytes = append(peerBytes, []byte(singlePeerBytes)...)
		}

		output = append(output, peerBytes...)
	}

	return output, nil
}

// ToReference converts bytes to reference
func (app *adapter) ToReference(content []byte) (Reference, error) {
	contentLength := len(content)
	if contentLength < minReferenceSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Reference instance, %d provided", minReferenceSize, contentLength)
		return nil, errors.New(str)
	}

	contentKeysBytesLengthDelimiter := uint64(8)
	contentKeysBytesLength := binary.LittleEndian.Uint64(content[:contentKeysBytesLengthDelimiter])
	contentKeysBytesDelimiter := int(contentKeysBytesLength + contentKeysBytesLengthDelimiter)
	if contentLength < contentKeysBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the ContentKeys size of the Reference instance, %d provided", contentKeysBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	contentKeys, err := app.contentKeysAdapter.ToContentKeys(content[contentKeysBytesLengthDelimiter:contentKeysBytesDelimiter])
	if err != nil {
		return nil, err
	}

	commitBytesLengthDelimiter := uint64(contentKeysBytesDelimiter + 8)
	commitBytesLength := binary.LittleEndian.Uint64(content[contentKeysBytesDelimiter:commitBytesLengthDelimiter])
	commitBytesDelimiter := int(commitBytesLength + commitBytesLengthDelimiter)
	if contentLength < commitBytesDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to retrieve the Commits size of the Reference instance, %d provided", commitBytesDelimiter, contentLength)
		return nil, errors.New(str)
	}

	commits, err := app.commitsAdapter.ToCommits(content[commitBytesLengthDelimiter:commitBytesDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[commitBytesDelimiter:]
	builder := app.builder.Create().WithContentKeys(contentKeys).WithCommits(commits)
	if len(remaining) > 0 {
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

		if len(peersList) > 0 {
			builder.WithPeers(peersList)
		}
	}

	return builder.Now()
}
