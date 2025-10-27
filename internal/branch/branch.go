/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package branch

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/adaviloper/gco/internal/git"
	"github.com/judedaryl/go-arrayutils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
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

	shellCommand := exec.Command("git", ops...)

	stdOut, err := shellCommand.Output()

	if err != nil {
		fmt.Printf("Failed to get branches: [%s]", err.Error())
		return
	}

	branches := PrepareBranches(string(stdOut), target)

	if target == "-" {
		git.CheckoutBranch(target)
	} else if len(branches) == 1 {
		git.CheckoutBranch(branches[0])
	} else if len(branches) > 1 {
		prompt := promptui.Select{
			Label: "Which branch would you like to switch to?",
			Items: branches,
			HideSelected: true,
		}

		_, branch, err := prompt.Run()

		if err != nil {
			fmt.Println("No branch selected")
			return
		}

		git.CheckoutBranch(branch)
	} else {
		fmt.Println("No matching branches found.")
	}
}

func PrepareBranches(stdOut string, target string) []string {
	branches := strings.Split(stdOut, "\n")

	branches = arrayutils.Filter(branches, func(branch string) bool {
		return strings.Contains(branch, target) && !strings.Contains(branch, "*")
	})

	branches = arrayutils.Map(branches, func(branch string) string {
		return strings.Trim(branch, " ")
	})

  return branches
}
