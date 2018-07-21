package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please specify which script you want to execute!")
		return
	}

	var detail string
	if len(args) > 2 {
		detail = args[2]
	}
	var powershellPath = GetPowershellPath()

	var scripts map[string]string

	_, err := os.Stat("psm.json")
	if err != nil {
		_, err = os.Stat("psm.yaml")
		if err != nil {
			fmt.Println("Cannot find psm.json and psm.yaml!")
			return
		}
		scripts = readYAML("psm.yaml")
	} else {
		scripts = readJSON("psm.json")
	}

	switch args[1] {
	case "-i":
		InitJSON()
	case "-ij":
		InitJSON()
	case "--Init-JSON":
		InitJSON()
	case "-iy":
		InitYAML()
	case "--Init-YAML":
		InitYAML()
	case "-s":
		SetPowershellPath(detail)
	default: // Execute script
		for k, v := range scripts {
			if k == args[1] {
				args[0] = "-Command"
				args[1] = v
				cmd := exec.Command(powershellPath, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
				return
			}
		}
		fmt.Println("Script is not available!")
		return
	}
}

func readJSON(filePath string) map[string]string {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	var s interface{}
	errJSON := json.Unmarshal(fileContent, &s)

	if errJSON != nil {
		panic(errJSON)
	}

	m := s.(map[string]interface{})

	var mapped = make(map[string]string)
	for k, v := range m {
		mapped[k] = v.(string)
	}

	return mapped
}

func readYAML(filePath string) map[string]string {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	var s interface{}
	errJSON := yaml.Unmarshal(fileContent, &s)

	if errJSON != nil {
		panic(errJSON)
	}
	m := s.(map[interface{}]interface{})

	var mapped = make(map[string]string)
	for k, v := range m {
		mapped[k.(string)] = v.(string)
	}

	return mapped
}

func InitJSON() {
	fmt.Println("Creating psm.json")
}

func InitYAML() {
	fmt.Println("Creating psm.yaml")
}

func SetPowershellPath(exePath string) string {
	if len(exePath) == 0 {
		exePath = "powershell"
	}
	ioutil.WriteFile(GetPSMPath()+"/powershell-path", []byte(exePath), 0644)
	return exePath
}

func GetPSMPath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	ex, err = filepath.Abs(filepath.Dir(ex))
	if err != nil {
		log.Fatal(err)
	}
	return ex
}

func GetPowershellPath() string {
	_, err := os.Stat(GetPSMPath() + "/powershell-path")
	if err == nil {
		storedPath, err := ioutil.ReadFile(GetPSMPath() + "/powershell-path")
		if err == nil {
			return string(storedPath)
		}
	}
	return SetPowershellPath("powershell")
}
