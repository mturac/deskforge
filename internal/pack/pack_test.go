package pack

import (
	"strings"
	"testing"
)

func TestFilesContainAgents(t *testing.T) {
	files, err := Files()
	if err != nil {
		t.Fatal(err)
	}
	if len(files) < 10 {
		t.Fatalf("expected embedded pack, got %d files", len(files))
	}
	need := map[string]bool{
		"ARCHITECTURE.md":    false,
		"WATCHLIST.md":       false,
		"APPROVAL.md":        false,
		"HOWTO.md":           false,
		"agents/MAINER.md":   false,
		"agents/PULSE.md":    false,
		"agents/SCOUT.md":    false,
		"agents/JUDGE.md":    false,
		"agents/BUILDER.md":  false,
		"agents/CRITIC.md":   false,
		"desks/DEMAND.md":    false,
		"desks/PREFLIGHT.md": false,
	}
	for _, f := range files {
		p := filepathSlash(f.Path)
		if _, ok := need[p]; ok {
			need[p] = true
		}
		if strings.HasPrefix(p, "agents/") && !strings.Contains(string(f.Data), "```") {
			t.Errorf("%s: missing fenced Description block for Grok Bot paste", p)
		}
	}
	for p, ok := range need {
		if !ok {
			t.Errorf("missing embedded file %s", p)
		}
	}
}

func filepathSlash(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}
