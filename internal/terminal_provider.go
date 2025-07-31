package internal

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

var ErrorCMDNotFound = errors.New("терминал не найден для текущей системы")
var ErrorOSNotHandle = errors.New("операционная система не поддерживается")

type Executor interface {
	Exec(string)
}

type PowershellExecutor struct {
	cmd *exec.Cmd
}

func (executor PowershellExecutor) Exec(scriptPath string) {
	executor.cmd.Args = append(executor.cmd.Args, scriptPath)
	executor.cmd.Start()
}

func GetTerminal() (Executor, error) {
	var executor Executor

	switch runtime.GOOS {
	case "windows":
		var cmd = exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		executor = PowershellExecutor{cmd: cmd}
	case "linux", "darwin":
		// cmd = exec.Command("/bin/sh", "-c")
		return nil, ErrorOSNotHandle
	default:
		return nil, ErrorCMDNotFound
	}

	return executor, nil
}
