package gogenlib

import (
	"fmt"
	"os"
)

const (
	gogenConfigFolderName  = "/gogen"
	configFileName         = "/config.yaml"
	configFileCompletePath = gogenConfigFolderName + configFileName
	templatesFolderName    = "/templates"
)

func GetGogenBasePath() string {
	dir, err := os.UserConfigDir()

	exitOnError(err, "Could not determine the user config directory.")
	return dir + gogenConfigFolderName
}

func GetGogenConfigFolderPath() string {
	dir := GetGogenBasePath()
	return dir + gogenConfigFolderName
}

func GetGogenConfigFilePath() string {
	dir := GetGogenBasePath()
	return dir + gogenConfigFolderName + configFileName
}

func GetGogenTemplatesFolderPath() string {
	dir := GetGogenBasePath()
	return dir + gogenConfigFolderName + templatesFolderName
}

func EnsureConfigFoldersExists() {
	gogenDir := GetGogenConfigFolderPath()

	if !ensureNamedFolderExists("gogen", gogenDir) {
		downloadTemplateConfigFile()
	}

	templatesDir := GetGogenTemplatesFolderPath()

	if !ensureNamedFolderExists("templates", templatesDir) {
		fmt.Println("The templates folder was just created. The folder is empty. You should put some templates in there!")
	}
}

func ensureNamedFolderExists(name string, path string) bool {
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, os.ModePerm)
		exitOnError(err, fmt.Sprintf("Could not create %s folder. This might be an permissions issue.\n", name))

		return false
	}

	return true
}
