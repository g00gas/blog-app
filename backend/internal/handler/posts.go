package handler

import (
	"blog-app-backend/internal/models"
	"blog-app-backend/internal/repository/postsRepository"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostsHandler struct{}

func (h *PostsHandler) RegisterRoutes(r *gin.RouterGroup) {
	postsGroup := r.Group("/api/posts")
	{
		postsGroup.GET("/", h.getAllPosts)
		postsGroup.GET("/:id", h.getPostById)
		postsGroup.POST("/", h.createPost)
		postsGroup.DELETE("/:id", h.deletePost)
	}
}

func (h *PostsHandler) getAllPosts(ctx *gin.Context) {
	posts, err := postsRepository.GetAllPosts(ctx)
	if err != nil {
		log.Panicf("Error retrieving posts: %v", err)
		return
	}
	ctx.JSON(200, posts)
}

func (h *PostsHandler) getPostById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		post, err := postsRepository.GetPostById(ctx, id)
		if err != nil {
			log.Panicf("Error retrieving post: %v", err)
			return
		}
		ctx.JSON(200, post)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Message with %s not found!", idParam)})
	}
}

func (h *PostsHandler) createPost(ctx *gin.Context) {
	var newPost models.CreatePostRequest
	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		log.Panicf("Error parsing JSON %v\n", err)
		return
	} else {
		post, err := postsRepository.CreateNewPost(ctx, newPost)
		if err != nil {
			log.Panicf("Error creating new post: %v", err)
			return
		}
		ctx.JSON(http.StatusCreated, post)
	}
}

func (h *PostsHandler) deletePost(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if id, err := strconv.Atoi(idParam); err == nil {
		res, err := postsRepository.DeletePostById(ctx, id)
		if err != nil {
			log.Panicf("Error deleting a post: %v", err)
			return
		}
		if res {
			ctx.Status(http.StatusNoContent)
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Post with %s not found!", idParam)})
	}
}
