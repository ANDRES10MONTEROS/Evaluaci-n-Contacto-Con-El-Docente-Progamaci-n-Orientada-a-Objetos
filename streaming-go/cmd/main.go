package main

import (
	"streaming/internal/handlers"
	"streaming/internal/repos"
	"streaming/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := repos.NewMemoryRepo()
	service := services.NewStreamingService(repo, repo)

	r := gin.Default()

	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("templates/*")

	handlers.RegisterRoutes(r, service)

	r.Run(":8082")
}
