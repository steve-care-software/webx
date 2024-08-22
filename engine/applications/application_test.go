package applications

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/steve-care-software/webx/engine/domain/programs"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
)

func TestApplication_interpret_steveCare_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.root;
		# .SPACE .TAB. EOL;

		root: ._addTwoNumbers .numbers .PLUS_SIGN .numbers;
		numbers: .oneNumber+;
		oneNumber: .N_ZERO
					| .N_ONE
					| .N_TWO
					| .N_THREE
					| .N_FOUR
					| .N_FIVE
					| .N_SIX
					| .N_SEVEN
					| .N_HEIGHT
					| .N_NINE
					;

		_addTwoNumbers: @math_operation_arithmetic_add .numbers[0]:first .numbers[1]:second;

		N_ZERO: "0";
		N_ONE: "1";
		N_TWO: "2";
		N_THREE: "3";
		N_FOUR: "4";
		N_FIVE: "5";
		N_SIX: "6";
		N_SEVEN: "7";
		N_HEIGHT: "8";
		N_NINE: "9";
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";
		PLUS_SIGN: "+";
		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	programRemaining := []byte("this is a remaining")
	programInput := append([]byte(`
		12 + 345`), programRemaining...)

	grammarParserAdapter := grammars.NewParserAdapter()
	retGrammar, retRemaining, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal([]byte(""), retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	parserAdapter := programs.NewParserAdapter()
	retProgram, retRemaining, err := parserAdapter.ToProgram(retGrammar, programInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(programRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	application := NewApplication()
	_, err = application.Interpret(retProgram)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	//panic(errors.New("stop"))

	fmt.Printf("\n%v\n", retProgram)
}

func TestApplication_interpret_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.additionInParenthesis;
		# .SPACE .TAB .EOL;

		additionInParenthesis: .OPEN_PARENTHESIS .addition .CLOSE_PARENTHESIS;
		addition: ._addTwoNumbers .firstNumber .PLUS_SIGN .secondNumber
				| ._addTwoNumbers .firstNumber .PLUS_WORD .secondNumber
				;

		secondNumber: .N_THREE .N_FOUR .N_FIVE;
		firstNumber: .N_ONE .N_TWO;

		_addTwoNumbers: @math_operation_arithmetic_add .firstNumber:first .secondNumber:second;

		PLUS_WORD: "plus";
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
		TAB: "	";
		EOL: "
";
	`)

	programRemaining := []byte("this is a remaining")
	programInput := append([]byte(`
		( 12 + 345 )`), programRemaining...)

	grammarParserAdapter := grammars.NewParserAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	parserAdapter := programs.NewParserAdapter()
	retProgram, retRemaining, err := parserAdapter.ToProgram(retGrammar, programInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(programRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	application := NewApplication()
	_, err = application.Interpret(retProgram)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%v\n", retProgram)
}

func TestApplication_grammar_composeBlock_withReplacement_Success(t *testing.T) {
	input := []byte(`
		v1;
		>.additionInParenthesis;
		# .SPACE .TAB .EOL;

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

	application := NewApplication()
	retGrammar, _, err := application.ParseGrammar(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retValue, err := application.ComposeBlock(retGrammar, "additionInParenthesis")
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
