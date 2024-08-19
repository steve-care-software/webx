package grammars

import (
	"bytes"
	"testing"
)

func TestApplication_blocks_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myFirst:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement;mySecond:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement---myFirst:.myElement.mySecond:@.myThird.mySecondTest:.myFourth.myTest:@.myElement.;`), remaining...)

	application := NewApplication().(*application)
	retBlocks, retRemaining, err := application.bytesToBlocks(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	list := retBlocks.List()
	if len(list) != 2 {
		t.Errorf("the block was expected to contain %d suite instances, %d returned", 4, len(list))
		return
	}
}

func TestApplication_blocks_withoutBlocks_returnsError(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(``), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToBlocks(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withSuites_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myBlock:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement---myFirst:.myElement.mySecond:@.myThird.mySecondTest:.myFourth.myTest:@.myElement.;`), remaining...)

	application := NewApplication().(*application)
	retBlock, retRemaining, err := application.bytesToBlock(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retBlock.Name() != "myBlock" {
		t.Errorf("the block name was expected to be (%s), (%s) returned", "myBlock", retBlock.Name())
		return
	}

	list := retBlock.Lines().List()
	if len(list) != 3 {
		t.Errorf("the lines was expected to contain %d suite instances, %d returned", 3, len(list))
		return
	}

	if !retBlock.HasSuites() {
		t.Errorf("the block was expected to contain suites")
		return
	}

	suitesList := retBlock.Suites().List()
	if len(suitesList) != 4 {
		t.Errorf("the suites was expected to contain %d suite instances, %d returned", 4, len(suitesList))
		return
	}
}

func TestApplication_block_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myBlock:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement;`), remaining...)

	application := NewApplication().(*application)
	retBlock, retRemaining, err := application.bytesToBlock(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retBlock.Name() != "myBlock" {
		t.Errorf("the block name was expected to be (%s), (%s) returned", "myBlock", retBlock.Name())
		return
	}

	if retBlock.HasSuites() {
		t.Errorf("the block was expected to NOT contain suites")
		return
	}

	list := retBlock.Lines().List()
	if len(list) != 3 {
		t.Errorf("the lines was expected to contain %d suite instances, %d returned", 3, len(list))
		return
	}
}

func TestApplication_block_withoutSuffix_returnsError(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myBlock:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement---myFirst:.myElement.mySecond:@.myThird.mySecondTest:.myFourth.myTest:@.myElement.`), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withoutSuffix_withoutRemaining_returnsError(t *testing.T) {
	input := []byte(`myBlock:.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement---myFirst:.myElement.mySecond:@.myThird.mySecondTest:.myFourth.myTest:@.myElement.`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withoutSuffix_withInvalidBlockDefinition_returnsError(t *testing.T) {
	input := []byte(`myBlock.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement---myFirst:.myElement.mySecond:@.myThird.mySecondTest:.myFourth.myTest:@.myElement.;`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_block_withoutSuffix_withoutLines_returnsError(t *testing.T) {
	input := []byte(`myBlock:---myFirst:.myElement.mySecond:@.myThird.mySecondTest:.myFourth.myTest:@.myElement.;`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToBlock(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_suites_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`---myTest:.myElement.myTest:@.myElement.`), remaining...)

	application := NewApplication().(*application)
	retSuites, retRemaining, err := application.bytesToSuites(input)
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
	remaining := []byte("!this is some remaining")
	input := append([]byte(``), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToSuites(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myTest:.myElement.`), remaining...)

	application := NewApplication().(*application)
	retSuite, retRemaining, err := application.bytesToSuite(input)
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

	if retSuite.Element().Name() != "myElement" {
		t.Errorf("the suite's element name was expected to be (%s), (%s) returned", "myElement", retSuite.Element().Name())
		return
	}

	if retSuite.IsFail() {
		t.Errorf("the suite was expected to NOT fail")
		return
	}
}

func TestApplication_suite_isFail_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myTest:@.myElement.`), remaining...)

	application := NewApplication().(*application)
	retSuite, retRemaining, err := application.bytesToSuite(input)
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

	if retSuite.Element().Name() != "myElement" {
		t.Errorf("the suite's element name was expected to be (%s), (%s) returned", "myElement", retSuite.Element().Name())
		return
	}

	if !retSuite.IsFail() {
		t.Errorf("the suite was expected to fail")
		return
	}
}

func TestApplication_suite_withInvalidElement_returnsError(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myTest:myElement`), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_withInvalidBlockNameDefinition_returnsError(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`#myTest:.myElement`), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_withoutSuiteLineSuffix_returnsError(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myTest:.myElement`), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_suite_withoutSuiteLineSuffix_withoutRemainingBytes_returnsError(t *testing.T) {
	input := []byte(`myTest:.myElement`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestApplication_lines_withOneLine_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT`), remaining...)

	application := NewApplication().(*application)
	retLines, retRemaining, err := application.bytesToLines(input)
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
	remaining := []byte("!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth|.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement`), remaining...)

	application := NewApplication().(*application)
	retLines, retRemaining, err := application.bytesToLines(input)
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
	remaining := []byte("!this is some remaining")
	input := append([]byte(`not a line`), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToLines(input)
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_line_withExecution_withReplacement_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT`), remaining...)

	application := NewApplication().(*application)
	retLine, retRemaining, err := application.bytesToLine(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retLine.HasReplacement() {
		t.Errorf("the execution was expected to contain a replacement")
		return
	}

	if !retLine.HasExecution() {
		t.Errorf("the execution was expected to contain an execution")
		return
	}
}

func TestApplication_line_withExecution_withReplacement_reversed_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth`), remaining...)

	application := NewApplication().(*application)
	retLine, retRemaining, err := application.bytesToLine(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retLine.HasReplacement() {
		t.Errorf("the execution was expected to contain a replacement")
		return
	}

	if !retLine.HasExecution() {
		t.Errorf("the execution was expected to contain an execution")
		return
	}
}

func TestApplication_line_withExecution_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth`), remaining...)

	application := NewApplication().(*application)
	retLine, retRemaining, err := application.bytesToLine(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retLine.HasReplacement() {
		t.Errorf("the execution was expected to NOT contain a replacement")
		return
	}

	if !retLine.HasExecution() {
		t.Errorf("the execution was expected to contain an execution")
		return
	}
}

func TestApplication_line_withReplacement_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]-.myReplacement`), remaining...)

	application := NewApplication().(*application)
	retLine, retRemaining, err := application.bytesToLine(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retLine.HasReplacement() {
		t.Errorf("the execution was expected to contain a replacement")
		return
	}

	if retLine.HasExecution() {
		t.Errorf("the execution was expected to NOT contain an execution")
		return
	}
}

func TestApplication_withoutTokens_returnsError(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth-.MY_REPLACEMENT`), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToLine(input)
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_execution_withElements_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth`), remaining...)

	application := NewApplication().(*application)
	retExecution, retRemaining, err := application.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retExecution.HasElements() {
		t.Errorf("the execution was expected to contain elements")
		return
	}

	list := retExecution.Elements().List()
	if len(list) != 5 {
		t.Errorf("the tokens list was expected to contain %d tokens, %d returned", 5, len(list))
		return
	}
}

