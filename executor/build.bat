@echo off
set EXECUTABLE_NAME=executor
REM 设置项目目录和可执行文件名
set PROJECT_DIR=C:\Users\yfy2001\yfy\Learn\Projects\dominant\executor
go env -w GOPROXY=https://goproxy.cn
REM 切换到项目目录
cd %PROJECT_DIR%
REM 编译 Windows 版本
echo Compiling for Windows...
set GOOS=windows
set GOARCH=amd64
echo %EXECUTABLE_NAME%.go
go build -o %EXECUTABLE_NAME%.exe %EXECUTABLE_NAME%.go
REM 编译 macOS 版本
echo Compiling for macOS...
set GOOS=darwin
set GOARCH=amd64
go build -o %EXECUTABLE_NAME%-mac %EXECUTABLE_NAME%_darwin.go
REM 编译 Linux 版本
echo Compiling for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o %EXECUTABLE_NAME%-linux %EXECUTABLE_NAME%_linux.go
echo Compilation finished.
pause