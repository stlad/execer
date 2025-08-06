package repository

import (
	"errors"
	"execer/internal/core"
	"log"
	"os"
	"path/filepath"
)

var ErrorAliasNotFound = errors.New("псевдоним не найден")

type AliasRepository interface {
	FindAlias(name string) (*core.Alias, error)
}

func GetAliasRepository() AliasRepository {
	return &LocalFileAliasRepository{}
}

type LocalFileAliasRepository struct{}

func (repo LocalFileAliasRepository) FindAlias(name string) (*core.Alias, error) {
	log.Printf("Ищем команду по псеводниму \"%s\"", name)
	var path, _ = os.Executable()
	dir := filepath.Dir(path)

	//TODO Должны быть получение директории в зависимости от конфига
	var lib = filepath.Join(dir, "resources")

	var scriptPath, err = findFilePath(lib, name)

	if err != nil {
		return nil, err
	}

	log.Printf("Команда с псеводнонимом \"%s\"", name)
	return &core.Alias{Name: name, ScriptPath: scriptPath}, nil
}

func findFilePath(lib, name string) (string, error) {

	if _, err := os.Stat(lib); os.IsNotExist(err) {
		log.Printf("Не найден каталог с скриптами. Создаем %s", lib)
		// Создаем директорию с правами 0755 (rwxr-xr-x)
		_ = os.Mkdir(lib, 0755)
	}

	//TODO СЕЙЧАС ЕСТЬ ЗАВИСИМОСТЬ .ps1 для powershell. Необходимо сделать так, чтобы файл находился, не зная его расширения
	fullPath := filepath.Join(lib, name+".ps1")

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", ErrorAliasNotFound
	}

	return fullPath, nil
}
