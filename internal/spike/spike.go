/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package spike

import (
	"fmt"
	"strings"

	"github.com/adaviloper/gco/internal/git"
	"github.com/adaviloper/gco/internal/str_utils"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	prefix, err := git.GetPrefixForRepo("spike")
	if err != nil {
	  return 
	}
	branch := str_utils.Slugify(strings.Join(args, "-"))
	fmt.Printf("Branch: %s/%s\n", prefix, branch)
	git.CheckoutBranch(fmt.Sprintf("%s/%s", prefix, branch))
}

