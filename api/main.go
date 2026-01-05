package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type errorResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

var (
	users = make(map[string]User)
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/createuser", addUser)
	http.HandleFunc("/user", getUsers)
	http.HandleFunc("/health", health)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResp{
			StatusCode: 400,
			Message:    "Invalid payload",
		})
		return
	}

	mu.Lock()
	users[user.Name] = user
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
