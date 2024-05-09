package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home/", func(w http.ResponseWriter, r *http.Request) { fmt.Print("It works!") })

    http.ListenAndServe(":5432", nil)
}
