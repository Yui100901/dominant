package main

import (
	"dominant/mq/api"
	"dominant/server"
	"fmt"
	"log"
)

func main() {
	r := server.NewServer()
	r.POST("/newMessage", api.NewMessage)
	r.GET("/getClientList", api.GetClientList)
	//执行器相关接口
	r.GET("/getMessage", api.GetMessage)
	r.POST("/register", api.Register)
	err := r.Run(fmt.Sprintf(":%s", server.Port))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
