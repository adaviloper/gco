package cmd

import (
    "bytes"
    "testing"

    "github.com/spf13/cobra"
)

func TestRoot_Delegation_TableDriven(t *testing.T) {
    tests := []struct {
        name        string
        args        []string
        wantCalled  string // "bug" or "switch"
    }{
        {name: "story arg delegates to story", args: []string{"story"}, wantCalled: "story"},
        {name: "bug arg delegates to bug", args: []string{"bug"}, wantCalled: "bug"},
        {name: "non-bug delegates to switch", args: []string{"GC-23"}, wantCalled: "switch"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            called := ""
            origStoryRun := storyCmd.Run
            origBugRun := bugCmd.Run
            origSwitchRun := switchCmd.Run
            t.Cleanup(func() {
                storyCmd.Run = origStoryRun
                bugCmd.Run = origBugRun
                switchCmd.Run = origSwitchRun
            })

            storyCmd.Run = func(cmd *cobra.Command, args []string) { called = "story" }
            bugCmd.Run = func(cmd *cobra.Command, args []string) { called = "bug" }
            switchCmd.Run = func(cmd *cobra.Command, args []string) { called = "switch" }

            rootCmd.SetArgs(tt.args)
            var buf bytes.Buffer
            rootCmd.SetOut(&buf)
            rootCmd.SetErr(&buf)
            if err := rootCmd.Execute(); err != nil {
                t.Fatalf("execute error: %v", err)
            }
            if called != tt.wantCalled {
                t.Fatalf("expected %s to be called, got %q", tt.wantCalled, called)
            }
        })
    }
}
