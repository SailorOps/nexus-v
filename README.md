# NEXUS-V

**A modern, zero-install VS Code extension scaffolder — built in Go, outputs TypeScript.**

NEXUS-V replaces the legacy Yeoman `yo code` generator with a single static binary that produces clean, dependency-light VS Code extension projects. No global installs, no ecosystem rot, no hidden state — just a portable tool that does one thing well.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Remote Plugins](#remote-templates-plugins)
- [Template Variants](#template-variants)
- [Configuration](#configuration)
- [Hooks](#hooks)
- [Environment Health](#check-environment)
- [Telemetry](#telemetry)
- [Roadmap](#next-steps--roadmap)
- [License](#license)

---

## Overview

The VS Code extension ecosystem depends heavily on Yeoman — a scaffolding tool with deep Node.js dependency trees, global installs, and a maintenance surface area that drifts over time. NEXUS-V takes a different approach:

- **Written in Go** — compiles to a single static binary with zero runtime dependencies.
- **Outputs TypeScript** — every scaffolded project is a clean, modern TypeScript VS Code extension ready for development.
- **Interactive TUI** — uses **Bubble Tea** for a rich, visual template selection experience.
- **Portable by design** — no PATH pollution, no global packages, no persistent state.

---

## Features

| Feature | Description |
|---|---|
| **Single binary** | One executable, zero runtime dependencies — works on Windows, macOS, and Linux |
| **Interactive TUI** | Searchable, stylized menu for choosing template variants |
| **Remote Plugins** | Scaffold directly from any GitHub repository (`--template-dir <URL>`) |
| **Doctor Command** | Diagnostic tool to verify your local environment (`node`, `npm`, `vsce`) |
| **GoReleaser Pipeline** | Automated builds and distribution via Homebrew, Scoop, and Winget |
| **Self-Update** | Built-in `update` command to fetch the latest version from GitHub |
| **Hook system** | Pre- and post-scaffold hooks for custom automation (`--install`, `--git`) |
| **Opt-in telemetry** | Anonymous, minimal usage telemetry — off by default |

---

## Installation

### Homebrew (macOS / Linux)
```bash
brew tap billy-kidd-dev/nexusv
brew install nexus-v
```

### Scoop (Windows)
```powershell
scoop bucket add nexusv https://github.com/billy-kidd-dev/scoop-bucket
scoop install nexus-v
```

### Winget (Windows)
```powershell
winget install BillyKidd.NexusV
```

### One-liner (Unix)
```bash
curl -fsSL https://raw.githubusercontent.com/billy-kidd-dev/nexus-v/main/install.sh | bash
```

---

## Usage

### Interactive Mode (TUI)

```bash
nexus-v init  # or: nexus-v i
```

NEXUS-V will launch a beautiful interactive menu for selecting your extension type.

### Remote Templates (Plugins)

NEXUS-V supports scaffolding directly from remote Git repositories. This allows you to use community-created templates as plugins:

```bash
nexus-v init --template-dir https://github.com/user/my-custom-template
```

### Check Environment

Ensure you have all the necessary tools installed for VS Code development:

```bash
nexus-v doctor  # or: nexus-v dr
```

### Update NEXUS-V

```bash
nexus-v update  # or: nexus-v u
```

---

## Template Variants

| Variant | Description |
|---|---|
| `command` | Basic extension with a registered command and activation event |
| `webview` | Extension with a webview panel boilerplate |
| `language` | Language support with syntax highlighting and config |
| `theme` | Color theme extension with a base theme JSON |

---

## Configuration

Place a `.nexusvrc.yaml` in your home directory or project root:

```yaml
defaults:
  publisher: "my-org"
  variant: "command"
  git: true
  license: "MIT"

hooks:
  post_scaffold:
    - "npm install"
    - "code ."
```

---

## Next Steps & Roadmap

### Completed ✅
- [x] **TUI Mode** — Rich terminal UI for variant selection
- [x] **Plugin System** — Remote Git template support
- [x] **`nexus-v doctor`** — Environment diagnostic tool
- [x] **Multi-Channel Distribution** — Homebrew, Scoop, and Winget
- [x] **Self-Update** — Built-in `update` command
- [x] **CI/CD Pipeline** — GoReleaser + GitHub Actions

### Planned additions
- [ ] **Monorepo variant** — Multi-extension monorepo support
- [ ] **Scaffold history** — `nexus-v history` / `nexus-v replay`
- [ ] **VS Code Meta Extension** — Native UI wrapper for NEXUS-V

---

## License

MIT © [Billy Kidd](https://github.com/billy-kidd-dev)