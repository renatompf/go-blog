package main

import (
	"github.com/renatompf/initializers"
	"github.com/renatompf/models"
	"log"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToPostgres()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Something wrong when migration data to the database")
		return
	}
	log.Println("Migrations ran with no problems")
}
