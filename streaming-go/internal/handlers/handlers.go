package handlers

import (
	"fmt"
	"net/http"

	"streaming/internal/models"
	"streaming/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.StreamingService
}

const (
	usersPath           = "/users"
	contentPath         = "/content"
	playPath            = "/play"
	historyPath         = "/history/:id"
	editContentPath     = "/content/:id/edit"
	updateContentPath   = "/content/:id/update"
	deleteContentPath   = "/content/:id/delete"
	indexPath           = "/"
	errorTemplate       = "error.html"
	usersTemplate       = "users.html"
	contentTemplate     = "content.html"
	playTemplate        = "play.html"
	historyTemplate     = "history.html"
	indexTemplate       = "index.html"
	editContentTemplate = "edit_content.html"
)

func RegisterRoutes(r *gin.Engine, service *services.StreamingService) {

	h := &Handler{service: service}

	// Web routes
	r.GET(indexPath, h.Index)
	r.GET(usersPath, h.UsersPage)
	r.GET(contentPath, h.ContentPage)
	r.GET(playPath, h.PlayPage)
	r.GET(historyPath, h.HistoryPage)
	r.GET(editContentPath, h.EditContentPage)
	r.POST(updateContentPath, h.UpdateContentWeb)
	r.POST(deleteContentPath, h.DeleteContentWeb)

	r.POST(usersPath, h.CreateUserWeb)
	r.POST(contentPath, h.CreateContentWeb)
	r.POST(playPath, h.PlayContentWeb)

	api := r.Group("/api")

	api.POST(usersPath, h.CreateUser)
	api.GET(usersPath, h.GetUsers)

	api.POST(contentPath, h.CreateContent)
	api.GET(contentPath, h.GetContent)

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

func (h *Handler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, indexTemplate, gin.H{})
}

func (h *Handler) UsersPage(c *gin.Context) {
	users, _ := h.service.GetUsers()
	var userResponses []models.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, u.ToResponse())
	}
	c.HTML(http.StatusOK, usersTemplate, gin.H{"Users": userResponses})
}

func (h *Handler) ContentPage(c *gin.Context) {
	content, _ := h.service.GetContent()
	var contentResponses []models.ContentResponse
	for _, c := range content {
		contentResponses = append(contentResponses, c.ToResponse())
	}
	c.HTML(http.StatusOK, contentTemplate, gin.H{"Content": contentResponses})
}

func (h *Handler) PlayPage(c *gin.Context) {
	c.HTML(http.StatusOK, playTemplate, gin.H{})
}

func (h *Handler) HistoryPage(c *gin.Context) {
	id := c.Param("id")
	history, _ := h.service.GetHistory(id)

	var historyResponses []models.ContentResponse
	for _, c := range history {
		historyResponses = append(historyResponses, c.ToResponse())
	}

	c.HTML(http.StatusOK, historyTemplate, gin.H{"UserID": id, "History": historyResponses})
}

func (h *Handler) EditContentPage(c *gin.Context) {
	id := c.Param("id")
	content, err := h.service.GetContentByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, errorTemplate, gin.H{"error": "Contenido no encontrado"})
		return
	}
	c.HTML(http.StatusOK, editContentTemplate, gin.H{"Content": content.ToResponse()})
}

func (h *Handler) UpdateContentWeb(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	typeStr := c.PostForm("type")
	durationStr := c.PostForm("duration")

	typeInt := 0
	fmt.Sscanf(typeStr, "%d", &typeInt)
	ctype := models.ContentType(typeInt)

	duration := 0
	fmt.Sscanf(durationStr, "%d", &duration)

	if title != "" {
		h.service.UpdateContentTitle(id, title)
	}
	if duration > 0 {
		h.service.UpdateContentDuration(id, duration)
	}
	h.service.UpdateContentType(id, ctype)

	c.Redirect(http.StatusFound, contentPath)
}

func (h *Handler) DeleteContentWeb(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteContent(id)
	if err != nil {
		c.HTML(http.StatusBadRequest, errorTemplate, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, contentPath)
}

func (h *Handler) CreateUserWeb(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	email := c.PostForm("email")

	_, err := h.service.RegisterUser(id, name, email)
	if err != nil {
		c.HTML(http.StatusBadRequest, errorTemplate, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, usersPath)
}

func (h *Handler) CreateContentWeb(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	typeStr := c.PostForm("type")
	durationStr := c.PostForm("duration")

	typeInt := 0
	fmt.Sscanf(typeStr, "%d", &typeInt)
	ctype := models.ContentType(typeInt)

	duration := 0
	fmt.Sscanf(durationStr, "%d", &duration)

	_, err := h.service.AddContent(id, title, ctype, duration)
	if err != nil {
		c.HTML(http.StatusBadRequest, errorTemplate, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, contentPath)
}

func (h *Handler) PlayContentWeb(c *gin.Context) {
	userID := c.PostForm("user_id")
	contentID := c.PostForm("content_id")

	_, err := h.service.Play(userID, contentID)
	if err != nil {
		c.HTML(http.StatusBadRequest, errorTemplate, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, indexPath)
}
