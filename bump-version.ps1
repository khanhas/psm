param (
    [Parameter(Mandatory=$true)][int16]$major,
    [Parameter(Mandatory=$true)][int16]$minor,
    [Parameter(Mandatory=$true)][int16]$patch
)

$ver = "$($major).$($minor).$($patch)"

(Get-Content ".\src\psm.go") -replace "version = `"[\d\.]*`"", "version = `"$($ver)`"" |
    Set-Content ".\src\psm.go"

(Get-Content ".\installer.iss") -replace "#define MyAppVersion `"[\d\.]*`"", "#define MyAppVersion `"$($ver)`"" |
    Set-Content ".\installer.iss"

(Get-Content ".\dist.ps1") -replace "version=`"[\d\.]*`"", "version=`"$($ver)`"" |
    Set-Content ".\dist.ps1"