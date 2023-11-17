package repository

import (
	"database/sql"

	"github.com/LaQuannT/shark-api/types"
)

type SharkRepo struct {
	db *sql.DB
}

func fromRowToShark(rows *sql.Rows) (*types.Shark, error) {
	var name, t, ocean string
	var id, leng, spd, atks int

	if err := rows.Scan(&id, &name, &t, &leng, &ocean, &spd, &atks); err != nil {
		return &types.Shark{}, err
	}
	return &types.Shark{
		Id:             id,
		Name:           name,
		Type:           t,
		MaxLength:      leng,
		Ocean:          ocean,
		TopSpeed:       spd,
		AttacksPerYear: atks,
	}, nil
}

func (r *SharkRepo) CreateShark(s types.Shark) (int, error) {
	id := 0
	err := r.db.QueryRow(`
    INSERT INTO shark (name, type, max_length, ocean, top_speed, attacks_per_year)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURING id`,
		s.Name, s.Type, s.MaxLength, s.Ocean, s.TopSpeed, s.AttacksPerYear).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *SharkRepo) GetSharkByName(name string) (*types.Shark, error) {
	rows, err := r.db.Query(`SELECT * FROM shark WHERE name = $1`, name)
	if err != nil {
		return &types.Shark{}, err
	}
	defer rows.Close()

	for rows.Next() {
		return fromRowToShark(rows)
	}
	return nil, nil
}

func (r *SharkRepo) GetSharkById(id int) (*types.Shark, error) {
	rows, err := r.db.Query(`SELECT * FROM shark WHERE id = $1`, id)
	if err != nil {
		return &types.Shark{}, err
	}
	defer rows.Close()

	for rows.Next() {
		return fromRowToShark(rows)
	}
	return nil, nil
}

func (r *SharkRepo) UpdateShark(s types.Shark) error {
	_, err := r.db.Exec(`
    UPDATE shark
    SET name = $1, type = $2, max_length = $3, ocean = $4, top_speed = $5, attacks_per_year = $6,
    WHERE id = $7`,
		s.Name, s.Type, s.MaxLength, s.Ocean, s.TopSpeed, s.AttacksPerYear, s.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SharkRepo) DeleteShark(id int) error {
	_, err := r.db.Exec(`DROP FROM shark WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func NewSharkRepo(db *sql.DB) *SharkRepo {
	return &SharkRepo{db}
}
