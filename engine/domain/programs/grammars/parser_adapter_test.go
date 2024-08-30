package grammars

import (
	"bytes"
	"testing"
)

func TestParserAdapter_withOmissions_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`
		v1;
		>.myRoot;
		#.first.second.third;

		myFirst: !.myFirst[1] .mySecond* .myThird+ .myFourth? .myFifth[1,] - myFuncName_secondSection .myFirst:first .mySecond:second .myThird:third .myFourth:fourth .myFifth:fifth
				 | (another_syscall) .myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,] - .MY_REPLACEMENT
				 | .myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,] - .myReplacement
				 | .myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,]
				 ---
				 	firstTest:!"this is some value";
					secondTest:!"this is some value";
				 ;

		mySecond: (this_is_a_syscall .myFirst:first.mySecond[5]:second) .myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,] - myFuncName_secondSection .myFirst:first .mySecond:second .myThird:third .myFourth:fourth .myFifth:fifth
				 | .myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,] - myFuncName_secondSection .myFirst:first .mySecond:second .myThird:third .myFourth:fourth .myFifth:fifth
				 | .myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,] - .MY_REPLACEMENT
				 ;

		FIRST: "this \" with escape";
		SECOND: "some value";
		`), remaining...)

	retAdapter := NewParserAdapter()
	retGrammar, retRemaining, err := retAdapter.ToGrammar(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retGrammar.Version() != 1 {
		t.Errorf("the version was expected to be %d, %d returned", 1, retGrammar.Version())
		return
	}

	if retGrammar.Root().Name() != "myRoot" {
		t.Errorf("the root was expected to be %s, %s returned", "myRoot", retGrammar.Root())
		return
	}

	retBlocks := retGrammar.Blocks().List()
	if len(retBlocks) != 2 {
		t.Errorf("the grammar was expected to contain %d block instances, %d returned", 2, len(retBlocks))
		return
	}

	retRules := retGrammar.Rules().List()
	if len(retRules) != 2 {
		t.Errorf("the grammar was expected to contain %d rule instances, %d returned", 2, len(retRules))
		return
	}

	if !retGrammar.HasOmissions() {
		t.Errorf("the grammar was expected to contain omissions")
		return
	}

	retOmissions := retGrammar.Rules().List()
	if len(retOmissions) != 2 {
		t.Errorf("the grammar was expected to contain %d omission elements, %d returned", 2, len(retOmissions))
		return
	}
}

