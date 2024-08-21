package grammars

import "testing"

func TestComposeAdapter_withReplacement_Success(t *testing.T) {
	input := []byte(`
		v1;
		>.additionInParenthesis;
		# .SPACE .TAB. EOL;

		additionInParenthesis: .OPEN_PARENTHESIS .addition .CLOSE_PARENTHESIS - .addition;
		addition: .firstNumber .PLUS_SIGN .secondNumber - .myReplacement;
		secondNumber: .N_THREE .N_FOUR .N_FIVE;
		firstNumber: .N_ONE .N_TWO;
		myReplacement: .N_ONE .N_THREE;

		replacedNumber: .N_TWO .N_FOUR;

		N_ZERO: "0";
		N_ONE: "1";
		N_TWO: "2";
		N_THREE: "3";
		N_FOUR: "4";
		N_FIVE: "5";
		N_SIX: "6";
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";
		PLUS_SIGN: "+";
		SPACE: " ";
		TAB: "\t\t";
		EOL: "\n";
	`)

	parserAdapter := NewParserAdapter()
	retGrammar, _, err := parserAdapter.ToGrammar(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	composeAdapter := NewComposeAdapter()
	retValue, err := composeAdapter.ToBytes(retGrammar, "additionInParenthesis")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := "13"
	if string(retValue) != expected {
		t.Errorf("the returned value is invalid")
		return
	}
}

func TestComposeAdapter_withFunc_withReplacement_Success(t *testing.T) {
	input := []byte(`
		v1;
		>.addition;
		# .SPACE .TAB. EOL;

		addition: .firstNumber .PLUS_SIGN .secondNumber - .myReplacement;
		secondNumber: .N_THREE .N_FOUR .N_FIVE;
		firstNumber: .N_ONE .N_TWO;
		myReplacement: .N_ONE .N_THREE;

		replacedNumber: .N_TWO .N_FOUR;

		N_ZERO: "0";
		N_ONE: "1";
		N_TWO: "2";
		N_THREE: "3";
		N_FOUR: "4";
		N_FIVE: "5";
		N_SIX: "6";
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";
		PLUS_SIGN: "+";
		SPACE: " ";
		TAB: "\t\t";
		EOL: "\n";
	`)

	parserAdapter := NewParserAdapter()
	retGrammar, _, err := parserAdapter.ToGrammar(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	composeAdapter := NewComposeAdapter()
	retValue, err := composeAdapter.ToBytes(retGrammar, "addition")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := "13"
	if string(retValue) != expected {
		t.Errorf("the returned value is invalid")
		return
	}
}

func TestComposeAdapter_withFunc_Success(t *testing.T) {
	input := []byte(`
		v1;
		>.addition;
		# .SPACE .TAB. EOL;
		
		addition: .firstNumber .PLUS_SIGN .secondNumber - math_operation_arithmetic_addition .firstNumber[0]:first .secondNumber[0]:second;
		secondNumber: .N_THREE .N_FOUR .N_FIVE;
		firstNumber: .N_ONE .N_TWO;

		replacedNumber: .N_TWO .N_FOUR;

		N_ZERO: "0";
		N_ONE: "1";
		N_TWO: "2";
		N_THREE: "3";
		N_FOUR: "4";
		N_FIVE: "5";
		N_SIX: "6";
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";	
		PLUS_SIGN: "+";
		SPACE: " ";
		TAB: "\t\t";
		EOL: "\n";
	`)

	parserAdapter := NewParserAdapter()
	retGrammar, _, err := parserAdapter.ToGrammar(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	composeAdapter := NewComposeAdapter()
	retValue, err := composeAdapter.ToBytes(retGrammar, "addition")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := "357"
	if string(retValue) != expected {
		t.Errorf("the returned value is invalid")
		return
	}
}
