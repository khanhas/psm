Tab completion is currently supported in three shells: PowerShell, bash, and zsh.  
These scripts assume that `psm` is on your path. You can verify that you have psm set up correctly by running:
```bash
> psm -c "-"
```

It should list out all psm default options
```bash
> psm -c "-"
--init
--complete
--help
--list-script
--set-path
```

### PowerShell

To enable tab completion in PowerShell, edit your PowerShell profile:
```ps
notepad $PROFILE
```

Add the contents of `register-completion.ps1` to this file and save.

### bash

To enable tab completion in bash, edit your .bashrc file to add the contents of `register-completion.bash`.

### zsh

To enable tab completion in zsh, edit your .zshrc file to add the contents of `register-completions.zsh`.
