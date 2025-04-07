package main

import (
	"go-app/config"
	"go-app/routes"
	"net/http"
)

func main() {
	config.ConnectDB()

	mux := http.NewServeMux()

	// Servir os arquivos HTML e estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	// Rotas da API
	routes.RegisterRoutes(mux)

	http.ListenAndServe(":8000", mux)
}
