#!/usr/bin/env python3
"""
Claude Code - Azure Foundry Configuration Manager
Gestiona la configuración de Claude Code para Azure Foundry en Windows
"""

import os
import sys
import json
import winreg
from datetime import datetime
from pathlib import Path
from typing import Dict, Optional, List

# Configuración
BACKUP_DIR = Path.home() / ".claude-code-backups"
CLAUDE_CODE_VARS = [
    'CLAUDE_CODE_USE_FOUNDRY',
    'ANTHROPIC_FOUNDRY_RESOURCE',
    'ANTHROPIC_FOUNDRY_BASE_URL',
    'ANTHROPIC_FOUNDRY_API_KEY',
    'ANTHROPIC_DEFAULT_SONNET_MODEL',
    'ANTHROPIC_DEFAULT_HAIKU_MODEL',
    'ANTHROPIC_DEFAULT_OPUS_MODEL'
]

# Colores ANSI para terminal
class Colors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKCYAN = '\033[96m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'

def is_admin():
    """Verifica si el script se está ejecutando con privilegios de administrador"""
    try:
        import ctypes
        return ctypes.windll.shell32.IsUserAnAdmin()
    except:
        return False

def print_colored(message: str, color: str = Colors.OKCYAN):
    """Imprime mensaje con color"""
    print(f"{color}{message}{Colors.ENDC}")

def print_success(message: str):
    """Imprime mensaje de éxito"""
    print_colored(f"✓ {message}", Colors.OKGREEN)

def print_error(message: str):
    """Imprime mensaje de error"""
    print_colored(f"❌ {message}", Colors.FAIL)

def print_warning(message: str):
    """Imprime mensaje de advertencia"""
    print_colored(f"⚠️  {message}", Colors.WARNING)

def print_info(message: str):
    """Imprime mensaje informativo"""
    print_colored(f"ℹ️  {message}", Colors.OKCYAN)

def show_banner():
    """Muestra el banner del programa"""
    os.system('cls' if os.name == 'nt' else 'clear')
    print()
    print_colored("╔════════════════════════════════════════════════════════════╗", Colors.OKCYAN)
    print_colored("║     Claude Code - Azure Foundry Configuration Manager     ║", Colors.OKCYAN)
    print_colored("║                      Python Edition                        ║", Colors.OKCYAN)
    print_colored("╚════════════════════════════════════════════════════════════╝", Colors.OKCYAN)
    print()

def get_env_var(name: str) -> Optional[str]:
    """Obtiene una variable de entorno del sistema (Machine level)"""
    try:
        with winreg.OpenKey(winreg.HKEY_LOCAL_MACHINE, 
                           r"SYSTEM\CurrentControlSet\Control\Session Manager\Environment",
                           0, winreg.KEY_READ) as key:
            try:
                value, _ = winreg.QueryValueEx(key, name)
                return value
            except FileNotFoundError:
                return None
    except Exception as e:
        print_error(f"Error al leer variable {name}: {e}")
        return None

def set_env_var(name: str, value: Optional[str]):
    """Establece o elimina una variable de entorno del sistema (Machine level)"""
    try:
        with winreg.OpenKey(winreg.HKEY_LOCAL_MACHINE,
                           r"SYSTEM\CurrentControlSet\Control\Session Manager\Environment",
                           0, winreg.KEY_WRITE) as key:
            if value is None:
                try:
                    winreg.DeleteValue(key, name)
                except FileNotFoundError:
                    pass
            else:
                winreg.SetValueEx(key, name, 0, winreg.REG_SZ, value)
        
        # Notificar al sistema del cambio
        import ctypes
        HWND_BROADCAST = 0xFFFF
        WM_SETTINGCHANGE = 0x1A
        SMTO_ABORTIFHUNG = 0x0002
        result = ctypes.c_long()
        ctypes.windll.user32.SendMessageTimeoutW(
            HWND_BROADCAST, WM_SETTINGCHANGE, 0, "Environment", 
            SMTO_ABORTIFHUNG, 5000, ctypes.byref(result)
        )
    except Exception as e:
        print_error(f"Error al establecer variable {name}: {e}")
        raise

def get_current_config() -> Dict[str, str]:
    """Obtiene la configuración actual"""
    config = {}
    for var in CLAUDE_CODE_VARS:
        value = get_env_var(var)
        if value:
            config[var] = value
    return config

