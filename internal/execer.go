package internal

import (
	"os/exec"
)


func execute(alias *Alias, cmd *exec.Cmd) (error) {

	cmd.Args = append(cmd.Args, alias.scriptPath)

	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil;
}
