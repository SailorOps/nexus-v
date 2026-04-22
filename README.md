```markdown
# NEXUS-V

**A modern, zero-install VS Code extension scaffolder — built in Go, outputs TypeScript.**

NEXUS-V replaces the legacy Yeoman `yo code` generator with a single static binary that produces clean, dependency-light VS Code extension projects. No global installs, no ecosystem rot, no hidden state — just a portable tool that does one thing well.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Template Variants](#template-variants)
- [Configuration](#configuration)
- [Hooks](#hooks)
- [Telemetry](#telemetry)
- [Project Structure](#project-structure)
- [Next Steps & Roadmap](#next-steps--roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

The VS Code extension ecosystem depends heavily on Yeoman — a scaffolding tool with deep Node.js dependency trees, global installs, and a maintenance surface area that drifts over time. NEXUS-V takes a different approach:

- **Written in Go** — compiles to a single static binary with zero runtime dependencies.
- **Outputs TypeScript** — every scaffolded project is a clean, modern TypeScript VS Code extension ready for development.
- **Modular template engine** — variant-based template system that renders projects from embedded Go templates, not fragile file-copy heuristics.
- **Portable by design** — no PATH pollution, no global packages, no persistent state. Drop it in a directory and run it.

NEXUS-V is built for developers who want a scaffolding tool that stays out of the way and never breaks between uses.

---

## Features

| Feature | Description |
|---|---|
| **Single binary** | One executable, zero runtime dependencies — works on Windows, macOS, and Linux |
| **Embedded templates** | All project templates are compiled into the binary via Go's `embed` package |
| **Template variants** | Modular variant system for different extension types (commands, webviews, language support, etc.) |
| **Filesystem walker** | Intelligent `fswalk` module that traverses and renders template trees with context-aware filtering |
| **Template renderer** | Dedicated `render` module that processes Go `text/template` files with project-specific data |
| **Interactive prompts** | CLI prompts collect project metadata (name, publisher, description) at scaffold time |
| **Zero-install** | No `npm install -g`, no Yeoman, no generator packages — download and run |
| **Drift-free** | No lockfiles or dependency trees to rot between uses |
| **Hook system** | Pre- and post-scaffold hooks for custom automation |
| **Opt-in telemetry** | Anonymous, minimal usage telemetry — off by default |

---

## Installation

### Download the binary

Grab the latest release for your platform from the [Releases](https://github.com/geriatric-sailor/nexus-v/releases) page.

```bash
# Linux / macOS
curl -L https://github.com/geriatric-sailor/nexus-v/releases/latest/download/nexus-v-$(uname -s)-amd64 -o nexus-v
chmod +x nexus-v
```

```powershell
# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/geriatric-sailor/nexus-v/releases/latest/download/nexus-v-windows-amd64.exe" -OutFile "nexus-v.exe"
```

### Build from source

Requires Go 1.22+.

```bash
git clone https://github.com/geriatric-sailor/nexus-v.git
cd nexus-v
go build -o nexus-v ./cmd/nexus-v
```

### Verify installation

```bash
nexus-v --version
```

> **Note:** NEXUS-V is a standalone binary. No need to add it to your system PATH — run it from wherever you keep your tools.

---

## Usage

### Interactive mode (default)

```bash
nexus-v init
```

NEXUS-V walks you through a series of prompts:

```
? Extension name: my-extension
? Display name: My Extension
? Description: A helpful VS Code extension
? Publisher: your-publisher-id
? Template variant: command
? Initialize git repo? Yes

✔ Scaffolded my-extension using variant "command"
```

### Non-interactive mode

Pass all values as flags for scripting and CI use:

```bash
nexus-v init \
  --name my-extension \
  --display-name "My Extension" \
  --description "A helpful VS Code extension" \
  --publisher your-publisher-id \
  --variant command \
  --git
```

### Specify an output directory

```bash
nexus-v init --out ./projects/my-extension
```

### List available template variants

```bash
nexus-v variants
```

### Dry run (preview without writing files)

```bash
nexus-v init --dry-run
```

This prints the file tree that would be generated without writing anything to disk.

---

## Template Variants

NEXUS-V uses a modular variant system. Each variant is a self-contained template set embedded in the binary that produces a different type of VS Code extension.

| Variant | Description |
|---|---|
| `command` | Basic extension with a registered command and activation event |
| `webview` | Extension with a webview panel (HTML/CSS/JS inside VS Code) |
| `language` | Language support extension with syntax highlighting and language configuration |
| `snippet` | Snippet-only extension with a structured snippet file |
| `theme` | Color theme extension with a base theme JSON |
| `notebook` | Notebook renderer or controller extension |
| `test-suite` | Extension pre-configured with a full Mocha + `@vscode/test-electron` test harness |

### Variant selection

```bash
# Interactive — prompted automatically
nexus-v init

