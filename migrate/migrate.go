package main

import (
	"fmt"
	"log"

	"github.com/bebop/ark/initializers"
	"github.com/bebop/ark/models"
)

func init() {
	config, err := initializers.LoadConfig(".devcontainer")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
