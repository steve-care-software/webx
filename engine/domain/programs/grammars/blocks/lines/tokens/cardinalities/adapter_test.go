package cardinalities

import "testing"

func TestAdapter_Success(t *testing.T) {
	cardinality := NewCardinalityForTests(32)
	adapter := NewAdapter()
	retAST, err := adapter.ToNFT(cardinality)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retCardinality, err := adapter.ToInstance(retAST)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if cardinality.Min() != retCardinality.Min() {
		t.Errorf("the returned min was expected to be %d, %d returned", cardinality.Min(), retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestAdapter_withMax_Success(t *testing.T) {
	max := uint(32)
	cardinality := NewCardinalityWithMaxForTests(0, max)
	adapter := NewAdapter()
	retAST, err := adapter.ToNFT(cardinality)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retCardinality, err := adapter.ToInstance(retAST)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if cardinality.Min() != retCardinality.Min() {
		t.Errorf("the returned min was expected to be %d, %d returned", cardinality.Min(), retCardinality.Min())
		return
	}

	if !retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pRetMax := retCardinality.Max()
	if max != *pRetMax {
		t.Errorf("the returned max was expected to be %d, %d returned", max, *pRetMax)
		return
	}
}
