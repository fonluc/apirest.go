### API Restful em Go

Em resumo, uma API Restful é uma forma de criar uma comunicação padronizada entre sistemas diferentes na web. Ela usa URLs, verbos HTTP e formatos de dados comuns para permitir que sistemas se entendam e compartilhem informações de maneira eficiente.

```go
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
```

Neste exemplo, estamos criando uma API Restful simples que lida com a entidade "User". A rota "/users" é usada para obter todos os usuários cadastrados. Quando um cliente faz uma requisição GET para essa rota, a função "getUsers" é chamada. Essa função define o cabeçalho da resposta para "application/json" e então codifica o slice "users" em JSON e envia a resposta de volta para o cliente.

Para testar a API, você pode executar o código acima e fazer uma requisição GET para "**http://localhost:8080/users**". Você verá a lista de usuários retornada em formato JSON.
