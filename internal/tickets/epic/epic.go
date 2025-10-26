/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package epic

import (
	"strings"

	"github.com/adaviloper/gco/internal/git"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	git.GenerateBranchName("epic", strings.Join(args, "-"))
}

