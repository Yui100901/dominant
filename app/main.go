package main

import (
	"dominant/api/http_api"
	"dominant/api/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := server.NewServer()
	//创建一则新消息
	r.POST("/newMessage", http_api.NewMessage)
	r.GET("/getClientList", http_api.GetClientList)
	//获取节点状态
	r.GET("/getNodeStatusList", http_api.GetNodeStatusList)
	//获取节点状态-WebSocket
	r.GET("/wsGetNodeStatusList", func(c *gin.Context) {
		http_api.ServeWebSocket(c.Writer, c.Request)
	})
	//执行器相关接口
	r.GET("/getMessage", http_api.GetMessage)
	r.POST("/login", http_api.Login)
	r.POST("/connect", http_api.Connect)
	err := r.Run(fmt.Sprintf(":%s", server.Port))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
