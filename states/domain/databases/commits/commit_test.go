package commits

import (
	"testing"

	"github.com/steve-care-software/datastencil/states/domain/databases/commits/executions"
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

func TestCommitBuilder(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	commitBuilder := NewBuilder()

	executionBuilder := executions.NewExecutionBuilder()
	execution, err := executionBuilder.Create().WithBytes([]byte("some bytes")).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	executionsBuilder := executions.NewBuilder()
	executionsInstance, err := executionsBuilder.Create().WithList([]executions.Execution{execution}).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test creating Commit without parent
	commit, err := commitBuilder.Create().WithExecutions(executionsInstance).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !commit.Executions().Hash().Compare(executionsInstance.Hash()) {
		t.Errorf("Expected executions hash %v, got %v", executionsInstance.Hash(), commit.Executions().Hash())
	}

	if commit.HasParent() {
		t.Errorf("Expected HasParent to be false, got %v", commit.HasParent())
	}

	// Test creating Commit with parent
	parentHash, err := hashAdapter.FromBytes([]byte("parent hash"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commit, err = commitBuilder.Create().WithExecutions(executionsInstance).WithParent(*parentHash).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !commit.Executions().Hash().Compare(executionsInstance.Hash()) {
		t.Errorf("Expected executions hash %v, got %v", executionsInstance.Hash(), commit.Executions().Hash())
	}

	if !commit.HasParent() {
		t.Errorf("Expected HasParent to be true, got %v", commit.HasParent())
	}

	if !commit.Parent().Compare(*parentHash) {
		t.Errorf("Expected parent hash %v, got %v", *parentHash, commit.Parent())
	}

	// Test creating Commit without executions
	_, err = commitBuilder.Create().Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
