package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Azmi117/API-TV.git/internal/models"
	"github.com/Azmi117/API-TV.git/internal/pkg/apperror"
	"github.com/Azmi117/API-TV.git/internal/usecase"
)

type TvHandler struct {
	usecase *usecase.TvUsecase
}

func NewTvHandler(params *usecase.TvUsecase) *TvHandler {
	return &TvHandler{
		usecase: params,
	}
}

func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if appErr, ok := err.(*apperror.Apperror); ok {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error"})
}

func (h *TvHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.usecase.GetAll()

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TvHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	res, err := h.usecase.GetById(id)

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TvHandler) Create(w http.ResponseWriter, r *http.Request) {
	// 1. Decode struct from request body
	var input models.Tv
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&input)

	// 2. panggil usecase
	res, err := h.usecase.Create(input)

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TvHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	var input models.Tv
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&input)

	res, err := h.usecase.Update(id, input)

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *TvHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	err := h.usecase.Delete(id)

	if err != nil {
		SendError(w, err)
		return
	}

	fmt.Println("Delete Success!")
}
