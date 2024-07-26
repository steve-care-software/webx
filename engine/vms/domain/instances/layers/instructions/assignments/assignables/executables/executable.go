package executables

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

type executable struct {
	hash   hash.Hash
	local  string
	remote string
}

func createExecutableWithLocal(
	hash hash.Hash,
	local string,
) Executable {
	return createExecutableInternally(hash, local, "")
}

func createExecutableWithRemote(
	hash hash.Hash,
	remote string,
) Executable {
	return createExecutableInternally(hash, "", remote)
}

func createExecutableInternally(
	hash hash.Hash,
	local string,
	remote string,
) Executable {
	out := executable{
		hash:   hash,
		local:  local,
		remote: remote,
	}

	return &out
}

// Hash returns the hash
func (obj *executable) Hash() hash.Hash {
	return obj.hash
}

// IsLocal returns true if local, false otherwise
func (obj *executable) IsLocal() bool {
	return obj.local != ""
}

// Local returns the local path, if any
func (obj *executable) Local() string {
	return obj.local
}

// IsRemote returns true if remote, false otherwise
func (obj *executable) IsRemote() bool {
	return obj.remote != ""
}

// Remote returns the remote host, if any
func (obj *executable) Remote() string {
	return obj.remote
}
