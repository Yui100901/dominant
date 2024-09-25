package main

import (
	"dominant/infrastructure/utils/log_utils"
	"fmt"
	"os/exec"
	"syscall"
)

type ExecCommand interface {
	Exec()
}

type Command struct {
	cmdLine string
	result  string
}

func NewCommand(cmdLine string) *Command {
	return &Command{
		cmdLine: cmdLine,
		result:  "",
	}
}

// Exec 运行命令将结果保存到result中
func (c *Command) Exec() {
	cmd := exec.Command("cmd.exe")
	switch c.cmdLine {
	case "download":
	default:
		//系统调用执行命令，隐藏执行终端窗口
		cmd.SysProcAttr = &syscall.SysProcAttr{
			CmdLine:    fmt.Sprintf(`/c %s`, c.cmdLine),
			HideWindow: true,
		}
		res, err := cmd.Output()
		if err != nil {
			c.result = err.Error()
		}
		c.result = string(res)
		log_utils.Info.Printf(`
Execute Command:
%s
Command Result:
%s`, c.cmdLine, c.result)
	}
}