func TestApplication_execution_withElements_withoutRemaining_Success(t *testing.T) {
	remaining := []byte("")
	input := append([]byte(`myFuncName_secondSection.myFirst.mySecond.myThird.myFourth.myFifth`), remaining...)

	application := NewApplication().(*application)
	retExecution, retRemaining, err := application.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if !retExecution.HasElements() {
		t.Errorf("the execution was expected to contain elements")
		return
	}

	list := retExecution.Elements().List()
	if len(list) != 5 {
		t.Errorf("the tokens list was expected to contain %d elements, %d returned", 5, len(list))
		return
	}
}

func TestApplication_execution_withoutElements_Success(t *testing.T) {
	remaining := []byte("!this is some remaining")
	input := append([]byte(`myFuncName_secondSection`), remaining...)

	application := NewApplication().(*application)
	retExecution, retRemaining, err := application.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retExecution.HasElements() {
		t.Errorf("the execution was expected to NOT contain elements")
		return
	}
}

func TestApplication_execution_withoutElements_withoutRemaining_Success(t *testing.T) {
	remaining := []byte("")
	input := append([]byte(`myFuncName_secondSection`), remaining...)

	application := NewApplication().(*application)
	retExecution, retRemaining, err := application.bytesToExecution(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retExecution.HasElements() {
		t.Errorf("the execution was expected to NOT contain elements")
		return
	}
}

func TestApplication_tokens_Success(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`.myFirst[1].mySecond*.myThird+.myFourth.myFifth[1,]`), remaining...)

	application := NewApplication().(*application)
	retToken, retRemaining, err := application.bytesToTokens(input)
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

	application := NewApplication().(*application)
	retToken, retRemaining, err := application.bytesToToken(input)
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
	remaining := []byte("!this is some remaining")
	input := append([]byte(`.myToken`), remaining...)

	application := NewApplication().(*application)
	retToken, retRemaining, err := application.bytesToToken(input)
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
	input := append([]byte(`.MY_RULE[1]`), remaining...)

	application := NewApplication().(*application)
	retToken, retRemaining, err := application.bytesToToken(input)
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

	application := NewApplication().(*application)
	_, _, err := application.bytesToToken(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_token_withoutTokenReferenceByte_returnsError(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`myToken[1]`), remaining...)

	application := NewApplication().(*application)
	_, _, err := application.bytesToToken(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_token_withoutInput_returnsError(t *testing.T) {
	application := NewApplication().(*application)
	_, _, err := application.bytesToToken([]byte{})
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_rule_Success(t *testing.T) {
	expectedName := "MY_RULE"
	expectedValue := []byte(`this " with escape`)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`MY_RULE:"this \" with escape"this is some remaining`)

	application := NewApplication().(*application)
	retRule, retRemaining, err := application.bytesToRule(input)
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
	input := []byte(`_MY_RULE:"this \" with escape"this is some remaining`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_rule_withoutValue_returnsError(t *testing.T) {
	input := []byte(`MY_RULE:""this is some remaining`)
	application := NewApplication().(*application)
	_, _, err := application.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestApplication_cardinality_withoutMax_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`[1,]this is some remaining`)

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
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
	input := []byte(`[1,1]this is some remaining`)

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
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

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
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

	application := NewApplication().(*application)
	retCardinality, retRemaining, err := application.bytesToCardinality(input)
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
	application := NewApplication().(*application)
	_, _, err := application.bytesToCardinality(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
