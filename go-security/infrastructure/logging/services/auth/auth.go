package auth

func LoginUser(login, passwd string) (User, bool) {
	// FIXME: Use real auth database
	if login == "daffy" && passwd == "r4bb1ts3as0n" {
		return User{"daffy", "Daffy Duck"}, true
	}

	return User{}, false
}
