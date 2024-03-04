package sqllites

import "strings"

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
