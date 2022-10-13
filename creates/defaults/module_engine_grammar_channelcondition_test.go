package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_engineGrammarChannelCondition_withPrevious_withNext_Success(t *testing.T) {
	script := `
        module @engineGrammarSuite;;
        module @containerList;;
        module @engineGrammarSuites;;
        module @engineGrammarValue;;
        module @engineGrammarCardinality;;
        module @engineGrammarElement;;
        module @engineGrammarLine;;
        module @engineGrammarBlock;;
        module @engineGrammarToken;;
        module @engineGrammarChannelCondition;;

        <- $channelCondition;;

        // suite app:
		@engineGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:$valid $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $valid:$invalid $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        @containerList $suitesListApp;;
        attach $first:$0 $suitesListApp;;
        attach $second:$1 $suitesListApp;;
        $list = execute $suitesListApp;;

        // suites:
        @engineGrammarSuites $suitesApp;;
        attach $list:$suites $suitesApp;;
        $suites = execute $suitesApp;;

        // value app:
		@engineGrammarValue $valueApp;;
		$number = 157;;
        $name = myName;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

        // cardinality:
		@engineGrammarCardinality $cardinalityApp;;
        $myMin = 1;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@engineGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@engineGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@engineGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName = myToken;;
		@engineGrammarToken $tokenApp;;
        attach $tokenName:$name $tokenApp;;
        attach $suites:$suites $tokenApp;;
        attach $block:$block $tokenApp;;
        $token = execute $tokenApp;;

        // channel condition:
        @engineGrammarChannelCondition $channelConditionApp;;
        attach $token:$previous $channelConditionApp;;
        attach $token:$next $channelConditionApp;;
        $channelCondition = execute $channelConditionApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["channelCondition"].(grammars.ChannelCondition); ok {
		if !ins.HasPrevious() {
			t.Errorf("the channelCondition was expected to contain a Previous token")
			return
		}

		if !ins.HasNext() {
			t.Errorf("the channelCondition was expected to contain a Next token")
			return
		}
		return
	}

	t.Errorf("the channelCondition output was expected to contain a ChannelCondition instance")
	return
}

func TestModule_engineGrammarChannelCondition_withPrevious__Success(t *testing.T) {
	script := `
        module @engineGrammarSuite;;
        module @containerList;;
        module @engineGrammarSuites;;
        module @engineGrammarValue;;
        module @engineGrammarCardinality;;
        module @engineGrammarElement;;
        module @engineGrammarLine;;
        module @engineGrammarBlock;;
        module @engineGrammarToken;;
        module @engineGrammarChannelCondition;;

        <- $channelCondition;;

        // suite app:
		@engineGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:$valid $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $valid:$invalid $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        @containerList $suitesListApp;;
        attach $first:$0 $suitesListApp;;
        attach $second:$1 $suitesListApp;;
        $list = execute $suitesListApp;;

        // suites:
        @engineGrammarSuites $suitesApp;;
        attach $list:$suites $suitesApp;;
        $suites = execute $suitesApp;;

        // value app:
		@engineGrammarValue $valueApp;;
		$number = 157;;
        $name = myName;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

        // cardinality:
		@engineGrammarCardinality $cardinalityApp;;
        $myMin = 1;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@engineGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@engineGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@engineGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName = myToken;;
		@engineGrammarToken $tokenApp;;
        attach $tokenName:$name $tokenApp;;
        attach $suites:$suites $tokenApp;;
        attach $block:$block $tokenApp;;
        $token = execute $tokenApp;;

        // channel condition:
        @engineGrammarChannelCondition $channelConditionApp;;
        attach $token:$previous $channelConditionApp;;
        $channelCondition = execute $channelConditionApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["channelCondition"].(grammars.ChannelCondition); ok {
		if !ins.HasPrevious() {
			t.Errorf("the channelCondition was expected to contain a Previous token")
			return
		}

		if ins.HasNext() {
			t.Errorf("the channelCondition was expected to NOT contain a Next token")
			return
		}
		return
	}

	t.Errorf("the channelCondition output was expected to contain a ChannelCondition instance")
	return
}

func TestModule_engineGrammarChannelCondition_withNext_Success(t *testing.T) {
	script := `
        module @engineGrammarSuite;;
        module @containerList;;
        module @engineGrammarSuites;;
        module @engineGrammarValue;;
        module @engineGrammarCardinality;;
        module @engineGrammarElement;;
        module @engineGrammarLine;;
        module @engineGrammarBlock;;
        module @engineGrammarToken;;
        module @engineGrammarChannelCondition;;

        <- $channelCondition;;

        // suite app:
		@engineGrammarSuite $suiteApp;;

        // first suite:
		$valid = 157;;
		attach $valid:$valid $suiteApp;;
        $first = execute $suiteApp;;

        // second suite:
		$invalid = 234;;
		attach $valid:$invalid $suiteApp;;
        $second = execute $suiteApp;;

        // suites list:
        @containerList $suitesListApp;;
        attach $first:$0 $suitesListApp;;
        attach $second:$1 $suitesListApp;;
        $list = execute $suitesListApp;;

        // suites:
        @engineGrammarSuites $suitesApp;;
        attach $list:$suites $suitesApp;;
        $suites = execute $suitesApp;;

        // value app:
		@engineGrammarValue $valueApp;;
		$number = 157;;
        $name = myName;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

        // cardinality:
		@engineGrammarCardinality $cardinalityApp;;
        $myMin = 1;;
		attach $myMin:$min $cardinalityApp;;
        $cardinality = execute $cardinalityApp;;

        // element:
		@engineGrammarElement $elementApp;;
        attach $cardinality:$cardinality $elementApp;;
        attach $value:$value $elementApp;;
        $element = execute $elementApp;;

        // elements:
        @containerList $listApp;;
        attach $element:$0 $listApp;;
        $elements = execute $listApp;;

        // line:
		@engineGrammarLine $lineApp;;
        attach $elements:$elements $lineApp;;
        $line = execute $lineApp;;

        // lines:
        attach $line:$0 $listApp;;
        $lines = execute $listApp;;

        // block:
		@engineGrammarBlock $blockApp;;
        attach $lines:$lines $blockApp;;
        $block = execute $blockApp;;

        // token:
        $tokenName = myToken;;
		@engineGrammarToken $tokenApp;;
        attach $tokenName:$name $tokenApp;;
        attach $suites:$suites $tokenApp;;
        attach $block:$block $tokenApp;;
        $token = execute $tokenApp;;

        // channel condition:
        @engineGrammarChannelCondition $channelConditionApp;;
        attach $token:$next $channelConditionApp;;
        $channelCondition = execute $channelConditionApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["channelCondition"].(grammars.ChannelCondition); ok {
		if ins.HasPrevious() {
			t.Errorf("the channelCondition was expected to NOT contain a Previous token")
			return
		}

		if !ins.HasNext() {
			t.Errorf("the channelCondition was expected to contain a Next token")
			return
		}
		return
	}

	t.Errorf("the channelCondition output was expected to contain a ChannelCondition instance")
	return
}
