package defaults

import (
	"testing"

	"github.com/steve-care-software/syntax/applications/engines"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

func TestModule_engineGrammarChannels_Success(t *testing.T) {
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
        module @engineGrammarChannels;;

        <- $channels;;

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

        // channels list:
        attach $channel:$0 $listApp;;
        $channelsList = execute $listApp;;

        // channels:
        @engineGrammarChannels $channelsApp;;
        attach $channelsList:$channels $channelsApp;;
        $channels = execute $channelsApp;;

	`

	output, _, err := engines.NewApplication(NewApplication()).ParseThenInterpret(map[string]interface{}{}, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins, ok := output["channels"].(grammars.Channels); ok {
		list := ins.List()
		if len(list) != 1 {
			t.Errorf("%d channels were expected, %d returned", 1, len(list))
			return
		}

		return
	}

	t.Errorf("the channels output was expected to contain a Channels instance")
	return
}
