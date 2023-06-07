package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	// Adicione alguns usuários para teste
	users = append(users, User{ID: 1, Name: "John", Email: "john@example.com", Password: "123456"})
	users = append(users, User{ID: 2, Name: "Jane", Email: "jane@example.com", Password: "abcdef"})

	// Configuração das rotas
	http.HandleFunc("/users", getUsers)

	// Inicia o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
