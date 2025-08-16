package main

import (
	"log"

	"phonebook/internal/router"
	pkg "phonebook/pkg/db"

	_ "phonebook/docs"
)

// @title           Phonebook API
// @version         1.0
// @description     Simple phonebook API (Gin + PostgreSQL).
// @host            localhost:8080
// @BasePath        /
func main() {
	// Initialize the database
	pkg.ConnectDB()

	// 2) build gin router with routes + swagger
	r := router.SetupRouter()

	log.Println("server running on http://:8080")

	// 3) start http server
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
