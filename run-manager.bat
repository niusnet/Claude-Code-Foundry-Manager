@echo off
:: Claude Code - Azure Foundry Manager (Python)
:: Ejecuta el gestor de configuración con privilegios de administrador

setlocal

:: Obtener la ruta del script
set "SCRIPT_DIR=%~dp0"
set "PY_SCRIPT=%SCRIPT_DIR%claude_foundry_manager.py"

:: Verificar que el script existe
if not exist "%PY_SCRIPT%" (
    echo.
    echo ERROR: No se encontro claude_foundry_manager.py
    echo Asegurate de que este archivo .bat esta en la misma carpeta que el script .py
    echo.
    pause
    exit /b 1
)

:: Verificar si Python esta instalado
python --version >nul 2>&1
if %errorlevel% neq 0 (
    echo.
    echo ERROR: Python no esta instalado o no esta en el PATH
    echo.
    echo Por favor instala Python desde https://www.python.org/downloads/
    echo Asegurate de marcar "Add Python to PATH" durante la instalacion
    echo.
    pause
    exit /b 1
)

:: Verificar si ya se esta ejecutando como administrador
net session >nul 2>&1
if %errorlevel% == 0 (
    :: Ya es administrador, ejecutar el script directamente
    echo Ejecutando Claude Code Azure Foundry Manager...
    echo.
    python "%PY_SCRIPT%"
) else (
    :: No es administrador, solicitar elevacion
    echo Solicitando privilegios de administrador...
    echo.
    powershell -Command "Start-Process python -ArgumentList '\"%PY_SCRIPT%\"' -Verb RunAs"
)

exit /b 0
