package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mturac/deskforge/internal/pack"
)

const version = "0.1.0"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "deskforge: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		printHelp()
		return nil
	}
	switch args[0] {
	case "version", "-v", "--version":
		fmt.Println(version)
		return nil
	case "help", "-h", "--help":
		printHelp()
		return nil
	case "init":
		dir := "."
		if len(args) > 1 {
			dir = args[1]
		}
		return cmdInit(dir)
	case "agents":
		dir := "."
		if len(args) > 1 {
			dir = args[1]
		}
		return cmdAgents(dir)
	default:
		return fmt.Errorf("unknown command %q (try: deskforge help)", args[0])
	}
}

func printHelp() {
	fmt.Print(`deskforge — your agent agency, in your repo

Usage:
  deskforge init [dir]     Write ARCHITECTURE, desks, vault, and 6 GrokBot agent contracts
  deskforge agents [dir]   List agent Description files to paste into Grok Bot
  deskforge version        Print version
  deskforge help           Show this help

First value:
  deskforge init my-desk && cd my-desk && deskforge agents
  Then paste each agents/*.md fenced block into a Grok Bot Description.
`)
}

func cmdInit(dir string) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(abs, 0o755); err != nil {
		return err
	}
	entries, err := pack.Files()
	if err != nil {
		return err
	}
	written := 0
	skipped := 0
	for _, e := range entries {
		target := filepath.Join(abs, e.Path)
		if _, err := os.Stat(target); err == nil {
			skipped++
			continue
		} else if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(target, e.Data, 0o644); err != nil {
			return err
		}
		written++
	}
	fmt.Printf("deskforge init: %s\n  wrote %d files", abs, written)
	if skipped > 0 {
		fmt.Printf(", skipped %d existing", skipped)
	}
	fmt.Println()
	fmt.Println("Next: deskforge agents")
	fmt.Println("Then paste agent Description fences into Grok Bot (see HOWTO.md).")
	return nil
}

func cmdAgents(dir string) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	agentsDir := filepath.Join(abs, "agents")
	names := []string{
		"MAINER.md",
		"PULSE.md",
		"SCOUT.md",
		"JUDGE.md",
		"BUILDER.md",
		"CRITIC.md",
	}
	fmt.Println("Paste each fenced Description into the matching Grok Bot agent:")
	for i, name := range names {
		path := filepath.Join(agentsDir, name)
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("%s missing — run deskforge init first", path)
		}
		role := strings.TrimSuffix(name, ".md")
		fmt.Printf("  %d. %s  ←  %s\n", i+1, role, path)
	}
	fmt.Println("Kickoff text: HOWTO.md § Group kickoff")
	return nil
}
