package contents

import (
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/commands/contents/actions"
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors"
)

// Content represents content
type Content interface {
	Cursor() cursors.Cursor
	Action() actions.Action
	CreatedOn() time.Time
}
