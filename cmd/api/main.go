package main

import (
	"fmt"
	"net/http"

	"github.com/Azmi117/API-TV.git/internal/config"
	delivery "github.com/Azmi117/API-TV.git/internal/delivery/http"
	"github.com/Azmi117/API-TV.git/internal/repository"
	"github.com/Azmi117/API-TV.git/internal/usecase"
)

func main() {
	db := config.ConnectDB()

	tvRepo := repository.NewTvRepository(db)

	tvUsecase := usecase.NewTvUsecase(tvRepo)

	tvHandler := delivery.NewTvHandler(tvUsecase)

	mux := http.NewServeMux()

	delivery.MapRoutes(mux, tvHandler)

	port := ":8080"

	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Failed run server : ", err)
	}
}
