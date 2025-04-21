package animal

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

func (r *Repo) GetAllByCategory(category string) ([]types.Animal, error) {
	query := "SELECT id, status, name, img_url FROM animals WHERE category = ?"

	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, fmt.Errorf("Animal|GetAllByCategory: %w", err)
	}
	defer rows.Close()

	var animals []types.Animal
	for rows.Next() {
		var a types.Animal
		err = rows.Scan(
			&a.ID,
			&a.Status,
			&a.Name,
			&a.ImageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("Abimal|GetAllByCategory: %w", err)
		}
		animals = append(animals, a)
	}
	return animals, nil
}

func (r *Repo) GetAllByUser(userId int) ([]types.Animal, error) {
	query := "SELECT id, status, name, img_url FROM animals WHERE user_id = ?"

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("Animal|GetAllByUser: %w", err)
	}
	defer rows.Close()

	var animals []types.Animal
	for rows.Next() {
		var a types.Animal
		err = rows.Scan(
			&a.ID,
			&a.Status,
			&a.Name,
			&a.ImageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("Animal|GetAllByUser: %w", err)
		}
		animals = append(animals, a)
	}
	return animals, nil
}

func (r *Repo) Get(id int) (*types.Animal, error) {
	query := "SELECT * FROM animals WHERE id = ?"
	var a types.Animal
	err := r.db.QueryRow(query, id).Scan(
		&a.ID,
		&a.UserID,
		&a.Status,
		&a.Name,
		&a.Sex,
		&a.Breed,
		&a.Size,
		&a.AgeInMonth,
		&a.Category,
		&a.ImageURL,
		&a.Description,
		&a.ContactInfo,
		&a.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("Animal|Get: %w", err)
	}
	return &a, nil
}

func (r *Repo) Create(a *types.Animal) (int, error) {
	query := "INSERT INTO animals (user_id, status, name, sex, breed, size, age_in_month, category, img_url, description, contact_info) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query,
		a.UserID,
		a.Status,
		a.Name,
		a.Sex,
		a.Breed,
		a.Size,
		a.AgeInMonth,
		a.Category,
		a.ImageURL,
		a.Description,
		a.ContactInfo,
	)
	if err != nil {
		return 0, fmt.Errorf("Animal|Create: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Animal|Create: %w", err)
	}
	return int(id), nil
}

func (r *Repo) Delete(userId int, id int) error {
	query := `DELETE FROM animals WHERE user_id = ? AND id = ?`
	res, err := r.db.Exec(query, userId, id)
	if err != nil {
		return fmt.Errorf("Animal|Delete: %w", err)
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Animal|Delete: %w", err)
	}
	if aff == 0 {
		return sql.ErrNoRows
	}
	return nil
}
