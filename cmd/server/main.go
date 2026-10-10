package main

import (
	"fmt"

	"github.com/ColleBoll/LabControl/internal/server/api"
)

func main() {
	api.NewRouter()

	fmt.Println("LabControl running on http://localhost:8080")

}
