package main

import (
	"github.com/gin-gonic/gin"
	"streaming/internal/handlers"
	"streaming/internal/repos"
	"streaming/internal/services"
)

func main() {

	repo := repos.NewMemoryRepo()
	service := services.NewStreamingService(repo, repo)

	r := gin.Default()

	handlers.RegisterRoutes(r, service)

	r.Run(":8080")
}
