package defaults

import (
	"testing"

	grammar_application "github.com/steve-care-software/webx/applications/grammars"
	selector_application "github.com/steve-care-software/webx/applications/selectors"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/webx/domain/grammars/values"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
	"github.com/steve-care-software/webx/domain/selectors"
)

func TestSelector_Success(t *testing.T) {
	grammarCreateIns, err := createGrammar(
		grammars.NewBuilder(),
		grammars.NewChannelsBuilder(),
		grammars.NewChannelBuilder(),
		grammars.NewInstanceBuilder(),
		grammars.NewEverythingBuilder(),
		grammars.NewTokensBuilder(),
		grammars.NewTokenBuilder(),
		grammars.NewSuitesBuilder(),
		grammars.NewSuiteBuilder(),
		grammars.NewBlockBuilder(),
		grammars.NewLineBuilder(),
		grammars.NewElementBuilder(),
		grammar_values.NewBuilder(),
		cardinalities.NewBuilder(),
	).Execute()

	if err != nil {
		t.Errorf("the error was expecte to be nil, error returned: %s", err.Error())
		return
	}

	selectorIns, err := createSelector(
		selectors.NewBuilder(),
		selectors.NewSelectorFnBuilder(),
		selectors.NewTokenBuilder(),
		selectors.NewElementBuilder(),
		selectors.NewInsideBuilder(),
		selectors.NewFetchersBuilder(),
		selectors.NewFetcherBuilder(),
		selectors.NewContentFnBuilder(),
		instructions.NewBuilder(),
		instructions.NewInstructionBuilder(),
		applications.NewBuilder(),
		parameters.NewBuilder(),
		attachments.NewBuilder(),
		attachments.NewVariableBuilder(),
		instructions.NewAssignmentBuilder(),
		instructions.NewValueBuilder(),
	).Execute()

	if err != nil {
		t.Errorf("the error was expecte to be nil, error returned: %s", err.Error())
		return
	}

	grammarApp := grammar_application.NewApplication()
	selectorApp := selector_application.NewApplication()

	script := `
        module @myModule;;
		@myModule $myApp;;

		-> $myInput;;
		<- $myOutput;;

		$constantVariable = this is a constant;;
		$assignedVariable = $myInput;;
		$instructions = {
			-> $first;;
			<- $output;;

			$output = $first;;
		};;

		attach $myInput:$internal $myApp;;

		execute $myApp;;
		$appExec = execute $myApp;;
	`
	treeIns, err := grammarApp.Execute(grammarCreateIns, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if treeIns.HasRemaining() {
		t.Errorf("the tree was expected to not contain remaining data")
		return
	}

	instructionsIns, isValid, err := selectorApp.Execute(selectorIns, treeIns)
	if err != nil {
		t.Errorf("the error was expecte to be nil, error returned: %s", err.Error())
		return
	}

	if !isValid {
		t.Errorf("the selection was expected to be valid")
		return
	}

	list := instructionsIns.(instructions.Instructions).List()
	if len(list) != 10 {
		t.Errorf("%d instructions were expected, %d returned", 10, len(list))
		return
	}

	for idx, ins := range list {
		switch idx {
		case 0:
			if !ins.IsModule() {
				t.Errorf("the instruction (index: %d) was expected to contain a Module", idx)
				return
			}

			name := string(ins.Module())
			if name != "myModule" {
				t.Errorf("the module name was expected to be '%s', '%s' returned", "myModule", name)
				return
			}

			break
		case 1:
			if !ins.IsApplication() {
				t.Errorf("the instruction (index: %d) was expected to contain an Application", idx)
				return
			}

			application := ins.Application()
			module := string(application.Module())
			if module != "myModule" {
				t.Errorf("the module name was expected to be '%s', '%s' returned", "myModule", module)
				return
			}

			name := string(application.Name())
			if name != "myApp" {
				t.Errorf("the parameter's name was expected to be '%s', '%s' returned", "myApp", name)
				return
			}

			break
		case 2:
			if !ins.IsParameter() {
				t.Errorf("the instruction (index: %d) was expected to contain a Parameter", idx)
				return
			}

			parameter := ins.Parameter()
			if !parameter.IsInput() {
				t.Errorf("the parameter was expected to be an input")
				return
			}

			name := string(parameter.Name())
			if name != "myInput" {
				t.Errorf("the parameter's name was expected to be '%s', '%s' returned", "myInput", name)
				return
			}
			break
		case 3:
			if !ins.IsParameter() {
				t.Errorf("the instruction (index: %d) was expected to contain a Parameter", idx)
				return
			}

			parameter := ins.Parameter()
			if parameter.IsInput() {
				t.Errorf("the parameter was expected to be an output")
				return
			}

			name := string(parameter.Name())
			if name != "myOutput" {
				t.Errorf("the parameter's name was expected to be '%s', '%s' returned", "myOutput", name)
				return
			}
			break
		case 4:
			if !ins.IsAssignment() {
				t.Errorf("the instruction (index: %d) was expected to contain an Assignment", idx)
				return
			}

			assignment := ins.Assignment()
			variable := string(assignment.Variable())
			if variable != "constantVariable" {
				t.Errorf("the variable name was expected to be '%s', '%s' returned", "constantVariable", variable)
				return
			}

			value := assignment.Value()
			if !value.IsConstant() {
				t.Errorf("the assignment was expected to be a constant")
				return
			}

			constant := string(value.Constant())
			if constant != " this is a constant" {
				t.Errorf("the value's constant name was expected to be '%s', '%s' returned", " this is a constant", constant)
				return
			}
			break
		case 5:
			if !ins.IsAssignment() {
				t.Errorf("the instruction (index: %d) was expected to contain an Assignment", idx)
				return
			}

			assignment := ins.Assignment()
			variable := string(assignment.Variable())
			if variable != "assignedVariable" {
				t.Errorf("the variable name was expected to be '%s', '%s' returned", "assignedVariable", variable)
				return
			}

			value := assignment.Value()
			if !ins.Assignment().Value().IsVariable() {
				t.Errorf("the assignment was expected to be a variable")
				return
			}

			valueVariable := string(value.Variable())
			if valueVariable != "myInput" {
				t.Errorf("the value's variable was expected to be '%s', '%s' returned", "myInput", valueVariable)
				return
			}
			break
		case 6:
			if !ins.IsAssignment() {
				t.Errorf("the instruction (index: %d) was expected to contain an Assignment", idx)
				return
			}

			assignment := ins.Assignment()
			variable := string(assignment.Variable())
			if variable != "instructions" {
				t.Errorf("the variable name was expected to be '%s', '%s' returned", "instructions", variable)
				return
			}

			value := assignment.Value()
			if !value.IsInstructions() {
				t.Errorf("the assignment was expected to be instructions")
				return
			}

			list := value.Instructions().List()
			if len(list) != 3 {
				t.Errorf("the instructions assigned was expecting %d instructions, %d returned", 3, len(list))
				return
			}
			break
		case 7:
			if !ins.IsAttachment() {
				t.Errorf("the instruction (index: %d) was expected to contain an Attachment", idx)
				return
			}

			attachment := ins.Attachment()
			variable := attachment.Variable()
			current := string(variable.Current())
			if current != "myInput" {
				t.Errorf("the assignment's current variable was expected to be '%s', '%s' returned", "myInput", current)
				return
			}

			target := string(variable.Target())
			if target != "internal" {
				t.Errorf("the assignment's target variable was expected to be '%s', '%s' returned", "internal", target)
				return
			}

			application := string(attachment.Application())
			if application != "myApp" {
				t.Errorf("the assignment's application was expected to be '%s', '%s' returned", "myApp", application)
				return
			}

			break
		case 8:
			if !ins.IsExecution() {
				t.Errorf("the instruction (index: %d) was expected to contain an Execution", idx)
				return
			}

			execution := string(ins.Execution())
			if execution != "myApp" {
				t.Errorf("the execution was expected to be '%s', '%s' returned", "myApp", execution)
				return
			}

			break
		case 9:
			if !ins.IsAssignment() {
				t.Errorf("the instruction (index: %d) was expected to contain an Assignment", idx)
				return
			}

			assignment := ins.Assignment()
			variable := string(assignment.Variable())
			if variable != "appExec" {
				t.Errorf("the variable name was expected to be '%s', '%s' returned", "appExec", variable)
				return
			}

			value := assignment.Value()
			if !value.IsExecution() {
				t.Errorf("the assignment was expected to be an execution")
				return
			}

			execution := string(value.Execution())
			if execution != "myApp" {
				t.Errorf("the assignment's execution was expected to be '%s', '%s' returned", "myApp", execution)
				return
			}
			break
		}
	}

}
