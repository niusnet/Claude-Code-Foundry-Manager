# 🚀 Cómo Empezar

Este documento te guía paso a paso para usar el Claude Foundry Manager.

---

## ¿Qué necesitas hacer ahora?

### Tienes 3 opciones (de más fácil a más técnica):

---

## 📦 OPCIÓN 1: Usar GitHub Releases (Recomendada - Sin instalar nada)

Esta es la forma más fácil. GitHub compilará todo por ti.

### Pasos:

1. **Sube el código a GitHub:**

   ```powershell
   # Si no tienes un repositorio en GitHub:
   # - Ve a https://github.com/new
   # - Crea un repositorio llamado "claude-foundry-manager"
   # - No inicialices con README

   # Conecta tu repo local:
   git remote add origin https://github.com/TU_USUARIO/claude-foundry-manager.git
   git push -u origin main
   ```

2. **Crea un tag de versión para activar la compilación automática:**

   ```powershell
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **Espera 5-10 minutos** mientras GitHub Actions compila los binarios

4. **Descarga tu ejecutable:**
   - Ve a: `https://github.com/TU_USUARIO/claude-foundry-manager/releases`
   - Descarga: `claude-foundry-manager-windows-amd64.exe`
   - Guárdalo donde quieras
   - ¡Listo! Ejecuta el archivo

**✅ Ventajas:**
- No necesitas instalar Go
- Obtienes binarios para Windows, Linux y macOS
- Proceso automatizado

---

## 🔨 OPCIÓN 2: Compilar Localmente con Scripts de Ayuda

Si prefieres compilar en tu PC, usa los scripts que creé.

### Pasos:

1. **Ejecuta el script de setup:**

   ```powershell
   # En PowerShell (modo administrador recomendado)
   .\setup-windows.ps1
   ```

   El script:
   - Verifica si Go está instalado
   - Si no lo está, te ofrece descargarlo
   - Si ya lo tienes, compila el proyecto automáticamente
   - Te dice exactamente qué hacer en cada paso

2. **Resultado:**
   - Archivo creado: `claude-foundry-manager.exe`
   - Listo para usar

**✅ Ventajas:**
- Control total del proceso
- Puedes modificar el código y recompilar
- No dependes de GitHub

---

## ⚙️ OPCIÓN 3: Compilación Manual (Para Desarrolladores)

Si ya tienes Go instalado y quieres hacerlo manualmente:

```powershell
# Descargar dependencias
go mod download

# Compilar
go build -ldflags="-s -w" -o claude-foundry-manager.exe .

# Ejecutar
.\claude-foundry-manager.exe
```

---

## 🎯 Después de Tener el Ejecutable

### Primer Uso (Modo Interactivo):

```powershell
# Ejecutar como administrador (clic derecho → Ejecutar como administrador)
.\claude-foundry-manager.exe
```

Verás un menú con opciones:
```
[1] Configure Azure Foundry
[2] Rollback to Default (Direct Anthropic)
[3] View Current Configuration
[4] List Available Backups
[5] Restore from Backup
[6] Save Manual Backup
[7] Exit
```

### Uso Rápido (CLI):

```powershell
# Configurar Azure Foundry
.\claude-foundry-manager.exe configure --resource=my-foundry --api-key=sk-xxx

# Ver configuración actual
.\claude-foundry-manager.exe show

# Hacer rollback a Anthropic directo
.\claude-foundry-manager.exe rollback

# Ver todos los comandos disponibles
.\claude-foundry-manager.exe --help
```

---

## 📝 Resumen de Archivos Útiles

| Archivo | Propósito |
|---------|-----------|
| `setup-windows.ps1` | Script de instalación automática (PowerShell) |
| `build.bat` | Script de compilación simple (Batch) |
| `INSTALL.md` | Guía detallada de instalación |
| `README.md` | Documentación completa del proyecto |
| `GET-STARTED.md` | Este archivo (inicio rápido) |

---

## ❓ ¿Problemas?

### "Go no está instalado"
→ Ejecuta `setup-windows.ps1` y sigue las instrucciones

### "Access denied"
→ Ejecuta el .exe como Administrador (clic derecho)

### "No tengo tiempo para configurar Go"
→ Usa la OPCIÓN 1 (GitHub Releases) - la más fácil

### "El binario no funciona"
→ Verifica que ejecutas como Administrador (necesario en Windows)

---

## 🎉 ¿Qué Sigue?

Una vez que tengas el ejecutable funcionando:

1. **Configura Azure Foundry:**
   ```powershell
   .\claude-foundry-manager.exe configure --resource=TU_RECURSO
   ```

2. **Reinicia tu terminal** para que los cambios tomen efecto

3. **Prueba Claude Code:**
   ```powershell
   claude-code --version
   ```

4. **Si algo sale mal:**
   - Haz rollback: `.\claude-foundry-manager.exe rollback`
   - O restaura un backup: `.\claude-foundry-manager.exe backup restore`

---

## 💡 Recomendación

**Para la mayoría de usuarios:** Usa la **OPCIÓN 1** (GitHub Releases)
- Es la más simple
- No necesitas instalar nada
- Funciona inmediatamente

**Si eres desarrollador:** Usa la **OPCIÓN 2** (Compilación local)
- Puedes modificar el código
- Recompilas cuando quieras
- Control completo

---

**¿Listo para empezar?** 🚀

Elige una opción y sigue los pasos. Si tienes dudas, revisa `INSTALL.md` para más detalles.
