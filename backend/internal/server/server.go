package server

import (
	"blog-app-backend/internal/config"
	"blog-app-backend/internal/handler"
	"blog-app-backend/internal/middleware"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	db := config.InitDB()
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.UseDatabase(db))
	registerRoutes(r, &handler.PostsHandler{})
	err := r.Run("localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
}

func registerRoutes(r *gin.Engine, h handler.Handler) {
	handlerType := reflect.TypeOf(h)

	handlerValue := reflect.New(handlerType).Elem().Interface().(handler.Handler)

	handlerValue.RegisterRoutes(r.Group(""))
}
