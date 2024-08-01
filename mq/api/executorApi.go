package api

import (
	"dominant/mq/message"
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
	id := c.Query("id")
	fmt.Println(ip)
	msg := b.GetMessage(id)
	c.JSON(http.StatusOK, msg)
}

func Register(c *gin.Context) {
	ip := c.ClientIP()
	body := make(map[string]any)
	if err := c.ShouldBind(&body); err == nil {
		//获取请求体中json数据
		id := body["id"].(string)
		b.Register(id, ip)
		msg := message.NewMessage("Server", []string{id}, "Alive Success!")
		c.JSON(http.StatusOK, msg)
	}
}
