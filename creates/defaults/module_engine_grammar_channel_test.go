package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_engineGrammarChannel_Success(t *testing.T) {
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
        module @engineGrammarChannel;;

        <- $channel;;

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

		// cast to uint app:
		module @castToUint;;
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:$value $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
		@engineGrammarValue $valueApp;;
        $name = myName;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:$value $castToUintApp;;
		$myMin = execute $castToUintApp;;

        // cardinality:
		@engineGrammarCardinality $cardinalityApp;;
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

        // channel:
        $channelName = myChannel;;
        @engineGrammarChannel $channelApp;;
        attach $channelName:$name $channelApp;;
        attach $token:$token $channelApp;;
        $channel = execute $channelApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["channel"].(grammars.Channel); ok {
		name := ins.Name()
		if name != "myChannel" {
			t.Errorf("the name was expected to be '%s', '%s' returned", "myChannel", name)
			return
		}

		if ins.HasCondition() {
			t.Errorf("the channel was expected to NOT contain a condition")
			return
		}

		return
	}

	t.Errorf("the channel output was expected to contain a Channel instance")
	return
}

func TestModule_engineGrammarChannel_withCondition_Success(t *testing.T) {
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
        module @engineGrammarChannel;;
        module @engineGrammarChannelCondition;;

        <- $channel;;

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

		// cast to uint app:
		module @castToUint;;
		@castToUint $castToUintApp;;

		// number casting to uint:
		$myNumber = 157;;
		attach $myNumber:$value $castToUintApp;;
		$number = execute $castToUintApp;;

        // value app:
		@engineGrammarValue $valueApp;;
        $name = myName;;
		attach $number:$number $valueApp;;
		attach $name:$name $valueApp;;
        $value = execute $valueApp;;

		// min casting to uint:
		$myMinStr = 1;;
		attach $myMinStr:$value $castToUintApp;;
		$myMin = execute $castToUintApp;;

        // cardinality:
		@engineGrammarCardinality $cardinalityApp;;
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

        // channel:
        $channelName = myChannel;;
        @engineGrammarChannel $channelApp;;
        attach $channelName:$name $channelApp;;
        attach $token:$token $channelApp;;
        attach $channelCondition:$condition $channelApp;;
        $channel = execute $channelApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["channel"].(grammars.Channel); ok {
		name := ins.Name()
		if name != "myChannel" {
			t.Errorf("the name was expected to be '%s', '%s' returned", "myChannel", name)
			return
		}

		if !ins.HasCondition() {
			t.Errorf("the channel was expected to contain a condition")
			return
		}

		return
	}

	t.Errorf("the channel output was expected to contain a Channel instance")
	return
}
