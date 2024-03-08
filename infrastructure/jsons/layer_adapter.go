package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/executions"
	structs "github.com/steve-care-software/datastencil/infrastructure/jsons/structs/libraries/layers"
)

type layerAdapter struct {
	hashAdapter         hash.Adapter
	builder             layers.Builder
	layerBuilder        layers.LayerBuilder
	outputBuilder       layers.OutputBuilder
	kindBuilder         layers.KindBuilder
	instructionsBuilder layers.InstructionsBuilder
	instructionBuilder  layers.InstructionBuilder
	conditionBuilder    layers.ConditionBuilder
	assignmentBuilder   layers.AssignmentBuilder
	assignableBuilder   layers.AssignableBuilder
	constantBuilder     constants.ConstantBuilder
	executionBuilder    executions.Builder
	bytesBuilder        bytes.Builder
}

func createLayerAdapter(
	hashAdapter hash.Adapter,
	builder layers.Builder,
	layerBuilder layers.LayerBuilder,
	outputBuilder layers.OutputBuilder,
	kindBuilder layers.KindBuilder,
	instructionsBuilder layers.InstructionsBuilder,
	instructionBuilder layers.InstructionBuilder,
	conditionBuilder layers.ConditionBuilder,
	assignmentBuilder layers.AssignmentBuilder,
	assignableBuilder layers.AssignableBuilder,
	constantBuilder constants.ConstantBuilder,
	executionBuilder executions.Builder,
	bytesBuilder bytes.Builder,
) layers.LayerAdapter {
	out := layerAdapter{
		hashAdapter:         hashAdapter,
		builder:             builder,
		layerBuilder:        layerBuilder,
		outputBuilder:       outputBuilder,
		kindBuilder:         kindBuilder,
		instructionsBuilder: instructionsBuilder,
		instructionBuilder:  instructionBuilder,
		conditionBuilder:    conditionBuilder,
		assignmentBuilder:   assignmentBuilder,
		assignableBuilder:   assignableBuilder,
		constantBuilder:     constantBuilder,
		executionBuilder:    executionBuilder,
		bytesBuilder:        bytesBuilder,
	}

	return &out
}

