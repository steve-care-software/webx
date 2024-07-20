package executions

import (
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

func TestExecutionsBuilder(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()

	executionBuilder := createExecutionBuilder(hashAdapter)
	execution, err := executionBuilder.Create().WithBytes([]byte("some bytes")).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test creating Executions with all mandatory fields
	executionsInstance, err := builder.Create().WithList([]Execution{execution}).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(executionsInstance.List()) != 1 {
		t.Errorf("Expected list of 1 execution, got %v", len(executionsInstance.List()))
	}

	// Test creating Executions without mandatory fields
	_, err = builder.Create().Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
