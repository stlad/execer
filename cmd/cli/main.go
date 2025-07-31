package main

import (
	"execer/internal"
)

func main(){
	var alias = internal.Alias{Name: "foo", ScriptPath: "C:\\Projects\\execer\\foo.bat"}
	var executor, _ = internal.GetTerminal()
	executor.Exec(alias.ScriptPath)
}