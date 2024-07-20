package chunks

import (
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

func TestBuilder(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	chunkBuilder := createBuilder(hashAdapter)

	fingerPrint, err := hashAdapter.FromBytes([]byte("This is some bytes"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test creating a chunk with all mandatory fields
	chunk, err := chunkBuilder.WithPath([]string{"path", "to", "file"}).WithFingerPrint(*fingerPrint).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !chunk.FingerPrint().Compare(*fingerPrint) {
		t.Errorf("Expected fingerprint %v, got %v", *fingerPrint, chunk.FingerPrint())
	}

	if len(chunk.Path()) != 3 || chunk.Path()[0] != "path" || chunk.Path()[1] != "to" || chunk.Path()[2] != "file" {
		t.Errorf("Expected path [path, to, file], got %v", chunk.Path())
	}

	// Test creating a chunk without mandatory fields
	_, err = chunkBuilder.WithPath(nil).WithFingerPrint(*fingerPrint).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	_, err = chunkBuilder.WithPath([]string{"path", "to", "file"}).WithFingerPrint(nil).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
