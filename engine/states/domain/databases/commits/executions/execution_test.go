package executions

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

func TestExecutionBuilder(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	builder := NewExecutionBuilder()

	fingerPrint, err := hashAdapter.FromBytes([]byte("This is some bytes"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	chunkBuilder := chunks.NewBuilder()
	chunk, err := chunkBuilder.WithPath([]string{"path", "to", "file"}).WithFingerPrint(*fingerPrint).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test creating Execution with bytes
	expectedBytes := []byte("some bytes")
	execution, err := builder.Create().WithBytes(expectedBytes).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !execution.IsBytes() {
		t.Errorf("Expected IsBytes to be true, got %v", execution.IsBytes())
	}

	if !bytes.Equal(execution.Bytes(), expectedBytes) {
		t.Errorf("Expected Bytes to be %v, got %v", expectedBytes, execution.Bytes())
	}

	if execution.IsChunk() {
		t.Errorf("Expected IsChunk to be false, got %v", execution.IsChunk())
	}

	// Test creating Execution with chunk
	execution, err = builder.Create().WithChunk(chunk).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !execution.IsChunk() {
		t.Errorf("Expected IsChunk to be true, got %v", execution.IsChunk())
	}

	if execution.IsBytes() {
		t.Errorf("Expected IsBytes to be false, got %v", execution.IsBytes())
	}

	if !execution.Chunk().Hash().Compare(chunk.Hash()) {
		t.Errorf("Expected chunk %v, got %v", chunk.Hash(), execution.Chunk().Hash())
	}

	// Test creating Execution without mandatory fields
	_, err = builder.Create().Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	// Test creating Execution with both mandatory fields
	_, err = builder.Create().WithBytes(expectedBytes).WithChunk(chunk).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
