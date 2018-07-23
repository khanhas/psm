# PowerShell parameter completion shim for psm
Register-ArgumentCompleter -Native -CommandName psm -ScriptBlock {
    param($commandName, $wordToComplete, $cursorPosition)
        psm -c "$wordToComplete" | ForEach-Object {
           [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
        }
}