package prompts

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Answers struct {
    Name        string
    Identifier  string
    Description string
    CommandName string
}

func AskQuestions() (*Answers, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Welcome to the VS Code Extension Scaffolder")
    fmt.Println("-------------------------------------------")

    name := ask(reader, "Extension name", "My Extension")
    identifier := ask(reader, "Extension identifier", sanitizeIdentifier(name))
    description := ask(reader, "Description", "A VS Code extension")
    commandName := ask(reader, "Command name", "myExtension.helloWorld")

    return &Answers{
        Name:        name,
        Identifier:  identifier,
        Description: description,
        CommandName: commandName,
    }, nil
}

func ask(reader *bufio.Reader, label, defaultValue string) string {
    fmt.Printf("%s (%s): ", label, defaultValue)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    if input == "" {
        return defaultValue
    }
    return input
}

func sanitizeIdentifier(name string) string {
    id := strings.ToLower(name)
    id = strings.ReplaceAll(id, " ", "-")
    return id
}
