# Claude Code - Azure Foundry Manager

Cross-platform CLI tool to switch [Claude Code](https://github.com/anthropics/claude-code) between Azure AI Foundry and direct Anthropic API.

![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-blue)
![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)

---

## Features

- 🌐 **Cross-platform**: Windows, Linux, macOS
- 📦 **Single binary**: No dependencies required
- 🔄 **Easy switching**: Toggle between Azure Foundry and Anthropic
- 💾 **Auto-backup**: Safe configuration changes with rollback
- 🎨 **Dual mode**: Interactive menu or CLI commands
- 🔐 **Secure**: API key masking, timestamped backups

---

## Quick Start

### Download

Get the latest binary from [Releases](https://github.com/niusnet/Claude-Code-Foundry-Manager/releases/latest):

**Windows:**
```powershell
# Download claude-foundry-manager-windows-amd64.exe
# Run as Administrator (right-click → Run as administrator)
```

**Linux/macOS:**
```bash
# Download appropriate binary and make executable
chmod +x claude-foundry-manager-*
sudo mv claude-foundry-manager-* /usr/local/bin/claude-foundry-manager
```

### Interactive Mode

```bash
# Run without arguments for interactive menu
claude-foundry-manager
```

### CLI Mode

```bash
# Configure Azure Foundry
claude-foundry-manager configure --resource=my-foundry --api-key=sk-xxx

# View current config
claude-foundry-manager show

# Rollback to default
claude-foundry-manager rollback

# Manage backups
claude-foundry-manager backup list
claude-foundry-manager backup restore <filename>
```

---

## Commands

| Command | Description |
|---------|-------------|
| `claude-foundry-manager` | Interactive menu (default) |
| `configure` | Set up Azure Foundry configuration |
| `rollback` | Restore default Anthropic configuration |
| `show` | Display current configuration |
| `backup list` | List all available backups |
| `backup create` | Create manual backup |
| `backup restore` | Restore from backup |

### Configure Options

```bash
claude-foundry-manager configure \
  --resource=<azure-resource>      # Required: Azure Foundry resource name
  --api-key=<key>                  # Optional: API key (uses Entra ID if omitted)
  --sonnet-model=<deployment>      # Optional: Sonnet deployment (default: claude-sonnet-4-5)
  --haiku-model=<deployment>       # Optional: Haiku deployment (default: claude-haiku-4-5)
  --opus-model=<deployment>        # Optional: Opus deployment (default: claude-opus-4-1)
```

---

## Environment Variables

Manages these Claude Code environment variables:

- `CLAUDE_CODE_USE_FOUNDRY` - Enable/disable Azure Foundry
- `ANTHROPIC_FOUNDRY_RESOURCE` - Azure resource name
- `ANTHROPIC_FOUNDRY_BASE_URL` - Base URL (auto-generated)
- `ANTHROPIC_FOUNDRY_API_KEY` - API key (optional)
- `ANTHROPIC_DEFAULT_SONNET_MODEL` - Sonnet deployment
- `ANTHROPIC_DEFAULT_HAIKU_MODEL` - Haiku deployment
- `ANTHROPIC_DEFAULT_OPUS_MODEL` - Opus deployment

---

## Platform Specifics

**Windows:**
- Modifies system registry (`HKLM\SYSTEM\...\Environment`)
- Requires Administrator privileges
- Broadcasts change notifications

**Linux/macOS:**
- Modifies shell profiles (`.bashrc`, `.zshrc`, `.bash_profile`, `.profile`)
- No admin required (user-level)
- Restarts shell for changes to apply

**Backups:**
- Location: `~/.claude-code-backups/`
- Format: JSON with timestamp
- Contains all environment variables

---

## Building from Source

```bash
# Clone repository
git clone https://github.com/niusnet/Claude-Code-Foundry-Manager.git
cd Claude-Code-Foundry-Manager

# Build
go build -ldflags="-s -w" -o claude-foundry-manager .

# Or use provided script (Windows)
.\build.bat
```

### Cross-Compile

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o claude-foundry-manager-linux-amd64

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o claude-foundry-manager-darwin-arm64
```

---

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package
go test ./internal/config
```

---

## Project Structure

```
claude-foundry-manager/
├── cmd/                    # CLI commands (Cobra)
│   ├── root.go            # Main command + interactive mode
│   ├── configure.go       # Configure command
│   ├── rollback.go        # Rollback command
│   ├── show.go            # Show command
│   └── backup.go          # Backup commands
├── internal/
│   ├── config/            # Environment variable management
│   │   ├── manager.go             # Common logic
│   │   ├── manager_windows.go    # Windows registry
│   │   └── manager_unix.go       # Unix shell profiles
│   ├── backup/            # Backup system
│   │   └── backup.go
│   └── ui/                # Interactive interface
│       └── interactive.go
├── legacy/                # Python implementation (reference)
├── docs/                  # Additional documentation
├── .github/workflows/     # CI/CD
└── main.go               # Entry point
```

---

## Documentation

- **[Installation Guide](docs/INSTALL.md)** - Detailed installation instructions
- **[Getting Started](docs/GET-STARTED.md)** - Step-by-step usage guide
- **[Legacy Python Version](legacy/)** - Original Windows-only implementation

---

## Troubleshooting

**"Access denied" (Windows)**
→ Run as Administrator

**"Command not found" (Linux/macOS)**
→ Ensure binary is in PATH or use `./claude-foundry-manager`

**Changes not taking effect**
→ Restart terminal/shell after configuration

**Permission denied**
→ Make binary executable: `chmod +x claude-foundry-manager`

---

## Contributing

1. Fork the repository
2. Create feature branch: `git checkout -b feature/my-feature`
3. Commit changes: `git commit -m "Add feature"`
4. Push to branch: `git push origin feature/my-feature`
5. Open Pull Request

---

## License

MIT License - See LICENSE file for details

---

## Acknowledgments

- Built for [Claude Code](https://github.com/anthropics/claude-code) community
- Powered by [Cobra](https://github.com/spf13/cobra) CLI framework

---

**Need help?** [Open an issue](https://github.com/niusnet/Claude-Code-Foundry-Manager/issues)
