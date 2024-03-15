package sqllites

import (
	"database/sql"
	"strings"

	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
)

type transaction struct {
	actionsList []actions.Action
	pTX         *sql.Tx
}

type table struct {
	name     string
	key      field
	fields   []field
	children []table
}

type field struct {
	name     string
	kind     kind
	canBeNil bool
}

type kind struct {
	pSingle     *uint8
	pList       *list
	pForeignKey *foreignKey
	pConnection *connection
}

type list struct {
	value     uint8
	delimiter string
}

type foreignKey struct {
	localField       field
	foreignTableName string
	foreignField     field
}

type connection struct {
	middleTableName string
	from            foreignKey
	to              foreignKey
}

func createConnectionTableName(
	fromPath []string,
	toPath []string,
) string {
	fromTableName := strings.Join(fromPath, resourceNameDelimiter)
	toTableName := strings.Join(toPath, resourceNameDelimiter)
	return strings.Join(
		[]string{
			fromTableName,
			toTableName,
		},
		connectionNameDelimiter,
	)
}
