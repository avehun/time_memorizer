package provider

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type Time struct {
	Amount     uint64 `db:"amount"`
	CategoryID uint64 `db:"category_id"`
	UserID     uint64 `db:"user_id"`
}

type TimeProvider struct {
	db *sqlx.DB
}

func NewTimeProvider(db *sqlx.DB) *TimeProvider {
	return &TimeProvider{
		db: db,
	}
}

func (tp *TimeProvider) Add(ctx context.Context, amount uint64, categoryID uint64, userID uint64) error {
	// Проверяем, существует ли уже запись времени для этой категории и пользователя
	existingTime, err := tp.GetByCategoryAndUser(ctx, categoryID, userID)
	if err != nil {
		log.Println("Error checking if time exists:", err)
		return err
	}

	if existingTime != nil {
		// Если запись времени существует, обновляем значение amount
		newAmount := existingTime.Amount + amount
		err := tp.updateAmount(ctx, categoryID, userID, newAmount)
		if err != nil {
			log.Println("Error updating time amount:", err)
			return err
		}
	} else {
		// Создаем новую запись времени
		_, err = tp.db.ExecContext(ctx, "INSERT INTO times (amount, category_id, user_id) VALUES ($1, $2, $3)", amount, categoryID, userID)
		if err != nil {
			log.Println("Error adding time:", err)
			return err
		}
	}

	return nil
}

func (tp *TimeProvider) Subtract(ctx context.Context, amount uint64, categoryID uint64, userID uint64) error {
	existingTime, err := tp.GetByCategoryAndUser(ctx, categoryID, userID)
	if err != nil {
		log.Println("Error checking if time exists:", err)
		return err
	}

	if existingTime == nil {
		return errors.New("time record not found")
	}

	if existingTime.Amount < amount {
		return errors.New("insufficient time to subtract")
	}

	newAmount := existingTime.Amount - amount
	err = tp.updateAmount(ctx, categoryID, userID, newAmount)
	if err != nil {
		log.Println("Error updating time amount:", err)
		return err
	}

	return nil
}

func (tp *TimeProvider) GetByCategoryAndUser(ctx context.Context, categoryID, userID uint64) (*Time, error) {
	var time Time
	err := tp.db.GetContext(ctx, &time, "SELECT amount FROM times WHERE category_id = $1 AND user_id = $2", categoryID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error getting time by category and user:", err)
		return nil, err
	}
	return &time, nil
}

func (tp *TimeProvider) updateAmount(ctx context.Context, categoryID, userID uint64, amount uint64) error {
	_, err := tp.db.ExecContext(ctx, "UPDATE times SET amount = $1 WHERE category_id = $2 AND user_id = $3", amount, categoryID, userID)
	if err != nil {
		log.Println("Error updating time amount:", err)
		return err
	}
	return nil
}

func (tp *TimeProvider) GetByUserId(ctx context.Context, userID uint64) ([]*Category, error) {
	var categories []*Category
	err := tp.db.SelectContext(ctx, &categories, `
		SELECT c.id, c.name, t.amount
		FROM times t
		JOIN categories c ON t.category_id = c.id
		WHERE t.user_id = $1
	`, userID)
	if err != nil {
		log.Println("Error getting categories by user ID:", err)
		return nil, err
	}
	return categories, nil
}
