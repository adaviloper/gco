/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package branch

import (
	"fmt"
	"strings"

	"github.com/adaviloper/gco/internal/git"
	"github.com/judedaryl/go-arrayutils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// switchCmd represents the switch command
func Run(cmd *cobra.Command, args []string) {
	hasUncommittedWork := git.CheckForUncommittedWork()

	if hasUncommittedWork {
		fmt.Println("Please commit any work before switching branches")
		return
	}

 	remote, _ := cmd.Root().Flags().GetBool("remote")
	target := args[0]

	ops := []string{
		"branch",
	}

	if remote {
		ops = append(ops, "-a")
	}

	stdOut, err := git.Run(ops...)

	if err != nil {
		fmt.Printf("Failed to get branches: [%s]", err.Error())
		return
	}

	branches := PrepareBranches(string(stdOut), target)

	targetBranch := ""
	if target == "-" {
		targetBranch = target
	} else if len(branches) == 1 {
		targetBranch = branches[0]
	} else if len(branches) > 1 {
		prompt := promptui.Select{
			Label: "Which branch would you like to switch to?",
			Items: branches,
			HideSelected: true,
		}

		_, selectedBranch, err := prompt.Run()

		if err != nil {
			fmt.Println("No branch selected")
			return
		}

		targetBranch = selectedBranch
	}
	

	git.CheckoutBranch(targetBranch)
}

func PrepareBranches(stdOut string, target string) []string {
  if viper.GetBool("debug") {
		fmt.Printf("Found branches: %s\n", stdOut)
		fmt.Printf("Target branch: %s\n", target)
  }

	branches := strings.Split(stdOut, "\n")

	branches = arrayutils.Filter(branches, func(branch string) bool {
		return strings.Contains(branch, target) && !strings.Contains(branch, "*")
	})

	branches = arrayutils.Map(branches, func(branch string) string {
		return strings.Trim(strings.Replace(branch, "remotes/origin/", "", 1), " ")
	})

  return branches
}
