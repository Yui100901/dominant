package main

import (
	"dominant/mq/message"
	"dominant/mq/node"
	"dominant/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

var b *node.Broker

func init() {
	b = node.NewBroker()
}

func main() {
	r := server.NewServer()
	r.POST("/newMessage", newMessage)
	r.GET("/getMessage", getMessage)
	r.GET("/getClientList", getClientList)
	r.POST("/register", register)
	err := r.Run(fmt.Sprintf(":%s", server.Port))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func newMessage(c *gin.Context) {
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

func getMessage(c *gin.Context) {
	ip := c.ClientIP()
	id := c.Query("id")
	fmt.Println(ip)
	msg := b.GetMessage(id)
	b.Register(id, ip)
	c.JSON(http.StatusOK, msg)
}

func getClientList(c *gin.Context) {
	aliveList := b.ListNodes()
	c.JSON(http.StatusOK, aliveList)
}

func register(c *gin.Context) {
	ip := c.ClientIP()
	id := c.Query("id")
	b.Register(id, ip)
	b.KeepAlive(id)
	//if _, ok := cluster.Load(id); ok {
	//	c.JSON(http.StatusOK, gin.H{"message": "Keep Alive Success!"})
	//} else {
	//	cluster.Store(id, node.NewNode(id))
	//	c.JSON(http.StatusOK, gin.H{"message": "Create Connection Success!"})
	//}
}
