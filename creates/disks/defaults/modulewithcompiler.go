package defaults

import (
	"errors"
	"fmt"

	compiler_applications "github.com/steve-care-software/webx/applications/compilers"
	creates_module "github.com/steve-care-software/webx/applications/creates/modules"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/programs/modules"
)

type moduleWithCompiler struct {
	compilerApplication       compiler_applications.Application
	builder                   modules.Builder
	moduleBuilder             modules.ModuleBuilder
	compilerBuilder           compilers.Builder
	compilerElementsBuilder   compilers.ElementsBuilder
	compilerElementBuilder    compilers.ElementBuilder
	compilerExecutionBuilder  compilers.ExecutionBuilder
	compilerParametersBuilder compilers.ParametersBuilder
	compilerParameterBuilder  compilers.ParameterBuilder
	compilerValueBuilder      compilers.ValueBuilder
	additionalModules         modules.Modules
}

func createModuleWithCompiler(
	compilerApplication compiler_applications.Application,
	builder modules.Builder,
	moduleBuilder modules.ModuleBuilder,
	compilerBuilder compilers.Builder,
	compilerElementsBuilder compilers.ElementsBuilder,
	compilerElementBuilder compilers.ElementBuilder,
	compilerExecutionBuilder compilers.ExecutionBuilder,
	compilerParametersBuilder compilers.ParametersBuilder,
	compilerParameterBuilder compilers.ParameterBuilder,
	compilerValueBuilder compilers.ValueBuilder,
	additionalModules modules.Modules,
) creates_module.Application {
	out := moduleWithCompiler{
		compilerApplication:       compilerApplication,
		builder:                   builder,
		moduleBuilder:             moduleBuilder,
		compilerBuilder:           compilerBuilder,
		compilerElementsBuilder:   compilerElementsBuilder,
		compilerElementBuilder:    compilerElementBuilder,
		compilerExecutionBuilder:  compilerExecutionBuilder,
		compilerParametersBuilder: compilerParametersBuilder,
		compilerParameterBuilder:  compilerParameterBuilder,
		compilerValueBuilder:      compilerValueBuilder,
		additionalModules:         additionalModules,
	}

	return &out
}

// Execute executes the application
func (app *moduleWithCompiler) Execute() (modules.Modules, error) {
	list := []modules.Module{}
	compiler, err := app.compiler()
	if err != nil {
		return nil, err
	}

	list = append(list, app.additionalModules.List()...)
	list = append(list, compiler...)
	return app.builder.Create().WithList(list).Now()
}

func (app *moduleWithCompiler) compiler() ([]modules.Module, error) {
	executeCompiler, err := app.executeCompiler()
	if err != nil {
		return nil, err
	}

	newCompiler, err := app.newCompiler()
	if err != nil {
		return nil, err
	}

	newCompilerElements, err := app.newCompilerElements()
	if err != nil {
		return nil, err
	}

	newCompilerElement, err := app.newCompilerElement()
	if err != nil {
		return nil, err
	}

	newCompilerExecution, err := app.newCompilerExecution()
	if err != nil {
		return nil, err
	}

	newCompilerParameters, err := app.newCompilerParameters()
	if err != nil {
		return nil, err
	}

	newCompilerParameter, err := app.newCompilerParameter()
	if err != nil {
		return nil, err
	}

	newCompilerValue, err := app.newCompilerValue()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		executeCompiler,
		newCompiler,
		newCompilerElements,
		newCompilerElement,
		newCompilerExecution,
		newCompilerParameters,
		newCompilerParameter,
		newCompilerValue,
	}, nil
}

func (app *moduleWithCompiler) executeCompiler() (modules.Module, error) {
	name := "executeCompiler"
	fn := func(input map[string]interface{}) (interface{}, error) {
		if compiler, ok := input["compiler"].(compilers.Compiler); ok {
			if script, ok := input["script"].([]byte); ok {
				return app.compilerApplication.Execute(compiler, script)
			}

			return nil, errors.New("the script is mandatory in order to execute a compiler")
		}

		return nil, errors.New("the compiler is mandatory in order to execute a compiler")
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompiler() (modules.Module, error) {
	name := "newCompiler"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerBuilder.Create()
		if elements, ok := input["elements"].(compilers.Elements); ok {
			builder.WithElements(elements)
		}

		if outputs, ok := input["outputs"].([]interface{}); ok {
			outputsList := []string{}
			for idx, oneOutput := range outputs {
				if casted, ok := oneOutput.(string); ok {
					outputsList = append(outputsList, casted)
					continue
				}

				str := fmt.Sprintf("the element (index: %d) was expected to contain a string", idx)
				return nil, errors.New(str)
			}

			builder.WithOutputs(outputsList)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerElements() (modules.Module, error) {
	name := "newCompilerElements"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerElementsBuilder.Create()
		if elements, ok := input["list"].([]interface{}); ok {
			elementsList := []compilers.Element{}
			for idx, oneInstance := range elements {
				if casted, ok := oneInstance.(compilers.Element); ok {
					elementsList = append(elementsList, casted)
					continue
				}

				str := fmt.Sprintf("the element (index: %d) was expected to contain an Element instance", idx)
				return nil, errors.New(str)
			}

			builder.WithList(elementsList)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerElement() (modules.Module, error) {
	name := "newCompilerElement"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerElementBuilder.Create()
		if grammar, ok := input["grammar"].(grammars.Grammar); ok {
			builder.WithGrammar(grammar)
		}

		if parameters, ok := input["parameters"].(compilers.Parameters); ok {
			builder.WithParameters(parameters)
		}

		if execution, ok := input["execution"].(compilers.Execution); ok {
			builder.WithExecution(execution)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerExecution() (modules.Module, error) {
	name := "newCompilerExecution"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerExecutionBuilder.Create()
		if parameter, ok := input["parameter"].(string); ok {
			builder.WithParameter(parameter)
		}

		if instructions, ok := input["instructions"].(instructions.Instructions); ok {
			builder.WithInstructions(instructions)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerParameters() (modules.Module, error) {
	name := "newCompilerParameters"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerParametersBuilder.Create()
		if parameters, ok := input["list"].([]interface{}); ok {
			parametersList := []compilers.Parameter{}
			for idx, oneInstance := range parameters {
				if casted, ok := oneInstance.(compilers.Parameter); ok {
					parametersList = append(parametersList, casted)
					continue
				}

				str := fmt.Sprintf("the parameter (index: %d) was expected to contain a Parameter instance", idx)
				return nil, errors.New(str)
			}

			builder.WithList(parametersList)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerParameter() (modules.Module, error) {
	name := "newCompilerParameter"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerParameterBuilder.Create()
		if name, ok := input["name"].(string); ok {
			builder.WithName(name)
		}

		if value, ok := input["value"].(compilers.Value); ok {
			builder.WithValue(value)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerValue() (modules.Module, error) {
	name := "newCompilerValue"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerValueBuilder.Create()
		if constant, ok := input["constant"].(string); ok {
			builder.WithConstant(constant)
		}

		if criteria, ok := input["criteria"].(criterias.Criteria); ok {
			builder.WithCriteria(criteria)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) module(name string, fn modules.ExecuteFn) (modules.Module, error) {
	return app.moduleBuilder.Create().WithName(name).WithFunc(fn).Now()
}
