package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"../parser"
	yaml "gopkg.in/yaml.v2"
)

// PSMConfig is type of config file content
type PSMConfig struct {
	PowerShellPath string
	GlobalCommands parser.Commands
}

type rawPSMConfig struct {
	powershellpath string
	globalcommands parser.Commands
}

// GetConfigPath returns path of `.psm.yaml`
func GetConfigPath() string {
	configPath := "/"
	if runtime.GOOS == "windows" {
		configPath = os.Getenv("USERPROFILE")
	} else if runtime.GOOS == "linux" {
		configPath = os.Getenv("HOME")
	}

	return filepath.Join(configPath, ".psm.yaml")
}

// ParseConfig read config file content and returns
// a PSMConfig map
func ParseConfig() PSMConfig {
	configPath := GetConfigPath()
	_, err := os.Stat(configPath)

	if err != nil {
		defaultConfig := getDefaultConfig()
		WriteConfig(defaultConfig)
		return defaultConfig
	}

	content, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	var raw interface{}

	err = yaml.Unmarshal(content, &raw)

	if err != nil {
		log.Fatal(err)
	}

	psmConfig := getDefaultConfig()

	mapped := raw.(map[interface{}]interface{})
	for k, v := range mapped {
		if k == "powershellpath" {
			psmConfig.PowerShellPath = v.(string)
		} else if k == "globalcommands" {
			psmConfig.GlobalCommands = parser.MapCommands(v)
		}
	}
	return psmConfig
}

// WriteConfig writes content to config file.
func WriteConfig(content PSMConfig) {
	marshalled, err := yaml.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(GetConfigPath(), marshalled, 0644)
}

func getDefaultConfig() PSMConfig {
	return PSMConfig{
		PowerShellPath: "powershell",
		GlobalCommands: make(parser.Commands)}
}
