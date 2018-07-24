<div align="center">
  <img src="https://github.com/khanhas/psm/blob/master/asset/icon.svg">
</div>

# psm - Powershell Script Manager 
[![CircleCI](https://circleci.com/gh/khanhas/psm/tree/master.svg?style=svg)](https://circleci.com/gh/khanhas/psm/tree/master)
 
Execute Powershell script/command or chain of scripts/command with pre-defined alias.  
  
psm exists just to save your time!  

No need to import scripts repeatedly every time setting up your working enviroment.   
No need to spread out simple functions to seperated scripts. From now on, just gather them all in one:  
```ps
function runFunction() {
...
}

function rebuildFunction() {
...
}
```

In psm config, simply import that script and call function, like this:
```yaml
run: (. ./my-scripts-collection.ps1);runFunction
rebuild: (. ./my-scripts-collection.ps1);rebuildFunction
```
In terminal, you just need to type and run:
```bash
> psm run

> psm rebuild
```

## Features:
- Execute PS scripts in other kind of shells
- Chains and pipe commands effortlessly
- Scans current working folder recursively to find PS scripts and [auto-generate](https://github.com/khanhas/psm/blob/master/README.md#-i-ext---init-ext) psm config file.
- Execute scripts in different versions of powershell with no fuss. Just use [`--set-path`](https://github.com/khanhas/psm/blob/master/README.md#-s-path---set-path-path).
- Supports autocomplete for: powershell, bash, zsh. Check out [register-completion scripts](https://github.com/khanhas/psm/tree/master/scripts/).
- Need more? Post an issue or make pull request!

## Install
Download distribution corresponding to your system in [release]() page
### Windows
- Via installer: After installing succesfully, run `refreshenv` in terminal at least one then you're good to go.
- Via zip: After unzipping `psm.exe`, appends its folder path to `PATH` enviroment variable and run `refreshenv` in terminal at least one.

### Linux
Unpacks gzip file.
Appends
```
alias psm=/path/to/psm
```
to your `.bashrc` or `.zshrc`.  
Restart your terminal.

## Usage:
### Synopsis
```bash
psm [-i <ext>] [-s <path>] [-c <keyword>] [-l] [-h] [-v] 

psm script_alias
```

### Options
#### `-i <ext>`, `--init <ext>`  
Generate config in current directory  
Supports: `yaml`, `json`  
Default value is `yaml`  

#### `-s <path>`, `--set-path <path>`
Set powershell path/command  
Default value is `powershell`  
**Notes:** In Windows, if you installed Powershell Core >= 6.0 and want to run script with it, you should set to `pwsh` or direct path to `pwsh.exe`

#### `-c <keyword>`, `--complete <keyword>`
Print possible script aliases that match with keyword

#### `-l`, `--list`
List all available script aliases

#### `-h`, `--help`
Print help and exit

#### `-v`, `--version`
Print version number and exit

## Example:
Put a `psm.json` or `psm.yaml` in current working folder:  
Or use [`--init`](https://github.com/khanhas/psm/blob/master/README.md#-i-ext---init-ext) option to auto-generate one:

![image](https://i.imgur.com/dRkgsOe.png)
  
Inside that file, declare an object with key is whatever shorthand alias you want and value is powershell script to be executed.   

#### `psm.json`
```json
{
  "configure": ". ./task.ps1;configure",
  "build": ". ./task.ps1;build",
  "clean": "Remove-Item -Recurse ./build/",
  "cleanAll": "Remove-Item -Recurse ./Publish/; Remove-Item -Recurse ./build/",
  "fresh": "psm cleanAll; psm configure; psm build"
}
```

or

#### `psm.yaml`
```yaml
configure: ". ./task.ps1;configure"
build: ". ./task.ps1;build"
clean: "Remove-Item -Recurse ./build/"
cleanAll: "Remove-Item -Recurse ./Publish/; Remove-Item -Recurse ./build/"
fresh: "psm cleanAll; psm configure; psm build"
```

Then in terminal, you just need to run:

```bash
> psm fresh
```

to wipe off `Publish` and `build` folders, configure project files then build project

## Development
Requires:
- [Golang](https://golang.org/dl/)
- Powershell >= 6.0

1. Clone:
```bash
git clone https://github.com/khanhas/psm.git

cd psm
```

2. Build
```bash
./build.ps1
```

3. Set enviroment variable so you can run `psm` everywhere:
```bash
./install.ps1
```

## Roadmap
- [x] Init command: Gather all scripts file in project folder and automatically generate a `psm.json`/`psm.yaml`
- [x] Make an installer
- [x] Linux support.
