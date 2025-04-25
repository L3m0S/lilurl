package api

import (
	"lilurl/api/handlers"
	"lilurl/api/repositories"
	"lilurl/database"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, db *database.DataBase) {

	urlRepo := repositories.NewUrlRepository(db.Instance)
	urlHandler := &handlers.UrlHandler{Repo: urlRepo}

	mux.HandleFunc("POST /short", urlHandler.Shortner)
	mux.HandleFunc("GET /{shortId}", handlers.UrlRedirectHandler)
}
