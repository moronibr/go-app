package main

import (
	"fmt"
	"log"
	"net/http"

	"go-app/db"
)

func main() {
	// Conectar ao banco
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Servir arquivos estáticos (HTML/CSS/JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("🚀 Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
