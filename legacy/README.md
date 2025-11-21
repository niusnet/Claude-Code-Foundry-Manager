# Legacy Python Implementation

This directory contains the original **Windows-only** Python implementation of the Claude Code Foundry Manager.

## ⚠️ Deprecated

This implementation has been **replaced** by the cross-platform Go version in the root directory. It is maintained here for reference purposes only.

## What's Here

| File | Description |
|------|-------------|
| `claude_foundry_manager.py` | Interactive menu-based manager (498 lines) |
| `claude_foundry_quick.py` | CLI quick configuration tool (179 lines) |
| `run-manager.bat` | Windows launcher with auto-elevation |
| `START-HERE.md` | Python version getting started guide |
| `README-Python.md` | Complete Python documentation |
| `QUICK-START.md` | Python quick start guide |

## Limitations (Python Version)

- ❌ **Windows only** - Uses `winreg` module for registry access
- ❌ **Requires Python 3.7+** - Not a standalone binary
- ❌ **Admin privileges required** - Modifies system registry
- ❌ **No cross-platform support** - Can't run on Linux/macOS

## Why We Moved to Go

The Go implementation provides:

- ✅ **Cross-platform**: Windows, Linux, macOS
- ✅ **Single binary**: No Python installation required
- ✅ **Fast**: 10x faster startup time
- ✅ **Easy distribution**: Download and run
- ✅ **Modern tooling**: Automated testing and CI/CD

## Migration

If you're using the Python version, switch to the Go version:

1. **Download** the appropriate binary from [Releases](https://github.com/niusnet/Claude-Code-Foundry-Manager/releases/latest)
2. **Run** it directly (no installation needed)
3. **Existing backups** are compatible - both versions use the same JSON format

## For Reference Only

This code is kept for:
- Historical reference
- Understanding the original implementation
- Comparing approaches between Python and Go

**For new usage, please use the Go version from the root directory.**

---

**Go version documentation:** [../README.md](../README.md)
