package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

type Material struct {
	Id   int
	Name string
	Desc string
	Url  string
}

var testMaterials []Material

func main() {
	port := 5432

	prepareTestData()

	router := mux.NewRouter()
	router.HandleFunc("/timetable/", getTimetable).Methods("GET")
	router.HandleFunc("/materials/", getMaterials).Methods("GET")
	router.HandleFunc("/materials/{id:[0-9]+}", getMaterial).Methods("GET")

	log.Printf("Server stated listening on port %d", port)
	addr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func prepareTestData() {
	testMaterials = append(testMaterials, Material{Id: 1, Name: "Book", Desc: "Interesting"})
	testMaterials = append(testMaterials, Material{Id: 2, Name: "Paper", Desc: "Wonder"})
	testMaterials = append(testMaterials, Material{Id: 3, Name: "Video", Desc: ""})
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

func getMaterials(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(testMaterials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getMaterial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		panic(parseErr)
	}

	materialId := slices.IndexFunc(testMaterials, func(m Material) bool { return m.Id == id })
	if materialId == -1 {
		http.NotFound(w, r)
		return
	}

	material := testMaterials[materialId]
	err := json.NewEncoder(w).Encode(material)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
