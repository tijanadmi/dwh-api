package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) OneSignal(id int) (*Signal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, naziv, status
			  from zsi_info
			  where id=:1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var signal Signal

	err := row.Scan(
		&signal.ID,
		&signal.Name,
		&signal.Status,
	)

	if err != nil {
		return nil, err
	}

	return &signal, err
}

// Authenticate authenticates a user
func (m *DBModel) Authenticate(username, testPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	/*var id int
	var hashedPassword string*/

	var user User

	query := `select id, username, password from tis_services_users where username = :1`

	row := m.DB.QueryRowContext(ctx, query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password")
	} else if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetUserByUsername(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where username = :1`

	var user User
	row := m.DB.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *DBModel) GetUserByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where id = :1`

	var user User
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
