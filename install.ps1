$installPath = $MyInvocation.MyCommand.Definition
$psmRoot = Split-Path -parent $installPath
$psmBin = Join-Path $psmRoot "bin"

if ($Host.Version.Major -lt 6) {
    Write-Output "Requires Powershell version >= 6."
    Write-Output "Exit."
    return
}

$envKey = "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment\"

if ($IsWindows) {
    If (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
        $arguments = "-File " + $installPath
        $shellPath = (Get-Process -Id $PID).Name
        # Execute this script again but elevated
        Start-Process $shellPath -Verb runAs -ArgumentList $arguments
    } else {
        $path = [Microsoft.Win32.Registry]::GetValue($envKey, "PATH", "")
        if ($path -like "*$($psmBin)*") {
            break
        }
        [Microsoft.Win32.Registry]::SetValue(
            $envKey,
            "PATH",
            "$($path);$($psmBin)",
            [Microsoft.Win32.RegistryValueKind]::ExpandString
        )
    }
}

Write-Output "Succeed!"
