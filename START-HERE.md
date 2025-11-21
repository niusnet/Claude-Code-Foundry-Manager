# 🚀 EMPIEZA AQUÍ - Claude Code Azure Foundry Manager

## ⚡ Inicio Rápido (2 minutos)

### 1️⃣ Verifica Python
Abre CMD y ejecuta:
```cmd
python --version
```

**¿No tienes Python?** 
→ Descarga de: https://www.python.org/downloads/ 
→ ⚠️ IMPORTANTE: Marca "Add Python to PATH" durante instalación

### 2️⃣ Ejecuta el Manager
**Opción A** (Más fácil): Doble clic en → `run-manager.bat`

**Opción B**: Desde CMD como Administrador:
```cmd
python claude_foundry_manager.py
```

### 3️⃣ Configura
- Selecciona `[1]` Configurar Azure Foundry
- Ingresa tu recurso de Azure
- Presiona Enter para Entra ID (o ingresa API Key)
- Confirma con `S`

### 4️⃣ Reinicia tu terminal
**¡MUY IMPORTANTE!** Cierra y abre una nueva terminal.

### ✅ ¡Listo!
Claude Code ahora usa Azure Foundry.

---

## 📦 Archivos Incluidos

### 🐍 **Python (RECOMENDADO - Sin problemas de execution policy)**

| Archivo | Descripción | Uso |
|---------|-------------|-----|
| `claude_foundry_manager.py` | Script principal con menú interactivo | `python claude_foundry_manager.py` |
| `claude_foundry_quick.py` | Script CLI rápido | `python claude_foundry_quick.py --resource tu-recurso` |
| `run-manager.bat` | Ejecuta el manager fácilmente | Doble clic |
| `install-check.bat` | Verifica Python y requisitos | Doble clic |
| `README-Python.md` | Documentación completa Python | Lee esto si tienes dudas |
| `QUICK-START.md` | Guía rápida con ejemplos | Ejemplos de uso |

### 📜 **PowerShell (Alternativo)**

| Archivo | Descripción | Uso |
|---------|-------------|-----|
| `ClaudeCode-AzureFoundry-Manager.ps1` | Script PowerShell con menú | Como Admin: `.\ClaudeCode...ps1` |
| `ClaudeCode-AzureFoundry-Quick.ps1` | PowerShell CLI rápido | Con parámetros |
| `Run-Manager.bat` | Ejecuta PowerShell script | Doble clic |
| `README.md` | Documentación PowerShell | Documentación completa |

### 📋 **Referencia**

| Archivo | Descripción |
|---------|-------------|
| `configuration-examples.json` | Ejemplos de configuración y referencia |

---

## 🎯 ¿Cuál usar? Python o PowerShell

### ✅ Usa Python si:
- ✅ Tuviste problemas con PowerShell execution policy
- ✅ Prefieres algo más estándar y portable
- ✅ Tienes o puedes instalar Python fácilmente
- ✅ Quieres evitar problemas de firma digital

### ⚠️ Usa PowerShell si:
- ✅ No quieres instalar Python
- ✅ Ya tienes PowerShell configurado con RemoteSigned
- ✅ Prefieres scripts nativos de Windows

**Recomendación**: Python es más fácil y directo.

---

## 🔥 Comandos Más Usados

### Ver configuración actual
```cmd
python claude_foundry_manager.py
# Selecciona [3]
```

### Configurar Azure Foundry
```cmd
python claude_foundry_quick.py --resource mi-recurso-azure
```

### Volver a Anthropic default
```cmd
python claude_foundry_quick.py --rollback
```

### Ver backups
```cmd
python claude_foundry_manager.py
# Selecciona [4]
```

---

## ⚠️ ¡IMPORTANTE!

### Después de CUALQUIER cambio:
1. **Cierra TODAS tus terminales**
2. **Abre una nueva terminal**
3. **Verifica**: `echo %CLAUDE_CODE_USE_FOUNDRY%`

Las variables de entorno no se actualizan hasta que reinicias la terminal.

---

## 🆘 Solución Rápida de Problemas

### "Python no está instalado"
```
1. https://www.python.org/downloads/
2. Instala (marca "Add to PATH")
3. Reinicia terminal
4. Verifica: python --version
```

### "Permission Denied" / "Access Denied"
```
Clic derecho en CMD → Ejecutar como administrador
O usa: run-manager.bat (eleva automáticamente)
```

### "Los cambios no funcionan"
```
1. Cierra TODAS las terminales
2. Abre nueva terminal
3. echo %CLAUDE_CODE_USE_FOUNDRY%
4. Debe mostrar: 1
```

---

## 📖 Documentación

- **Python**: Lee `README-Python.md` para documentación completa
- **PowerShell**: Lee `README.md` para la versión PowerShell
- **Ejemplos**: Lee `QUICK-START.md` para ejemplos de uso
- **Referencia**: Revisa `configuration-examples.json` para opciones

---

## 🎬 Flujo Recomendado Primera Vez

```
┌─────────────────────────────────┐
│ 1. Ejecuta: install-check.bat  │
└───────────┬─────────────────────┘
            │
            ▼
┌─────────────────────────────────┐
│ 2. Si OK: run-manager.bat      │
└───────────┬─────────────────────┘
            │
            ▼
┌─────────────────────────────────┐
│ 3. Opción [1] Configurar        │
└───────────┬─────────────────────┘
            │
            ▼
┌─────────────────────────────────┐
│ 4. Ingresa tu recurso Azure     │
└───────────┬─────────────────────┘
            │
            ▼
┌─────────────────────────────────┐
│ 5. Confirma con S               │
└───────────┬─────────────────────┘
            │
            ▼
┌─────────────────────────────────┐
│ 6. REINICIA TERMINAL            │
└───────────┬─────────────────────┘
            │
            ▼
┌─────────────────────────────────┐
│ 7. Ejecuta Claude Code          │
└───────────┬─────────────────────┘
            │
            ▼
       ✅ ¡LISTO! 🎉
```

---

## 🌟 Features Principales

✅ Configuración completa de Azure Foundry  
✅ Rollback a Anthropic con un comando  
✅ Sistema de backups automático  
✅ Soporte API Key o Entra ID  
✅ Interfaz con colores y fácil de usar  
✅ Configuración global (sistema)  
✅ Restaurar backups antiguos  
✅ Sin problemas de execution policy (Python)  

---

## 💡 Tips Pro

1. **Crea backup antes de experimentar**
   - Opción [6] en el manager
   - Guárdalo con nombre descriptivo

2. **Para múltiples configs**
   - Crea backup de cada una
   - Restaura la que necesites (Opción [5])

3. **Verifica siempre después de cambios**
   ```cmd
   echo %CLAUDE_CODE_USE_FOUNDRY%
   ```

---

## 📞 Soporte

¿Problemas? Sigue este orden:

1. ✅ Lee `QUICK-START.md` - Ejemplos y troubleshooting
2. ✅ Lee `README-Python.md` - Documentación completa
3. ✅ Revisa `configuration-examples.json` - Opciones disponibles
4. ✅ Consulta docs oficiales: https://code.claude.com/docs/

---

**Versión**: 2.0  
**Stack**: Python 3.7+ o PowerShell 5.1+  
**Autor**: José Díaz  
**Fecha**: Noviembre 2025

---

## 🔗 Links Útiles

- [Claude Code Docs](https://code.claude.com/docs/en/azure-ai-foundry)
- [Azure AI Foundry](https://ai.azure.com/)
- [Python Download](https://www.python.org/downloads/)
- [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)

---

**¿Listo para empezar?** → Doble clic en `run-manager.bat` 🚀