def show_current_config():
    """Muestra la configuración actual"""
    print_info("Configuración actual:")
    print()
    
    config = get_current_config()
    
    if not config:
        print_warning("No hay configuración de Azure Foundry activa.")
        print_info("Claude Code está usando el provider default (Anthropic directo).")
    else:
        for var, value in config.items():
            if var == 'ANTHROPIC_FOUNDRY_API_KEY':
                masked_value = value[:8] + "..." if len(value) > 8 else "***"
                print_success(f"{var} = {masked_value}")
            else:
                print_success(f"{var} = {value}")
    
    print()
    return bool(config)

def save_backup(description: str = "Manual") -> Optional[Path]:
    """Guarda un backup de la configuración actual"""
    try:
        BACKUP_DIR.mkdir(parents=True, exist_ok=True)
        
        timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
        backup_file = BACKUP_DIR / f"backup_{timestamp}.json"
        
        backup = {
            "timestamp": timestamp,
            "description": description,
            "variables": get_current_config()
        }
        
        with open(backup_file, 'w', encoding='utf-8') as f:
            json.dump(backup, f, indent=2, ensure_ascii=False)
        
        print_success(f"Backup guardado: {backup_file}")
        return backup_file
    except Exception as e:
        print_error(f"Error al guardar backup: {e}")
        return None

def list_backups() -> List[Path]:
    """Lista los backups disponibles"""
    if not BACKUP_DIR.exists():
        return []
    
    return sorted(BACKUP_DIR.glob("backup_*.json"), reverse=True)

def show_backups():
    """Muestra los backups disponibles"""
    print_info("Backups disponibles:")
    print()
    
    backups = list_backups()
    
    if not backups:
        print_info("No hay backups disponibles.")
        return
    
    for idx, backup_file in enumerate(backups, 1):
        try:
            with open(backup_file, 'r', encoding='utf-8') as f:
                content = json.load(f)
            
            timestamp = datetime.strptime(content['timestamp'], "%Y%m%d_%H%M%S")
            date_str = timestamp.strftime("%Y-%m-%d %H:%M:%S")
            
            print_info(f"[{idx}] {date_str} - {content['description']}")
            print_colored(f"    Archivo: {backup_file.name}", Colors.OKCYAN)
        except Exception as e:
            print_warning(f"Error al leer backup: {backup_file.name}")
    
    print()
    print_info(f"Los backups se guardan en: {BACKUP_DIR}")

