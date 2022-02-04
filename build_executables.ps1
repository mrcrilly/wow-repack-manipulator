cd .\cmd\v1
$env:GOARCH = "amd64"
$env:GOOS = "windows"
go build -o builds/wow-repack-manipulator.exe

$env:GOOS = "linux"
go build -o builds/wow-repack-manipulator
cd ..\..