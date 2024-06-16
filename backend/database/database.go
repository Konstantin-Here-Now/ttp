package database

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gookit/config/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ttp/timing"
)

var DBConn *sqlx.DB

func Connect() *sqlx.DB {
	db, err := sqlx.Connect("pgx", "host=ttp-db port=5432 user=admin password=admin dbname=ttpdb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
	return db
}

func FillDefaultOccupation(db *sqlx.DB) error {
	dates := timing.GetNextSevenDaysDates(time.Now().Local())
	for i := range dates {
		occup := DefaultOccupation{
			Date: dates[i],
			Day:  dates[i].Weekday().String(),
			At:   config.String("DefaultOccupationAt"),
		}
		err := SetDefaultOccupation(db, occup)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func SetDefaultOccupation(db *sqlx.DB, occup DefaultOccupation) error {
	_, err := db.NamedExec("INSERT INTO default_occupation (\"day\", \"at\", \"date\") "+
		"VALUES (:day, :at, :date) "+
		"ON CONFLICT (\"day\") DO UPDATE "+
		"SET \"day\" = :day, \"at\" = :at, \"date\" = :date;",
		&occup)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetDefaultOccupation(db *sqlx.DB, dayName string) (DefaultOccupation, error) {
	occupation := DefaultOccupation{}
	err := db.Get(&occupation, "SELECT * FROM default_occupation WHERE \"day\" = $1", dayName)
	if err != nil {
		log.Println(err)
	}
	return occupation, err
}

func GetAllDefaultOccupations(db *sqlx.DB) ([]DefaultOccupation, error) {
	occupations := []DefaultOccupation{}
	err := db.Select(&occupations, "SELECT * FROM default_occupation")
	if err != nil {
		log.Println(err)
	}
	return occupations, err
}

func AddDefaultOccupationChange(db *sqlx.DB, occup DefaultOccupationChange) error {
	_, err := db.NamedExec("INSERT INTO default_occupation_changes (\"at\", \"date\") "+
		"VALUES (:at, :date) "+
		"ON CONFLICT (\"day\") DO UPDATE "+
		"SET \"at\" = :at, \"date\" = :date;",
		&occup)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetDefaultOccupationChange(db *sqlx.DB, date time.Time) (DefaultOccupationChange, error) {
	occupation := DefaultOccupationChange{}
	err := db.Get(&occupation, "SELECT * FROM default_occupation_changes WHERE \"date\" = $1", date)
	if err != nil {
		log.Println(err)
	}
	return occupation, err
}

func AddOccupationType(db *sqlx.DB, occType string) error {
	value, err := db.Exec("INSERT INTO occupation_type (type) VALUES ($1)", occType)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(value)
	return err
}

func GetOccupationType(db *sqlx.DB, id int) (OccupationType, error) {
	o := OccupationType{}
	err := db.Get(&o, "SELECT * FROM occupation_type WHERE \"id\" = $1", id)
	if err != nil {
		log.Println(err)
	}
	return o, err
}

func GetAllOccupationTypes(db *sqlx.DB) ([]OccupationType, error) {
	t := []OccupationType{}
	err := db.Select(&t, "SELECT * FROM occupation_type")
	if err != nil {
		log.Println(err)
	}
	return t, err
}

func AddOccupation(db *sqlx.DB, o Occupation) error {
	_, err := db.NamedExec("INSERT INTO occupation "+
		"(id, type_id, \"date\", start, end, desc, created_at) "+
		"VALUES (:id, :type_id, :date, :start, :end, :desc, :created_at);",
		&o)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetOccupation(db *sqlx.DB, id uuid.UUID) (Occupation, error) {
	o := Occupation{}
	err := db.Get(&o, "SELECT * FROM occupation WHERE \"id\" = $1", id)
	if err != nil {
		log.Println(err)
	}
	return o, err
}

func GetAllOccupations(db *sqlx.DB) ([]Occupation, error) {
	o := []Occupation{}
	err := db.Select(&o, "SELECT * FROM occupation")
	if err != nil {
		log.Println(err)
	}
	return o, err
}

func AddMaterial(db *sqlx.DB, m Material) error {
	_, err := db.NamedExec("INSERT INTO material "+
		"(id, name, desc, url) "+
		"VALUES (:id, :name, :desc, :url);",
		&m)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetMaterial(db *sqlx.DB, name string) (Material, error) {
	m := Material{}
	err := db.Get(&m, "SELECT * FROM material WHERE \"name\" = $1", name)
	if err != nil {
		log.Println(err)
	}
	return m, err
}

func GetAllMaterials(db *sqlx.DB) ([]Material, error) {
	m := []Material{}
	err := db.Select(&m, "SELECT * FROM material")
	if err != nil {
		log.Println(err)
	}
	return m, err
}
