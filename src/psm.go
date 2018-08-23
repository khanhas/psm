package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"

	"./config"
	"./inits"
	"./parser"
	"./utils"
)

const (
	version = "1.1.0"
)

var supportedFile = []string{
	"psm.yaml",
	"psm.json",
}

var psmConfig = config.ParseConfig()

func main() {
	args := os.Args

	if len(args) < 2 {
		help()
		return
	}

	var detail string
	if len(args) > 2 {
		detail = args[2]
	}

	switch args[1] {
	case "-i", "--init":
		if len(detail) > 0 {
			switch detail {
			case "yaml":
				inits.YAML()
			case "json":
				inits.JSON()
			}
		} else {
			inits.YAML()
		}
	case "-s", "--set-path":
		utils.SetPowershellPath(detail)
	case "-l", "--list-script":
		names := getScriptNames()
		list(names)
	case "-c", "--complete":
		names := getScriptNames()
		names = append(names,
			"--init",
			"--complete",
			"--help",
			"--list-script",
			"--set-path",
		)

		if len(detail) > 0 {
			var position = strings.LastIndex(detail, " ")
			if position > 0 {
				position++
			} else {
				position = 0
			}
			wordToComplete := detail[position:len(detail)]
			for _, v := range names {
				if strings.Index(v, wordToComplete) == 0 {
					fmt.Println(v)
				}
			}
		} else {
			list(names)
		}
	case "-h", "--help":
		help()
	case "-v", "--version":
		fmt.Println(version)
	default: // Execute script
		scripts := gatherCommands()
		command := scripts[args[1]]
		if len(command) != 0 {
			args[0] = "-Command"
			args[1] = command
			cmd := exec.Command(psmConfig.PowerShellPath, args...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Println("Command '" + args[1] + "' not found.")
			matches := fuzzy.RankFind(args[1], getScriptNames())
			if matches.Len() > 0 {
				sort.Sort(matches)
				fmt.Println("Did you mean '" + matches[0].Target + "'?")
			}
		}
	}
	return
}

func gatherCommands() parser.Commands {
	commands := make(parser.Commands, 0)

	for cmd, script := range psmConfig.GlobalCommands {
		commands[cmd] = script
	}

	for _, v := range supportedFile {
		_, err := os.Stat(v)
		if err == nil {
			localCommands := parser.Parse(v)
			for cmd, script := range localCommands {
				commands[cmd] = script
			}
			break
		}
	}

	if len(commands) == 0 {
		fmt.Println("Could not find any commands.")
		os.Exit(1)
	}

	return commands
}

func getScriptNames() []string {
	scripts := gatherCommands()
	results := make([]string, 0)
	for k := range scripts {
		results = append(results, k)
	}
	return results
}

func list(scripts []string) {
	for _, v := range scripts {
		fmt.Println(v)
	}
}

func help() {
	fmt.Print(`
SYNOPSIS
psm [-i [ext]] [-s [path]] [-c [word]] [-l] [-h] script_alias

DESCRIPTION
Run powershell script with shorthand alias.

OPTIONS
-i <ext>, --init <ext>                Generate config in current directory (yaml, json) (default = yaml)
-s <path>, --set-path <path>          Set powershell path/command (default = powershell)
-c <keyword>, --complete <keyword>    Print possible script aliases that match with keyword
-l, --list                            List all available script aliases
-h, --help                            Print this help and exit
-v, --version                         Print version number and exit

EXAMPLES
# Generate a psm.json
psm -i json

# Set pwsh as default shell
psm -s pwsh

# Set powershell.exe in D:\my-powershell-fork as default shell
psm --set-path "D:\my-powershell-fork\powershell.exe"

# Run "fresh" script
psm fresh

`)
}