// ToData converts layer to bytes
func (app *layerAdapter) ToData(ins layers.Layer) ([]byte, error) {
	str := app.toStructLayer(ins)
	data, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ToInstance converts bytes to layer
func (app *layerAdapter) ToInstance(data []byte) (layers.Layer, error) {
	ins := structs.Layer{}
	err := json.Unmarshal(data, &ins)
	if err != nil {
		return nil, err
	}

	return app.toInstanceLayer(ins)
}

func (app *layerAdapter) toStructLayer(ins layers.Layer) structs.Layer {
	return structs.Layer{
		Instructions: app.toStructInstructions(ins.Instructions()),
		Output:       app.toStructOutput(ins.Output()),
		Input:        ins.Input(),
	}
}

func (app *layerAdapter) toInstanceLayer(str structs.Layer) (layers.Layer, error) {
	instructions, err := app.toInstanceInstructions(str.Instructions)
	if err != nil {
		return nil, err
	}

	output, err := app.toInstanceOutput(str.Output)
	if err != nil {
		return nil, err
	}

	return app.layerBuilder.Create().
		WithInstructions(instructions).
		WithOutput(output).
		WithInput(str.Input).
		Now()
}

func (app *layerAdapter) toStructOutput(ins layers.Output) structs.Output {
	output := structs.Output{
		Variable: ins.Variable(),
		Kind:     app.toStructKind(ins.Kind()),
	}

	if ins.HasExecute() {
		output.Execute = ins.Execute()
	}

	return output
}

func (app *layerAdapter) toInstanceOutput(str structs.Output) (layers.Output, error) {
	kind, err := app.toInstanceKind(str.Kind)
	if err != nil {
		return nil, err
	}

	builder := app.outputBuilder.Create().
		WithKind(kind).
		WithVariable(str.Variable)

	if str.Execute != "" {
		builder.WithExecute(str.Execute)
	}

	return builder.Now()
}

func (app *layerAdapter) toStructKind(ins layers.Kind) structs.Kind {
	return structs.Kind{
		Prompt:   ins.IsPrompt(),
		Continue: ins.IsContinue(),
	}
}

func (app *layerAdapter) toInstanceKind(str structs.Kind) (layers.Kind, error) {
	builder := app.kindBuilder.Create()
	if str.Prompt {
		builder.IsPrompt()
	}

	if str.Continue {
		builder.IsContinue()
	}

	return builder.Now()
}

func (app *layerAdapter) toStructInstructions(ins layers.Instructions) []structs.Instruction {
	list := ins.List()
	output := []structs.Instruction{}
	for _, oneInstruction := range list {
		str := app.toStructInstruction(oneInstruction)
		output = append(output, str)
	}

	return output
}

func (app *layerAdapter) toInstanceInstructions(list []structs.Instruction) (layers.Instructions, error) {
	output := []layers.Instruction{}
	for _, oneStr := range list {
		ins, err := app.toInstanceInstruction(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.instructionsBuilder.Create().
		WithList(output).
		Now()
}

func (app *layerAdapter) toStructInstruction(ins layers.Instruction) structs.Instruction {
	output := structs.Instruction{}
	if ins.IsStop() {
		output.Stop = true
	}

	if ins.IsRaiseError() {
		output.RaiseError = ins.RaiseError()
	}

	if ins.IsCondition() {
		condition := app.toStructCondition(ins.Condition())
		output.Condition = &condition
	}

	if ins.IsAssignment() {
		assignment := app.toStructAssignment(ins.Assignment())
		output.Assignment = &assignment
	}

	return output
}

func (app *layerAdapter) toInstanceInstruction(str structs.Instruction) (layers.Instruction, error) {
	builder := app.instructionBuilder.Create()
	if str.Stop {
		builder.IsStop()
	}

	if str.RaiseError != 0 {
		builder.WithRaiseError(str.RaiseError)
	}

	if str.Condition != nil {
		condition, err := app.toInstanceCondition(*str.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	if str.Assignment != nil {
		assignment, err := app.toInstanceAssignment(*str.Assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(assignment)
	}

	return builder.Now()
}

func (app *layerAdapter) toStructCondition(ins layers.Condition) structs.Condition {
	return structs.Condition{
		Variable:     ins.Variable(),
		Instructions: app.toStructInstructions(ins.Instructions()),
	}
}

func (app *layerAdapter) toInstanceCondition(str structs.Condition) (layers.Condition, error) {
	instructions, err := app.toInstanceInstructions(str.Instructions)
	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().
		WithVariable(str.Variable).
		WithInstructions(instructions).
		Now()
}

func (app *layerAdapter) toStructAssignment(ins layers.Assignment) structs.Assignment {
	return structs.Assignment{
		Name:       ins.Name(),
		Assignable: app.toStructAssignable(ins.Assignable()),
	}
}

func (app *layerAdapter) toInstanceAssignment(str structs.Assignment) (layers.Assignment, error) {
	assignable, err := app.toInstanceAssignable(str.Assignable)
	if err != nil {
		return nil, err
	}

	return app.assignmentBuilder.Create().
		WithName(str.Name).
		WithAssignable(assignable).
		Now()
}

func (app *layerAdapter) toStructAssignable(ins layers.Assignable) structs.Assignable {
	output := structs.Assignable{}
	if ins.IsBytes() {
		str := app.toStructBytes(ins.Bytes())
		output.Bytes = &str
	}

	if ins.IsExecution() {
		execution := app.toStructExecution(ins.Execution())
		output.Execution = &execution
	}

	return output
}

func (app *layerAdapter) toInstanceAssignable(str structs.Assignable) (layers.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if str.Bytes != nil {
		bytes, err := app.toInstanceBytes(*str.Bytes)
		if err != nil {
			return nil, err
		}

		builder.WithBytes(bytes)
	}

	if str.Execution != nil {
		execution, err := app.toInstanceExecution(*str.Execution)
		if err != nil {
			return nil, err
		}

		builder.WithExecution(execution)
	}

	return builder.Now()
}

func (app *layerAdapter) toStructConstant(ins constants.Constant) structs.Constant {
	output := structs.Constant{}
	if ins.IsBool() {
		output.Bool = ins.Bool()
	}

	if ins.IsBytes() {
		output.Bytes = ins.Bytes()
	}

	return output
}

func (app *layerAdapter) toInstanceConstant(str structs.Constant) (constants.Constant, error) {
	builder := app.constantBuilder.Create()
	if str.Bool != nil {
		builder.WithBool(*str.Bool)
	}

	if str.Bytes != nil {
		builder.WithBytes(str.Bytes)
	}

	return builder.Now()
}

func (app *layerAdapter) toStructExecution(ins executions.Execution) structs.Execution {
	output := structs.Execution{
		Input: ins.Input(),
	}

	if ins.HasLayer() {
		output.Layer = ins.Layer()
	}

	return output
}

func (app *layerAdapter) toInstanceExecution(str structs.Execution) (executions.Execution, error) {
	builder := app.executionBuilder.Create().
		WithInput(str.Input)

	if str.Layer != "" {
		builder.WithLayer(str.Layer)
	}

	return builder.Now()
}

func (app *layerAdapter) toStructBytes(ins bytes.Bytes) structs.Bytes {
	output := structs.Bytes{}
	if ins.IsJoin() {
		output.Join = ins.Join()
	}

	if ins.IsCompare() {
		output.Compare = ins.Compare()
	}

	if ins.IsHashBytes() {
		output.Hash = ins.HashBytes()
	}

	return output
}

func (app *layerAdapter) toInstanceBytes(str structs.Bytes) (bytes.Bytes, error) {
	builder := app.bytesBuilder.Create()
	if str.Join != nil {
		builder.WithJoin(str.Join)
	}

	if str.Compare != nil {
		builder.WithCompare(str.Compare)
	}

	if str.Hash != "" {
		builder.WithHashBytes(str.Hash)
	}

	return builder.Now()
}
