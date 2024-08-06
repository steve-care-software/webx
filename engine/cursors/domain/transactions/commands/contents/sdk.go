package contents

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors"
	"github.com/steve-care-software/webx/engine/cursors/domain/transactions/commands/contents/actions"
)

// Contents represents contents
type Contents interface {
	List() []Content
}

// Content represents content
type Content interface {
	Cursor() cursors.Cursor
	Action() actions.Action
}
