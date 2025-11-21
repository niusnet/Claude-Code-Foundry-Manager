# Claude Code Azure Foundry - Guía de Inicio Rápido

## 🚀 Primeros Pasos (5 minutos)

### 1. Verificar Python
```cmd
python --version
```
Si no tienes Python: https://www.python.org/downloads/ (marca "Add to PATH")

### 2. Ejecutar el Manager
Doble clic en: `run-manager.bat`

O desde CMD como Administrador:
```cmd
python claude_foundry_manager.py
```

### 3. Configurar Azure Foundry
- Selecciona opción `[1]`
- Ingresa tu recurso de Azure
- Presiona Enter para usar Entra ID (o ingresa tu API Key)
- Confirma con `S`
- **¡Reinicia tu terminal!**

¡Listo! Claude Code ahora usa Azure Foundry.

---

## 📝 Ejemplos de Uso

### Ejemplo 1: Configuración básica con Entra ID
```cmd
python claude_foundry_quick.py --resource contoso-foundry
```

### Ejemplo 2: Configuración con API Key
```cmd
python claude_foundry_quick.py --resource contoso-foundry --api-key sk-ant-api03-...
```

### Ejemplo 3: Configuración con deployments personalizados
```cmd
python claude_foundry_quick.py --resource contoso-foundry ^
    --sonnet-model my-sonnet-prod ^
    --haiku-model my-haiku-prod ^
    --opus-model my-opus-prod
```

### Ejemplo 4: Volver a configuración default
```cmd
python claude_foundry_quick.py --rollback
```

### Ejemplo 5: Ver ayuda completa
```cmd
python claude_foundry_quick.py --help
```

---

## 🔍 Verificar Configuración

Después de configurar, verifica en una terminal nueva:

```cmd
echo %CLAUDE_CODE_USE_FOUNDRY%
echo %ANTHROPIC_FOUNDRY_RESOURCE%
echo %ANTHROPIC_DEFAULT_SONNET_MODEL%
```

Deberías ver:
```
1
tu-recurso-azure
claude-sonnet-4-5
```

---

## 🎯 Escenarios Comunes

### Escenario 1: Primera vez configurando Azure Foundry
```cmd
# 1. Ejecuta el manager
run-manager.bat

# 2. Selecciona [1] Configurar Azure Foundry
# 3. Ingresa tu recurso: contoso-foundry
# 4. API Key: [Enter para Entra ID]
# 5. Modelos: [Enter para defaults]
# 6. Confirma: S
# 7. Reinicia tu terminal
```

### Escenario 2: Cambiar entre Azure Foundry y Anthropic directo
```cmd
# Ver configuración actual
python claude_foundry_manager.py
# Selecciona [3]

# Hacer rollback a Anthropic
python claude_foundry_quick.py --rollback
# Reinicia terminal

# Volver a Azure Foundry (restaurar último backup)
run-manager.bat
# Selecciona [5] y elige el último backup
# Reinicia terminal
```

### Escenario 3: Trabajar con múltiples configuraciones
```cmd
# Crear backup de configuración actual
run-manager.bat
# Selecciona [6] y nombra: "Config Producción"

# Configurar para desarrollo
python claude_foundry_quick.py --resource dev-foundry
# Reinicia terminal

# Crear backup de desarrollo
run-manager.bat
# Selecciona [6] y nombra: "Config Desarrollo"

# Cambiar entre ellas:
run-manager.bat
# Selecciona [5] y elige la que necesites
```

---

## 🔧 Troubleshooting Rápido

### Problema: "Python no está instalado"
**Solución**:
1. Descarga: https://www.python.org/downloads/
2. Durante instalación: ✅ "Add Python to PATH"
3. Reinicia terminal
4. Verifica: `python --version`

### Problema: "Access Denied" o "Permission Denied"
**Solución**:
- Clic derecho en CMD → "Ejecutar como administrador"
- O usa `run-manager.bat` (solicita privilegios automáticamente)

### Problema: "Los cambios no se aplican en Claude Code"
**Solución**:
1. Cierra TODAS las terminales/CMD/PowerShell
2. Abre una nueva
3. Verifica: `echo %CLAUDE_CODE_USE_FOUNDRY%`
4. Si sale "1", está configurado correctamente

### Problema: "Failed to get token from azureADTokenProvider"
**Solución**:
```cmd
# Opción 1: Login con Azure CLI
az login

# Opción 2: Usar API Key en su lugar
python claude_foundry_quick.py --resource tu-recurso --api-key tu-key
```

---

## 📋 Checklist de Configuración

- [ ] Python 3.7+ instalado
- [ ] Script ejecutado como Administrador
- [ ] Recurso de Azure configurado
- [ ] Modelos configurados (o defaults aceptados)
- [ ] Backup automático creado
- [ ] Terminal reiniciada
- [ ] Variables verificadas con `echo %CLAUDE_CODE_USE_FOUNDRY%`
- [ ] Claude Code ejecutado y conectado a Azure Foundry

---

## 🆘 Comandos de Emergencia

### Ver todo lo configurado
```cmd
run-manager.bat
# Selecciona [3]
```

### Volver a default (emergencia)
```cmd
python claude_foundry_quick.py --rollback
```

### Ver backups disponibles
```cmd
run-manager.bat
# Selecciona [4]
```

### Restaurar última configuración que funcionaba
```cmd
run-manager.bat
# Selecciona [5]
# Elige el backup más reciente que funcionaba
```

---

## 💡 Tips Pro

1. **Crea un backup manual antes de experimentar**
   ```cmd
   run-manager.bat
   # [6] → "Antes de probar nueva config"
   ```

2. **Usa nombres descriptivos en backups**
   - "Config Producción"
   - "Config con API Key"
   - "Config Entra ID"

3. **Verifica siempre después de cambios**
   ```cmd
   # En una terminal NUEVA
   echo %CLAUDE_CODE_USE_FOUNDRY%
   ```

4. **Mantén un backup "working"**
   - Cuando tengas una config que funcione perfecto
   - Crea un backup manual: "WORKING CONFIG - NO BORRAR"

---

## 📞 Flujo Recomendado para Primera Vez

```
1. Descarga todos los archivos
   ↓
2. Ejecuta: install-check.bat
   ↓
3. Si todo OK → run-manager.bat
   ↓
4. Selecciona [1] Configurar Azure Foundry
   ↓
5. Ingresa tu información
   ↓
6. Confirma con S
   ↓
7. REINICIA TU TERMINAL
   ↓
8. Verifica: echo %CLAUDE_CODE_USE_FOUNDRY%
   ↓
9. Ejecuta Claude Code
   ↓
10. ¡Listo! 🎉
```

---

## 🔗 Links Útiles

- [Documentación Claude Code con Azure](https://code.claude.com/docs/en/azure-ai-foundry)
- [Azure AI Foundry Portal](https://ai.azure.com/)
- [Descargar Python](https://www.python.org/downloads/)
- [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)

---

**¿Problemas?** Revisa el README-Python.md completo para más detalles.
