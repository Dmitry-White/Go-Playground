package dal

func GetUser(login string) *User {
	// Dummy data
	return &User{
		Login:   login,
		Age:     21,
		Address: "23519 West, Civic Center Way, Malibu, CA 90265",
	}
}

func GetUserFriends(user *User) ([]*User, error) {
	// Dummy data
	return nil, &CustomError{"friends not enabled", user}
}
