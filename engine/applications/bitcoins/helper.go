package bitcoins

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
)

// deriveKeyFromPath derives a child key from the master key using the given derivation path.
func deriveKeyFromPath(masterKey *hdkeychain.ExtendedKey, path string) (*hdkeychain.ExtendedKey, error) {
	// Ensure the path starts with 'm'
	if !strings.HasPrefix(path, "m/") {
		return nil, errors.New("invalid derivation path, must start with 'm/'")
	}

	// Remove the 'm/' prefix
	path = path[2:]

	// Split the path into components
	components := strings.Split(path, "/")

	// Start with the master key
	key := masterKey

	// Iterate over each component and derive the corresponding child key
	for _, component := range components {
		// Check if the component is hardened (ends with "'")
		hardened := strings.HasSuffix(component, "'")
		if hardened {
			// Remove the trailing "'"
			component = component[:len(component)-1]
		}

		// Convert component to uint32
		index, err := strconv.ParseUint(component, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid path component %s: %v", component, err)
		}

		// Derive the child key
		if hardened {
			index += hdkeychain.HardenedKeyStart // Hardened keys have index >= 2^31
		}

		key, err = key.Derive(uint32(index))
		if err != nil {
			return nil, fmt.Errorf("failed to derive child key at index %d: %v", index, err)
		}
	}

	return key, nil
}
