package router

import (
	"github.com/gin-gonic/gin"
  "chat-diary-server/interface/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	h := handler.NewHelloHandler()
	r.GET("/hello", h.Hello)
	return r
}
