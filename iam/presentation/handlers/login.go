package handlers

import (
	"katsan/iam/presentation/pages"
	"net/http"
)

func HandleGetLoginPage(w http.ResponseWriter, r *http.Request) {
	pages.LoginPage().Render(r.Context(), w)
}
