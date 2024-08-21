package grammars

import (
	"testing"
)

func TestApplication_grammar_compileGrammar_decompile_Success(t *testing.T) {
	input := []byte(`
		v1;
		>.numberInParenthesis;
		# .SPACE .TAB. EOL;

		numberInParenthesis: 	.number - my_func .number:number
								| .OPEN_PARENTHESIS .numberInParenthesis .CLOSE_PARENTHESIS
								---
				 					firstTest:@.N_ONE.
									secondTest:.N_TWO.
								;

		number: .anyNumber+;
		anyNumber: 	.N_ZERO
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

	nftAdapter := NewNFTAdapter()
	_, err = nftAdapter.ToNFT(retGrammar)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	//fmt.Printf("\n%v\n", retNFT)
}
