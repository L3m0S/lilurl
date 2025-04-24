package api

import (
	"database/sql"
	"log"
	"net/http"
)

func StartHttpServer(db *sql.DB) {
	mux := http.NewServeMux()
	SetupRoutes(mux, db)
	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
