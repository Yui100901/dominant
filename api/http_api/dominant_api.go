package http_api

import (
	"dominant/api/server"
	"dominant/domain/broker"
	"dominant/infrastructure/messaging/mq"
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
	msg := &mq.Message{}
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
	msg := mq.NewMessage("", "", "Server", []string{}, aliveList)
	c.JSON(http.StatusOK, msg)
}

func ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	ws := server.NewWebSocket(w, r)
	defer ws.Close()

	go ws.OnMessage(nil)
	go ws.PushMessage(broker.GlobalBroker.GetAliveNodeMessage)

	<-ws.Done
}

func GetNodeStatusList(c *gin.Context) {
	msgList := broker.GlobalBroker.GetAliveNodeMessage()
	var messageList []*mq.Message
	for _, msg := range msgList {
		stringMessage := msg.(string)
		mqttMessage := new(mq.Message)
		json.Unmarshal([]byte(stringMessage), mqttMessage)
		messageList = append(messageList, mqttMessage)
	}
	c.JSON(http.StatusOK, msgList)
}
