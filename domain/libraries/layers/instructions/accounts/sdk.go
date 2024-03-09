package accounts

// Account represents an account instruction
type Account interface {
	IsUpdateCriteria() bool
	UpdateCriteria() UpdateCriteria
	IsUpdate() bool
	Update() Update
	IsDelete() bool
	Delete() Delete
}

// UpdateCriteria represents an update criteria
type UpdateCriteria interface {
	ChangeSigner() bool
	ChangeEncryptor() bool
	Credentials() string
}

// Update represents an update
type Update interface {
	Credentials() string
	Account() string
	Criteria() string
}

// Delete represents a delete
type Delete interface {
	Credentials() string
	Account() string
}
