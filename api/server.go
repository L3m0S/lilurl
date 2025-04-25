package api

import (
	"lilurl/database"
	"log"
	"net/http"
)

func StartHttpServer(db *database.DataBase) {
	mux := http.NewServeMux()
	SetupRoutes(mux, db)
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
