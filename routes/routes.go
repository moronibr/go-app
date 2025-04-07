package routes

import (
	"go-app/controllers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/users", controllers.GetUsers)
	mux.HandleFunc("/api/add-user", controllers.AddUser) // <- nova rota POST
}
