package api

import (
	"dominant/mq/message"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

//
// @Author yfy2001
// @Date 2024/8/1 12 40
//

func NewMessage(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	msg := &message.Message{}
	err := msg.MessageJsonUnMarshal(body)
	if err != nil {
		return
	}
	b.Distribute(msg)
	fmt.Println(msg)
	c.JSON(http.StatusOK, gin.H{"message": "Success!"})
}

func GetClientList(c *gin.Context) {
	aliveList := b.ListNodes()
	msg := message.NewMessage("Server", []string{}, aliveList)
	c.JSON(http.StatusOK, msg)
}
