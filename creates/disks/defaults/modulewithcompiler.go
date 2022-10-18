package defaults

import (
	"errors"
	"fmt"

	compiler_applications "github.com/steve-care-software/syntax/applications/engines/compilers"
	creates_module "github.com/steve-care-software/syntax/applications/engines/creates/modules"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"
)

type moduleWithCompiler struct {
	compilerApplication         compiler_applications.Application
	builder                     modules.Builder
	moduleBuilder               modules.ModuleBuilder
	compilerBuilder             compilers.Builder
	compilerElementBuilder      compilers.ElementBuilder
	compilerCompositionBuilder  compilers.CompositionBuilder
	compilerReplacementsBuilder compilers.ReplacementsBuilder
	compilerReplacementBuilder  compilers.ReplacementBuilder
	additionalModules           modules.Modules
}

func createModuleWithCompiler(
	compilerApplication compiler_applications.Application,
	builder modules.Builder,
	moduleBuilder modules.ModuleBuilder,
	compilerBuilder compilers.Builder,
	compilerElementBuilder compilers.ElementBuilder,
	compilerCompositionBuilder compilers.CompositionBuilder,
	compilerReplacementsBuilder compilers.ReplacementsBuilder,
	compilerReplacementBuilder compilers.ReplacementBuilder,
	additionalModules modules.Modules,
) creates_module.Application {
	out := moduleWithCompiler{
		compilerApplication:         compilerApplication,
		builder:                     builder,
		moduleBuilder:               moduleBuilder,
		compilerBuilder:             compilerBuilder,
		compilerElementBuilder:      compilerElementBuilder,
		compilerCompositionBuilder:  compilerCompositionBuilder,
		compilerReplacementsBuilder: compilerReplacementsBuilder,
		compilerReplacementBuilder:  compilerReplacementBuilder,
		additionalModules:           additionalModules,
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

	newCompilerElement, err := app.newCompilerElement()
	if err != nil {
		return nil, err
	}

	newCompilerComposition, err := app.newCompilerComposition()
	if err != nil {
		return nil, err
	}

	newCompilerReplacements, err := app.newCompilerReplacements()
	if err != nil {
		return nil, err
	}

	newCompilerReplacement, err := app.newCompilerReplacement()
	if err != nil {
		return nil, err
	}

	return []modules.Module{
		executeCompiler,
		newCompiler,
		newCompilerElement,
		newCompilerComposition,
		newCompilerReplacements,
		newCompilerReplacement,
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
		if elements, ok := input["elements"].([]interface{}); ok {
			elementsList := []compilers.Element{}
			for idx, oneInstance := range elements {
				if casted, ok := oneInstance.(compilers.Element); ok {
					elementsList = append(elementsList, casted)
					continue
				}

				str := fmt.Sprintf("the element (index: %d) was expected to contain an Element instance", idx)
				return nil, errors.New(str)
			}

			builder.WithElements(elementsList)
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

		if composition, ok := input["composition"].(compilers.Composition); ok {
			builder.WithComposition(composition)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerComposition() (modules.Module, error) {
	name := "newCompilerComposition"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerCompositionBuilder.Create()
		if prefix, ok := input["prefix"].([]byte); ok {
			builder.WithPrefix(prefix)
		}

		if suffix, ok := input["suffix"].([]byte); ok {
			builder.WithSuffix(suffix)
		}

		if pattern, ok := input["pattern"].([]byte); ok {
			builder.WithPattern(pattern)
		}

		if replacements, ok := input["replacements"].(compilers.Replacements); ok {
			builder.WithReplacements(replacements)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerReplacements() (modules.Module, error) {
	name := "newCompilerReplacements"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerReplacementsBuilder.Create()
		if list, ok := input["list"].([]interface{}); ok {
			replacementsList := []compilers.Replacement{}
			for idx, oneInstance := range list {
				if casted, ok := oneInstance.(compilers.Replacement); ok {
					replacementsList = append(replacementsList, casted)
					continue
				}

				str := fmt.Sprintf("the element (index: %d) was expected to contain a Replacement instance", idx)
				return nil, errors.New(str)
			}

			builder.WithList(replacementsList)
		}

		return builder.Now()
	}

	return app.module(name, fn)
}

func (app *moduleWithCompiler) newCompilerReplacement() (modules.Module, error) {
	name := "newCompilerReplacement"
	fn := func(input map[string]interface{}) (interface{}, error) {
		builder := app.compilerReplacementBuilder.Create()
		if name, ok := input["name"].([]byte); ok {
			builder.WithName(name)
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
