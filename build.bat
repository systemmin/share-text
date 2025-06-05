@REM 2025年 author：Mr.Fang
@echo off
chcp 65001

setlocal enabledelayedexpansion

set GOOS=windows
set GOARCH=amd64

go build -o share-text.exe main.go

@REM set GOOS=linux
@REM set GOARCH=amd64
@REM
@REM go build -o share-text main.go

echo Build completed!
endlocal
exit /b 0
