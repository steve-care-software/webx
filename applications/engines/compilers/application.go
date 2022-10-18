package compilers

import (
	"bytes"

	"github.com/steve-care-software/syntax/applications/engines/creates"
	criteria_application "github.com/steve-care-software/syntax/applications/engines/criterias"
	grammar_application "github.com/steve-care-software/syntax/applications/engines/grammars"
	programs_application "github.com/steve-care-software/syntax/applications/engines/programs"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/compilers/outputs"
)

type application struct {
	grammarApp    grammar_application.Application
	criteriaApp   criteria_application.Application
	programApp    programs_application.Application
	createApp     creates.Application
	outputBuilder outputs.Builder
}

func createApplication(
	grammarApp grammar_application.Application,
	criteriaApp criteria_application.Application,
	programApp programs_application.Application,
	createApp creates.Application,
	outputBuilder outputs.Builder,
) Application {
	out := application{
		grammarApp:    grammarApp,
		criteriaApp:   criteriaApp,
		programApp:    programApp,
		createApp:     createApp,
		outputBuilder: outputBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(compiler compilers.Compiler, script []byte) (outputs.Output, error) {
	remaining := script
	compiledScript := []byte{}
	elementsList := compiler.Elements()
	for _, oneElement := range elementsList {
		grammar := oneElement.Grammar()
		tree, err := app.grammarApp.Execute(grammar, remaining)
		if err != nil {
			return nil, err
		}

		composition := oneElement.Composition()
		prefix := composition.Prefix()
		suffix := composition.Suffix()
		value := composition.Pattern()
		replacementsList := composition.Replacements().List()
		for _, oneReplacement := range replacementsList {
			name := oneReplacement.Name()
			criteria := oneReplacement.Criteria()
			found, err := app.criteriaApp.Execute(criteria, tree)
			if err != nil {
				return nil, err
			}

			replacement := bytes.Join([][]byte{
				prefix,
				name,
				suffix,
			}, []byte{})

			value = bytes.ReplaceAll(value, replacement, found)
		}

		remaining = tree.Remaining()
		compiledScript = append(compiledScript, value...)
	}

	grammarIns, err := app.createApp.Grammar().Execute()
	if err != nil {
		return nil, err
	}

	commandIns, err := app.createApp.Command().Execute()
	if err != nil {
		return nil, err
	}

	programIns, compiledRemaining, err := app.programApp.Execute(grammarIns, commandIns, compiledScript)
	if err != nil {
		return nil, err
	}

	builder := app.outputBuilder.Create().WithProgram(programIns)
	if remaining != nil {
		builder.WithScript(remaining)
	}

	if compiledRemaining != nil {
		builder.WithEngine(compiledRemaining)
	}

	return builder.Now()
}
