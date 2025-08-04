package core

type Alias struct {
	Name   string
	Script string
}

func NewAlias(newName string, script string) *Alias {
	return &Alias{Name: newName, Script: script}
}
