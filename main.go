package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Day struct {
	At   string
	Date time.Time
}

type Timetable struct {
	Mon Day
	Tue Day
	Wed Day
	Thu Day
	Fri Day
	Sat Day
	Sun Day
}

func main() {
	port := 5432
	http.HandleFunc("/timetable/", timetableHandler)

	log.Printf("Server stated listening on port %d", port)
	addr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func timetableHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTimetable(w, r)
	default:
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
	}
}

func getTimetable(w http.ResponseWriter, r *http.Request) {
	testDay := Day{At: "12:00 - 13:00", Date: time.Now()}
	testTT := Timetable{Mon: testDay, Tue: testDay, Wed: testDay, Thu: testDay, Fri: testDay, Sat: testDay, Sun: testDay}
	err := json.NewEncoder(w).Encode(testTT)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
