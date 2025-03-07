# Mock Invoke-RestMethod
function Invoke-RestMethod {
    return @{
        assets = @(
            @{
                name = "aicraft-cli-windows-latest-alpha"
                browser_download_url = "https://github.com/aicraftlab-dev/aicraft-cli/releases/download/v0.1.0/aicraft-cli-windows-latest-alpha"
            }
        )
    }
}

# Mock Invoke-WebRequest
function Invoke-WebRequest {
    param($Uri, $OutFile)
    New-Item -Path $OutFile -ItemType File -Force | Out-Null
}

# Mock environment variables
$env:TEMP = "C:\Temp"
$env:ProgramFiles = "C:\ProgramFiles"

# Run the install script
. ..\install.ps1

# Verify installation
if (-not (Get-Command aicraft -ErrorAction SilentlyContinue)) {
    Write-Error "Test failed: aicraft not found in PATH"
    exit 1
}

Write-Host "All tests passed!"