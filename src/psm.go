package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
	err = json.Unmarshal(fileContent, &s)

	if err != nil {
		panic(err)
	}

	m := s.(map[string]interface{})

	var mapped = make(map[string]string)
	for k, v := range m {
		switch v.(type) {
		case string:
			mapped[k] = v.(string)
		}
	}

	return mapped
}

func readYAML(filePath string) map[string]string {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	var s interface{}
	err = yaml.Unmarshal(fileContent, &s)

	if err != nil {
		log.Fatal(err)
	}

	m := s.(map[interface{}]interface{})

	var mapped = make(map[string]string)
	for k, v := range m {
		switch k.(type) {
		case string:
			switch k.(type) {
			case string:
				fmt.Println(k, v)
				mapped[k.(string)] = v.(string)
			}
		}
	}

	return mapped
}

// ScanFolderRecursively return array of file path that has `.ps1`
// extension
func ScanFolderRecursively(folderPath string) []string {
	dir, err := ioutil.ReadDir(folderPath)

	var collection = make([]string, 0)

	if err != nil {
		return collection
	}

	for _, v := range dir {
		currentPath := filepath.Join(folderPath, v.Name())
		if v.IsDir() {
			collection = append(collection, ScanFolderRecursively(currentPath)...)
		} else if strings.ToLower(filepath.Ext(v.Name())) == ".ps1" {
			currentPath = "." + string(os.PathSeparator) + currentPath
			collection = append(collection, currentPath)
		}
	}

	return collection
}

// FileListToMappedObject convert array of files path to an object
// that has keys are file name and values are file path.
// Duplicated file name is postfixed its index
func FileListToMappedObject(fileList []string) map[string]string {
	results := make(map[string]string)
	duplicatedName := make(map[string]int)
	for _, v := range fileList {
		// Remove extention in file name
		baseName := strings.Replace(filepath.Base(v), filepath.Ext(v), "", 1)
		if results[baseName] != "" {
			duplicatedName[baseName]++
			baseName = fmt.Sprintf("%s%d", baseName, duplicatedName[baseName])
		}
		results[baseName] = v
	}
	return results
}

// InitJSON scans current folder recursively, maps all `ps1`
// files with their name and path. Then marshal them to json
// format and write to `psm.json` file
func InitJSON() {
	fmt.Println("Auto-generate psm.json")
	_, err := os.Stat("psm.json")
	if err == nil {
		if !ReadAnswer("psm.json is found in current folder. Overwrite? [y/N]: ", false) {
			return
		}
	}
	list := ScanFolderRecursively(".")
	mapped := FileListToMappedObject(list)
	marshal, err := json.MarshalIndent(mapped, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("psm.json", marshal, 0644)
	fmt.Println("psm.json is created successfully!")
}

// InitYAML scans current folder recursively, maps all `ps1`
// files with their name and path. Then marshal them to yaml
// format and write to `psm.yaml` file
func InitYAML() {
	fmt.Println("Auto-generate psm.yaml")
	_, err := os.Stat("psm.yaml")
	if err == nil {
		if !ReadAnswer("psm.yaml is found in current folder. Overwrite? [y/N]: ", false) {
			return
		}
	}
	list := ScanFolderRecursively(".")
	mapped := FileListToMappedObject(list)
	marshal, err := yaml.Marshal(mapped)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("psm.yaml", marshal, 0644)
	fmt.Println("psm.yaml is created successfully!")
}

// GetPSMPath returns `psm` root folder
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

// GetPowershellPathStorageFilePath returns `powershell-path` file path.
func GetPowershellPathStorageFilePath() string {
	return GetPSMPath() + "/powershell-path"
}

// SetPowershellPath store custom powershell executable path in
// `powershell-path` file and returns that path.
func SetPowershellPath(exePath string) string {
	if len(exePath) == 0 {
		exePath = "powershell"
	}
	ioutil.WriteFile(GetPowershellPathStorageFilePath(), []byte(exePath), 0644)
	return exePath
}

// GetPowershellPath returns stored powershell host path or name
// in `powershell-path` file.
// If no path is set or powershell-path file is not found, returns `powershell`
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

// ReadAnswer prints out a yes/no form with string from `info`
// and returns boolean value based on user input (`y` or `n`) or
// return `defaultAnswer` if input is omitted
// If input is neither of them, print form again.
func ReadAnswer(info string, defaultAnswer bool) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(info)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(strings.ToLower(text), "\r\n", "", 1)
	if text == "" {
		return defaultAnswer
	} else if text == "y" {
		return true
	} else if text == "n" {
		return false
	}
	return ReadAnswer(info, defaultAnswer)
}
