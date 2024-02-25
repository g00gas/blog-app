package server

import (
	"blog-app-backend/internal/handler"
	"reflect"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	registerRoutes(r, &handler.HelloHandler{})
	r.Run()
}

func registerRoutes(r *gin.Engine, h handler.Handler) {
	handlerType := reflect.TypeOf(h)

	handlerValue := reflect.New(handlerType).Elem().Interface().(handler.Handler)

	handlerValue.RegisterRoutes(r.Group(""))
}
