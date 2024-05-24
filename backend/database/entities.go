package database

import (
	"time"
)

type DefaultOccupation struct {
	Id int `db:"id"`
	Day string `db:"day"`
	At string `db:"at"`
	Date time.Time `db:"date"`
}