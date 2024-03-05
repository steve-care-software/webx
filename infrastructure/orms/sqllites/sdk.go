package sqllites

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/datastencil/domain/orms"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

// listInstanceToElementHashesListFn takes a list instance and return its element's hashes
type listInstanceToElementHashesListFn func(input interface{}) ([]hash.Hash, error)

// elementsToListInstanceFn takes a list of elements and returns a list instance
type elementsToListInstanceFn func(input []interface{}) (orms.Instance, error)

// buildInstanceFn takes values and build an Instance instance
type buildInstanceFn func(values map[string]interface{}) (orms.Instance, error)

const resourceNameDelimiter = "_"
const endOfLine = "\n"
const connectionNameDelimiter = "$"

// NewOrmRepository creates a new orm repository
func NewOrmRepository(
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	hashAdapter := hash.NewAdapter()
	buildInstances := map[string]buildInstanceFn{
		"library": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewBuilder()
			return builder.Now()
		},
		"library_link": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewLinkBuilder()
			return builder.Now()
		},
		"library_link_element": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewElementBuilder()
			return builder.Now()
		},
		"library_link_element_condition": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewConditionBuilder()
			return builder.Now()
		},
		"library_link_element_condition_value": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewConditionValueBuilder()
			return builder.Now()
		},
		"library_link_element_condition_resource": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewConditionResourceBuilder()
			return builder.Now()
		},
		"library_link_origin": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewOriginBuilder()
			return builder.Now()
		},
		"library_link_origin_value": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewOriginValueBuilder()
			return builder.Now()
		},
		"library_link_origin_resource": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewOriginResourceBuilder()
			return builder.Now()
		},
		"library_link_origin_operator": func(values map[string]interface{}) (orms.Instance, error) {
			builder := links.NewOperatorBuilder()
			return builder.Now()
		},
		"library_layer": func(values map[string]interface{}) (orms.Instance, error) {
			builder := layers.NewLayerBuilder()
			if value, ok := values["instructions"]; ok {
				if ins, ok := value.(layers.Instructions); ok {
					builder.WithInstructions(ins.(layers.Instructions))
				}
			}

			if value, ok := values["output"]; ok {
				if pIns, ok := value.(*orms.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithOutput(ins.(layers.Output))
					}
				}
			}

			if value, ok := values["input"]; ok {
				if value != nil {
					builder.WithInput(value.(string))
				}
			}

			return builder.Now()
		},
		"library_layer_output": func(values map[string]interface{}) (orms.Instance, error) {
			builder := layers.NewOutputBuilder()
			if value, ok := values["variable"]; ok {
				builder.WithVariable(value.(string))
			}

			if value, ok := values["kind"]; ok {
				if pIns, ok := value.(*orms.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithKind(ins.(layers.Kind))
					}
				}
			}

			if value, ok := values["execute"]; ok {
				if value != nil {
					builder.WithExecute(value.(string))
				}
			}

			return builder.Now()
		},
		"library_layer_output_kind": func(values map[string]interface{}) (orms.Instance, error) {
			builder := layers.NewKindBuilder()
			if value, ok := values["prompt"]; ok {
				if value.(int64) != 0 {
					builder.IsPrompt()
				}
			}

			if value, ok := values["continue"]; ok {
				if value.(int64) != 0 {
					builder.IsContinue()
				}
			}

			return builder.Now()
		},
		"library_layer_instruction": func(values map[string]interface{}) (orms.Instance, error) {
			builder := layers.NewInstructionBuilder()
			if value, ok := values["assignment"]; ok {
				if pIns, ok := value.(*orms.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithAssignment(ins.(layers.Assignment))
					}
				}
			}

			return builder.Now()
		},
		"library_layer_instruction_assignment": func(values map[string]interface{}) (orms.Instance, error) {
			builder := layers.NewAssignmentBuilder()
			if value, ok := values["name"]; ok {
				builder.WithName(value.(string))
			}

			if value, ok := values["assignable"]; ok {
				if pIns, ok := value.(*orms.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithAssignable(ins.(layers.Assignable))
					}
				}
			}

			return builder.Now()
		},
		"library_layer_instruction_assignment_assignable": func(values map[string]interface{}) (orms.Instance, error) {
			builder := layers.NewAssignableBuilder()
			if value, ok := values["bytes"]; ok {
				if pIns, ok := value.(*orms.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithBytes(ins.(layers.Bytes))
					}
				}
			}

			return builder.Now()
		},
		"library_layer_instruction_assignment_assignable_bytes": func(values map[string]interface{}) (orms.Instance, error) {
			builder := layers.NewBytesBuilder()
			if value, ok := values["joins"]; ok {
				if value != nil {
					builder.WithJoin(value.([]string))
				}

			}

			if value, ok := values["compares"]; ok {
				if value != nil {
					builder.WithCompare(value.([]string))
				}

			}

			if value, ok := values["hash_bytes"]; ok {
				if value != nil {
					builder.WithHashBytes(value.(string))
				}
			}

			return builder.Now()
		},
	}

	listInstances := map[string]elementsToListInstanceFn{
		"library_links": func(input []interface{}) (orms.Instance, error) {
			output := []links.Link{}
			for _, oneIns := range input {
				output = append(output, oneIns.(links.Link))
			}

			return links.NewBuilder().Create().
				WithList(output).
				Now()
		},
		"link_elements": func(input []interface{}) (orms.Instance, error) {
			output := []links.Element{}
			for _, oneIns := range input {
				output = append(output, oneIns.(links.Element))
			}

			return links.NewElementsBuilder().Create().
				WithList(output).
				Now()
		},
		"library_layers": func(input []interface{}) (orms.Instance, error) {
			output := []layers.Layer{}
			for _, oneIns := range input {
				output = append(output, oneIns.(layers.Layer))
			}

			return layers.NewBuilder().Create().
				WithList(output).
				Now()
		},
		"layer_instructions": func(input []interface{}) (orms.Instance, error) {
			output := []layers.Instruction{}
			for _, oneIns := range input {
				output = append(output, oneIns.(layers.Instruction))
			}

			return layers.NewInstructionsBuilder().Create().
				WithList(output).
				Now()
		},
	}

	return createOrmRepository(
		hashAdapter,
		buildInstances,
		listInstances,
		skeleton,
		dbPtr,
	)
}