# Explicit
nexus-v init --variant webview
```

Each variant produces a complete, buildable project with:

- `package.json` with correct `activationEvents`, `contributes`, and `engines`
- `src/extension.ts` entry point
- `tsconfig.json` tuned for VS Code extension development
- `.vscode/launch.json` with Extension Host debug configuration
- Variant-specific boilerplate (webview HTML, grammar files, theme JSON, etc.)

---

## Configuration

NEXUS-V can be configured at three levels, in order of precedence:

### 1. CLI flags (highest precedence)

```bash
nexus-v init --publisher my-org --variant webview --no-git
```

### 2. Environment variables

```bash
export NEXUSV_PUBLISHER="my-org"
export NEXUSV_DEFAULT_VARIANT="command"
export NEXUSV_TELEMETRY="off"
```

### 3. Configuration file (lowest precedence)

Place a `.nexusvrc.yaml` in your home directory or project root:

```yaml
# ~/.nexusvrc.yaml

defaults:
  publisher: "my-org"
  variant: "command"
  git: true
  license: "MIT"

telemetry:
  enabled: false

hooks:
  post_scaffold:
    - "npm install"
    - "git add -A && git commit -m 'initial scaffold'"
```

### Configuration reference

| Key | Type | Default | Description |
|---|---|---|---|
| `defaults.publisher` | string | _(prompted)_ | Default publisher ID |
| `defaults.variant` | string | `command` | Default template variant |
| `defaults.git` | bool | `true` | Initialize a git repository |
| `defaults.license` | string | `MIT` | License to include in scaffolded project |
| `telemetry.enabled` | bool | `false` | Enable anonymous usage telemetry |
| `hooks.pre_scaffold` | string[] | `[]` | Commands to run before scaffold |
| `hooks.post_scaffold` | string[] | `[]` | Commands to run after scaffold |

---

## Hooks

NEXUS-V supports lifecycle hooks that run shell commands at defined points during scaffolding.

### Hook stages

| Stage | Trigger | Typical use |
|---|---|---|
| `pre_scaffold` | After prompts resolve, before any files are written | Validate environment, check prerequisites |
| `post_scaffold` | After all files are written to disk | Install dependencies, initialize git, open VS Code |

### Defining hooks

**In `.nexusvrc.yaml`:**

```yaml
hooks:
  pre_scaffold:
    - "echo 'Starting scaffold...'"
  post_scaffold:
    - "npm install"
    - "code ."
```

**Via CLI flags:**

```bash
nexus-v init --post-hook "npm install" --post-hook "code ."
```

### Hook behavior

- Hooks execute sequentially in the order defined.
- Each hook runs in a shell (`sh` on Unix, `cmd` on Windows) with the scaffolded project directory as the working directory.
- If a hook exits with a non-zero code, NEXUS-V prints the error and halts subsequent hooks. The scaffolded files remain on disk.
- Hooks are skipped entirely in `--dry-run` mode.

### Built-in hook shortcuts

For common post-scaffold actions, NEXUS-V provides convenience flags:

```bash
nexus-v init --install    # runs "npm install" after scaffold
nexus-v init --open       # runs "code ." after scaffold
nexus-v init --git        # runs "git init && git add -A" after scaffold
```

---

## Telemetry

NEXUS-V includes **optional, anonymous** usage telemetry to help guide development priorities. **Telemetry is off by default.**

### What is collected (when opted in)

- NEXUS-V version
- OS and architecture
- Template variant selected
- Scaffold success or failure (boolean)
- Timestamp

### What is never collected

- Project names, file contents, or paths
- Publisher IDs or personal identifiers
- IP addresses (requests are not logged server-side)
- Environment variables or system configuration

### Controlling telemetry

```bash
# Opt in
nexus-v config set telemetry.enabled true

# Opt out (or just never opt in — it's off by default)
nexus-v config set telemetry.enabled false
```

```bash
# Environment variable override (takes precedence)
export NEXUSV_TELEMETRY=off
```

```yaml
# .nexusvrc.yaml
telemetry:
  enabled: false
