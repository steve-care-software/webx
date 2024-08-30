package instructions

type instruction struct {
	block   string
	line    uint
	tokens  Tokens
	syscall Syscall
}

func createInstruction(
	block string,
	line uint,
	tokens Tokens,
) Instruction {
	return createInstructionInternally(
		block,
		line,
		tokens,
		nil,
	)
}

func createInstructionWithSyscall(
	block string,
	line uint,
	tokens Tokens,
	syscall Syscall,
) Instruction {
	return createInstructionInternally(
		block,
		line,
		tokens,
		syscall,
	)
}

func createInstructionInternally(
	block string,
	line uint,
	tokens Tokens,
	syscall Syscall,
) Instruction {
	out := instruction{
		block:   block,
		line:    line,
		tokens:  tokens,
		syscall: syscall,
	}

	return &out
}

// Block returns the block
func (obj *instruction) Block() string {
	return obj.block
}

// Line returns the line
func (obj *instruction) Line() uint {
	return obj.line
}

// Tokens returns the tokens
func (obj *instruction) Tokens() Tokens {
	return obj.tokens
}

// HasSyscall returns true if there is a syscall, false otherwise
func (obj *instruction) HasSyscall() bool {
	return obj.syscall != nil
}

// Syscall returns the syscall, if any
func (obj *instruction) Syscall() Syscall {
	return obj.syscall
}
