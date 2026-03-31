package http

import "net/http"

func MapRoutes(mux *http.ServeMux, h *TvHandler) {
	mux.HandleFunc("GET /tv", h.GetAll)
	mux.HandleFunc("GET /tv/{id}", h.GetById)
	mux.HandleFunc("POST /tv", h.Create)
	mux.HandleFunc("PATCH /tv/{id}", h.Update)
	mux.HandleFunc("DELETE /tv/{id}", h.Delete)
}
