# Guía de Instalación y Compilación

Este documento te guía para compilar y usar el Claude Foundry Manager.

## 🎯 Opciones de Instalación

### Opción A: Descargar Binario Pre-compilado (Más Fácil)

Una vez que subas el proyecto a GitHub y crees un release:

1. Ve a: `https://github.com/TU_USUARIO/claude-foundry-manager/releases`
2. Descarga el binario para Windows:
   - `claude-foundry-manager-windows-amd64.exe` (64-bit)
3. Guarda el archivo y ejecútalo directamente
4. ¡Listo! No necesitas instalar nada más

---

### Opción B: Compilar Localmente

#### Paso 1: Instalar Go

**Windows:**

1. Descarga Go desde: https://go.dev/dl/
   - Busca: `go1.21.x.windows-amd64.msi` (última versión)
   - Tamaño: ~130 MB

2. Ejecuta el instalador `.msi`
   - Sigue el asistente (instalación estándar)
   - Ubicación por defecto: `C:\Program Files\Go`

3. Verifica la instalación:
   ```powershell
   # Abre una NUEVA terminal PowerShell
   go version
   ```
   Deberías ver: `go version go1.21.x windows/amd64`

#### Paso 2: Compilar el Proyecto

```powershell
# Navega al directorio del proyecto
cd "C:\Users\gilbe\Desktop\Claude Code Foundry Manager"

# Descarga las dependencias
go mod download

# Compila el proyecto
go build -ldflags="-s -w" -o claude-foundry-manager.exe .
```

**Explicación de flags:**
- `-ldflags="-s -w"`: Reduce el tamaño del binario (elimina símbolos de debug)
- `-o claude-foundry-manager.exe`: Nombre del archivo de salida

#### Paso 3: Ejecutar

```powershell
# Modo interactivo
.\claude-foundry-manager.exe

# O con comandos CLI
.\claude-foundry-manager.exe configure --resource=my-foundry --api-key=sk-xxx
```

---

### Opción C: Usar GitHub Actions (Sin instalar Go)

Esta opción compila automáticamente binarios para 6 plataformas.

#### Paso 1: Subir a GitHub

```powershell
# Si aún no tienes un repositorio remoto en GitHub:
# 1. Ve a https://github.com/new
# 2. Crea un repositorio llamado "claude-foundry-manager"
# 3. No inicialices con README (ya lo tienes)

# Conecta tu repositorio local con GitHub
git remote add origin https://github.com/TU_USUARIO/claude-foundry-manager.git

# Sube el código
git push -u origin main
```

#### Paso 2: Crear un Release

```powershell
# Crea y sube un tag de versión
git tag v1.0.0
git push origin v1.0.0
```

#### Paso 3: Esperar la Compilación

1. Ve a tu repositorio en GitHub
2. Click en "Actions" → Verás el workflow "Build and Release" ejecutándose
3. Espera ~5-10 minutos a que termine
4. Ve a "Releases" → Verás `v1.0.0` con 6 binarios adjuntos

#### Paso 4: Descargar y Usar

Descarga `claude-foundry-manager-windows-amd64.exe` y ejecútalo.

---

## 🛠️ Compilación Multi-Plataforma (Avanzado)

Si quieres compilar para otras plataformas desde Windows:

```powershell
# Linux (64-bit)
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o claude-foundry-manager-linux-amd64 .

# macOS Intel (64-bit)
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o claude-foundry-manager-darwin-amd64 .

# macOS Apple Silicon (ARM64)
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o claude-foundry-manager-darwin-arm64 .

# Windows ARM64
$env:GOOS="windows"; $env:GOARCH="arm64"; go build -o claude-foundry-manager-windows-arm64.exe .

# Resetear variables de entorno
$env:GOOS=""; $env:GOARCH=""
```

---

## ❓ Problemas Comunes

### "go: command not found"

**Solución:** Reinicia tu terminal después de instalar Go. Las variables de entorno se actualizan en nuevas sesiones.

### "cannot find package github.com/spf13/cobra"

**Solución:** Ejecuta `go mod download` primero para descargar dependencias.

### "Access is denied" al ejecutar el .exe

**Solución:** En Windows, ejecuta el .exe como Administrador (clic derecho → Ejecutar como administrador).

### El binario es muy grande (>50 MB)

**Solución:** Usa flags de compilación optimizados:
```powershell
go build -ldflags="-s -w" -o claude-foundry-manager.exe .
```

---

## 📋 Checklist de Instalación

- [ ] Go instalado (verifica con `go version`)
- [ ] Dependencias descargadas (`go mod download`)
- [ ] Proyecto compilado (archivo `.exe` creado)
- [ ] Binario ejecutable (prueba con `.\claude-foundry-manager.exe`)
- [ ] Privilegios de administrador (necesario para modificar variables del sistema)

---

## 🚀 Siguiente Paso

Una vez que tengas el binario compilado, lee el [README.md](README.md) para aprender a usar la herramienta.

**Uso rápido:**
```powershell
# Modo interactivo (recomendado para primera vez)
.\claude-foundry-manager.exe

# Ayuda de comandos
.\claude-foundry-manager.exe --help
```
