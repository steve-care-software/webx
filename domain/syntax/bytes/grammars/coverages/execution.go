package coverages

import "github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"

type execution struct {
	suite grammars.Suite
	line  Line
}

func createExecution(
	suite grammars.Suite,
	line Line,
) Execution {
	out := execution{
		suite: suite,
		line:  line,
	}

	return &out
}

// Suite returns the suite
func (obj *execution) Suite() grammars.Suite {
	return obj.suite
}

// Line returns the line
func (obj *execution) Line() Line {
	return obj.line
}
