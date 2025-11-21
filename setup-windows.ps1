# Claude Foundry Manager - Windows Setup Script
# This script checks for Go installation and helps set up the project

Write-Host "============================================" -ForegroundColor Cyan
Write-Host "Claude Foundry Manager - Setup Script" -ForegroundColor Cyan
Write-Host "============================================" -ForegroundColor Cyan
Write-Host ""

# Check if Go is installed
function Test-GoInstalled {
    try {
        $goVersion = go version 2>$null
        return $true
    } catch {
        return $false
    }
}

# Main setup logic
if (Test-GoInstalled) {
    Write-Host "[✓] Go is already installed" -ForegroundColor Green
    go version
    Write-Host ""

    Write-Host "Would you like to build the project now? (Y/N)" -ForegroundColor Yellow
    $response = Read-Host

    if ($response -eq 'Y' -or $response -eq 'y') {
        Write-Host ""
        Write-Host "[1/3] Downloading dependencies..." -ForegroundColor Cyan
        go mod download

        if ($LASTEXITCODE -eq 0) {
            Write-Host "[✓] Dependencies downloaded" -ForegroundColor Green
            Write-Host ""

            Write-Host "[2/3] Building executable..." -ForegroundColor Cyan
            go build -ldflags="-s -w" -o claude-foundry-manager.exe .

            if ($LASTEXITCODE -eq 0) {
                Write-Host "[✓] Build successful!" -ForegroundColor Green
                Write-Host ""

                Write-Host "[3/3] Checking file size..." -ForegroundColor Cyan
                $fileSize = (Get-Item claude-foundry-manager.exe).Length
                $fileSizeMB = [math]::Round($fileSize / 1MB, 2)
                Write-Host "Executable size: $fileSizeMB MB" -ForegroundColor White
                Write-Host ""

                Write-Host "============================================" -ForegroundColor Green
                Write-Host "SETUP COMPLETE!" -ForegroundColor Green
                Write-Host "============================================" -ForegroundColor Green
                Write-Host ""
                Write-Host "To run the tool:" -ForegroundColor Yellow
                Write-Host "  .\claude-foundry-manager.exe" -ForegroundColor White
                Write-Host ""
                Write-Host "For help:" -ForegroundColor Yellow
                Write-Host "  .\claude-foundry-manager.exe --help" -ForegroundColor White
                Write-Host ""
                Write-Host "Note: Run as Administrator to modify system environment variables" -ForegroundColor Cyan
            } else {
                Write-Host "[✗] Build failed" -ForegroundColor Red
                exit 1
            }
        } else {
            Write-Host "[✗] Failed to download dependencies" -ForegroundColor Red
            exit 1
        }
    } else {
        Write-Host ""
        Write-Host "Setup cancelled. You can build manually later with:" -ForegroundColor Yellow
        Write-Host "  go build -o claude-foundry-manager.exe ." -ForegroundColor White
    }
} else {
    Write-Host "[✗] Go is not installed on this system" -ForegroundColor Red
    Write-Host ""
    Write-Host "To use this project, you have two options:" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Option A: Install Go and build locally" -ForegroundColor Cyan
    Write-Host "  1. Download Go from: https://go.dev/dl/" -ForegroundColor White
    Write-Host "  2. Install the .msi file (follow the wizard)" -ForegroundColor White
    Write-Host "  3. Restart this PowerShell terminal" -ForegroundColor White
    Write-Host "  4. Run this script again" -ForegroundColor White
    Write-Host ""
    Write-Host "Option B: Use GitHub Actions (no local Go needed)" -ForegroundColor Cyan
    Write-Host "  1. Push this code to GitHub" -ForegroundColor White
    Write-Host "  2. Create a release tag: git tag v1.0.0 && git push origin v1.0.0" -ForegroundColor White
    Write-Host "  3. GitHub will automatically build binaries for you" -ForegroundColor White
    Write-Host "  4. Download the .exe from the Releases page" -ForegroundColor White
    Write-Host ""
    Write-Host "For detailed instructions, see: INSTALL.md" -ForegroundColor Cyan
    Write-Host ""

    Write-Host "Would you like to open the Go download page in your browser? (Y/N)" -ForegroundColor Yellow
    $response = Read-Host

    if ($response -eq 'Y' -or $response -eq 'y') {
        Start-Process "https://go.dev/dl/"
        Write-Host ""
        Write-Host "Opening browser... After installing Go, restart PowerShell and run this script again." -ForegroundColor Green
    }
}

Write-Host ""
Write-Host "Press any key to exit..."
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
