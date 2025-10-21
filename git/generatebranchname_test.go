package git

import "testing"

func TestGenerateBranchName_WithPrefixAndSlug(t *testing.T) {
	// force prefix via env override
	t.Setenv("GCO_REPO_NAME", "redwood") // RW
	name := generateBranchName(TicketData{Type: "story", ID: 123, Description: "Add login flow"})
	if name != "story/RW-123-add-login-flow" {
		t.Fatalf("got %q", name)
	}
}

func TestGenerateBranchName_NoPrefix(t *testing.T) {
	// no mapping for this repo name
	t.Setenv("GCO_REPO_NAME", "unknown-repo")
	name := generateBranchName(TicketData{Type: "bugfix", ID: 7, Description: "Fix   crash!! @ startup"})
	if name != "bugfix/7-fix-crash-startup" {
		t.Fatalf("got %q", name)
	}
}

func TestGenerateBranchName_EmptyTypeDefaultsToFeature(t *testing.T) {
	t.Setenv("GCO_REPO_NAME", "ultimate-tic-tac-toe") // APP
	name := generateBranchName(TicketData{Type: "", ID: 45, Description: "  spaces  and $$ symbols  "})
	if name != "feature/UTTT-45-spaces-and-symbols" {
		t.Fatalf("got %q", name)
	}
}

func TestGenerateBranchName_Bug(t *testing.T) {
	t.Setenv("GCO_REPO_NAME", "redwood")
	name := Bug(123, "some description")
	if name != "bugfix/RW-123-some-description" {
		t.Fatalf("got %q", name)
	}
}
