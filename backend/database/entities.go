package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Entity DefaultOccupation

type DefaultOccupation struct {
	Id   int       `db:"id"`
	Day  string    `db:"day"`
	At   string    `db:"at"`
	Date time.Time `db:"date"`
}

type DefaultOccupationChange struct {
	Id   int       `db:"id"`
	At   string    `db:"at"`
	Date time.Time `db:"date"`
}

type OccupationType struct {
	Id   int    `db:"id"`
	Type string `db:"type"`
}

type Occupation struct {
	Id        uuid.UUID      `db:"id"`
	TypeId    int            `db:"type_id"`
	Date      time.Time      `db:"date"`
	Start     time.Time      `db:"start"`
	End       time.Time      `db:"end"`
	Desc      sql.NullString `db:"desc"`
	CreatedAt time.Time      `db:"created_at"`
}

type Material struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Desc string `db:"desc"`
	Url  string `db:"url"`
}
