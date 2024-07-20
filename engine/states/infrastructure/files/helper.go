package files

import "path/filepath"

func createFilePath(basePath []string, path []string) string {
	return filepath.Join(append(basePath, path...)...)
}
