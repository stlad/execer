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
	dir, _ := os.Getwd()

	//TODO Должны быть получение директории в зависимости от конфига
	var lib = dir + "/resources"

	var script, err = getScriptText(lib, name)

	if err != nil {
		return nil, err
	}

	log.Printf("Команда с псеводнонимом \"%s\"", name)
	return core.NewAlias(name, script), nil
}

func getScriptText(lib, name string) (string, error) {

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
	var content, _ = os.ReadFile(fullPath)
	log.Println("Текст полученного скрипта: ", string(content))
	return string(content), nil
}
