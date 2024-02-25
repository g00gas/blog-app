package handler

import "github.com/gin-gonic/gin"

type HelloHandler struct{}

func (h *HelloHandler) RegisterRoutes(r *gin.RouterGroup) {
	helloGroup := r.Group("/hello")
	{
		helloGroup.GET("/1", h.handleGet)
		helloGroup.GET("/2", h.handleGetSecond)
	}
}

func (h *HelloHandler) handleGet(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello World!"})
}

func (h *HelloHandler) handleGetSecond(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello World!!"})
}