func TestParserAdapter_withoutVersion_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`>.myRoot;myFirst:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,];mySecond:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;FIRST:"this \" with escape";SECOND:"some value";`), remaining...)

	retAdapter := NewParserAdapter()
	_, _, err := retAdapter.ToGrammar(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestParserAdapter_withNonNumericVersion_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`vDE;>.myRoot;myFirst:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,];mySecond:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;FIRST:"this \" with escape";SECOND:"some value";`), remaining...)

	retAdapter := NewParserAdapter()
	_, _, err := retAdapter.ToGrammar(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestParserAdapter_withoutRoot_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`v1;myFirst:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,];mySecond:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;FIRST:"this \" with escape";SECOND:"some value";`), remaining...)

	retAdapter := NewParserAdapter()
	_, _, err := retAdapter.ToGrammar(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestParserAdapter_withInvalidRootElementReference_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`v1;>myRoot;myFirst:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,];mySecond:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;FIRST:"this \" with escape";SECOND:"some value";`), remaining...)

	retAdapter := NewParserAdapter()
	_, _, err := retAdapter.ToGrammar(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestParserAdapter_withInvalidOmissionElementReference_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`v1;>.myRoot;#invalidReference;myFirst:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,];mySecond:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;FIRST:"this \" with escape";SECOND:"some value";`), remaining...)

	retAdapter := NewParserAdapter()
	_, _, err := retAdapter.ToGrammar(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestParserAdapter_withoutBlocks_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`v1;>.myRoot;FIRST:"this \" with escape";SECOND:"some value";`), remaining...)

	retAdapter := NewParserAdapter()
	_, _, err := retAdapter.ToGrammar(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestParserAdapter_withoutRules_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`v1;>.myRoot;myFirst:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,];mySecond:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;`), remaining...)

	retAdapter := NewParserAdapter()
	_, _, err := retAdapter.ToGrammar(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_blocks_withoutBlocks_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(``), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToBlocks(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withoutSuffix_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myBlock:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withoutSuffix_withoutRemaining_returnsError(t *testing.T) {
	input := []byte(`myBlock:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";`)
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withoutSuffix_withInvalidBlockDefinition_returnsError(t *testing.T) {
	input := []byte(`myBlock.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;`)
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withoutSuffix_withoutLines_returnsError(t *testing.T) {
	input := []byte(`myBlock:---myFirst:"somedata";mySecond:!"somedata";mySecondTest:"somedata";myTest:!"somedata";;`)
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_suites_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`---myTest:"somedata";myTest:!"somedata";`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retSuites, retRemaining, err := retAdapter.bytesToSuites(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	list := retSuites.List()
	if len(list) != 2 {
		t.Errorf("the suites was expected to contain %d suite instances, %d returned", 2, len(list))
		return
	}
}

func TestApplication_suites_withoutSuites_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(``), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToSuites(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:"somedata";`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retSuite, retRemaining, err := retAdapter.bytesToSuite(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retSuite.Name() != "myTest" {
		t.Errorf("the suite name was expected to be (%s), (%s) returned", "myTest", retSuite.Name())
		return
	}

	if retSuite.IsFail() {
		t.Errorf("the suite was expected to NOT fail")
		return
	}
}

func TestApplication_suite_isFail_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:!"somedata";`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retSuite, retRemaining, err := retAdapter.bytesToSuite(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retSuite.Name() != "myTest" {
		t.Errorf("the suite name was expected to be (%s), (%s) returned", "myTest", retSuite.Name())
		return
	}

	if !retSuite.IsFail() {
		t.Errorf("the suite was expected to fail")
		return
	}
}

func TestApplication_suite_withInvalidElement_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:myElement`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_withInvalidBlockNameDefinition_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`#myTest:.myElement`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_withoutSuiteLineSuffix_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:.myElement`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_withoutSuiteLineSuffix_withoutRemainingBytes_returnsError(t *testing.T) {
	input := []byte(`myTest:.myElement`)
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_lines_withOneLine_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retLines, retRemaining, err := retAdapter.bytesToLines(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	list := retLines.List()
	if len(list) != 1 {
		t.Errorf("the lines was expected to contain %d lines, %d returned", 1, len(list))
		return
	}
}

func TestApplication_lines_withMultipleLines_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst:first.mySecond:second.myThird:third.myFourth:fourth.myFifth:fifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retLines, retRemaining, err := retAdapter.bytesToLines(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	list := retLines.List()
	if len(list) != 3 {
		t.Errorf("the lines was expected to contain %d lines, %d returned", 3, len(list))
		return
	}
}

func TestApplication_lines_withoutLine_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`not a line`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToLines(input)
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_line_withExecution_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`.myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,] - myFuncName_secondSection .myFirst:first .mySecond:second .myThird:third .myFourth:fourth .myFifth:fifth`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retLine, retRemaining, err := retAdapter.bytesToLine(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retLine.HasProcessor() {
		t.Errorf("the execution was expected to contain a processor")
		return
	}
}

func TestApplication_line_withReplacement_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`.myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,] - .myReplacement`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retLine, retRemaining, err := retAdapter.bytesToLine(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retLine.HasProcessor() {
		t.Errorf("the execution was expected to contain a processor")
		return
	}
}

func TestApplication_withoutTokens_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myFuncName_secondSection .myFirst:first .mySecond:second .myThird:third .myFourth:fourth .myFifth:fifth - .MY_REPLACEMENT`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToLine(input)
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_execution_withElements_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myFuncName_secondSection .myFirst:first .mySecond:second .myThird:third .myFourth:fourth .myFifth:fifth`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retExecution, retRemaining, err := retAdapter.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retExecution.HasParameters() {
		t.Errorf("the execution was expected to contain paramters")
		return
	}

	list := retExecution.Parameters().List()
	if len(list) != 5 {
		t.Errorf("the paramters list was expected to contain %d paramters, %d returned", 5, len(list))
		return
	}
}

func TestApplication_execution_withElements_withoutRemaining_Success(t *testing.T) {
	remaining := []byte("")
	input := append([]byte(`myFuncName_secondSection .myFirst:first .mySecond:second .myThird:third .myFourth:fourth .myFifth:fifth`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retExecution, retRemaining, err := retAdapter.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retExecution.HasParameters() {
		t.Errorf("the execution was expected to contain paramters")
		return
	}

	list := retExecution.Parameters().List()
	if len(list) != 5 {
		t.Errorf("the tokens list was expected to contain %d paramters, %d returned", 5, len(list))
		return
	}
}

func TestApplication_execution_withoutElements_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myFuncName_secondSection`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retExecution, retRemaining, err := retAdapter.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retExecution.HasParameters() {
		t.Errorf("the execution was expected to NOT contain parameters")
		return
	}
}

func TestApplication_execution_withoutElements_withoutRemaining_Success(t *testing.T) {
	remaining := []byte("")
	input := append([]byte(`myFuncName_secondSection`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retExecution, retRemaining, err := retAdapter.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retExecution.HasParameters() {
		t.Errorf("the execution was expected to NOT contain parameters")
		return
	}
}

func TestApplication_tokens_Success(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`.myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,]`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retToken, retRemaining, err := retAdapter.bytesToTokens(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	list := retToken.List()
	if len(list) != 5 {
		t.Errorf("the tokens list was expected to contain %d tokens, %d returned", 5, len(list))
		return
	}
}

func TestApplication_token_withBlockName_withCardinality_Success(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`.myToken[1]`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retToken, retRemaining, err := retAdapter.bytesToToken(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retToken.Name() != "myToken" {
		t.Errorf("the token name is invalid, expected (%s), returned (%s)", "myToken", retToken.Name())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	cardinality := retToken.Cardinality()
	if cardinality.Min() != 1 {
		t.Errorf("the cardinality min was expected to be (%d), (%d) returned", 1, cardinality.Min())
		return
	}

	if !cardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pMax := cardinality.Max()
	max := *pMax
	if max != 1 {
		t.Errorf("the cardinality max was expected to be (%d), (%d) returned", 1, max)
		return
	}
}

func TestApplication_token_withBlockName_withoutCardinality_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(` . myToken`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retToken, retRemaining, err := retAdapter.bytesToToken(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retToken.Name() != "myToken" {
		t.Errorf("the token name is invalid, expected (%s), returned (%s)", "myToken", retToken.Name())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	cardinality := retToken.Cardinality()
	if cardinality.Min() != 1 {
		t.Errorf("the cardinality min was expected to be (%d), (%d) returned", 1, cardinality.Min())
		return
	}

	if !cardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pMax := cardinality.Max()
	max := *pMax
	if max != 1 {
		t.Errorf("the cardinality max was expected to be (%d), (%d) returned", 1, max)
		return
	}
}

func TestApplication_token_withRuleName_withCardinality_Success(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`. MY_RULE [1]`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retToken, retRemaining, err := retAdapter.bytesToToken(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retToken.Name() != "MY_RULE" {
		t.Errorf("the token name is invalid, expected (%s), returned (%s)", "MY_RULE", retToken.Name())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	cardinality := retToken.Cardinality()
	if cardinality.Min() != 1 {
		t.Errorf("the cardinality min was expected to be (%d), (%d) returned", 1, cardinality.Min())
		return
	}

	if !cardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pMax := cardinality.Max()
	max := *pMax
	if max != 1 {
		t.Errorf("the cardinality max was expected to be (%d), (%d) returned", 1, max)
		return
	}
}

func TestApplication_token_withoutBlockName_withoutRuleName_returnsError(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`.___`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToToken(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_token_withoutTokenReferenceByte_returnsError(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`myToken [ 1 ]`), remaining...)

	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToToken(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_token_withoutInput_returnsError(t *testing.T) {
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToToken([]byte{})
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_rule_Success(t *testing.T) {
	expectedName := "MY_RULE"
	expectedValue := []byte(`this " with escape`)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`MY_RULE: "this \" with escape";this is some remaining`)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retRule, retRemaining, err := retAdapter.bytesToRule(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retRule.Name() != expectedName {
		t.Errorf("the name was expected to be %s, %s returned", expectedName, retRule.Name())
		return
	}

	if !bytes.Equal(expectedValue, retRule.Bytes()) {
		t.Errorf("the expected value was (%s), returned (%s)", expectedValue, retRule.Bytes())
		return
	}
}

func TestApplication_rule_withInvalidName_returnsError(t *testing.T) {
	input := []byte(`_MY_RULE: "this \" with escape";this is some remaining`)
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_rule_withoutValue_returnsError(t *testing.T) {
	input := []byte(`MY_RULE: "";this is some remaining`)
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_cardinality_withoutMax_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`[1, ]this is some remaining`)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestApplication_cardinality_withMax_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedMax := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`[ 1, 1 ] this is some remaining`)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if !retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pRetMax := retCardinality.Max()
	if *pRetMax != expectedMax {
		t.Errorf("the max was expected to be %d, %d returned", expectedMax, *pRetMax)
		return
	}
}

func TestApplication_cardinality_withZeroPlus_Success(t *testing.T) {
	expectedMin := uint(0)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`*this is some remaining`)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestApplication_cardinality_withOnePlus_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`+this is some remaining`)

	retAdapter := NewParserAdapter().(*parserAdapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestApplication_cardinality_withInvalidInput_returnsError(t *testing.T) {
	input := []byte(`this is some invalid input`)
	retAdapter := NewParserAdapter().(*parserAdapter)
	_, _, err := retAdapter.bytesToCardinality(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
