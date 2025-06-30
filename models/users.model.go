package models

import (
	"context"
	"time"

	"github.com/fulsep/go-backend/utils"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id        int
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
