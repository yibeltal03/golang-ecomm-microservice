// user_service.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

var users = make(map[string]string) // Simple in-memory user storage

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func registerUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    users[user.Username] = user.Password
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "User %s registered successfully", user.Username)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]
    password, exists := users[username]
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"username": username, "password": password})
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/register", registerUser).Methods("POST")
    r.HandleFunc("/user/{username}", getUser).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", r))
}
