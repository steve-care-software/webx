package defaults

const bitrateForTests = 1024
const basePathForTests = "./test_files"
const delimiterForTests = "."
const extensionForTests = "identity"

const fullScriptForTests = `
    -> $script;;
    <- $output;;

    $createGrammarValueCode = {
        -> $name;;
        -> $numberStr;;
        <- $value;;

        module @castToInt;;
        @castToInt $castToIntApp;;
        attach $numberStr:$value $castToIntApp;;
        $number = execute $castToIntApp;;

        module @newGrammarValue;;
        @newGrammarValue $valueApp;;
        attach $number:$number $valueApp;;
        attach $name:$name $valueApp;;
        $value = execute $valueApp;;
    };;


    $justAssignment = $script;;

    module @containerMapWithStringKeynames;;
    @containerMapWithStringKeynames $paramsApp;;
    $nameStr = dollarSign;;
    $valueStr = 36;;
    attach $nameStr:$name $paramsApp;;
    attach $valueStr:$number $valueApp;;
    $params = execute $paramsApp;;

    module @parseThenInterpret;;
    @parseThenInterpret $interpreterApp;;
    attach $params:$params $interpreterApp;;
    attach $createGrammarValueCode:$script $interpreterApp;;
    $output = execute $interpreterApp;;
    execute $interpreterApp;;
`
