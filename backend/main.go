package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gookit/config/v2"
	"github.com/gorilla/mux"
	db "github.com/ttp/database"
	"github.com/ttp/handlers"
)

func main() {
	config.AddDriver(config.JSONDriver)
	err := config.LoadFiles("settings.json")
	if err != nil {
		panic(err)
	}

	db.DBConn = db.Connect()
	err = db.FillDefaultOccupation(db.DBConn)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/timetable", handlers.GetTimetable).Methods("GET")
	router.HandleFunc("/occtype", handlers.OccTypeHandler).Methods("GET", "POST")
	router.HandleFunc("/occtype/{id}", handlers.OccTypeHandler).Methods("GET", "PUT")

	port := config.Int("Port")
	addr := ":" + strconv.Itoa(port)
	log.Printf("Server stated listening on port %d", port)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
