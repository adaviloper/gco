/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"strings"

	"github.com/adaviloper/gco/internal/git"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	git.GenerateBranchName("task", strings.Join(args, "-"))
}

