package main

import (
	"github.com/joho/godotenv"
	"github.com/therobby/simpleblog/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	gin.SetMode(os.Getenv("GIN_MODE"))
	
	r := gin.Default()
	routes.HandleRoutes(r)
	err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}