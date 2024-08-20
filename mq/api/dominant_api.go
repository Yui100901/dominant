package api

import (
	"dominant/broker"
	"dominant/mq/message"
	"dominant/server"
	"encoding/json"
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
	err := json.Unmarshal(body, msg)
	if err != nil {
		return
	}
	broker.GlobalBroker.MainMQ.Enqueue(msg)
	fmt.Println(msg)
	c.JSON(http.StatusOK, gin.H{"message": "Success!"})
}

func GetClientList(c *gin.Context) {
	aliveList, _ := broker.GlobalBroker.GetAliveNodeIDList()
	msg := message.NewMessage("", "", "Server", []string{}, aliveList)
	c.JSON(http.StatusOK, msg)
}

func ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	ws := server.NewWebSocket(w, r)
	defer ws.Close()

	go ws.OnMessage(nil)
	go ws.PushMessage(broker.GlobalBroker.GetAliveNodeMessage)

	<-ws.Done
}
