package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	CreaatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Get(ctx context.Context, userId int64) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user User
	row := s.db.QueryRowContext(ctx, "SELECT id, username, email, created_at FROM users WHERE id = $1", userId)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreaatedAt)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
	    INSERT INTO users (username, email, password) VALUES($1, $2, $3)
	    RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreaatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
