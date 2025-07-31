package main

import (
	"execer/internal/command"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ожидается подкоманда")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCmd := flag.NewFlagSet("run", flag.ExitOnError)
		name := runCmd.String("a", "", "имя псевдонима скрипта")
		runCmd.Parse(os.Args[2:])

		command.CommandRun(*name)
	default:
		fmt.Printf("Неизвестная команда: %s\n", os.Args[1])
		os.Exit(1)
	}
}
