package handlers

import (
	"katsan/traceability/presentation/pages"
	"net/http"
)

func HandleGetShipmentsPage(w http.ResponseWriter, r *http.Request) {
	pages.ShipmentsPage().Render(r.Context(), w)
}
