@echo off
REM Build script for Claude Foundry Manager
REM Compiles the Go project into a Windows executable

echo ========================================
echo Claude Foundry Manager - Build Script
echo ========================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed!
    echo.
    echo Please install Go from: https://go.dev/dl/
    echo After installation, restart this terminal and run this script again.
    echo.
    pause
    exit /b 1
)

echo [1/3] Checking Go version...
go version
echo.

echo [2/3] Downloading dependencies...
go mod download
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to download dependencies
    pause
    exit /b 1
)
echo.

echo [3/3] Building executable...
go build -ldflags="-s -w" -o claude-foundry-manager.exe .
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Build failed
    pause
    exit /b 1
)

echo.
echo ========================================
echo BUILD SUCCESSFUL!
echo ========================================
echo.
echo Executable created: claude-foundry-manager.exe
echo Size:
dir claude-foundry-manager.exe | findstr "claude-foundry-manager.exe"
echo.
echo To run the tool:
echo   .\claude-foundry-manager.exe
echo.
echo For interactive mode, just double-click the .exe file
echo (remember to run as Administrator)
echo.
pause
