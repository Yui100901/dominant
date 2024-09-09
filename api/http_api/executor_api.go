package http_api

import (
	"dominant/broker"
	"dominant/mq"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

func Login(c *gin.Context) {
	ipAddr := c.ClientIP()
	body := make(map[string]any)
	if err := c.ShouldBind(&body); err == nil {
		//获取请求体中json数据
		id := body["id"].(string)
		log.Println("接收到id:", id)
		token := broker.GlobalBroker.Login(id, ipAddr, []byte("Login test"))
		msg := mq.NewMessage("", "", "Server", []string{id}, token)
		c.JSON(http.StatusOK, msg)
	}
}

func Verify(c *gin.Context) {
	//ipAddr := c.ClientIP()
	var cmd VerifyCommand
	if err := c.ShouldBind(&cmd); err == nil {
		//获取请求体中json数据
		id := cmd.ID
		token := cmd.Token
		flag := broker.GlobalBroker.Verify(id, token, []byte("Verify test"))
		msg := mq.NewMessage("", "", "Server", []string{id}, flag)
		c.JSON(http.StatusOK, msg)
	}
}
