package main

import (
	"fmt"

	"github.com/martinsmessias/go_rest_api/database"
	"github.com/martinsmessias/go_rest_api/routes"
)

func main() {
	fmt.Println("Starting the application...")
	database.ConnectDatabase()
	fmt.Println("Iniciando servidor em http://localhost:8080")
	routes.HandleRequests()
}
