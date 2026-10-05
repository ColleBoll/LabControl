package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ColleBoll/LabControl/internal/server/api"
)

func main() {
	router := api.NewRouter()

	fmt.Println("LabControl running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
