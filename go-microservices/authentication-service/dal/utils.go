package dal

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ResetPassword is the method we will use to change a user's password.
func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), PASSWORD_COST)
	if err != nil {
		return err
	}

	query := getQuery(QUERIES.UPDATE_PASSWORD)

	_, err = db.ExecContext(ctx, query, hashedPassword, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func getQuery(name string) string {
	queryBytes, err := queriesFolder.ReadFile(name)
	if err != nil {
		log.Panic(err)
	}

	return string(queryBytes)
}
