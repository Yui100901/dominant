package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	IP      = "127.0.0.1" //控制端连接IP
	Port    = "27777"     //控制端连接端口
	BaseUrl = "http://" + IP + ":" + Port
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}

func NewServer() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	return r
}
