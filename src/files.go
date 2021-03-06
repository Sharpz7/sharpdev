package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Creates default sharpdev.yml file
func genFile() {
	scriptsEx := make(map[string]string)
	scriptsEx["example"] = "echo 'Hello World!'"

	testfile := config{
		Version: 0.1,
		Scripts: scriptsEx,
		EnvFile: ".env"}

	err := saveFile(testfile)
	check(err, "Failed to generate sharpdev.yml")

	fmt.Println("Created sharpdev.yml")
}

// Loads a sharpdev file
func loadFile() (config, error) {
	f, readErr := ioutil.ReadFile("./sharpdev.yml")
	var devFile config
	marshErr := yaml.Unmarshal(f, &devFile)
	if marshErr != nil || readErr != nil {
		return config{}, errors.New("failed to load file")
	}

	return devFile, nil
}

// Saves a sharpdev file
func saveFile(devFile config) error {
	yamlData, marshErr := yaml.Marshal(devFile)
	check(marshErr, "Failed to Convert to Yaml")
	writeErr := ioutil.WriteFile("./sharpdev.yml", yamlData, 0644)

	if marshErr != nil || writeErr != nil {
		return errors.New("failed to save file")
	}

	return nil
}
