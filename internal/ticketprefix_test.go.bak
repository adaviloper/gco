package git

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTicketPrefix_FromEnvOverride(t *testing.T) {
	t.Setenv("GCO_REPO_NAME", "redwood")
	if got := ticketPrefix(); got != "RW" {
		t.Fatalf("expected RW, got %q", got)
	}
}

func TestTicketPrefix_FromGitTopLevel(t *testing.T) {
	tmp := setupFakeGit(t)
	oldPath := os.Getenv("PATH")
	t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)

	tmpRepo := filepath.Join(t.TempDir(), "ultimate-tic-tac-toe")
	// Our fake git script in setupFakeGit will echo $GIT_FAKE_OUTPUT
	t.Setenv("GIT_FAKE_OUTPUT", tmpRepo)

	if got := ticketPrefix(); got != "UTTT" {
		t.Fatalf("expected APP, got %q", got)
	}
}

func TestTicketPrefix_FallsBackToWd(t *testing.T) {
	// When git rev-parse fails, we should fallback to cwd
	tmp := setupFakeGit(t)
	oldPath := os.Getenv("PATH")
	t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)
	t.Setenv("GIT_FAKE_EXIT", "1") // cause Run to fail

	wd := t.TempDir()
	if err := os.Mkdir(filepath.Join(wd, "ultimate-tic-tac-toe"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.Chdir(filepath.Join(wd, "ultimate-tic-tac-toe")); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	if got := ticketPrefix(); got != "UTTT" {
		t.Fatalf("expected UTTT, got %q", got)
	}
}
