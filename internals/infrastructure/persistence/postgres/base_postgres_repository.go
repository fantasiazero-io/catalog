package postgres

import (
	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func CreateConnection() (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", ConnectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func CloseConnection(db *sqlx.DB) {
	if db != nil {
		error := db.Close()
		if error != nil {
			log.Fatal(error)
		}
	}
}
