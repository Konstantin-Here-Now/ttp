package database

import (
	"log"
	"time"

	"github.com/google/uuid"
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

func GetAllDefaultOccupations(db *sqlx.DB) []DefaultOccupation {
	occupations := []DefaultOccupation{}
	err := db.Select(&occupations, "SELECT * FROM default_occupation")
	if err != nil {
		log.Println(err)
	}
	return occupations
}

func AddDefaultOccupationChange(db *sqlx.DB, occup DefaultOccupationChange) {
	_, err := db.NamedExec("INSERT INTO default_occupation_changes (\"at\", \"date\") "+
		"VALUES (:at, :date) "+
		"ON CONFLICT (\"day\") DO UPDATE "+
		"SET \"at\" = :at, \"date\" = :date;",
		&occup)
	if err != nil {
		log.Println(err)
	}
}

func GetDefaultOccupationChange(db *sqlx.DB, date time.Time) DefaultOccupationChange {
	occupation := DefaultOccupationChange{}
	err := db.Get(&occupation, "SELECT * FROM default_occupation_changes WHERE \"date\" = $1", date)
	if err != nil {
		log.Println(err)
	}
	return occupation
}

func AddOccupationType(db *sqlx.DB, t OccupationType) {
	_, err := db.NamedExec("INSERT INTO occupation_type (\"type\") VALUES (:type);", &t)
	if err != nil {
		log.Println(err)
	}
}

func GetAllOccupationTypes(db *sqlx.DB) []OccupationType {
	t := []OccupationType{}
	err := db.Select(&t, "SELECT * FROM occupation_type")
	if err != nil {
		log.Println(err)
	}
	return t
}

func AddOccupation(db *sqlx.DB, o Occupation) {
	_, err := db.NamedExec("INSERT INTO occupation "+
		"(id, type_id, \"date\", start, end, desc, created_at) "+
		"VALUES (:id, :type_id, :date, :start, :end, :desc, :created_at);",
		&o)
	if err != nil {
		log.Println(err)
	}
}

func GetOccupation(db *sqlx.DB, id uuid.UUID) Occupation {
	o := Occupation{}
	err := db.Get(&o, "SELECT * FROM occupation WHERE \"id\" = $1", id)
	if err != nil {
		log.Println(err)
	}
	return o
}

func GetAllOccupations(db *sqlx.DB) []Occupation {
	o := []Occupation{}
	err := db.Select(&o, "SELECT * FROM occupation")
	if err != nil {
		log.Println(err)
	}
	return o
}

func AddMaterial(db *sqlx.DB, m Material) {
	_, err := db.NamedExec("INSERT INTO material "+
		"(id, name, desc, url) "+
		"VALUES (:id, :name, :desc, :url);",
		&m)
	if err != nil {
		log.Println(err)
	}
}

func GetMaterial(db *sqlx.DB, name string) Material {
	m := Material{}
	err := db.Get(&m, "SELECT * FROM material WHERE \"name\" = $1", name)
	if err != nil {
		log.Println(err)
	}
	return m
}

func GetAllMaterials(db *sqlx.DB) []Material {
	m := []Material{}
	err := db.Select(&m, "SELECT * FROM material")
	if err != nil {
		log.Println(err)
	}
	return m
}
