/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/judedaryl/go-arrayutils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func prepareBranches(stdOut string, target string) []string {
	branches := strings.Split(stdOut, "\n")

	branches = arrayutils.Filter(branches, func(branch string) bool {
		return strings.Contains(branch, target) && !strings.Contains(branch, "*")
	})

	branches = arrayutils.Map(branches, func(branch string) string {
		return strings.Trim(branch, " ")
	})

  return branches
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gco",
	Short: "Convenient git branch switcher",
	Long: `Shell utility for quickly switching between git branches by pattern-matching:

	gco full/branch-name
	gco <some-substring>
	gco <ticket_number>`,
	Args: cobra.ExactArgs(1),
	// Uncoment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		hasUncommittedWork := checkForUncommittedWork()

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
			checkoutBranch(target)
		} else if len(branches) == 1 {
			fmt.Printf("Switching to [%s]", branches[0])
			checkoutBranch(branches[0])
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
			checkoutBranch(branch)
		} else {
			fmt.Println("No matching branches found.")
		}
	},
}

func git(args ...string) ([]byte, error) {
	checkoutCommand := exec.Command("git", args...)
	output, err := checkoutCommand.Output()

	if err != nil {
		return nil, err
	}

	return output, nil
}

func checkForUncommittedWork() bool {
	hasWork, err := git("status", "--short")

	if err != nil {
		fmt.Println("An error occurred while trying to check the status.")
		return true
	}

	return len(hasWork) > 0
}

func checkoutBranch(branch string) {
	fmt.Printf("Switching to [%s]", branch)
	_, err := git("git", "checkout", branch)

	if err != nil {
		return
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gco.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("remote", "r", false, "Include remote branches")
}


