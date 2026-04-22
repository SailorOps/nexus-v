package cli

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"

    "nexus-v/internal/config"
    "nexus-v/internal/git"
    "nexus-v/internal/hooks"
    "nexus-v/internal/templates"
    "nexus-v/internal/telemetry"
)

func Run() {
    var (
        name        = flag.String("name", "", "Human-readable project name")
        identifier  = flag.String("id", "", "Extension identifier (e.g. my-extension)")
        description = flag.String("desc", "", "Short description")
        template    = flag.String("template", "", "Template variant (default, minimal, web, multi)")
        force       = flag.Bool("force", false, "Overwrite existing files")
        dryRun      = flag.Bool("dry-run", false, "Preview files without writing them")
        list        = flag.Bool("list-templates", false, "List available template variants")
        teleFlag    = flag.String("telemetry", "", "Telemetry modes (none, session, local, project, all)")
        noHooks     = flag.Bool("no-hooks", false, "Disable post-generation hooks")
    )

    flag.Usage = func() {
        fmt.Println("Usage:")
        fmt.Println("  nexus-v [options] <target-directory>")
        fmt.Println()
        fmt.Println("Options:")
        flag.PrintDefaults()
        fmt.Println()
        fmt.Println("Examples:")
        fmt.Println("  nexus-v --name \"My Extension\" --id my-ext ./my-ext")
        fmt.Println("  nexus-v --list-templates")
    }

    flag.Parse()

    // --list-templates short-circuit
    if *list {
        templatesList, err := templates.ListTemplates()
        if err != nil {
            Error("Failed to list templates: " + err.Error())
            os.Exit(1)
        }
        Success("Available templates:")
        for _, t := range templatesList {
            fmt.Println(" -", t)
        }
        return
    }

    if flag.NArg() < 1 {
        Error("Missing target directory")
        os.Exit(1)
    }

    targetDir := flag.Arg(0)

    // Load config (user + project)
    cfg, _ := config.LoadConfig(targetDir)

    // Merge CLI flags with config defaults
    finalName := firstNonEmpty(*name, cfg.Name)
    finalID := firstNonEmpty(*identifier, cfg.Identifier)
    finalDesc := firstNonEmpty(*description, cfg.Description)
    finalTemplate := firstNonEmpty(*template, cfg.Template)
    if finalTemplate == "" {
        finalTemplate = "default"
    }

    if finalName == "" || finalID == "" {
        Error("Both --name and --id are required (or set in nexusv.json)")
        os.Exit(1)
    }

    // Telemetry mode resolution
    sess, loc, proj := telemetry.ParseModes(*teleFlag)
    if *teleFlag == "" {
        sess = cfg.Telemetry.Session
        loc = cfg.Telemetry.Local
        proj = cfg.Telemetry.Project
    }

    tel := telemetry.Telemetry{
        SessionEnabled: sess,
        LocalEnabled:   loc,
        ProjectEnabled: proj,
        SessionSink:    &telemetry.SessionSink{},
        LocalSink:      &telemetry.LocalSink{},
        ProjectSink:    &telemetry.ProjectSink{},
    }

    ctx := templates.Context{
        Name:        finalName,
        Identifier:  finalID,
        Description: finalDesc,
        CommandName: finalID + ".hello",
        Template:    finalTemplate,
        Force:       *force,
        DryRun:      *dryRun,
    }

    // Generation
    spin := NewSpinner()
    spin.Start("Generating project...")
    err := templates.GenerateProject(ctx, targetDir)
    spin.Stop()

    // Telemetry event
    ev := telemetry.Event{
        Template:   finalTemplate,
        DryRun:     *dryRun,
        Force:      *force,
        ProjectDir: targetDir,
    }
    tel.Record(ev)

    if err != nil {
        Error(err.Error())
        os.Exit(1)
    }

    if ctx.DryRun {
        Success("Dry run complete — no files were written")
        return
    }

    Success("Project created at " + filepath.Clean(targetDir))

    // Git integration
    if git.Available() {
        Info("Initializing Git repository...")
        if err := git.InitRepo(targetDir); err == nil {
            git.AddAll(targetDir)
            git.FirstCommit(targetDir)
            Success("Git repository initialized")
        } else {
            Warn("Git is installed but initialization failed")
        }
    } else {
        Warn("Git not found — skipping repository initialization")
    }

    // Post-generation hooks
    if !*noHooks && len(cfg.Hooks.Post) > 0 {
        Info("Running post-generation hooks...")
        if err := hooks.RunPostHooks(targetDir, cfg.Hooks.Post); err != nil {
            Warn("Some hooks failed: " + err.Error())
        } else {
            Success("All hooks completed")
        }
    }

    Info("Run `npm install` then press F5 to launch the extension")
}

func firstNonEmpty(a, b string) string {
    if a != "" {
        return a
    }
    return b
}
