package files

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

func prepareFilePath(hashAdapter hash.Adapter, dbPath []string, basePath []string, endPath []string) (string, error) {
	if len(dbPath) <= 0 {
		return "", errors.New("the database path is invalid")
	}

	path := []string{}
	path = append(path, basePath...)
	path = append(path, dbPath[0:len(dbPath)-1]...)
	path = append(path, endPath...)

	dirPath := filepath.Join(path...)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, readWritePermissionBits)
		if err != nil {
			return "", nil
		}
	}

	dbPathStr := filepath.Join(dbPath...)
	pHash, err := hashAdapter.FromBytes([]byte(dbPathStr))
	if err != nil {
		return "", nil
	}

	path = append(path, pHash.String())
	return filepath.Join(path...), nil
}
