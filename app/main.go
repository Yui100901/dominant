package main

import (
	"dominant/api/http_api"
	"dominant/api/mqtt_api"
	"dominant/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := server.NewServer()
	r.POST("/newMessage", http_api.NewMessage)
	r.GET("/getClientList", http_api.GetClientList)
	//获取节点状态
	r.GET("/getNodeStatusList", mqtt_api.GetNodeStatusList)
	//获取节点状态-WebSocket
	r.GET("/wsGetNodeStatusList", func(c *gin.Context) {
		http_api.ServeWebSocket(c.Writer, c.Request)
	})
	//执行器相关接口
	r.GET("/getMessage", http_api.GetMessage)
	r.POST("/register", http_api.Register)
	err := r.Run(fmt.Sprintf(":%s", server.Port))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
