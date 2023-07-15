package dal

import (
	"context"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// GetAll returns a slice of all users, sorted by last name
func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	query := getQuery(QUERIES.GET_ALL)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail returns one user by email
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	query := getQuery(QUERIES.GET_BY_EMAIL)

	var user User
	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *User) GetOne(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	query := getQuery(QUERIES.GET_BY_ID)

	var user User
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates one user in the database,
// using the information stored in the receiver u
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	query := getQuery(QUERIES.UPDATE_BY_ID)

	_, err := db.ExecContext(ctx, query,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Active,
		time.Now(),
		u.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes one user from the database, by User.ID
func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	query := getQuery(QUERIES.DELETE_BY_ID)

	_, err := db.ExecContext(ctx, query, u.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *User) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	query := getQuery(QUERIES.DELETE_BY_ID)

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), PASSWORD_COST)
	if hashErr != nil {
		return 0, hashErr
	}

	query := getQuery(QUERIES.INSERT)

	row := db.QueryRowContext(ctx, query,
		user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now(),
	)

	var newID int
	err := row.Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}
