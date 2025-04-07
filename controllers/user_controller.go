package controllers

import (
	"encoding/json"
	"go-app/config"
	"go-app/models"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (name, email, password, status) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.Status)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário adicionado com sucesso"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			continue
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