```

Telemetry respects the `DO_NOT_TRACK=1` environment variable convention.

---

## Project Structure

NEXUS-V's internal architecture is modular and intentionally minimal:

```
nexus-v/
├── cmd/
│   └── nexus-v/
│       └── main.go              # CLI entry point and flag parsing
├── internal/
│   ├── templates/
│   │   ├── templates.go         # Template registry and embedded FS
│   │   ├── command/             # "command" variant template files
│   │   ├── webview/             # "webview" variant template files
│   │   ├── language/            # "language" variant template files
│   │   └── ...                  # Additional variant directories
│   ├── fswalk/
│   │   └── fswalk.go            # Filesystem walker — traverses template trees
│   ├── render/
│   │   └── render.go            # Template renderer — processes and writes output
│   ├── config/
│   │   └── config.go            # Configuration loading and precedence resolution
│   └── hooks/
│       └── hooks.go             # Pre/post-scaffold hook execution
├── .nexusvrc.yaml               # Example configuration file
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

### Module responsibilities

| Module | File | Role |
|---|---|---|
| **templates** | `templates.go` | Registers variant template sets and embeds them into the binary via `//go:embed` |
| **fswalk** | `fswalk.go` | Walks the embedded filesystem tree for a selected variant, applying path-level filtering |
| **render** | `render.go` | Executes Go `text/template` rendering on each file with scaffold context (name, publisher, etc.) |
| **config** | `config.go` | Loads and merges configuration from CLI flags, env vars, and `.nexusvrc.yaml` |
| **hooks** | `hooks.go` | Runs shell commands at pre/post-scaffold lifecycle stages |

---

## Next Steps & Roadmap

### Near-term priorities

- [ ] **CI/CD pipeline** — GitHub Actions workflow for automated builds, tests, and multi-platform releases
- [ ] **Cross-compilation matrix** — Prebuilt binaries for `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`
- [ ] **`nexus-v update`** — Self-update command that pulls the latest binary from GitHub Releases
- [ ] **Integration tests** — End-to-end tests that scaffold each variant and verify the output compiles and passes `vsce package`

### Planned additions

- [ ] **Custom template support** — Load user-defined template directories alongside built-in variants (`nexus-v init --template ./my-templates`)
- [ ] **Workspace multi-root scaffolding** — Scaffold multiple related extensions into a VS Code workspace configuration
- [ ] **Plugin system** — Lightweight plugin interface for community-contributed variants distributed as standalone template packs
- [ ] **`nexus-v doctor`** — Diagnostic command that checks for Node.js, npm, and `vsce` availability and prints environment health
- [ ] **Monorepo variant** — Template variant for multi-extension monorepo projects with shared build configuration
- [ ] **Changelog and versioning automation** — Built-in `CHANGELOG.md` generation and semantic version bumping for scaffolded projects

### Stretch goals

- [ ] **TUI mode** — Rich terminal UI with selection menus, progress indicators, and tree previews (powered by Bubble Tea or similar)
- [ ] **Scaffold history** — Local log of past scaffolds for quick re-runs (`nexus-v history` / `nexus-v replay`)
- [ ] **VS Code extension for NEXUS-V** — Meta extension that runs NEXUS-V from the VS Code command palette

---

## Contributing

NEXUS-V is maintained under the [Geriatric Sailor](https://github.com/geriatric-sailor) organization.

1. Fork the repository
2. Create a feature branch (`git checkout -b feat/my-feature`)
3. Commit your changes (`git commit -m 'Add my feature'`)
4. Push to the branch (`git push origin feat/my-feature`)
5. Open a Pull Request

Please open an issue before starting work on large features to discuss the approach.

---

## License

MIT © [Geriatric Sailor](https://github.com/geriatric-sailor)
```

Every section maps to the architecture we've built — `templates.go`, `fswalk.go`, `render.go`, the variant system, the hook lifecycle, and the zero-install philosophy. A few notes on what I included and why:

- **Telemetry** is framed as opt-in/off-by-default with a `DO_NOT_TRACK` convention — standard practice for developer tools and a good trust signal for open-source users.
- **Hooks** are designed as sequential shell commands scoped to the scaffolded directory, with explicit failure behavior documented.
- **Roadmap** separates near-term (CI, cross-compile, self-update) from longer-term (custom templates, plugin system, TUI) so contributors can see where to jump in.
- **Variant table** is a placeholder set — adjust the variants to match exactly what you've built or plan to ship.

Copy this straight into your repo root. Let me know if you want to tune any section or add anything else.