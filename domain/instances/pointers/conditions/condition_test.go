package conditions

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions/operators"
)

func TestCondition_resourceMatch_comparisonMatch_andOperator_returnsTrue(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, false),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithAndForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, false),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if !condition.Match(history) {
		t.Errorf("the condition was expected to match!")
		return
	}

}

func TestCondition_resourceMatch_comparisoDoNotnMatch_andOperator_returnsFalse(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, false),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithAndForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, true),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if condition.Match(history) {
		t.Errorf("the condition was expected to NOT match!")
		return
	}

}

func TestCondition_resourceMatch_comparisoDoNotnMatch_orOperator_returnsTrue(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, false),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithOrForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, true),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if !condition.Match(history) {
		t.Errorf("the condition was expected to match!")
		return
	}

}

func TestCondition_resourceDoNotMatch_comparisoDoNotnMatch_orOperator_returnsFalse(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, true),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithOrForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, true),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if condition.Match(history) {
		t.Errorf("the condition was expected to NOT match!")
		return
	}

}

func TestCondition_resourceMatch_comparisonMatch_xorOperator_returnsFalse(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, false),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithXorForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, false),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if condition.Match(history) {
		t.Errorf("the condition was expected to NOT match!")
		return
	}

}

func TestCondition_resourceDoNotMatch_comparisonMatch_xorOperator_returnsTrue(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, true),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithXorForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, false),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if !condition.Match(history) {
		t.Errorf("the condition was expected to match!")
		return
	}

}

func TestCondition_resourceMatch_comparisonDoNotMatch_xorOperator_returnsTrue(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, false),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithXorForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, true),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if !condition.Match(history) {
		t.Errorf("the condition was expected to match!")
		return
	}

}

func TestCondition_resourceDoNotMatch_comparisonDoNotMatch_xorOperator_returnsFalse(t *testing.T) {
	condition := NewConditionWithComparisonsForTests(
		NewResourceForTests([]string{"this", "is", "a", "path"}, true),
		NewComparisonsForTests([]Comparison{
			NewComparisonForTests(
				operators.NewOperatorWithXorForTests(),
				NewConditionForTests(
					NewResourceForTests([]string{"this", "is", "a", "second", "path"}, true),
				),
			),
		}),
	)

	history := [][]string{
		[]string{"first", "path"},
		[]string{"second", "path"},
		[]string{"last", "path"},
	}

	if condition.Match(history) {
		t.Errorf("the condition was expected to NOT match!")
		return
	}

}
