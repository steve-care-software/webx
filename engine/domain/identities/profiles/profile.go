package profiles

type profile struct {
	name        string
	description string
	packages    []string
}

func createProfile(
	name string,
	description string,
) Profile {
	return createProfileInternally(name, description, nil)
}

func createProfileWithPackages(
	name string,
	description string,
	packages []string,
) Profile {
	return createProfileInternally(name, description, packages)
}

func createProfileInternally(
	name string,
	description string,
	packages []string,
) Profile {
	out := profile{
		name:        name,
		description: description,
		packages:    packages,
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

// HasPackages returns true if there is packages, false otherwise
func (obj *profile) HasPackages() bool {
	return obj.packages != nil
}

// Packages returns the packages, if any
func (obj *profile) Packages() []string {
	return obj.packages
}
