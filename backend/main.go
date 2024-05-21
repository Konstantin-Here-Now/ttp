package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ttp/timing"
)

type Material struct {
	Id   int
	Name string
	Desc string
	Url  string
}

var testMaterials []Material
var testTimetable timing.Timetable

func main() {
	port := 7777

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

func readJsonTimetable() {
	dat, readErr := os.ReadFile("./timetable.json")
	if readErr != nil {
		log.Fatal("No timetable.json found")
	}
	parseErr := json.Unmarshal(dat, &testTimetable)
	if parseErr != nil {
		log.Fatal("Invalid json", parseErr)
	}
}

func getTimetable(w http.ResponseWriter, r *http.Request) {
	readJsonTimetable()
	err := json.NewEncoder(w).Encode(testTimetable)
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
	id, _ := strconv.Atoi(vars["id"])

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
