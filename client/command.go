package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"syscall"
)

var (
	runtimeEnvironment string
)

func init() {
	//获取当前运行的系统环境
	runtimeEnvironment = runtime.GOOS
}

type Command struct {
	cmdLine string
	env     string
	result  string
}

func NewCommand(cmdLine string) *Command {
	return &Command{
		cmdLine: cmdLine,
		env:     runtimeEnvironment,
		result:  "",
	}
}

// Exec 运行命令将结果保存到result中
func (c *Command) Exec() {
	switch c.env {
	case "windows":
		cmd := exec.Command("cmd.exe")
		switch c.cmdLine {
		case "download":
		default:
			log.Printf("Execute Command:%s", c.cmdLine)
			//系统调用执行命令
			//隐藏执行终端窗口
			cmd.SysProcAttr = &syscall.SysProcAttr{
				CmdLine:    fmt.Sprintf(`/c %s`, c.cmdLine),
				HideWindow: true,
			}
			res, err := cmd.Output()
			if err != nil {
				c.result = err.Error()
			}
			log.Printf("Result:%s", res)
			c.result = string(res)
		}
	case "linux":

	case "darwin":

	default:

	}
}
