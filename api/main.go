package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", getUsers)
	fmt.Println("Api is running")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]Users{
		{ID: 1, Name: "John"},
	})
}