def configure_azure_foundry():
    """Configura Azure Foundry para Claude Code"""
    show_banner()
    print_colored("🔧 Configuración de Azure Foundry para Claude Code", Colors.OKCYAN)
    print()
    print_info("Por favor, proporciona la siguiente información:")
    print()
    
    # Recopilar información
    resource = input(f"{Colors.BOLD}Nombre del recurso de Azure Foundry: {Colors.ENDC}").strip()
    
    if not resource:
        print_error("El nombre del recurso es obligatorio.")
        return False
    
    print()
    print_info("API Key de Azure (opcional, presiona Enter para usar Entra ID):")
    api_key = input(f"{Colors.BOLD}API Key: {Colors.ENDC}").strip()
    
    print()
    print_info("Nombres de deployment de modelos (presiona Enter para usar defaults):")
    
    sonnet_model = input(f"{Colors.BOLD}Sonnet Model (default: claude-sonnet-4-5): {Colors.ENDC}").strip()
    sonnet_model = sonnet_model or 'claude-sonnet-4-5'
    
    haiku_model = input(f"{Colors.BOLD}Haiku Model (default: claude-haiku-4-5): {Colors.ENDC}").strip()
    haiku_model = haiku_model or 'claude-haiku-4-5'
    
    opus_model = input(f"{Colors.BOLD}Opus Model (default: claude-opus-4-1): {Colors.ENDC}").strip()
    opus_model = opus_model or 'claude-opus-4-1'
    
    # Confirmación
    print()
    print_colored("━" * 60, Colors.OKCYAN)
    print_info("Resumen de configuración:")
    print()
    print_info(f"  • Recurso Azure: {resource}")
    print_info(f"  • API Key: {'Configurada (oculta)' if api_key else 'No configurada (usará Entra ID)'}")
    print_info(f"  • Sonnet Model: {sonnet_model}")
    print_info(f"  • Haiku Model: {haiku_model}")
    print_info(f"  • Opus Model: {opus_model}")
    print_colored("━" * 60, Colors.OKCYAN)
    print()
    
    confirm = input(f"{Colors.BOLD}¿Confirmas esta configuración? (S/N): {Colors.ENDC}").strip().upper()
    
    if confirm != 'S':
        print_warning("Configuración cancelada.")
        return False
    
    # Guardar backup
    print()
    print_info("💾 Guardando backup de configuración actual...")
    save_backup("Antes de configurar Azure Foundry")
    
    # Aplicar configuración
    print()
    print_info("🔄 Aplicando configuración...")
    
    try:
        set_env_var('CLAUDE_CODE_USE_FOUNDRY', '1')
        print_success("CLAUDE_CODE_USE_FOUNDRY habilitado")
        
        set_env_var('ANTHROPIC_FOUNDRY_RESOURCE', resource)
        print_success("ANTHROPIC_FOUNDRY_RESOURCE configurado")
        
        if api_key:
            set_env_var('ANTHROPIC_FOUNDRY_API_KEY', api_key)
            print_success("ANTHROPIC_FOUNDRY_API_KEY configurada")
        
        set_env_var('ANTHROPIC_DEFAULT_SONNET_MODEL', sonnet_model)
        print_success("ANTHROPIC_DEFAULT_SONNET_MODEL configurado")
        
        set_env_var('ANTHROPIC_DEFAULT_HAIKU_MODEL', haiku_model)
        print_success("ANTHROPIC_DEFAULT_HAIKU_MODEL configurado")
        
        set_env_var('ANTHROPIC_DEFAULT_OPUS_MODEL', opus_model)
        print_success("ANTHROPIC_DEFAULT_OPUS_MODEL configurado")
        
        print()
        print_colored("━" * 60, Colors.OKGREEN)
        print_success("Configuración aplicada exitosamente!")
        print_colored("━" * 60, Colors.OKGREEN)
        print()
        print_warning("IMPORTANTE: Reinicia tu terminal o sesión para que los cambios surtan efecto.")
        print()
        
        return True
    except Exception as e:
        print_error(f"Error al aplicar configuración: {e}")
        return False

def rollback_to_default():
    """Hace rollback a la configuración default de Claude Code"""
    show_banner()
    print_info("🔄 Rollback a configuración default de Claude Code")
    print()
    
    has_config = show_current_config()
    
    if not has_config:
        print_info("Ya estás usando la configuración default. No hay nada que revertir.")
        return True
    
    print()
    print_warning("Esta acción eliminará toda la configuración de Azure Foundry")
    print_warning("y volverá a usar el provider default de Anthropic.")
    print()
    
    confirm = input(f"{Colors.BOLD}¿Deseas continuar? (S/N): {Colors.ENDC}").strip().upper()
    
    if confirm != 'S':
        print_warning("Rollback cancelado.")
        return False
    
    # Guardar backup
    print()
    print_info("💾 Guardando backup de configuración actual...")
    save_backup("Antes de rollback a default")
    
    # Eliminar variables
    print()
    print_info("🗑️  Eliminando configuración de Azure Foundry...")
    
    try:
        for var in CLAUDE_CODE_VARS:
            if get_env_var(var):
                set_env_var(var, None)
                print_success(f"{var} eliminada")
        
        print()
        print_colored("━" * 60, Colors.OKGREEN)
        print_success("Rollback completado exitosamente!")
        print_colored("━" * 60, Colors.OKGREEN)
        print()
        print_info("Claude Code ahora usará el provider default de Anthropic.")
        print_warning("IMPORTANTE: Reinicia tu terminal o sesión para que los cambios surtan efecto.")
        print()
        
        return True
    except Exception as e:
        print_error(f"Error durante rollback: {e}")
        return False

