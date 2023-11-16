package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/LaQuannT/shark-api/database/migration"
)

func connect(path string) (*sql.DB, error) {
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Init(path string) (*sql.DB, error) {
	log.Println("Connection to database...")
	db, err := connect(path)
	if err != nil {
		return nil, err
	}

	log.Println("Building/Checking tables...")
	err = migration.Up(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
