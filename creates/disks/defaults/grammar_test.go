package defaults

import (
	"bytes"
	"fmt"
	"testing"

	grammar_applications "github.com/steve-care-software/webx/applications/grammars"
	selection_applications "github.com/steve-care-software/webx/applications/selections"
	"github.com/steve-care-software/webx/domain/grammars"
)

func TestGrammar_singleCharacters_Success(t *testing.T) {
	valid := [][]byte{
		[]byte("abcdefghijklmnopqrstuvwxyz"),
		[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		[]byte("0123456789"),
	}

	createGrammarApp := NewApplication(
		bitrateForTests,
		basePathForTests,
		delimiterForTests,
		extensionForTests,
	).Grammar().(*grammar)

	rootsList := []grammars.Token{
		createGrammarApp.LowerCaseLetter(),
		createGrammarApp.UpperCaseLetter(),
		createGrammarApp.AnyCaseLetter(),
		createGrammarApp.AnyNumber(),
	}

	channels := createGrammarApp.Channels()
	grammarApp := grammar_applications.NewApplication()
	selectionApp := selection_applications.NewApplication()
	for idx, oneRange := range valid {
		grammarIns, err := grammars.NewBuilder().WithRoot(rootsList[idx]).WithChannels(channels).Now()
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		for _, oneElement := range oneRange {
			input := []byte{
				oneElement,
			}

			tree, err := grammarApp.Execute(grammarIns, input)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			converted, err := selectionApp.Convert(tree, false)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retBytes := converted.Bytes()
			if bytes.Compare(input, retBytes) != 0 {
				t.Errorf("the returned bytes is invalid (index: %d), %v", idx, retBytes)
				return
			}
		}

	}
}

func TestGrammar_multipleCharacters_Success(t *testing.T) {
	valid := [][][]byte{
		[][]byte{
			[]byte("myName"),
			[]byte("01"),
		},
		[][]byte{
			[]byte("module"),
		},
		[][]byte{
			[]byte(";;"),
		},
		[][]byte{
			[]byte("@"),
		},
		[][]byte{
			[]byte("@myModule"),
			[]byte("@01"),
		},
		[][]byte{
			[]byte("module @myModule;;"),
			[]byte("module @01;;"),
		},
		[][]byte{
			[]byte("$"),
		},
		[][]byte{
			[]byte("$myName"),
			[]byte("$13"),
		},
		[][]byte{
			[]byte("@myModule $myApplication;;"),
			[]byte("@12 $myApplication;;"),
			[]byte("@myModule $01;;"),
			[]byte("@12 $1;;"),
		},
		[][]byte{
			[]byte("->"),
		},
		[][]byte{
			[]byte("<-"),
		},
		[][]byte{
			[]byte("-> $myInput;;"),
			[]byte("-> $12;;"),
		},
		[][]byte{
			[]byte("<- $myInput;;"),
			[]byte("<- $12;;"),
		},
		[][]byte{
			[]byte("="),
		},
		[][]byte{
			[]byte("attach"),
		},
		[][]byte{
			[]byte(":"),
		},
		[][]byte{
			[]byte("attach $current:$local $myApp;;"),
			[]byte("attach $1:$local $myApp;;"),
			[]byte("attach $current:$2 $myApp;;"),
			[]byte("attach $current:$local $3;;"),
			[]byte("attach $1:$2 $3;;"),
		},
		[][]byte{
			[]byte("$myVariable = ANY VALUE EXCEPT \\;; NON-ESCAPED SEMI-COLON;;"),
			[]byte("$myOutput = execute $myAppVariable;;"),
			[]byte("$myValue = $myOtherVariable;;"),
		},
		[][]byte{
			[]byte("module @myModule;;"),
			[]byte("@myModule $myApplication;;"),
			[]byte("-> $myInput;;"),
			[]byte("<- $myOutput;;"),
			[]byte("$myVariable = ANY VALUE EXCEPT \\;; NON-ESCAPED SEMI-COLON;;"),
			[]byte("attach $myDataVariable:$data $myAppVariable;;"),
			[]byte("$myOutput = execute $myAppVariable;;"),
			[]byte("$myValue = $myOtherVariable;;"),
		},
		[][]byte{
			[]byte("-> $script;;"),
			[]byte(`
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


				module @containerMapWithStringKeynames;;
				@containerMapWithStringKeynames $paramsApp;;
				$nameStr = dollarSign;;
				$valueStr = 36;
				attach $nameStr:$name $paramsApp;;
				attach $valueStr:$number $valueApp;;
				$params = execute $paramsApp;;

				module @parseThenInterpret;;
				@parseThenInterpret $interpreterApp;;
				attach $params:$params $interpreterApp;;
				attach $createGrammarValueCode:$script $interpreterApp;;
				$output = execute $interpreterApp;;

			`),
		},
	}

	createGrammarApp := NewApplication(
		bitrateForTests,
		basePathForTests,
		delimiterForTests,
		extensionForTests,
	).Grammar().(*grammar)

	rootsList := []grammars.Token{
		createGrammarApp.Name(),
		createGrammarApp.ModuleString(),
		createGrammarApp.EndOfLine(),
		createGrammarApp.CommercialAChar(),
		createGrammarApp.ModuleName(),
		createGrammarApp.ModuleDeclaration(),
		createGrammarApp.DollarSignChar(),
		createGrammarApp.VariableName(),
		createGrammarApp.ApplicationDeclaration(),
		createGrammarApp.InputDirection(),
		createGrammarApp.OutputDirection(),
		createGrammarApp.InputParameter(),
		createGrammarApp.OutputParameter(),
		createGrammarApp.EqualChar(),
		createGrammarApp.AttachString(),
		createGrammarApp.SemiColonChar(),
		createGrammarApp.Attach(),
		createGrammarApp.Assignment(),
		createGrammarApp.Instruction(),
		createGrammarApp.Root(),
	}

	channels := createGrammarApp.Channels()
	grammarApp := grammar_applications.NewApplication()
	selectionApp := selection_applications.NewApplication()
	for idx, oneElement := range valid {
		for _, input := range oneElement {
			grammarIns, err := grammars.NewBuilder().WithRoot(rootsList[idx]).WithChannels(channels).Now()
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			tree, err := grammarApp.Execute(grammarIns, input)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			converted, err := selectionApp.Convert(tree, true)
			if err != nil {
				t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
				return
			}

			retBytes := converted.Bytes()
			inputWithoutEscapeChars := bytes.ReplaceAll(input, []byte("\\"), []byte{})
			if bytes.Compare(inputWithoutEscapeChars, retBytes) != 0 {
				t.Errorf("the returned bytes is invalid (index: %d), expected: '%s', returned: '%s'", idx, input, retBytes)
				return
			}
		}

	}
}

func TestGrammar_Success(t *testing.T) {
	grammarIns, err := NewApplication(bitrateForTests, basePathForTests, delimiterForTests, extensionForTests).Grammar().Execute()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	grammarApp := grammar_applications.NewApplication()
	coverages, err := grammarApp.Coverages(grammarIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if coverages != nil {
		list := coverages.List()
		for _, oneCoverage := range list {
			tokenName := oneCoverage.Token().Name()
			executionsList := oneCoverage.Executions().List()
			for execIdx, oneExecution := range executionsList {
				expectation := oneExecution.Expectation()
				content := expectation.Content()
				result := oneExecution.Result()
				expected := expectation.IsValid()

				path := fmt.Sprintf("%s.%d=%v", tokenName, execIdx, string(content))
				if expected {
					if result.IsTree() {
						if result.Tree().Block().HasSuccessful() {
							continue
						}

						t.Errorf("the test suite expected to be successful, but there was no successful result, path: %s", path)
						continue
					}

					t.Errorf("the test suite expected to be successful, but there was an error while executing the test suite (path: %s): %s", path, result.Error())
					continue
				}

				if result.IsError() {
					continue
				}

				t.Errorf("the test suite expected to be unsuccessful, but the result was successful, path: %s", path)
			}

		}
	}

	uncovered, err := grammarApp.Uncovered(grammarIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	for tokenName, lines := range uncovered {
		for lineIdx, line := range lines {
			for elIdx, element := range line {
				t.Errorf("token: %s, line: %d, element: %d, name: %s - is uncovered", tokenName, lineIdx, elIdx, element)
			}
		}
	}
}
