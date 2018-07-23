package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func Parse(file string) map[string]string {
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

func parseJSON(content []byte) map[string]string {
	var s interface{}
	err := json.Unmarshal(content, &s)

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

func parseYAML(content []byte) map[string]string {
	var s interface{}
	err := yaml.Unmarshal(content, &s)

	if err != nil {
		log.Fatal(err)
	}

	m := s.(map[interface{}]interface{})

	var mapped = make(map[string]string)
	for k, v := range m {
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
