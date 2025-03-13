package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// User represents a user
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = map[string]User{
	"1": {ID: "1", Username: "john", Email: "john@example.com"},
	"2": {ID: "2", Username: "jane", Email: "jane@example.com"},
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updateUser User
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := users[id]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	users[id] = updateUser

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", patchUser).Methods("PATCH")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":9000", r)
}
