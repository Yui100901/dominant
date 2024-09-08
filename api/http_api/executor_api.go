package http_api

import (
	"dominant/broker"
	"dominant/mq"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// @Author yfy2001
// @Date 2024/8/1 12 39
//

func GetMessage(c *gin.Context) {
	ip := c.ClientIP()
	nodeId := c.Query("nodeId")
	fmt.Println(ip)
	msg := broker.GlobalBroker.GetMessage(nodeId)
	c.JSON(http.StatusOK, msg)
}

func Register(c *gin.Context) {
	ipAddr := c.ClientIP()
	body := make(map[string]any)
	if err := c.ShouldBind(&body); err == nil {
		//获取请求体中json数据
		id := body["id"].(string)
		broker.GlobalBroker.Register(id, ipAddr, []byte(""))
		msg := mq.NewMessage("", "", "Server", []string{id}, "Alive Success!")
		c.JSON(http.StatusOK, msg)
	}
}
