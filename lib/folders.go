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
	ensureNamedFolderExists("gogen", GetGogenConfigFolderPath())
	ensureNamedFolderExists("templates", GetGogenTemplatesFolderPath())
	downloadTemplateConfigFile()
}

func ensureNamedFolderExists(name string, path string) bool {
	if _, err := os.Stat(path); err != nil {
		verbosePrint(fmt.Sprintf("Attempting to ensure folder %s exists with compelte path \n%s", name, path))
		err := os.MkdirAll(path, os.ModePerm)
		exitOnError(err, fmt.Sprintf("Could not create %s folder. This might be an permissions issue.\n", name))
	}

	verbosePrint(fmt.Sprintf("Successfully ensured that the folder with name %s exists. Full path: \n%s", name, path))
	return true
}
