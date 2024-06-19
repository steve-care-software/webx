package pointers

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
)

func TestPointers_isActive_containsNoLoader_containsNoCanceller_returnsOnePointer(t *testing.T) {
	path := []string{"this", "is", "a", "path"}
	pointers := NewPointersForTests([]Pointer{
		NewPointerForTests(path, true),
	})

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	matches := pointers.Match(history)
	if len(matches) != 1 {
		t.Errorf("%d match was expected, %d returned", 1, len(matches))
		return
	}
}

func TestPointers_isNotActive_containsNoLoader_containsNoCanceller_returnsZeroPointer(t *testing.T) {
	path := []string{"this", "is", "a", "path"}
	pointers := NewPointersForTests([]Pointer{
		NewPointerForTests(path, false),
	})

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	matches := pointers.Match(history)
	if len(matches) != 0 {
		t.Errorf("%d match was expected, %d returned", 0, len(matches))
		return
	}
}

func TestPointers_isActive_containsNoLoader_containsCanceller_historyhMatchesCanceller_returnsZeroPointer(t *testing.T) {
	cancellerPath := []string{"first", "path"}
	path := []string{"this", "is", "a", "path"}
	pointers := NewPointersForTests([]Pointer{
		NewPointerWithCancellerForTests(
			path,
			true,
			conditions.NewConditionForTests(
				conditions.NewResourceForTests(cancellerPath, false),
			),
		),
	})

	history := [][]string{
		[]string{"first", "path"},
		cancellerPath,
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	matches := pointers.Match(history)
	if len(matches) != 0 {
		t.Errorf("%d match was expected, %d returned", 0, len(matches))
		return
	}
}

func TestPointers_isActive_containsLoader_containsNoCanceller_historyhMatchesLoader_returnsOnePointer(t *testing.T) {
	loaderPath := []string{"first", "path"}
	path := []string{"this", "is", "a", "path"}
	pointers := NewPointersForTests([]Pointer{
		NewPointerWithLoaderForTests(
			path,
			true,
			conditions.NewConditionForTests(
				conditions.NewResourceForTests(loaderPath, false),
			),
		),
	})

	history := [][]string{
		[]string{"first", "path"},
		loaderPath,
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	matches := pointers.Match(history)
	if len(matches) != 1 {
		t.Errorf("%d match was expected, %d returned", 1, len(matches))
		return
	}
}

func TestPointers_isActive_containsLoader_containsCanceller_historyhMatchesLoader_historyMatchesCanceller_returnsZeroPointer(t *testing.T) {
	loaderAndCancellerPath := []string{"first", "path"}
	path := []string{"this", "is", "a", "path"}
	pointers := NewPointersForTests([]Pointer{
		NewPointerWithLoaderAndCancellerForTests(
			path,
			true,
			conditions.NewConditionForTests(
				conditions.NewResourceForTests(loaderAndCancellerPath, false),
			),
			conditions.NewConditionForTests(
				conditions.NewResourceForTests(loaderAndCancellerPath, false),
			),
		),
	})

	history := [][]string{
		[]string{"first", "path"},
		loaderAndCancellerPath,
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	matches := pointers.Match(history)
	if len(matches) != 0 {
		t.Errorf("%d match was expected, %d returned", 0, len(matches))
		return
	}
}
