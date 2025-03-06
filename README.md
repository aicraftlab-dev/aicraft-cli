# AICraft CLI

[![codecov](https://codecov.io/gh/aicraftlab-dev/aicraft-cli/branch/main/graph/badge.svg?token=YOUR_CODECOV_TOKEN)](https://codecov.io/gh/aicraftlab-dev/aicraft-cli)

AICraft CLI is a powerful command-line interface for AI-augmented workflows, providing unified access to AI models, code execution, and web intelligence.

## Installation

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