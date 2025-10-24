/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package bug

import (
	"fmt"
	"strings"

	"github.com/adaviloper/gco/internal/git"
	"github.com/adaviloper/gco/internal/str_utils"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Printf("args: %v\n", args)
	branch := str_utils.Slugify(strings.Join(args, "-"))
	git.CheckoutBranch(fmt.Sprintf("bugfix/%s", branch))
}

