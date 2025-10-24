package branch

import (
	"reflect"
	"testing"
)

func TestPrepareBranches_FiltersAndTrims(t *testing.T) {
	stdOut := "* main\n  feature/GC-23-some-description\n  bugfix/GC-23-fix\n  chore/other\n"
	got := prepareBranches(stdOut, "GC-23")
	want := []string{"feature/GC-23-some-description", "bugfix/GC-23-fix"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("prepareBranches returned %v; want %v", got, want)
	}
}

func TestPrepareBranches_NoMatches(t *testing.T) {
	stdOut := "* main\n  chore/other\n  docs/readme\n"
	got := prepareBranches(stdOut, "GC-23")
	if len(got) != 0 {
		t.Fatalf("expected no matches, got %v", got)
	}
}

func TestPrepareBranches_SubstringMatch(t *testing.T) {
	stdOut := "  feat/ABC-1\n  feat/ABC-12\n  feat/XYZ-1\n"
	got := prepareBranches(stdOut, "ABC-1")
    want := []string{"feat/ABC-1", "feat/ABC-12"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
