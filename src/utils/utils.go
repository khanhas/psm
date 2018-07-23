package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ReadAnswer prints out a yes/no form with string from `info`
// and returns boolean value based on user input (y/Y or n/N) or
// return `defaultAnswer` if input is omitted
// If input is neither of them, print form again.
func ReadAnswer(info string, defaultAnswer bool) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(info)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r", "", 1)
	text = strings.Replace(text, "\n", "", 1)
	if len(text) == 0 {
		return defaultAnswer
	} else if text == "y" || text == "Y" {
		return true
	} else if text == "n" || text == "N" {
		return false
	}
	return ReadAnswer(info, defaultAnswer)
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
	return filepath.Join(GetPSMPath(), "powershell-path")
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
