package instructions

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	json_accounts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts"
	json_assignments "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/databases"
)

// Adapter represents an instructions adapter
type Adapter struct {
	accountAdapter     *json_accounts.Adapter
	assignmnetAdapter  *json_assignments.Adapter
	databaseAdapter    *json_databases.Adapter
	conditionBuilder   instructions.ConditionBuilder
	instructionBuilder instructions.InstructionBuilder
	builder            instructions.Builder
}

func createAdapter(
	accountAdapter *json_accounts.Adapter,
	assignmnetAdapter *json_assignments.Adapter,
	databaseAdapter *json_databases.Adapter,
	conditionBuilder instructions.ConditionBuilder,
	instructionBuilder instructions.InstructionBuilder,
	builder instructions.Builder,
) instructions.Adapter {
	out := Adapter{
		accountAdapter:     accountAdapter,
		assignmnetAdapter:  assignmnetAdapter,
		databaseAdapter:    databaseAdapter,
		conditionBuilder:   conditionBuilder,
		instructionBuilder: instructionBuilder,
		builder:            builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins instructions.Instructions) ([]byte, error) {
	ptr, err := app.InstructionsToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (instructions.Instructions, error) {
	ins := new([]Instruction)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToInstructions(*ins)
}

// InstructionsToStruct converts an instructions to struct
func (app *Adapter) InstructionsToStruct(ins instructions.Instructions) ([]Instruction, error) {
	out := []Instruction{}
	list := ins.List()
	for _, oneIns := range list {
		ptr, err := app.InstructionToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructToInstructions converts a struct to instructions
func (app *Adapter) StructToInstructions(list []Instruction) (instructions.Instructions, error) {
	output := []instructions.Instruction{}
	for _, oneStr := range list {
		ins, err := app.StructToInstruction(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().
		WithList(output).
		Now()
}

// InstructionToStruct converts an instruction to struct
func (app *Adapter) InstructionToStruct(ins instructions.Instruction) (*Instruction, error) {
	out := Instruction{}
	if ins.IsStop() {
		out.Stop = ins.IsStop()
	}

	if ins.IsRaiseError() {
		value := ins.RaiseError()
		out.RaiseError = &value
	}

	if ins.IsCondition() {
		ptr, err := app.ConditionToStruct(ins.Condition())
		if err != nil {
			return nil, err
		}

		out.Condition = ptr
	}

	if ins.IsAssignment() {
		ptr, err := app.assignmnetAdapter.AssignmentToStruct(ins.Assignment())
		if err != nil {
			return nil, err
		}

		out.Assignment = ptr
	}

	if ins.IsAccount() {
		ptr, err := app.accountAdapter.AccountToStruct(ins.Account())
		if err != nil {
			return nil, err
		}

		out.Account = ptr
	}

	if ins.IsDatabase() {
		ptr, err := app.databaseAdapter.DatabaseToStruct(ins.Database())
		if err != nil {
			return nil, err
		}

		out.Database = ptr
	}

	return &out, nil
}

// StructToInstruction converts a struct to instruction
func (app *Adapter) StructToInstruction(str Instruction) (instructions.Instruction, error) {
	builder := app.instructionBuilder.Create()
	if str.Stop {
		builder.IsStop()
	}

	if str.RaiseError != nil {
		builder.WithRaiseError(*str.RaiseError)
	}

	if str.Condition != nil {
		ins, err := app.StructToCondition(*str.Condition)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(ins)
	}

	if str.Assignment != nil {
		ins, err := app.assignmnetAdapter.StructToAssignment(*str.Assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(ins)
	}

	if str.Account != nil {
		ins, err := app.accountAdapter.StructToAccount(*str.Account)
		if err != nil {
			return nil, err
		}

		builder.WithAccount(ins)
	}

	if str.Database != nil {
		ins, err := app.databaseAdapter.StructToDatabase(*str.Database)
		if err != nil {
			return nil, err
		}

		builder.WithDatabase(ins)
	}

	return builder.Now()
}

// ConditionToStruct converts a condition to struct
func (app *Adapter) ConditionToStruct(ins instructions.Condition) (*Condition, error) {
	list, err := app.InstructionsToStruct(ins.Instructions())
	if err != nil {
		return nil, err
	}

	return &Condition{
		Variable:     ins.Variable(),
		Instructions: list,
	}, nil
}

// StructToCondition converts a struct to condition
func (app *Adapter) StructToCondition(str Condition) (instructions.Condition, error) {
	ins, err := app.StructToInstructions(str.Instructions)
	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().
		WithVariable(str.Variable).
		WithInstructions(ins).
		Now()
}
