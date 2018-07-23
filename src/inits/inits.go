package inits

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"../utils"
	yaml "gopkg.in/yaml.v2"
)

// scanFolderRecursively return array of file path that has `.ps1`
// extension
func scanFolderRecursively(folderPath string) []string {
	dir, err := ioutil.ReadDir(folderPath)

	var collection = make([]string, 0)

	if err != nil {
		return collection
	}

	for _, v := range dir {
		currentPath := filepath.Join(folderPath, v.Name())
		if v.IsDir() {
			collection = append(collection, scanFolderRecursively(currentPath)...)
		} else if strings.ToLower(filepath.Ext(v.Name())) == ".ps1" {
			currentPath = "." + string(os.PathSeparator) + currentPath
			collection = append(collection, currentPath)
		}
	}

	return collection
}

// fileListToMappedObject convert array of files path to an object
// that has keys are file name and values are file path.
// Duplicated file name is postfixed its index
func fileListToMappedObject(fileList []string) map[string]string {
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

// JSON scans current folder recursively, maps all `ps1`
// files with their name and path. Then marshal them to json
// format and write to `psm.json` file
func JSON() {
	fmt.Println("Auto-generate psm.json")
	_, err := os.Stat("psm.json")
	if err == nil {
		if !utils.ReadAnswer("psm.json is found in current folder. Overwrite? [y/N]: ", false) {
			return
		}
	}
	list := scanFolderRecursively(".")
	mapped := fileListToMappedObject(list)
	marshal, err := json.MarshalIndent(mapped, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("psm.json", marshal, 0644)
	fmt.Println("psm.json is created successfully!")
}

// YAML scans current folder recursively, maps all `ps1`
// files with their name and path. Then marshal them to yaml
// format and write to `psm.yaml` file
func YAML() {
	fmt.Println("Auto-generate psm.yaml")
	_, err := os.Stat("psm.yaml")
	if err == nil {
		if !utils.ReadAnswer("psm.yaml is found in current folder. Overwrite? [y/N]: ", false) {
			return
		}
	}
	list := scanFolderRecursively(".")
	mapped := fileListToMappedObject(list)
	marshal, err := yaml.Marshal(mapped)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("psm.yaml", marshal, 0644)
	fmt.Println("psm.yaml is created successfully!")
}
