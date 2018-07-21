# psm - Powershell Script Manager 
 
Execute Powershell script/command or chain of scripts/command with pre-defined alias.  

Simply put a `psm.json` or `psm.yaml` in current working folder:  

![image](https://i.imgur.com/dRkgsOe.png)
  
Inside that file, declare an object with key is whatever shorthand alias you want and value is powershell script to be executed.  

### Example:
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
psm fresh
```

to wipe off `Publish` and `build` folders, configure project files then build project

## Build
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
- [ ] Init command: Gather all scripts file in project folder and automatically generate a `psm.json`/`psm.yaml`
- [ ] Make an installer
- [ ] Linux support
