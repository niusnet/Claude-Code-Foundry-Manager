#!/usr/bin/env python3
"""
Claude Code - Azure Foundry Quick Setup
Script simplificado para configuración rápida desde línea de comandos
"""

import sys
import argparse
import winreg
from pathlib import Path

def is_admin():
    """Verifica si el script se está ejecutando con privilegios de administrador"""
    try:
        import ctypes
        return ctypes.windll.shell32.IsUserAnAdmin()
    except:
        return False

def set_env_var(name: str, value):
    """Establece o elimina una variable de entorno del sistema"""
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
        
        # Notificar al sistema
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
        print(f"❌ Error: {e}")
        raise

def configure(args):
    """Configura Azure Foundry"""
    print("\n🔧 Configurando Azure Foundry para Claude Code...")
    print()
    
    try:
        set_env_var('CLAUDE_CODE_USE_FOUNDRY', '1')
        print("✓ CLAUDE_CODE_USE_FOUNDRY habilitado")
        
        set_env_var('ANTHROPIC_FOUNDRY_RESOURCE', args.resource)
        print(f"✓ ANTHROPIC_FOUNDRY_RESOURCE = {args.resource}")
        
        if args.api_key:
            set_env_var('ANTHROPIC_FOUNDRY_API_KEY', args.api_key)
            print("✓ ANTHROPIC_FOUNDRY_API_KEY configurada")
        else:
            print("ℹ️  Usando autenticación Entra ID")
        
        set_env_var('ANTHROPIC_DEFAULT_SONNET_MODEL', args.sonnet_model)
        print(f"✓ ANTHROPIC_DEFAULT_SONNET_MODEL = {args.sonnet_model}")
        
        set_env_var('ANTHROPIC_DEFAULT_HAIKU_MODEL', args.haiku_model)
        print(f"✓ ANTHROPIC_DEFAULT_HAIKU_MODEL = {args.haiku_model}")
        
        set_env_var('ANTHROPIC_DEFAULT_OPUS_MODEL', args.opus_model)
        print(f"✓ ANTHROPIC_DEFAULT_OPUS_MODEL = {args.opus_model}")
        
        print()
        print("━" * 60)
        print("✅ Configuración aplicada exitosamente!")
        print("━" * 60)
        print()
        print("⚠️  IMPORTANTE: Reinicia tu terminal para que los cambios surtan efecto")
        print()
        
    except Exception as e:
        print(f"\n❌ Error al configurar: {e}")
        sys.exit(1)

def rollback():
    """Hace rollback a configuración default"""
    print("\n🔄 Haciendo rollback a configuración default...")
    print()
    
    vars_to_remove = [
        'CLAUDE_CODE_USE_FOUNDRY',
        'ANTHROPIC_FOUNDRY_RESOURCE',
        'ANTHROPIC_FOUNDRY_BASE_URL',
        'ANTHROPIC_FOUNDRY_API_KEY',
        'ANTHROPIC_DEFAULT_SONNET_MODEL',
        'ANTHROPIC_DEFAULT_HAIKU_MODEL',
        'ANTHROPIC_DEFAULT_OPUS_MODEL'
    ]
    
    try:
        removed = 0
        for var in vars_to_remove:
            try:
                set_env_var(var, None)
                print(f"✓ {var} eliminada")
                removed += 1
            except:
                pass
        
        if removed == 0:
            print("ℹ️  No había configuración de Azure Foundry activa")
        else:
            print()
            print("━" * 60)
            print("✅ Rollback completado exitosamente!")
            print("━" * 60)
            print()
            print("ℹ️  Claude Code ahora usa el provider default de Anthropic")
            print("⚠️  IMPORTANTE: Reinicia tu terminal para que los cambios surtan efecto")
        print()
        
    except Exception as e:
        print(f"\n❌ Error durante rollback: {e}")
        sys.exit(1)

def main():
    if not is_admin():
        print("\n❌ Este script requiere privilegios de administrador")
        print("   Ejecuta como Administrador (clic derecho > Ejecutar como administrador)")
        input("\nPresiona Enter para salir...")
        sys.exit(1)
    
    parser = argparse.ArgumentParser(
        description='Claude Code - Azure Foundry Quick Setup',
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Ejemplos:
  # Configurar con Entra ID
  python claude_foundry_quick.py --resource mi-recurso-foundry
  
  # Configurar con API Key
  python claude_foundry_quick.py --resource mi-recurso --api-key tu-key
  
  # Configurar con modelos personalizados
  python claude_foundry_quick.py --resource mi-recurso --sonnet-model my-deployment
  
  # Hacer rollback
  python claude_foundry_quick.py --rollback
        """
    )
    
    parser.add_argument('--resource', help='Nombre del recurso de Azure Foundry')
    parser.add_argument('--api-key', help='API Key de Azure (opcional, usa Entra ID si no se proporciona)')
    parser.add_argument('--sonnet-model', default='claude-sonnet-4-5', help='Deployment de Sonnet (default: claude-sonnet-4-5)')
    parser.add_argument('--haiku-model', default='claude-haiku-4-5', help='Deployment de Haiku (default: claude-haiku-4-5)')
    parser.add_argument('--opus-model', default='claude-opus-4-5', help='Deployment de Opus (default: claude-opus-4-5)')
    parser.add_argument('--rollback', action='store_true', help='Hacer rollback a configuración default')
    
    args = parser.parse_args()
    
    print()
    print("╔════════════════════════════════════════════════════════════╗")
    print("║     Claude Code - Azure Foundry Quick Setup               ║")
    print("╚════════════════════════════════════════════════════════════╝")
    
    if args.rollback:
        rollback()
    elif args.resource:
        configure(args)
    else:
        parser.print_help()
        print("\n❌ Error: Debes especificar --resource o --rollback")
        sys.exit(1)

if __name__ == "__main__":
    main()
