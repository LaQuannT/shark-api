package repository

import (
	"database/sql"

	"github.com/LaQuannT/shark-api/types"
)

type UserRepo struct {
	db *sql.DB
}

func fromRowToUser(row *sql.Rows) (*types.User, error) {
	var name, email, key string
	var id, permissionLvl int

	if err := row.Scan(&id, &name, &email, &key, &permissionLvl); err != nil {
		return &types.User{}, err
	}

	return &types.User{
		Id:              id,
		Name:            name,
		Email:           email,
		PermissionLevel: permissionLvl,
		ApiKey:          key,
	}, nil
}

func (r *UserRepo) CreateUser(u types.User) (int, error) {
	id := 0
	err := r.db.QueryRow(`
    INSERT INTO 
    "user"(name, email, permission_level, api_key)
    VALUES ($1, $2, $3, $4)
    RETURNING id`,
		u.Name, u.Email, u.PermissionLevel, u.ApiKey).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *UserRepo) GetUser(id int) (*types.User, error) {
	rows, err := r.db.Query(`SELECT * FROM "user" WHERE id = $1`, id)
	if err != nil {
		return &types.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		return fromRowToUser(rows)
	}

	return nil, nil
}

func (r *UserRepo) UpdateUser(u types.User) error {
	_, err := r.db.Exec(`
    UPDATE "user"
    SET name = $1, email = $2, permission_level = $3, api_key = $4,
    WHERE id = $5`,
		u.Name, u.Email, u.PermissionLevel, u.ApiKey, u.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteUser(id int) error {
	_, err := r.db.Exec(`DELETE FROM "users" WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}