def restore_from_backup():
    """Restaura la configuración desde un backup"""
    show_banner()
    print_info("🔄 Restaurar desde backup")
    print()
    
    backups = list_backups()
    
    if not backups:
        print_info("No hay backups disponibles.")
        return False
    
    show_backups()
    
    print()
    selection = input(f"{Colors.BOLD}Selecciona el número del backup a restaurar (0 para cancelar): {Colors.ENDC}").strip()
    
    if selection == '0' or not selection:
        print_warning("Restauración cancelada.")
        return False
    
    try:
        idx = int(selection) - 1
        if idx < 0 or idx >= len(backups):
            print_error("Número de backup inválido.")
            return False
        
        selected_backup = backups[idx]
        
        with open(selected_backup, 'r', encoding='utf-8') as f:
            content = json.load(f)
        
        print()
        print_info("📋 Contenido del backup:")
        for var, value in content['variables'].items():
            if var == 'ANTHROPIC_FOUNDRY_API_KEY':
                masked_value = value[:8] + "..." if len(value) > 8 else "***"
                print_info(f"  • {var} = {masked_value}")
            else:
                print_info(f"  • {var} = {value}")
        
        print()
        confirm = input(f"{Colors.BOLD}¿Confirmas la restauración? (S/N): {Colors.ENDC}").strip().upper()
        
        if confirm != 'S':
            print_warning("Restauración cancelada.")
            return False
        
        # Guardar backup del estado actual
        print()
        print_info("💾 Guardando backup del estado actual...")
        save_backup("Antes de restaurar backup")
        
        # Limpiar variables actuales
        print()
        print_info("🗑️  Limpiando configuración actual...")
        for var in CLAUDE_CODE_VARS:
            set_env_var(var, None)
        
        # Restaurar variables del backup
        print_info("🔄 Restaurando configuración...")
        for var, value in content['variables'].items():
            set_env_var(var, value)
            print_success(f"{var} restaurada")
        
        print()
        print_colored("━" * 60, Colors.OKGREEN)
        print_success("Backup restaurado exitosamente!")
        print_colored("━" * 60, Colors.OKGREEN)
        print()
        print_warning("IMPORTANTE: Reinicia tu terminal o sesión para que los cambios surtan efecto.")
        print()
        
        return True
    except (ValueError, IndexError):
        print_error("Número de backup inválido.")
        return False
    except Exception as e:
        print_error(f"Error al restaurar backup: {e}")
        return False

def show_menu():
    """Muestra el menú principal"""
    print()
    print_colored("Selecciona una opción:", Colors.BOLD)
    print()
    print_info("  [1] Configurar Azure Foundry")
    print_info("  [2] Rollback a configuración default (Anthropic)")
    print_info("  [3] Ver configuración actual")
    print_info("  [4] Listar backups disponibles")
    print_info("  [5] Restaurar desde backup")
    print_info("  [6] Guardar backup manual")
    print_info("  [0] Salir")
    print()
    return input(f"{Colors.BOLD}Opción: {Colors.ENDC}").strip()

def main():
    """Función principal"""
    # Verificar privilegios de administrador
    if not is_admin():
        print_error("Este script requiere privilegios de administrador.")
        print_warning("Por favor, ejecuta como Administrador (clic derecho > Ejecutar como administrador)")
        print()
        input("Presiona Enter para salir...")
        sys.exit(1)
    
    # Loop principal del menú
    while True:
        show_banner()
        show_current_config()
        option = show_menu()
        
        if option == '1':
            configure_azure_foundry()
            input("\nPresiona Enter para continuar...")
        elif option == '2':
            rollback_to_default()
            input("\nPresiona Enter para continuar...")
        elif option == '3':
            show_banner()
            show_current_config()
            input("\nPresiona Enter para continuar...")
        elif option == '4':
            show_banner()
            show_backups()
            input("\nPresiona Enter para continuar...")
        elif option == '5':
            restore_from_backup()
            input("\nPresiona Enter para continuar...")
        elif option == '6':
            show_banner()
            print_info("💾 Guardando backup manual...")
            description = input(f"{Colors.BOLD}Descripción del backup: {Colors.ENDC}").strip()
            if not description:
                description = "Backup manual"
            save_backup(description)
            input("\nPresiona Enter para continuar...")
        elif option == '0':
            print_success("\n👋 ¡Hasta luego!")
            print()
            break
        else:
            print_error("\nOpción inválida. Por favor, selecciona una opción válida.")
            import time
            time.sleep(2)

if __name__ == "__main__":
    main()
