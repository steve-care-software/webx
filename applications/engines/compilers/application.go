package compilers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/applications/engines/creates"
	criteria_application "github.com/steve-care-software/syntax/applications/engines/criterias"
	grammar_application "github.com/steve-care-software/syntax/applications/engines/grammars"
	interpreter_application "github.com/steve-care-software/syntax/applications/engines/interpreters"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/compilers/outputs"
)

type application struct {
	grammarApp     grammar_application.Application
	criteriaApp    criteria_application.Application
	interpreterApp interpreter_application.Application
	createApp      creates.Application
	outputBuilder  outputs.Builder
}

func createApplication(
	grammarApp grammar_application.Application,
	criteriaApp criteria_application.Application,
	interpreterApp interpreter_application.Application,
	createApp creates.Application,
	outputBuilder outputs.Builder,
) Application {
	out := application{
		grammarApp:     grammarApp,
		criteriaApp:    criteriaApp,
		interpreterApp: interpreterApp,
		createApp:      createApp,
		outputBuilder:  outputBuilder,
	}

	return &out
}

// Execute eecutes a compiler
func (app *application) Execute(compiler compilers.Compiler, script []byte) (outputs.Output, error) {
	remaining := script
	elementsList := compiler.Elements().List()
	outputs := map[string]interface{}{}
	for _, oneElement := range elementsList {
		grammar := oneElement.Grammar()
		tree, err := app.grammarApp.Execute(grammar, remaining)
		if err != nil {
			return nil, err
		}

		input := map[string]interface{}{}
		parameters := oneElement.Parameters().List()
		for _, oneParameter := range parameters {
			keyname := oneParameter.Name()
			value := oneParameter.Value()
			if value.IsConstant() {
				input[keyname] = value.Constant()
				continue
			}

			criteria := value.Criteria()
			found, err := app.criteriaApp.Execute(criteria, tree)
			if err != nil {
				return nil, err
			}

			input[keyname] = found
		}

		program := oneElement.Program()
		output, err := app.interpreterApp.Execute(input, program)
		if err != nil {
			return nil, err
		}

		for name, value := range output {
			outputs[name] = value
		}

		remaining = tree.Remaining()
	}

	builder := app.outputBuilder.Create()
	if remaining != nil {
		builder.WithRemaining(remaining)
	}

	if compiler.HasOutputs() {
		names := compiler.Outputs()
		compilerOutputs := map[string]interface{}{}
		for _, oneName := range names {
			if value, ok := outputs[oneName]; ok {
				compilerOutputs[oneName] = value
				continue
			}

			str := fmt.Sprintf("the compiler requested an output variable (name %s) that is undeclared after executing the compiler", oneName)
			return nil, errors.New(str)
		}

		builder.WithValues(compilerOutputs)
	}

	return builder.Now()
}
