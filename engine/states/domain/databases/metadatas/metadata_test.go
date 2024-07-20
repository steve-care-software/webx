package metadatas

import (
	"testing"
)

func TestMetadataBuilder(t *testing.T) {
	metadataBuilder := NewBuilder()

	path := []string{"path", "to", "metadata"}
	name := "metadata name"
	description := "metadata description"

	// Test creating MetaData
	metadata, err := metadataBuilder.Create().WithPath(path).WithName(name).WithDescription(description).Now()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if metadata.Name() != name {
		t.Errorf("Expected name %v, got %v", name, metadata.Name())
	}

	if metadata.Description() != description {
		t.Errorf("Expected description %v, got %v", description, metadata.Description())
	}

	if len(metadata.Path()) != len(path) {
		t.Errorf("Expected path length %v, got %v", len(path), len(metadata.Path()))
	}

	for i, p := range metadata.Path() {
		if p != path[i] {
			t.Errorf("Expected path %v at index %v, got %v", path[i], i, p)
		}
	}

	// Test creating MetaData without mandatory fields
	_, err = metadataBuilder.Create().WithName(name).WithDescription(description).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	_, err = metadataBuilder.Create().WithPath(path).WithDescription(description).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	_, err = metadataBuilder.Create().WithPath(path).WithName(name).Now()
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
