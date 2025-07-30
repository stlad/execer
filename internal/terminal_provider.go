package internal

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

var ErrorCMDNotFound = errors.New("терминал не найден для текущей системы")

func GetTerminal() (*exec.Cmd, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell.exe", "-File")
	case "linux", "darwin":
		cmd = exec.Command("/bin/sh", "-c")
	default:
		return nil, ErrorCMDNotFound
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd, nil
}
