package auth

type User struct {
	Login string
	Name  string
	// TODO: More fields
}

const (
	FormUser     = "user"
	FormPassword = "passwd"
)
