package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/ttp/database"
)

func GetOccupationType(w http.ResponseWriter, r *http.Request, id int) {
	occType, err := db.GetOccupationType(db.DBConn, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(occType)
}

func GetAllOccupationTypes(w http.ResponseWriter, r *http.Request) {
	occTypes, err := db.GetAllOccupationTypes(db.DBConn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(occTypes)
}

func AddOccupationType(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Host)
	fmt.Println(r.URL.Fragment)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.String())
	fmt.Println(r.URL.User.Username())
	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)
	
	var occType string
	err := json.NewDecoder(r.Body).Decode(&occType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = db.AddOccupationType(db.DBConn, occType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	location := r.Host + r.URL.Path + "/" + ""
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}
