package layers

// Layer represents a layer
type Layer struct {
	Instructions []Instruction `json:"instructions"`
	Output       Output        `json:"output"`
	Input        string        `json:"input"`
}

// Output represents an output
type Output struct {
	Variable bool   `json:"variable"`
	Kind     Kind   `json:"kind"`
	Execute  string `json:"execute"`
}

// Kind represents a kind
type Kind struct {
	Prompt   bool `json:"prompt"`
	Continue bool `json:"continue"`
}

// Instruction represents an instruction
type Instruction struct {
	Stop       bool        `json:"stop"`
	RaiseError uint        `json:"raise_error"`
	Condition  *Condition  `json:"condition"`
	Assignment *Assignment `json:"assignment"`
}

// Condition represents condition
type Condition struct {
	Variable     string        `json:"variable"`
	Instructions []Instruction `json:"instructions"`
}

// Assignment represents assignment
type Assignment struct {
	Name       string     `json:"name"`
	Assignable Assignable `json:"assignable"`
}

// Assignable represents assignable
type Assignable struct {
	Bytes     *Bytes     `json:"bytes"`
	Execution *Execution `json:"execution"`
}

// Constant represents constant
type Constant struct {
	Bool  *bool  `json:"bool"`
	Bytes []byte `json:"bytes"`
}

// Execution represents execution
type Execution struct {
	Input string `json:"input"`
	Layer string `json:"layer"`
}

// Bytes represents a bytes
type Bytes struct {
	Join    []string `json:"join"`
	Compare []string `json:"compare"`
	Hash    string   `json:"hash"`
}
