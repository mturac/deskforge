package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitAndAgents(t *testing.T) {
	dir := t.TempDir()
	if err := run([]string{"init", dir}); err != nil {
		t.Fatal(err)
	}
	need := []string{
		"ARCHITECTURE.md",
		"agents/MAINER.md",
		"agents/BUILDER.md",
		"agents/CRITIC.md",
		"desks/DEMAND.md",
	}
	for _, p := range need {
		if _, err := os.Stat(filepath.Join(dir, p)); err != nil {
			t.Fatalf("missing %s: %v", p, err)
		}
	}
	if err := run([]string{"init", dir}); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"agents", dir}); err != nil {
		t.Fatal(err)
	}
}
