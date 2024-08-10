package profiles

type profile struct {
	name        string
	description string
	namespaces  []string
}

func createProfile(
	name string,
	description string,
) Profile {
	return createProfileInternally(name, description, nil)
}

func createProfileWithNamespaces(
	name string,
	description string,
	namespaces []string,
) Profile {
	return createProfileInternally(name, description, namespaces)
}

func createProfileInternally(
	name string,
	description string,
	namespaces []string,
) Profile {
	out := profile{
		name:        name,
		description: description,
		namespaces:  namespaces,
	}

	return &out
}

// Name returns the name
func (obj *profile) Name() string {
	return obj.name
}

// Description returns the description
func (obj *profile) Description() string {
	return obj.description
}

// HasNamespaces returns true if there is namespaces, false otherwise
func (obj *profile) HasNamespaces() bool {
	return obj.namespaces != nil
}

// Namespaces returns the namespaces, if any
func (obj *profile) Namespaces() []string {
	return obj.namespaces
}
