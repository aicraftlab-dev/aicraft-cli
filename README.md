# AICraft CLI

[![codecov](https://codecov.io/gh/aicraftlab-dev/aicraft-cli/branch/main/graph/badge.svg?token=YOUR_CODECOV_TOKEN)](https://codecov.io/gh/aicraftlab-dev/aicraft-cli)

AICraft CLI is a powerful command-line interface for AI-augmented workflows, providing unified access to AI models, code execution, and web intelligence.

## Installation

### Quick Install

#### Linux/macOS
```bash
curl -sSL https://raw.githubusercontent.com/aicraftlab-dev/aicraft-cli/main/install.sh | bash
```

#### Windows
Run in PowerShell as Administrator:
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force
iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/aicraftlab-dev/aicraft-cli/main/install.ps1'))
```

### Manual Installation

### From Releases
1. Download the latest release for your platform from the [Releases page](https://github.com/aicraftlab-dev/aicraft-cli/releases)
2. Make the binary executable:
   ```bash
   chmod +x aicraft-cli-*
   ```
3. Move it to your PATH:
   ```bash
   sudo mv aicraft-cli-* /usr/local/bin/aicraft
   ```

### From Source
1. Clone the repository:
   ```bash
   git clone https://github.com/aicraftlab-dev/aicraft-cli.git
   cd aicraft-cli
   ```
2. Build the project:
   ```bash
   go build -o aicraft
   ```
3. Install:
   ```bash
   sudo mv aicraft /usr/local/bin/
   ```

## Basic Usage

### AI Operations
Generate text using AI models:
```bash
aicraft ai generate --model ollama:codellama "Write a Go HTTP server"
```

### Code Execution
Run code in a sandboxed environment:
```bash
aicraft code run --lang python script.py --sandbox docker
```

### Web Intelligence
Search the web:
```bash
aicraft web search "Python async best practices" --engine google
```

## Configuration

### Managing AI Providers
Set up AI providers using the config command:
```bash
aicraft config setup --provider deepseek --api-key YOUR_API_KEY
```

### Configuration File
The configuration is stored in `~/.aicraft/config.yaml`:
```yaml
ai:
  providers:
    deepseek:
      api_key: YOUR_API_KEY
      model: deepseek-chat
  default_model: ollama:llama2
sandbox:
  timeout: 10s
search:
  default_engine: duckduckgo
```

### Environment Variables
You can override configuration using environment variables:
- `AICRAFT_AI_DEFAULT_MODEL`
- `AICRAFT_SANDBOX_TIMEOUT`
- `AICRAFT_SEARCH_DEFAULT_ENGINE`
Create a config file at `~/.aicraft/config.yaml`:
```yaml
ai:
  default_model: ollama:llama2
sandbox:
  timeout: 10s
search:
  default_engine: duckduckgo
```

## Contributing
Contributions are welcome! Please see our [Contributing Guide](CONTRIBUTING.md).

## License
MIT License - See [LICENSE](LICENSE) for details.