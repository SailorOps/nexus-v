# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-04-21
### Added
- Initial release of NEXUS-V
- Template variants system (command, webview, language, theme)
- Dynamic template rendering with Go `text/template`
- Filesystem walker for embedded and local templates
- Interactive CLI prompts with `bufio`
- YAML-based configuration (`.nexusvrc.yaml`) with Environment Variable support
- Self-update functionality via GitHub Releases
- Multi-platform support (Windows, Linux, macOS)
- Opt-in telemetry system
- GitHub Actions for CI/CD and automated releases
