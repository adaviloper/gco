/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/adaviloper/gco/git"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Conveniently switch git branches by pattern",
	Long: `Shell utility for quickly switching between git branches by pattern-matching:
	gco full/branch-name
	gco <some-substring>
	gco <ticket_number>`,
	Run: func(cmd *cobra.Command, args []string) {
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

		branches := prepareBranches(string(stdOut), target)

		if target == "-" {
			git.CheckoutBranch(target)
		} else if len(branches) == 1 {
			fmt.Printf("Switching to [%s]", branches[0])
			git.CheckoutBranch(branches[0])
		} else if len(branches) > 1 {
			prompt := promptui.Select{
				Label: "Which branch would you like to switch to?",
				Items: branches,
			}

			_, branch, err := prompt.Run()

			if err != nil {
				fmt.Println("No branch selected")
				return
			}

			fmt.Printf("Switching to [%s]", branch)
			git.CheckoutBranch(branch)
		} else {
			fmt.Println("No matching branches found.")
		}

	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
