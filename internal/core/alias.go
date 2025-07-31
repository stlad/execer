package core



type Alias struct{
	Name string;
	ScriptPath string;
}


func newAlias(newName string, path string) *Alias{
	return &Alias{Name: newName, ScriptPath: path}
}

