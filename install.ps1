# Get latest release
$response = Invoke-RestMethod -Uri "https://api.github.com/repos/aicraftlab-dev/aicraft-cli/releases/latest"
$asset = $response.assets | Where-Object { $_.name -match "windows" } | Select-Object -First 1
$downloadUrl = $asset.browser_download_url

# Download the binary
Write-Host "Downloading AICraft CLI..."
$output = "$env:TEMP\aicraft-cli.exe"
Invoke-WebRequest -Uri $downloadUrl -OutFile $output

# Install to system path
Write-Host "Installing AICraft CLI..."
$installPath = "$env:ProgramFiles\AICraft\aicraft.exe"
New-Item -ItemType Directory -Path "$env:ProgramFiles\AICraft" -Force | Out-Null
Move-Item -Path $output -Destination $installPath -Force

# Add to PATH
$path = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
if ($path -notlike "*AICraft*") {
    [Environment]::SetEnvironmentVariable("Path", "$path;$env:ProgramFiles\AICraft", [EnvironmentVariableTarget]::Machine)
}

# Verify installation
Write-Host "Verifying installation..."
if (Get-Command aicraft -ErrorAction SilentlyContinue) {
    Write-Host "AICraft CLI installed successfully!"
    Write-Host "Version: $(aicraft --version)"
} else {
    Write-Host "Installation failed. Please check permissions."
    exit 1
}