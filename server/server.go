package server

import "github.com/gin-gonic/gin"

const (
	IP      = "127.0.0.1" //控制端连接IP
	Port    = "7700"      //控制端连接端口
	BaseUrl = "http://" + IP + ":" + Port
)

func NewServer() *gin.Engine {
	r := gin.Default()
	return r
}
