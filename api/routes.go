package api

import (
	"database/sql"
	"lilurl/api/handlers"
	"lilurl/api/repositories"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB) {

	urlRepo := repositories.NewUrlRepository(db)
	urlHandler := &handlers.UrlHandler{Repo: urlRepo}

	mux.HandleFunc("POST /short", urlHandler.Shortner)
	mux.HandleFunc("GET /{shortId}", handlers.UrlRedirectHandler)
}
