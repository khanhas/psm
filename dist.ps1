$version="1.1.0"
$name="psm"
$nameVersion="$($name)-$($version)"
$innoPath="C:\Program Files (x86)\Inno Setup 5\ISCC.exe"
Write-Output "Building linux binary:"
Set-Item -path env:GOOS -value linux

if (Test-Path "./bin/linux") {
    Remove-Item -Recurse "./bin/linux"
}

Set-Item -path env:GOARCH -value amd64
go build -o "./bin/linux/x64/psm" "./src/psm.go"
Write-Output "Linux x64: Done!"

Set-Item -path env:GOARCH -value 386
go build -o "./bin/linux/x86/psm" "./src/psm.go"
Write-Output "Linux x86: Done!"

Write-Output "Packing Linux distribute:"
7z a -bb0 ".\bin\linux\lunix.tar" ".\bin\linux\x64\*"
7z a -bb0 -sdel -mx9 ".\bin\$($nameVersion)-linux-x64.tar.gz" ".\bin\linux\lunix.tar"
Write-Output "Linux x64: Done!"

7z a -bb0 ".\bin\linux\lunix.tar" ".\bin\linux\x86\*"
7z a -bb0 -sdel -mx9 ".\bin\$($nameVersion)-linux-x86.tar.gz" ".\bin\linux\lunix.tar"
Write-Output "Linux x86: Done!"

Set-Item -path env:GOOS -value windows
Write-Output "Building Windows binary:"
if (Test-Path "./bin/windows") {
    Remove-Item -Recurse "./bin/windows"
}

Set-Item -path env:GOARCH -value amd64
go build -o "./bin/windows/x64/psm.exe" "./src/psm.go"
Write-Output "Windows x64: Done!"

Set-Item -path env:GOARCH -value 386
go build -o "./bin/windows/x86/psm.exe" "./src/psm.go"
Write-Output "Windows x86: Done!"

Write-Output "Packing Windows distribute:"
7z a -bb0 -mx9 ".\bin\$($nameVersion)-windows-x64.zip" ".\bin\windows\x64\*"
7z a -bb0 -mx9 ".\bin\$($nameVersion)-windows-x86.zip" ".\bin\windows\x86\*"

Start-Process $innoPath -ArgumentList ".\installer.iss"
