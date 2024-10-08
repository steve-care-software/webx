package programs

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
)

func TestParserAdapter_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.line;
		# .SPACE .TAB. EOL;
		
		line: !.additionInParenthesis .additionInParenthesis
			| .N_ZERO
			;

		additionInParenthesis: .OPEN_PARENTHESIS .addition .CLOSE_PARENTHESIS;
		addition: (my_syscall .firstNumber:first .secondNumber:second) .firstNumber .PLUS_SIGN .secondNumber;
		secondNumber: (my_syscall) .N_THREE .N_FOUR .N_FIVE;
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
		TAB: "	";
		EOL: "
";
	`)

	programRemaining := []byte("this is a remaining")
	programInput := append([]byte(`
		salut ( 12 + 345 )`), programRemaining...)

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

	parserAdapter := NewParserAdapter()
	retProgram, retRemaining, err := parserAdapter.ToProgram(retGrammar, programInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(programRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	fmt.Printf("\n%v\n", retProgram)

}
