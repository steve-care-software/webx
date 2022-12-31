package instances

import (
	"testing"

	grammar_application "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	selector_application "github.com/steve-care-software/webx/selectors/applications"
)

func TestSelector_Success(t *testing.T) {
	grammarIns := newGrammar()
	selectorIns := newSelector()

	grammarApp := grammar_application.NewApplication()
	selectorApp := selector_application.NewApplication()

	script := `
        module @myModule:0;;
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

		attach $myInput:0 $myApp;;

		execute $myApp;;
		$appExec = execute $myApp;;
	`
	treeIns, err := grammarApp.Execute(grammarIns, []byte(script))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if treeIns.HasRemaining() {
		t.Errorf("the tree was expected to not contain remaining data")
		return
	}

	instructionsIns, isValid, remaining, err := selectorApp.Execute(selectorIns, treeIns)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !isValid {
		t.Errorf("the selection was expected to be valid")
		return
	}

	if len(remaining) > 0 {
		t.Errorf("the remaining data was expected to be empty")
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

			moduleIns := ins.Module()
			name := string(moduleIns.Name())
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

			/*application := ins.Application()
			module := string(application.Module())
			if module != "myModule" {
				t.Errorf("the module name was expected to be '%s', '%s' returned", "myModule", module)
				return
			}

			name := string(application.Name())
			if name != "myApp" {
				t.Errorf("the parameter's name was expected to be '%s', '%s' returned", "myApp", name)
				return
			}*/

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

			if variable.Target() != 0 {
				t.Errorf("the assignment's target variable was expected to be '%d', '%d' returned", 0, variable.Target())
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
