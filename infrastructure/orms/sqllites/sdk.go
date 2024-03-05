package sqllites

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/datastencil/domain/orms"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

type toHashesFn func(input interface{}) ([]hash.Hash, error)
type toListInstance func(input []interface{}) (orms.Instance, error)

const resourceNameDelimiter = "_"
const endOfLine = "\n"
const connectionNameDelimiter = "$"

// NewOrmRepository creates a new orm repository
func NewOrmRepository(
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	hashAdapter := hash.NewAdapter()
	builders := map[string]interface{}{
		"library":                                               libraries.NewBuilder(),
		"library_link":                                          links.NewLinkBuilder(),
		"library_link_element":                                  links.NewElementBuilder(),
		"library_link_element_condition":                        links.NewConditionBuilder(),
		"library_link_element_condition_value":                  links.NewConditionValueBuilder(),
		"library_link_element_condition_resource":               links.NewConditionResourceBuilder(),
		"library_link_origin":                                   links.NewOriginBuilder(),
		"library_link_origin_resource":                          links.NewOriginResourceBuilder(),
		"library_link_origin_operator":                          links.NewOperatorBuilder(),
		"library_link_origin_value":                             links.NewOriginValueBuilder(),
		"library_layer":                                         layers.NewLayerBuilder(),
		"library_layer_output":                                  layers.NewOutputBuilder(),
		"library_layer_output_kind":                             layers.NewKindBuilder(),
		"library_layer_instruction":                             layers.NewInstructionBuilder(),
		"library_layer_instruction_assignment":                  layers.NewAssignmentBuilder(),
		"library_layer_instruction_assignment_assignable":       layers.NewAssignableBuilder(),
		"library_layer_instruction_assignment_assignable_bytes": layers.NewBytesBuilder(),
	}

	listInstances := map[string]toListInstance{
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
	}

	return createOrmRepository(
		hashAdapter,
		builders,
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
	toHashFns := map[string]toHashesFn{
		"instruction": func(input interface{}) ([]hash.Hash, error) {
			if ins, ok := input.(layers.Instructions); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneInstruction := range list {
					output = append(output, oneInstruction.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the input was expected to contain an Instructions instance")
		},
	}

	hashAdapter := hash.NewAdapter()
	return createOrmService(
		toHashFns,
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
