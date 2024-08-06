package actions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/commands/contents/actions/datas"
	"github.com/steve-care-software/webx/engine/cursors/domain/commands/contents/actions/moves"
	"github.com/steve-care-software/webx/engine/cursors/domain/commands/contents/actions/updates"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Action represents an action
type Action interface {
	IsPurgeAll() bool
	IsMerge() bool
	IsInsert() bool
	Insert() originals.Original
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() string
	IsPurge() bool
	Purge() string
	IsMove() bool
	Move() moves.Move
	IsData() bool
	Data() datas.Data
}
