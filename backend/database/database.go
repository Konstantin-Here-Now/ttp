package database

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Connect("pgx", "host=ttp-db port=5432 user=admin password=admin dbname=ttpdb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
	return db
}

func SetDefaultOccupation(db *sqlx.DB, occup DefaultOccupation) {
	_, err := db.NamedExec("INSERT INTO default_occupation (\"day\", \"at\", \"date\") "+
		"VALUES (:day, :at, :date) "+
		"ON CONFLICT (\"day\") DO UPDATE "+
		"SET \"day\" = :day, \"at\" = :at, \"date\" = :date;",
		&occup)
	if err != nil {
		log.Println(err)
	}
}

func GetDefaultOccupation(db *sqlx.DB, dayName string) DefaultOccupation {
	occupation := DefaultOccupation{}
	err := db.Get(&occupation, "SELECT * FROM default_occupation WHERE \"day\" = $1", dayName)
	if err != nil {
		log.Println(err)
	}
	return occupation
}

func GetDefaultOccupations(db *sqlx.DB) []DefaultOccupation {
	occupations := []DefaultOccupation{}
	err := db.Select(&occupations, "SELECT * FROM default_occupation")
	if err != nil {
		log.Println(err)
	}
	return occupations
}
