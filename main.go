package main

import (
	"bufio"
	"bytes"
	"dominant/mq/message"
	"dominant/server"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var client *http.Client

var DestinationAddr string

const BaseUrl = server.BaseUrl

var exitChan chan bool

func init() {
	client = &http.Client{}
	exitChan = make(chan bool)
}

func main() {
	go GetCommand()
	select {
	case <-exitChan:
		return
	}
}

func GetCommand() {
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	for {
		//从终端读取用户命令
		fmt.Print("Remote@", DestinationAddr, ">")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		line = strings.Trim(line, " \r\n")
		switch line {
		case "exit":
			//如果用户输入的是 exit就退出
			log.Println("客户端退出")
			exitChan <- true
		case "lse":
			clientList, _ := getClientList()
			fmt.Println(clientList)
			continue
		case "getm":
			msg := getMessage()
			fmt.Println(msg)
			continue
		case "":
			continue
		}
		clientList, _ := getClientList()
		msg := message.NewMessage("true", clientList, line)
		res := newMessage(msg)
		fmt.Println(res)
	}
}

func getClientList() ([]string, error) {
	url := fmt.Sprintf("%s/getClientList", BaseUrl)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}
	body, err := io.ReadAll(resp.Body)
	msg := new(message.Message)
	msg.MessageJsonUnMarshal(body)
	clientList := anyToStringSlice(msg.Body)
	return clientList, nil
}

func newMessage(msg *message.Message) string {
	url := fmt.Sprintf("%s/newMessage", BaseUrl)
	bytesMessage, err := msg.MessageJsonMarshal()
	if err != nil {
		fmt.Println("json序列化失败:", err.Error())
		return ""
	}
	reader := bytes.NewReader(bytesMessage)
	req, _ := http.NewRequest("POST", url, reader)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return "发送请求失败!"
	}
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func getMessage() string {
	url := fmt.Sprintf("%s/getMessage", BaseUrl)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return "发送请求失败!"
	}
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func anyToStringSlice(a any) []string {
	var strSlice []string
	if slice, ok := a.([]interface{}); ok {
		// 创建一个string切片
		// 将[]interface{}中的每个元素转换为string
		strSlice = make([]string, len(slice))
		for i, v := range slice {
			if str, ok := v.(string); ok {
				strSlice[i] = str
			} else {
				fmt.Println("元素不是string类型")
				return nil
			}
		}
	}
	return strSlice
}
