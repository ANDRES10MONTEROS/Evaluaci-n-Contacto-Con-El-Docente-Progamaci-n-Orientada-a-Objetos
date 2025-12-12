package handlers

import (
	"net/http"

	"streaming/internal/models"
	"streaming/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.StreamingService
}

func RegisterRoutes(r *gin.Engine, service *services.StreamingService) {

	h := &Handler{service: service}

	api := r.Group("/api")

	api.POST("/users", h.CreateUser)
	api.GET("/users", h.GetUsers)

	api.POST("/content", h.CreateContent)
	api.GET("/content", h.GetContent)

	api.POST("/play", h.PlayContent)
	api.GET("/history/:id", h.GetHistory)

	api.PUT("/content/:id/title", h.UpdateTitle)
	api.DELETE("/content/:id", h.DeleteContent)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.RegisterUser(req.ID, req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user.ToResponse())
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, _ := h.service.GetUsers()

	var out []models.UserResponse
	for _, u := range users {
		out = append(out, u.ToResponse())
	}

	c.JSON(http.StatusOK, out)
}

func (h *Handler) CreateContent(c *gin.Context) {
	var req struct {
		ID       string             `json:"id"`
		Title    string             `json:"title"`
		Type     models.ContentType `json:"type"`
		Duration int                `json:"duration"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := h.service.AddContent(req.ID, req.Title, req.Type, req.Duration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, content.ToResponse())
}

func (h *Handler) GetContent(c *gin.Context) {
	content, _ := h.service.GetContent()

	var out []models.ContentResponse
	for _, c := range content {
		out = append(out, c.ToResponse())
	}

	c.JSON(http.StatusOK, out)
}

func (h *Handler) PlayContent(c *gin.Context) {
	var req struct {
		UserID    string `json:"user_id"`
		ContentID string `json:"content_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.service.Play(req.UserID, req.ContentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *Handler) GetHistory(c *gin.Context) {
	id := c.Param("id")
	history, _ := h.service.GetHistory(id)

	var out []models.ContentResponse
	for _, c := range history {
		out = append(out, c.ToResponse())
	}

	c.JSON(http.StatusOK, out)
}

func (h *Handler) UpdateTitle(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Title string `json:"title"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateContentTitle(id, req.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "title updated"})
}

func (h *Handler) DeleteContent(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteContent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "content deleted"})
}
