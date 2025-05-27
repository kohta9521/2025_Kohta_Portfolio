package handlers

import (
	"net/http"

	"portfolio-backend/pkg/aws"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	s3Client *aws.S3Client
}

func NewBlogHandler(s3Client *aws.S3Client) *BlogHandler {
	return &BlogHandler{
		s3Client: s3Client,
	}
}

func (h *BlogHandler) GetBlogPosts(c *gin.Context) {
	posts, err := h.s3Client.GetBlogPosts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
} 