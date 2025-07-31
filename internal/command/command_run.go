package command

import (
	"execer/internal/core"
	"log"
)

func CommandRun(aliasName string) {
	log.Printf("Выполнение команды RUN для алиаса %s", aliasName)

	//TODO тут получение объекта alias из репозитория
	var alias = core.Alias{Name: aliasName, ScriptPath: "C:\\Projects\\execer\\foo.ps1"}

	var executor, _ = core.GetTerminal()
	executor.Exec(alias.ScriptPath)

}
