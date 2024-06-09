package provider

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       uint64 `db:"id"`
	Username string `db:"username"`
}

type UserProvider struct {
	db *sqlx.DB
}

func NewUserProvider(db *sqlx.DB) *UserProvider {
	return &UserProvider{
		db: db,
	}
}

func (up *UserProvider) GetByUsername(ctx context.Context, username string) (uint64, error) {
	var userID uint64
	err := up.db.GetContext(ctx, &userID, "SELECT id FROM users WHERE username = $1", username)
	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь не найден, создаем нового
			newUserID, err := up.createUser(ctx, username)
			if err != nil {
				log.Println("Error creating user:", err)
				return 0, err
			}
			return newUserID, nil
		}
		log.Println("Error getting user by username:", err)
		return 0, err
	}
	return userID, nil
}

func (up *UserProvider) createUser(ctx context.Context, username string) (uint64, error) {
	var userID uint64
	err := up.db.QueryRowContext(ctx, "INSERT INTO users (username) VALUES ($1) RETURNING id", username).Scan(&userID)
	if err != nil {
		log.Println("Error creating user:", err)
		return 0, err
	}
	return userID, nil
}