// NewOrmService creates a new orm service
func NewOrmService(
	repository orms.Repository,
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
	txPtr *sql.Tx,
) orms.Service {
	listInstanceToElementHashesListFn := map[string]listInstanceToElementHashesListFn{
		"layer": func(input interface{}) ([]hash.Hash, error) {
			if ins, ok := input.(layers.Layers); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the input was expected to contain a Layers instance")
		},
		"instruction": func(input interface{}) ([]hash.Hash, error) {
			if ins, ok := input.(layers.Instructions); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the input was expected to contain an Instructions instance")
		},
		"link": func(input interface{}) ([]hash.Hash, error) {
			if ins, ok := input.(links.Links); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the input was expected to contain a Links instance")
		},
		"element": func(input interface{}) ([]hash.Hash, error) {
			if ins, ok := input.(links.Elements); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the input was expected to contain an Elements instance")
		},
	}

	hashAdapter := hash.NewAdapter()
	return createOrmService(
		listInstanceToElementHashesListFn,
		repository,
		hashAdapter,
		skeleton,
		dbPtr,
		txPtr,
	)
}

// NewSkeletonFactory creates a new skeleton factory
func NewSkeletonFactory() skeletons.Factory {
	builder := skeletons.NewBuilder()
	resourcesBuilder := resources.NewBuilder()
	resourceBuilder := resources.NewResourceBuilder()
	fieldsBuilder := resources.NewFieldsBuilder()
	fieldBuilder := resources.NewFieldBuilder()
	builderInstructionBuilder := resources.NewBuilderInstructionBuilder()
	kindBuilder := resources.NewKindBuilder()
	nativeBuilder := resources.NewNativeBuilder()
	listBuilder := resources.NewListBuilder()
	connectionsBuilder := connections.NewBuilder()
	connectionBuilder := connections.NewConnectionBuilder()
	connectionFieldBuilder := connections.NewFieldBuilder()
	return createSkeletonFactory(
		builder,
		resourcesBuilder,
		resourceBuilder,
		fieldsBuilder,
		fieldBuilder,
		builderInstructionBuilder,
		kindBuilder,
		nativeBuilder,
		listBuilder,
		connectionsBuilder,
		connectionBuilder,
		connectionFieldBuilder,
	)
}
