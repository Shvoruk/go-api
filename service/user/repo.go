package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Shvoruk/go-api/types"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) ExistsByEmail(email string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)"
	var exists bool
	if err := r.db.QueryRow(query, email).Scan(&exists); err != nil {
		return false, fmt.Errorf("User|ExistsByEmail: %w", err)
	}
	return exists, nil
}

func (r *Repo) GetByEmail(email string) (*types.User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	var u types.User
	err := r.db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("User|GetByEmail: %w", err)
	}
	return &u, nil
}

func (r *Repo) Create(u *types.User) (int, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, u.Username, u.Email, u.Password)
	if err != nil {
		return 0, fmt.Errorf("User|Create: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("User|Create: %w", err)
	}
	return int(id), nil
}
