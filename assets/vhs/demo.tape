Output assets/demo.gif

Set FontSize 22
Set Width 1200
Set Height 600
Set Padding 40
Set WindowBar Colorful
Set Theme "Catppuccin Mocha"

# Global pacing (very slow)
Set TypingSpeed 200ms
Set Framerate 15

Sleep 3s

# Start init
Type "nexus-v init"
Enter
Sleep 4s

# Extension name
Type "MyAwesomeExtension"
Enter
Sleep 4s
Enter
Sleep 4s
Enter
Sleep 4s

# Template selection
Down
Sleep 3s
Enter
Sleep 4s

# License
Type "MIT"
Enter
Sleep 4s

# Show structure
Type "ls -R"
Enter
Sleep 5s

# Doctor
Type "nexus-v doctor"
Enter
Sleep 6s

# Final message
Type "echo '✔ Nexus-V ready to go!'"
Enter
Sleep 4s
