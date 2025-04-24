package handlers

import (
	"fmt"
	"lilurl/api/entities"
	"lilurl/api/repositories"
	"lilurl/utils"
	"net/http"
)

type UrlHandler struct {
	Repo *repositories.UrlRepository
}

func (h *UrlHandler) Shortner(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Query().Get("url") == "" {
		http.Error(w, "Url not informed", http.StatusBadRequest)
		return
	}

	generatedShortId := utils.GenerateShortId()

	generatedUrl := "http://localhost:8080/" + generatedShortId

	url := &entities.Url{
		ShortId: generatedShortId,
		Url:     r.URL.Query().Get("url"),
	}

	err := h.Repo.Create(url)

	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(generatedUrl))
}
