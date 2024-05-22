package controllers

import (
	"article-api/models"
	"article-api/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	PostService services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{service}
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	posts, err := pc.PostService.GetAllPosts(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (pc *PostController) GetPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := pc.PostService.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var postRequest models.PostCreateRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := postRequest.ValidateCreate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:     postRequest.Title,
		Content:   postRequest.Content,
		Category:  postRequest.Category,
		Status:    postRequest.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := pc.PostService.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID post tidak valid"})
		return
	}

	var postRequest models.PostEditRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := postRequest.ValidateUpdate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := pc.PostService.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if postRequest.Title != "" {
		post.Title = postRequest.Title
	}
	if postRequest.Content != "" {
		post.Content = postRequest.Content
	}
	if postRequest.Category != "" {
		post.Category = postRequest.Category
	}
	if postRequest.Status != "" {
		post.Status = postRequest.Status
	}
	post.UpdatedAt = time.Now()

	if err := pc.PostService.UpdatePost(uint(id), post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *PostController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	if err := pc.PostService.DeletePost(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
