# Claude Code - Azure Foundry Configuration Manager (Python)

Herramienta en Python para gestionar la configuración de Claude Code con Azure Foundry de manera sencilla en Windows.

## 🎯 Características

- ✅ Configuración completa de Azure Foundry para Claude Code
- 🔄 Rollback instantáneo a provider default de Anthropic
- 💾 Sistema automático de backups en JSON
- 🔐 Soporte para API Key o autenticación Entra ID
- 🌍 Configuración a nivel global (sistema)
- 📋 Interfaz interactiva con colores
- 🐍 Funciona en Python sin problemas de Execution Policy

## 📋 Requisitos

- Windows 10/11 o Windows Server
- Python 3.7 o superior
- Privilegios de Administrador

## 🚀 Instalación

### 1. Instalar Python

Si no tienes Python instalado:

1. Descarga Python desde: https://www.python.org/downloads/
2. Durante la instalación, **marca la casilla "Add Python to PATH"**
3. Completa la instalación

Para verificar que Python está instalado:
```cmd
python --version
```

### 2. Descargar los scripts

Descarga estos archivos a una carpeta (ej: `C:\claude-foundry-manager\`):
- `claude_foundry_manager.py` - Script principal con menú
- `claude_foundry_quick.py` - Script rápido para línea de comandos
- `run-manager.bat` - Launcher para ejecutar fácilmente

## 🎮 Uso

### Opción 1: Interfaz Interactiva (Recomendado)

**Forma más fácil** - Doble clic en `run-manager.bat`

O desde CMD/PowerShell como Administrador:
```cmd
python claude_foundry_manager.py
```

El menú te mostrará estas opciones:
```
[1] Configurar Azure Foundry
[2] Rollback a configuración default (Anthropic)
[3] Ver configuración actual
[4] Listar backups disponibles
[5] Restaurar desde backup
[6] Guardar backup manual
[0] Salir
```

### Opción 2: Línea de Comandos Rápida

Para configurar Azure Foundry con Entra ID:
```cmd
python claude_foundry_quick.py --resource mi-recurso-foundry
```

Para configurar con API Key:
```cmd
python claude_foundry_quick.py --resource mi-recurso-foundry --api-key tu-api-key
```

Para configurar con modelos personalizados:
```cmd
python claude_foundry_quick.py --resource mi-recurso ^
    --sonnet-model my-sonnet-deployment ^
    --haiku-model my-haiku-deployment ^
    --opus-model my-opus-deployment
```

Para hacer rollback a default:
```cmd
python claude_foundry_quick.py --rollback
```

Ver ayuda completa:
```cmd
python claude_foundry_quick.py --help
```

## 📖 Guía Paso a Paso

### Configurar Azure Foundry

1. Ejecuta `run-manager.bat` (doble clic)
2. Selecciona opción `[1]`
3. Ingresa tu información:
   - **Recurso de Azure**: Nombre de tu recurso en Azure (ej: `contoso-foundry`)
   - **API Key**: (Opcional) Tu API key, o Enter para usar Entra ID
   - **Modelos**: Nombres de tus deployments (o Enter para defaults)
4. Confirma con `S`
5. El script guardará un backup automático
6. **¡Reinicia tu terminal!** Para que los cambios surtan efecto

### Hacer Rollback a Default

1. Ejecuta el manager
2. Selecciona opción `[2]`
3. Confirma con `S`
4. El script eliminará toda la configuración de Azure Foundry
5. **¡Reinicia tu terminal!**

### Ver Configuración Actual

Selecciona opción `[3]` para ver todas las variables de entorno configuradas.

### Gestionar Backups

- **Listar backups**: Opción `[4]`
- **Restaurar backup**: Opción `[5]` - Elige el número del backup
- **Crear backup manual**: Opción `[6]` - Con descripción personalizada

## 🔧 Variables de Entorno

El script gestiona estas variables a nivel de sistema:

```
CLAUDE_CODE_USE_FOUNDRY=1
ANTHROPIC_FOUNDRY_RESOURCE=tu-recurso
ANTHROPIC_FOUNDRY_API_KEY=tu-api-key (opcional)
ANTHROPIC_DEFAULT_SONNET_MODEL=claude-sonnet-4-5
ANTHROPIC_DEFAULT_HAIKU_MODEL=claude-haiku-4-5
ANTHROPIC_DEFAULT_OPUS_MODEL=claude-opus-4-1
```

## 💾 Sistema de Backups

- **Ubicación**: `%USERPROFILE%\.claude-code-backups\`
- **Formato**: `backup_YYYYMMDD_HHMMSS.json`
- **Automático**: Se crea antes de cada cambio importante
- **Manual**: Crea backups cuando quieras con descripción personalizada

### Ejemplo de backup:

```json
{
  "timestamp": "20251120_143022",
  "description": "Antes de configurar Azure Foundry",
  "variables": {
    "CLAUDE_CODE_USE_FOUNDRY": "1",
    "ANTHROPIC_FOUNDRY_RESOURCE": "mi-recurso",
    "ANTHROPIC_DEFAULT_SONNET_MODEL": "claude-sonnet-4-5",
    ...
  }
}
```

## 🔐 Seguridad

- Las API Keys se muestran parcialmente enmascaradas
- Los backups contienen las API Keys completas (protege la carpeta de backups)
- El script requiere privilegios de administrador para modificar variables del sistema

## ⚠️ Importante

1. **Siempre reinicia tu terminal** después de hacer cambios
2. Las variables se configuran a nivel de **sistema** (Machine), no de usuario
3. Los backups son tu red de seguridad - consérvales

## 🐛 Troubleshooting

### Python no está instalado

**Síntoma**: `python: command not found` o error al ejecutar

**Solución**:
1. Instala Python desde https://www.python.org/downloads/
2. Durante instalación, marca "Add Python to PATH"
3. Reinicia tu terminal
4. Verifica: `python --version`

### Error de permisos

**Síntoma**: `PermissionError` o acceso denegado

**Solución**:
1. Clic derecho en `run-manager.bat`
2. Selecciona "Ejecutar como administrador"

O desde CMD/PowerShell:
- Clic derecho en CMD/PowerShell
- "Ejecutar como administrador"
- Navega a la carpeta y ejecuta el script

### Los cambios no se aplican

**Síntoma**: Las variables no aparecen en Claude Code

**Solución**:
1. **Cierra TODAS las ventanas de terminal/CMD/PowerShell**
2. Abre una nueva ventana
3. Verifica: `echo %CLAUDE_CODE_USE_FOUNDRY%`

### Error de autenticación con Azure

**Síntoma**: `Failed to get token from azureADTokenProvider`

**Solución**:
- Si usas Entra ID: Ejecuta `az login` en tu terminal
- Si prefieres API Key: Configura con `--api-key` tu key de Azure

### ImportError: No module named 'winreg'

**Síntoma**: Error al importar winreg

**Solución**: Asegúrate de estar usando Python en Windows. El módulo `winreg` es específico de Windows.

## 📚 Referencias

- [Claude Code - Azure AI Foundry Documentation](https://code.claude.com/docs/en/azure-ai-foundry)
- [Azure AI Foundry Documentation](https://learn.microsoft.com/en-us/azure/ai-foundry/)
- [Anthropic - Claude in Microsoft Foundry](https://docs.claude.com/en/docs/build-with-claude/claude-in-microsoft-foundry)

## 💡 Tips

### Múltiples configuraciones

Si trabajas con varios recursos de Azure:

1. Crea backups manuales con nombres descriptivos:
   - "Config Desarrollo"
   - "Config Producción"
   - "Config Testing"

2. Restaura el que necesites según tu contexto

### Autenticación Entra ID vs API Key

**Entra ID (Recomendado para empresas)**:
- ✅ Más seguro
- ✅ No requiere gestionar API keys
- ✅ Usa identidades de Azure AD
- ⚠️ Requiere Azure CLI: `az login`

**API Key (Más simple)**:
- ✅ Configuración directa
- ✅ Ideal para desarrollo local
- ⚠️ Debes gestionar la key manualmente
- 💡 Obtén la key desde Azure AI Foundry Portal

### Verificar configuración actual

Desde cualquier terminal:
```cmd
echo %CLAUDE_CODE_USE_FOUNDRY%
echo %ANTHROPIC_FOUNDRY_RESOURCE%
echo %ANTHROPIC_DEFAULT_SONNET_MODEL%
```

## 🆘 Soporte

Si tienes problemas:

1. ✅ Verifica que Python está instalado: `python --version`
2. ✅ Asegúrate de ejecutar como Administrador
3. ✅ Revisa la sección de Troubleshooting arriba
4. ✅ Consulta la documentación oficial de Claude Code

## 📝 Estructura de Archivos

```
claude-foundry-manager/
├── claude_foundry_manager.py    # Script principal con menú
├── claude_foundry_quick.py      # Script rápido CLI
├── run-manager.bat              # Launcher con auto-elevación
├── README.md                    # Este archivo
└── configuration-examples.json  # Ejemplos de configuración
```

## 🔄 Comparación con PowerShell

| Característica | Python | PowerShell |
|----------------|--------|------------|
| Execution Policy | ✅ Sin problemas | ❌ Puede bloquear |
| Portabilidad | ✅ Multiplataforma | ⚠️ Solo Windows |
| Facilidad | ✅ Fácil de ejecutar | ⚠️ Puede complicarse |
| Dependencias | Python 3.7+ | PowerShell 5.1+ |

---

**Versión**: 2.0 (Python Edition)  
**Autor**: José Díaz  
**Última actualización**: Noviembre 2025
