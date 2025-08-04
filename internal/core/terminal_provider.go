package core

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

func (executor PowershellExecutor) Exec(script string) {
	var tempDir, _ = os.Getwd()
	tmpfile, err := os.CreateTemp(tempDir, "script-*.ps1")
		if err != nil {
		log.Fatal(err)
	}
	//TODO так как cmd.Start() запускается в отдельном потоке (видимо) - defer os.Remove() удаляет временный файл до торго как ps успевает запустить скрипт
	defer os.Remove(tmpfile.Name())
	// Записываем скрипт в файл
	tmpfile.Write([]byte(script))
	tmpfile.Close()
	var tmpPath, _ = filepath.Abs(tmpfile.Name())
	log.Println("path:",tmpPath)
	executor.cmd.Args = append(executor.cmd.Args, tmpPath)
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
