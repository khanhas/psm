package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Commands map[string]string

// Parse detects supported file extension and use
// corresponding function to parse file content.
func Parse(file string) Commands {
	content := readContent(file)
	extension := filepath.Ext(file)
	extension = strings.ToLower(extension)
	switch extension {
	case ".yaml":
		return parseYAML(content)
	case ".json":
		return parseJSON(content)
	default:
		log.Fatal("Unsupported file!\nExit.")
	}

	return nil
}

func readContent(file string) []byte {
	fileContent, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return fileContent
}

func parseJSON(content []byte) Commands {
	var raw interface{}
	err := json.Unmarshal(content, &raw)

	if err != nil {
		panic(err)
	}

	return MapCommands(raw)
}

func parseYAML(content []byte) Commands {
	var raw interface{}
	err := yaml.Unmarshal(content, &raw)

	if err != nil {
		log.Fatal(err)
	}

	return MapCommands(raw)
}

func MapCommands(raw interface{}) Commands {
	mappedRaw := raw.(map[interface{}]interface{})
	mapped := make(Commands)
	for k, v := range mappedRaw {
		if isString(k) && isString(v) {
			mapped[k.(string)] = v.(string)
		}
	}
	return mapped
}

func isString(unk interface{}) bool {
	switch unk.(type) {
	case string:
		return true
	}
	return false
}
