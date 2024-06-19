package conditions

import (
	"testing"
)

func TestResource_match_mustBeLoaded_pathIsInHistory_returnsTrue(t *testing.T) {
	path := []string{"this", "is", "a", "path"}
	resource := NewResourceForTests(path, true)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		path,
		[]string{"last", "path"},
	}

	if !resource.Match(history) {
		t.Errorf("the resource was expected to match!")
		return
	}
}

func TestResource_match_mustBeLoaded_pathIsNOTInHistory_returnsFalse(t *testing.T) {
	path := []string{"this", "is", "a", "path"}
	resource := NewResourceForTests(path, true)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if resource.Match(history) {
		t.Errorf("the resource was expected to NOT match!")
		return
	}
}

func TestResource_match_pathIsNOTHistory_returnsTrue(t *testing.T) {
	path := []string{"this", "is", "a", "path"}
	resource := NewResourceForTests(path, false)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if !resource.Match(history) {
		t.Errorf("the resource was expected to match!")
		return
	}
}

func TestResource_match_pathIsInHistory_returnsTrue(t *testing.T) {
	path := []string{"this", "is", "a", "path"}
	resource := NewResourceForTests(path, false)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		path,
		[]string{"last", "path"},
	}

	if !resource.Match(history) {
		t.Errorf("the resource was expected to match!")
		return
	}
}
