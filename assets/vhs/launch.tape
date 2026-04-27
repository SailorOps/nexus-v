Output assets/launch.gif

Set Shell bash
Set FontSize 24
Set Width 1200
Set Height 650
Set Padding 50
Set WindowBar Colorful
Set Theme "Horizon"
Set TypingSpeed 75ms

# Setup environment silently
Hide
Type "export LANG=C.UTF-8 && chmod +x ./nexus-v-linux && rm -rf Nexus-Launch-Demo"
Enter
Type "clear"
Enter
Sleep 1s
Show
Sleep 1s

# No Git? No problem. Nexus-V adapts to any environment.
# 1. Show features
Sleep 0.5s
Type "./nexus-v-linux --help"
Sleep 1s
Enter
Sleep 3s

# 2. Show templates
Sleep 0.5s
Type "./nexus-v-linux list"
Sleep 1s
Enter
Sleep 3s

# 3. Interactive scaffolding
Sleep 0.5s
Type "./nexus-v-linux init"
Sleep 1s
Enter
Sleep 2s

# Fill in details
Type "Nexus-Launch-Demo"
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 1.5s

# Select 'webview' template - showing the options
Down 1
Sleep 0.4s
Down 1
Sleep 0.4s
Down 1
Sleep 0.4s
Down 1
Sleep 0.4s
Down 1
Sleep 0.4s
Down 1
Sleep 1s
Enter
Sleep 2s

# Handle 'Command Name' prompt
Type "openNexusPanel"
Sleep 1s
Enter
Sleep 1.5s

# Select License
Type "MIT"
Sleep 1s
Enter
Sleep 5s

# 4. Final reveal
Type "ls -R Nexus-Launch-Demo"
Sleep 1s
Enter
Sleep 10s