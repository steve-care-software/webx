package datas

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"

// Data represents data related action
type Data interface {
	IsInsert() bool
	Insert() delimiters.Delimiter
	IsDelete() bool
	Delete() delimiters.Delimiter
}
