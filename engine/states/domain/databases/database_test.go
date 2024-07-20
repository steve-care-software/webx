package databases

import (
	"testing"

	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions"
	"github.com/steve-care-software/webx/engine/states/domain/databases/metadatas"
)

func TestDatabaseBuilder(t *testing.T) {
	databaseBuilder := NewBuilder()

	// Create a commit
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

	commitBuilder := commits.NewBuilder()
	commit, err := commitBuilder.Create().WithExecutions(executionsInstance).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Create metadata
	path := []string{"path", "to", "metadata"}
	name := "metadata name"
	description := "metadata description"
	metadataBuilder := metadatas.NewBuilder()
	metadata, err := metadataBuilder.Create().WithPath(path).WithName(name).WithDescription(description).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test creating Database
	database, err := databaseBuilder.Create().WithHead(commit).WithMetaData(metadata).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !database.Head().Hash().Compare(commit.Hash()) {
		t.Errorf("Expected head hash %v, got %v", commit.Hash(), database.Head().Hash())
	}

	if !database.MetaData().Hash().Compare(metadata.Hash()) {
		t.Errorf("Expected metadata hash %v, got %v", metadata.Hash(), database.MetaData().Hash())
	}

	// Test creating Database without mandatory fields
	_, err = databaseBuilder.Create().WithMetaData(metadata).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	_, err = databaseBuilder.Create().WithHead(commit).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
