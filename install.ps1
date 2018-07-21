$installPath = $MyInvocation.MyCommand.Definition
$psmRoot = Split-Path -parent $installPath
$psmBin = Join-Path $psmRoot "bin"

if (IsWindows) {
    If (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
        $arguments = "& '" + $installPath + "'"
        # Execute this script again but elevated
        Start-Process powershell -Verb runAs -ArgumentList $arguments
    } else {
        [Microsoft.Win32.Registry]::SetValue(
            "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment\",
            "PSMROOT",
            $psmBin,
            [Microsoft.Win32.RegistryValueKind]::ExpandString
        )
    }
}