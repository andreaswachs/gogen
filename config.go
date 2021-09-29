package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Verbose             bool
	RemoveExistingFiles bool
}

type Generators struct {
	Name      string
	Filename  string
	ParseArgs []string
}

type YamlStructure struct {
	Settings   Settings     `yaml: "settings"`
	Generators []Generators `yaml: "generators"`
}

func GetConfig() YamlStructure {
	settings := YamlStructure{}

	config := readConfig()

	err := yaml.Unmarshal([]byte(config), &settings)
	if err != nil {
		fmt.Println("Something went wrong while trying to read the config file")
		fmt.Println("The gogen program could not understand it correctly.")
		os.Exit(5)
	}

	return settings
}

func readConfig() []byte {
	syscall.Umask(0)

	// read the config and return it
	// if not present, as in first run ever, download it and make the folder paths

	userConfigDir, err := os.UserConfigDir()
	appConfigDir := userConfigDir + "/gogen"
	configFileComplete := appConfigDir + "/config.yaml"

	if err != nil {
		fmt.Println("An error occured while attempting to determine the users application config directory.")
		fmt.Println(err)
		os.Exit(1)
	}

	// ensure gogen app config folder exists
	if _, err := os.Stat(appConfigDir); os.IsNotExist(err) {
		err = os.MkdirAll(appConfigDir, os.ModePerm)

		if err != nil {
			fmt.Println("An error occurred while attempting to create application config folder")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// ensure that the config file exists within the app config folder for gogen
	if _, err := os.Stat(configFileComplete); os.IsNotExist(err) {
		// download the default configuration
		fmt.Println("No config file was found locally, will download default config file.")
		resp, err := http.Get("https://raw.githubusercontent.com/andreaswachs/gogen/main/config.yaml")

		if err != nil {
			fmt.Println("Failed to download default config file.")
			os.Exit(2)
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("Failed to read the downloaded default config file.")
			os.Exit(3)
		}

		os.WriteFile(configFileComplete, body, os.ModePerm)
		return body
	}

	config, err := os.ReadFile(configFileComplete)

	if err != nil {
		fmt.Println("Could not read config file from user config directory.")
		fmt.Println("Please see if you have permission to read this file")
		os.Exit(4)
	}

	return config
}
