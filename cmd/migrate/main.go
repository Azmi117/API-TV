package main

import (
	"fmt"

	"github.com/Azmi117/API-TV.git/internal/config"
	"github.com/Azmi117/API-TV.git/internal/models"
)

func main() {
	db := config.ConnectDB()

	err := db.AutoMigrate(&models.Tv{})

	if err != nil {
		fmt.Printf("Migration failed : %s", err)
		return
	}

	fmt.Println("Migration Success!")
}
