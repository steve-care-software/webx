package connections

import "strings"

func createKeynameFromPaths(from []string, to []string) string {
	fromStr := strings.Join(from, "_")
	toStr := strings.Join(to, "_")
	separator := "+"
	return strings.Join([]string{
		fromStr,
		toStr,
	}, separator)
}
