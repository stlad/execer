package main

import (
	"execer/internal/command"
	"flag"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Ожидается подкоманда")
	}

	switch os.Args[1] {
	case "run":
		runCmd := flag.NewFlagSet("run", flag.ExitOnError)
		name := runCmd.String("a", "", "имя псевдонима скрипта")
		runCmd.Parse(os.Args[2:])

		command.CommandRun(*name)
	default:
		log.Fatalf("Неизвестная команда: %s\n", os.Args[1])
	}
}
