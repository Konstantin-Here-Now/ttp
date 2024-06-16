package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	db "github.com/ttp/database"
)

func GetTimetable(w http.ResponseWriter, r *http.Request) {
	defaultOccups, err := db.GetAllDefaultOccupations(db.DBConn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(defaultOccups)
}

func OccTypeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		strId := mux.Vars(r)["id"]
		if strId != "" {
			id, err := strconv.Atoi(strId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			GetOccupationType(w, r, id)
		} else {
			GetAllOccupationTypes(w, r)
		}
	case http.MethodPost:
		AddOccupationType(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}