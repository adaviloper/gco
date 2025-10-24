package git

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func setupFakeGit(t *testing.T) string {
	t.Helper()
	tmpDir := t.TempDir()
	scriptPath := filepath.Join(tmpDir, "git")
	script := "#!/usr/bin/env sh\n" +
		"# Fake git for tests\n" +
		"if [ \"$1\" = \"status\" ] && [ \"$2\" = \"--short\" ]; then\n" +
		"  if [ \"$GIT_FAKE_STATUS\" = \"error\" ]; then\n" +
		"    echo err 1>&2\n" +
		"    exit 1\n" +
		"  fi\n" +
		"  printf \"%s\" \"$GIT_FAKE_STATUS\"\n" +
		"  exit 0\n" +
		"fi\n" +
		"if [ \"$1\" = \"checkout\" ]; then\n" +
		"  if [ \"$GIT_FAKE_CHECKOUT\" = \"error\" ]; then\n" +
		"    echo err 1>&2\n" +
		"    exit 1\n" +
		"  fi\n" +
		"  exit 0\n" +
		"fi\n" +
		"if [ -n \"$GIT_FAKE_OUTPUT\" ]; then\n" +
		"  printf \"%s\" \"$GIT_FAKE_OUTPUT\"\n" +
		"fi\n" +
		"if [ -n \"$GIT_FAKE_EXIT\" ]; then\n" +
		"  exit \"$GIT_FAKE_EXIT\"\n" +
		"fi\n" +
		"exit 0\n"
	if err := os.WriteFile(scriptPath, []byte(script), 0755); err != nil {
		t.Fatalf("failed to write fake git script: %v", err)
	}
	return tmpDir
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	b, _ := io.ReadAll(r)
	return string(b)
}

func TestRun_ReturnsOutput(t *testing.T) {
	tmp := setupFakeGit(t)
  oldPath := os.Getenv("PATH")
  t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)
	t.Setenv("GIT_FAKE_OUTPUT", "hello")
	out, err := Run("whatever")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(out) != "hello" {
		t.Fatalf("expected output 'hello', got %q", string(out))
	}
}

func TestRun_ReturnsErrorOnNonZeroExit(t *testing.T) {
	tmp := setupFakeGit(t)
  oldPath := os.Getenv("PATH")
  t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)
	t.Setenv("GIT_FAKE_EXIT", "2")
	_, err := Run("whatever")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestCheckForUncommittedWork_TrueWhenHasWork(t *testing.T) {
	tmp := setupFakeGit(t)
  oldPath := os.Getenv("PATH")
  t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)
	t.Setenv("GIT_FAKE_STATUS", " M file.go\n")
	if got := CheckForUncommittedWork(); got != true {
		t.Fatalf("expected true, got %v", got)
	}
}

func TestCheckForUncommittedWork_FalseWhenClean(t *testing.T) {
	tmp := setupFakeGit(t)
  oldPath := os.Getenv("PATH")
  t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)
	t.Setenv("GIT_FAKE_STATUS", "")
	if got := CheckForUncommittedWork(); got != false {
		t.Fatalf("expected false, got %v", got)
	}
}

func TestCheckForUncommittedWork_TrueOnError(t *testing.T) {
	tmp := setupFakeGit(t)
  oldPath := os.Getenv("PATH")
  t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)
	t.Setenv("GIT_FAKE_STATUS", "error")
	if got := CheckForUncommittedWork(); got != true {
		t.Fatalf("expected true on error, got %v", got)
	}
}

func TestCheckoutBranch_PrintsMessage(t *testing.T) {
	tmp := setupFakeGit(t)
  oldPath := os.Getenv("PATH")
  t.Setenv("PATH", tmp+string(os.PathListSeparator)+oldPath)
	printed := captureStdout(t, func() {
		CheckoutBranch("feature/XYZ-1")
	})
	want := "Switching to [feature/XYZ-1]"
	if !reflect.DeepEqual(printed, want) && len(printed) > 0 && printed[:len(want)] != want {
		t.Fatalf("expected printed prefix %q, got %q", want, printed)
	}
}
