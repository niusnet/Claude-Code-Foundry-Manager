# Claude Code - Azure Foundry Configuration Manager

A **cross-platform CLI tool** to easily switch [Claude Code](https://github.com/anthropics/claude-code) between Azure AI Foundry and direct Anthropic API configurations.

![Platform Support](https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-blue)
![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)

---

## ✨ Features

- 🌐 **Cross-platform**: Works on Windows, Linux, and macOS
- 📦 **Single binary**: No dependencies, no runtime required
- 🔄 **Easy switching**: Toggle between Azure Foundry and Anthropic direct
- 💾 **Automatic backups**: Safe configuration changes with rollback support
- 🎨 **Two modes**: Interactive menu or quick CLI commands
- 🔐 **Secure**: Masks API keys in output, creates timestamped backups

---

## 📥 Installation

### Option 1: Download Pre-built Binary (Recommended)

Download the latest binary for your platform from the [Releases](https://github.com/gilbe/claude-foundry-manager/releases) page:

**Windows:**
```powershell
# Download claude-foundry-manager-windows-amd64.exe
# Run directly - no installation needed!
```

**Linux:**
```bash
# Download and make executable
curl -L -o claude-foundry-manager https://github.com/gilbe/claude-foundry-manager/releases/latest/download/claude-foundry-manager-linux-amd64
chmod +x claude-foundry-manager

# Optional: Move to PATH
sudo mv claude-foundry-manager /usr/local/bin/
```

**macOS:**
```bash
# For Intel Macs
curl -L -o claude-foundry-manager https://github.com/gilbe/claude-foundry-manager/releases/latest/download/claude-foundry-manager-darwin-amd64

# For Apple Silicon (M1/M2/M3)
curl -L -o claude-foundry-manager https://github.com/gilbe/claude-foundry-manager/releases/latest/download/claude-foundry-manager-darwin-arm64

chmod +x claude-foundry-manager
sudo mv claude-foundry-manager /usr/local/bin/
```

### Option 2: Build from Source

```bash
# Requires Go 1.21+
git clone https://github.com/gilbe/claude-foundry-manager.git
cd claude-foundry-manager
go build -o claude-foundry-manager .
```

---

## 🚀 Quick Start

### Interactive Mode

Simply run the tool without arguments to open the interactive menu:

```bash
claude-foundry-manager
```

You'll see a menu with options:
1. Configure Azure Foundry
2. Rollback to Default (Direct Anthropic)
3. View Current Configuration
4. List Available Backups
5. Restore from Backup
6. Save Manual Backup
7. Exit

### CLI Mode

For automation or quick configuration:

```bash
# Configure Azure Foundry
claude-foundry-manager configure \
  --resource=my-foundry \
  --api-key=sk-ant-xxx

# Configure with Entra ID (no API key)
claude-foundry-manager configure --resource=my-foundry

# Custom model deployments
claude-foundry-manager configure \
  --resource=my-foundry \
  --sonnet-model=claude-4-5 \
  --haiku-model=claude-haiku

# View current configuration
claude-foundry-manager show

# Rollback to default Anthropic
claude-foundry-manager rollback

# Backup management
claude-foundry-manager backup list
claude-foundry-manager backup create "My backup description"
claude-foundry-manager backup restore backup_20240115_143022.json
```

---

## 📖 Usage Examples

### Scenario 1: First-time Azure Foundry Setup

```bash
# Run interactive mode
claude-foundry-manager

# Select option [1] Configure Azure Foundry
# Enter your resource name: my-foundry-resource
# Enter API Key: sk-ant-api-xxxxx (or leave empty for Entra ID)
# Accept defaults for model names or customize

# Restart your terminal
# Test with: claude-code
```

### Scenario 2: Switch Between Configurations

```bash
# Save current config before switching
claude-foundry-manager backup create "Work Azure config"

# Switch to default Anthropic
claude-foundry-manager rollback

# Later, restore Azure config
claude-foundry-manager backup restore backup_20240115_143022.json
```

### Scenario 3: Automation in Scripts

```bash
#!/bin/bash
# setup-dev-environment.sh

# Configure for Azure Foundry
claude-foundry-manager configure \
  --resource=$AZURE_FOUNDRY_RESOURCE \
  --api-key=$AZURE_API_KEY \
  --sonnet-model=claude-sonnet-4-5

echo "✓ Claude Code configured for Azure Foundry"
```

---

## 🔧 How It Works

### Environment Variables

The tool manages these environment variables:

| Variable | Description |
|----------|-------------|
| `CLAUDE_CODE_USE_FOUNDRY` | Enable/disable Azure Foundry (`true`/`false`) |
| `ANTHROPIC_FOUNDRY_RESOURCE` | Azure Foundry resource name |
| `ANTHROPIC_FOUNDRY_BASE_URL` | Base URL (auto-generated from resource) |
| `ANTHROPIC_FOUNDRY_API_KEY` | API key (optional, uses Entra ID if not set) |
| `ANTHROPIC_DEFAULT_SONNET_MODEL` | Sonnet model deployment name |
| `ANTHROPIC_DEFAULT_HAIKU_MODEL` | Haiku model deployment name |
| `ANTHROPIC_DEFAULT_OPUS_MODEL` | Opus model deployment name |

### Platform-Specific Behavior

**Windows:**
- Modifies system environment variables in the registry (`HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`)
- Requires administrator privileges
- Broadcasts `WM_SETTINGCHANGE` to notify all processes

**Linux/macOS:**
- Modifies shell profile files (`.bashrc`, `.zshrc`, `.bash_profile`, or `.profile`)
- Adds a managed block between markers
- No admin privileges required (user-level configuration)

### Backup System

- **Location**: `~/.claude-code-backups/` (cross-platform)
- **Format**: JSON with timestamp, description, and variables
- **Auto-backup**: Created before every configuration change
- **Manual backup**: Create snapshots anytime

Example backup file:
```json
{
  "timestamp": "2024-01-15T14:30:22Z",
  "description": "Before configuring Azure Foundry",
  "variables": {
    "CLAUDE_CODE_USE_FOUNDRY": "true",
    "ANTHROPIC_FOUNDRY_RESOURCE": "my-foundry",
    ...
  }
}
```

---

## 🛠️ Building

### Local Development

```bash
# Clone repository
git clone https://github.com/gilbe/claude-foundry-manager.git
cd claude-foundry-manager

# Install dependencies
go mod download

# Run locally
go run .

# Build
go build -o claude-foundry-manager .

# Run tests (when available)
go test ./...
```

### Cross-Compilation

```bash
# Windows (from any platform)
GOOS=windows GOARCH=amd64 go build -o claude-foundry-manager-windows-amd64.exe .

# Linux
GOOS=linux GOARCH=amd64 go build -o claude-foundry-manager-linux-amd64 .

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o claude-foundry-manager-darwin-amd64 .

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o claude-foundry-manager-darwin-arm64 .
```

---

## 🔐 Security Considerations

### API Key Handling

- ✅ API keys are **masked** in terminal output (only first 8 characters shown)
- ⚠️ Backup files contain **plain-text** API keys
- 🔒 Ensure `~/.claude-code-backups/` has restricted permissions:
  ```bash
  # Linux/macOS
  chmod 700 ~/.claude-code-backups
  ```

### Administrator Privileges

- **Windows**: Requires admin to modify system registry
- **Linux/macOS**: No admin required (modifies user profile files)

### Best Practices

1. Use Entra ID authentication when possible (no API key storage)
2. Regularly review and clean old backups
3. Don't commit backup files to version control
4. Use environment-specific configurations for different contexts

---

## 🆚 Comparison: Python vs Go Version

| Feature | Python Version | Go Version |
|---------|----------------|------------|
| **Platform Support** | Windows only | Windows, Linux, macOS |
| **Dependencies** | Python 3.7+ runtime | None (single binary) |
| **Installation** | Script execution | Download & run |
| **Binary Size** | N/A | ~5-10 MB |
| **Startup Time** | ~500ms | ~50ms (10x faster) |
| **Distribution** | Requires Python | Single executable |
| **Admin Required** | Yes (Windows) | Yes (Windows), No (Unix) |

**Migration Note:** The Go version maintains 100% functional compatibility with the Python version. All backups created with the Python version are compatible with the Go version.

---

## 🐛 Troubleshooting

### Issue: "Access denied" on Windows

**Solution:** Run as administrator. Right-click the executable and select "Run as administrator".

### Issue: Command not found on Linux/macOS

**Solution:** Ensure the binary is in your PATH or use the full path:
```bash
./claude-foundry-manager
```

### Issue: Changes not taking effect

**Solution:** Restart your terminal/shell after configuration changes:
```bash
# Or source your profile manually
source ~/.bashrc  # or ~/.zshrc
```

### Issue: "Permission denied" when running binary

**Solution:** Make the binary executable:
```bash
chmod +x claude-foundry-manager
```

### Issue: Binary blocked on macOS

**Solution:** Allow the binary in System Preferences:
```bash
# Or remove quarantine attribute
xattr -d com.apple.quarantine claude-foundry-manager
```

---

## 📚 Additional Resources

- [Claude Code Official Repo](https://github.com/anthropics/claude-code)
- [Azure AI Foundry Documentation](https://learn.microsoft.com/en-us/azure/ai-studio/)
- **Legacy Python Documentation:**
  - [START-HERE.md](START-HERE.md) - Python version guide
  - [README-Python.md](README-Python.md) - Detailed Python docs
  - [QUICK-START.md](QUICK-START.md) - Python quick start

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

### Development Setup

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes and test on multiple platforms
4. Commit with descriptive messages
5. Push and create a pull request

---

## 📝 License

MIT License - See [LICENSE](LICENSE) file for details

---

## 🙏 Acknowledgments

- Built for the [Claude Code](https://github.com/anthropics/claude-code) community
- Powered by [Cobra](https://github.com/spf13/cobra) CLI framework
- Inspired by the original Python implementation

---

**Made with ❤️ for Claude Code users**

*Need help? [Open an issue](https://github.com/gilbe/claude-foundry-manager/issues)*
