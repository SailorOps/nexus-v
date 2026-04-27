#!/bin/bash

# Still Systems | Nexus-V Repository Setup Script
# This script uses the GitHub CLI (gh) to standardize repository settings, 
# labels, and topics to match the Still Systems brand.

REPO="stillsystems/nexus-v"

echo "⚓ Starting repository standardization for $REPO..."

# 1. Set Repository Topics
echo "Setting topics..."
gh repo edit $REPO --add-topic "vscode-extension,scaffolder,go,cli,still-systems,developer-tools"

# 2. Configure Brand Labels
echo "Configuring brand labels..."

# Define brand colors from GUIDE.md
declare -A labels
labels=(
  ["type:bug"]="#D73A4A"
  ["type:feature"]="#A2EEEF"
  ["good first issue"]="#7057FF"
  ["priority:high"]="#B60205"
  ["status:blocked"]="#000000"
  ["documentation"]="#0075CA"
  ["branding"]="#111827"
)

for label in "${!labels[@]}"; do
    color="${labels[$label]}"
    echo "  Syncing label: $label ($color)"
    gh label edit "$label" --color "$color" --repo $REPO || \
    gh label create "$label" --color "$color" --repo $REPO
done

# 3. Enable Discussions
echo "Enabling GitHub Discussions..."
gh repo edit $REPO --enable-discussions

# 4. Success
echo "✅ Repository standardization complete! Nexus-V is now 100% brand-compliant."
