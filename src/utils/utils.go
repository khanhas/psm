package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../config"
)

// ReadAnswer prints out a yes/no form with string from `info`
// and returns boolean value based on user input (y/Y or n/N) or
// return `defaultAnswer` if input is omitted.
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

// SetPowershellPath store custom powershell executable path in
// `powershell-path` file and returns that path.
func SetPowershellPath(exePath string) string {
	if len(exePath) == 0 {
		exePath = "powershell"
	}

	configContent := config.ParseConfig()
	configContent.PowerShellPath = exePath
	config.WriteConfig(configContent)
	return exePath
}
