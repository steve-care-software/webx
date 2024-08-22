package applications

import (
	"testing"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
)

func TestApplication_interpret_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.instructions;
		# .SPACE .TAB .EOL;

		instructions: .instruction*;
		instruction: .instructionPossibilities .SEMI_COLON;
		instructionPossibilities: .assignment;

		assignment: .boolAssignment
				  	| .uintAssignment
					| .intAssignment
					| .floatAssignment
					;

		boolAssignment: .typeBoolDefinition .EQUAL .boolValue;
		typeBoolDefinition: .typeBool .variableName;
		boolValue: .TRUE
				   | .FALSE
				   ;


		uintAssignment: .typeUintDefinition .EQUAL .numbers;
		intAssignment: .typeIntDefinition .EQUAL .intNumbers;
		floatAssignment: .typeFloatDefinition .EQUAL .floatNumbers;

		variableName: .oneLowerCaseLetter .letters+;
		
		typePrimitive: .typeBool
					   | .typeString
					   | .typeFloatDefinition
					   | .typeUintDefinition
					   | .typeIntDefinition
					   ;
		
		typeBool: .LL_B .LL_O[2] .LL_L;
		typeString: .LL_S .LL_T .LL_R .LL_I .LL_N .LL_G;

		typeFloatDefinition: .typeFloat64
				  			 | .typeFloat32
				  			 ;

		typeFloat64: .typeFloat .sixtyFour;
		typeFloat32: .typeFloat .thirtyTwo;
		typeFloat: .LL_F .LL_L .LL_O .LL_A .LL_T;

		typeUintDefinition: .typeUint64
				  			| .typeUint32
							| .typeUint16
							| .typeUint8
				  			;

		typeUint64: .typeUint .sixtyFour;
		typeUint32: .typeUint .thirtyTwo;
		typeUint16: .typeUint .sixteen;
		typeUint8: .typeUint .N_HEIGHT;
		typeUint: .LL_U .typeInt;

		typeIntDefinition: .typeInt64
				  		   | .typeInt32
						   | .typeInt16
						   | .typeInt8
				  		   ;

		typeInt64: .typeInt .sixtyFour;
		typeInt32: .typeInt .thirtyTwo;
		typeInt16: .typeInt .sixteen;
		typeInt8: .typeInt .N_HEIGHT;
		typeInt: .LL_I .LL_N .LL_T;

		sixtyFour: .N_SIX .N_FOUR;
		thirtyTwo: .N_THREE .N_TWO;
		sixteen: .N_ONE .N_SIX;

		letters: .uppercaseLetters
				 | .lowerCaseLetters
				 ;

		uppercaseLetters: .oneUpperCaseLetter+;
		oneUpperCaseLetter: .UL_A
							| .UL_B
							| .UL_C
							| .UL_D
							| .UL_E
							| .UL_F
							| .UL_G
							| .UL_H
							| .UL_I
							| .UL_J
							| .UL_K
							| .UL_L
							| .UL_M
							| .UL_N
							| .UL_O
							| .UL_P
							| .UL_Q
							| .UL_R
							| .UL_S
							| .UL_T
							| .UL_U
							| .UL_V
							| .UL_W
							| .UL_X
							| .UL_Y
							| .UL_Z
							;

		lowerCaseLetters: .oneLowerCaseLetter+;
		oneLowerCaseLetter: .LL_A
							| .LL_B
							| .LL_C
							| .LL_D
							| .LL_E
							| .LL_F
							| .LL_G
							| .LL_H
							| .LL_I
							| .LL_J
							| .LL_K
							| .LL_L
							| .LL_M
							| .LL_N
							| .LL_O
							| .LL_P
							| .LL_Q
							| .LL_R
							| .LL_S
							| .LL_T
							| .LL_U
							| .LL_V
							| .LL_W
							| .LL_X
							| .LL_Y
							| .LL_Z
							;

		
		floatNumbers: .negativeFloatNumber
					  | .floatNumber
					  ;

		negativeFloatNumber: .MINUS .floatNumber;
		floatNumber: .numbers .DOT .numbers;

		intNumbers: .negativeNumber
					| .numbers
					;

		negativeNumber: .MINUS .numbers;
		numbers: .oneNumber+
				---
					oneNumber: .N_ZERO.
					numberWithAllNumbers: .testNumberWithAllNumbers.
					oneLettter: @.LL_A.
				;

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


		testNumberWithAllNumbers: .N_ONE .N_ZERO .N_TWO .N_THREE .N_FOUR .N_FIVE .N_SIX .N_SEVEN .N_HEIGHT .N_NINE
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

		LL_A: "a";
		LL_B: "b";
		LL_C: "c";
		LL_D: "d";
		LL_E: "e";
		LL_F: "f";
		LL_G: "g";
		LL_H: "h";
		LL_I: "i";
		LL_J: "j";
		LL_K: "k";
		LL_L: "l";
		LL_M: "m";
		LL_N: "n";
		LL_O: "o";
		LL_P: "p";
		LL_Q: "q";
		LL_R: "r";
		LL_S: "s";
		LL_T: "t";
		LL_U: "u";
		LL_V: "v";
		LL_W: "w";
		LL_X: "x";
		LL_Y: "y";
		LL_Z: "z";
		
		UL_A: "A";
		UL_B: "B";
		UL_C: "C";
		UL_D: "D";
		UL_E: "E";
		UL_F: "F";
		UL_G: "G";
		UL_H: "H";
		UL_I: "I";
		UL_J: "J";
		UL_K: "K";
		UL_L: "L";
		UL_M: "M";
		UL_N: "N";
		UL_O: "O";
		UL_P: "P";
		UL_Q: "Q";
		UL_R: "R";
		UL_S: "S";
		UL_T: "T";
		UL_U: "U";
		UL_V: "V";
		UL_W: "W";
		UL_X: "X";
		UL_Y: "Y";
		UL_Z: "Z";

		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";

		PLUS: "+";
		MINUS: "-";
		DOT: ".";
		EQUAL: ".";
		SEMI_COLON: ";";

		TRUE: "true";
		FALSE: "false";

		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	//programRemaining := []byte("this is a remaining")
	//programInput := append([]byte(`2`), programRemaining...)

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

	application := NewApplication()
	err = application.Suites(retGrammar)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}

	/*parserAdapter := programs.NewParserAdapter()
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

	fmt.Printf("\n%v\n", retProgram)*/
}

/*
func TestApplication_grammar_composeBlock_withReplacement_Success(t *testing.T) {
	input := []byte(`
		v1;
		>.additionInParenthesis;
		# .SPACE .TAB .EOL;

		additionInParenthesis: .OPEN_PARENTHESIS .addition .CLOSE_PARENTHESIS - .addition;
		addition: .firstNumber .PLUS .secondNumber - .myReplacement;
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
		PLUS: "+";
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
*/
