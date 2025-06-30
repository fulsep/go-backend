package models

import (
	"context"
	"time"

	"github.com/fulsep/go-backend/utils"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FindAllUsers() ([]User, error) {
	db, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}

	defer func() {
		db.Conn().Close(context.Background())
	}()

	rows, err := db.Query(context.Background(), `
		SELECT id, email, password, created_at, updated_at FROM users
	`)

	if err != nil {
		return []User{}, err
	}

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)

	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func FindOneUser(id int) (User, error) {
	db, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}

	defer func() {
		db.Conn().Close(context.Background())
	}()

	rows, err := db.Query(
		context.Background(),
		`
		SELECT id, email, password, created_at, updated_at
		FROM users
		WHERE id = $1
		`,
		id,
	)

	if err != nil {
		return User{}, err
	}

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
