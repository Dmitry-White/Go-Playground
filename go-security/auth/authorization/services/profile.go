package services

type Profile struct {
	name   string
	status string
}

func GetProfile(user string) *Profile {
	profile := &Profile{
		name:   user,
		status: "Latest",
	}

	return profile
}

func UpdateProfile(user string) *Profile {
	profile := &Profile{
		name:   user,
		status: "Updated",
	}

	return profile
}
