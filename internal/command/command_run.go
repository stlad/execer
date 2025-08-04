package command

import (
	"execer/internal/core"
	"execer/internal/repository"
	"log"
)

func CommandRun(aliasName string) error {
	log.Printf("Выполнение команды RUN для алиаса \"%s\"", aliasName)

	var repository = repository.GetAliasRepository()
	var alias, notFoundError = repository.FindAlias(aliasName)

	if notFoundError != nil {
		log.Fatal(notFoundError)
	}

	var executor, terminalError = core.GetTerminal()
	if terminalError != nil {
		log.Fatal(terminalError)
	}

	log.Printf("Исполнение команды \"%s\"", alias.Name)
	executor.Exec(alias.Script)
	return nil
}
