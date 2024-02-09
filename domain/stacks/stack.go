package stacks

type stack struct {
	frames Frames
	body   Frames
	head   Frame
}

func createStack() Stack {
	return crateStackInternally(nil, nil, nil)
}

func createStackWithSingleFrame(
	frames Frames,
	head Frame,
) Stack {
	return crateStackInternally(frames, nil, head)
}

func createStackWithFrames(
	frames Frames,
	body Frames,
	head Frame,
) Stack {
	return crateStackInternally(frames, body, head)
}

func crateStackInternally(
	frames Frames,
	body Frames,
	head Frame,
) Stack {
	out := stack{
		frames: frames,
		body:   body,
		head:   head,
	}

	return &out
}

// HasFrames returns true if there is frames, false otherwise
func (obj *stack) HasFrames() bool {
	return obj.frames != nil
}

// Frames returns frames, if any
func (obj *stack) Frames() Frames {
	return obj.frames
}

// HasBody returns true if there is a body, false otherwise
func (obj *stack) HasBody() bool {
	return obj.body != nil
}

// Body returns body, if any
func (obj *stack) Body() Frames {
	return obj.body
}

// HasHead returns true if there is an head, false otherwise
func (obj *stack) HasHead() bool {
	return obj.head != nil
}

// Head returns head, if any
func (obj *stack) Head() Frame {
	return obj.head
}
