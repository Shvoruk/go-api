package animal

import (
	"database/sql"
	"errors"
	"github.com/Shvoruk/go-api/types"
	"log"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) GetAll() ([]types.Animal, error) {
	rows, err := r.db.Query("SELECT id, name, category FROM animals")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var animals []types.Animal
	for rows.Next() {
		var a types.Animal
		if err := rows.Scan(&a.ID, &a.Name, &a.Category); err != nil {
			return nil, err
		}
		animals = append(animals, a)
	}
	return animals, nil
}

func (r *Repo) Get(id string) (*types.Animal, error) {
	row := r.db.QueryRow("SELECT id, name, category FROM animals WHERE id = ?", id)

	var a types.Animal
	if err := row.Scan(&a.ID, &a.Name, &a.Category); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *Repo) Create(a *types.Animal) (*types.Animal, error) {
	res, err := r.db.Exec(
		"INSERT INTO animals (name, category) VALUES (?, ?)",
		a.Name, a.Category,
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	a.ID = int(id)
	return a, nil
}

func (r *Repo) Update(a *types.Animal) (*types.Animal, error) {
	res, err := r.db.Exec(
		"UPDATE animals SET name = ?, category = ? WHERE id = ?",
		a.Name, a.Category, a.ID,
	)
	if err != nil {
		return nil, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, sql.ErrNoRows
	}
	return a, nil
}

func (r *Repo) Delete(id string) error {
	res, err := r.db.Exec("DELETE FROM animals WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
