package outputs

type remaining struct {
	script []byte
	engine []byte
}

func createRemainingWithScript(
	script []byte,
) Remaining {
	return createRemainingInternally(script, nil)
}

func createRemainingWithEngine(
	engine []byte,
) Remaining {
	return createRemainingInternally(nil, engine)
}

func createRemainingWithScriptAndEngine(
	script []byte,
	engine []byte,
) Remaining {
	return createRemainingInternally(script, engine)
}

func createRemainingInternally(
	script []byte,
	engine []byte,
) Remaining {
	out := remaining{
		script: script,
		engine: engine,
	}

	return &out
}

// HasScript returns true if there is a script, false otherwise
func (obj *remaining) HasScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *remaining) Script() []byte {
	return obj.script
}

// HasEngine returns true if there is an engine, false otherwise
func (obj *remaining) HasEngine() bool {
	return obj.engine != nil
}

// Engine returns the engine, if any
func (obj *remaining) Engine() []byte {
	return obj.engine
}
