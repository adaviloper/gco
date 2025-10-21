/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/judedaryl/go-arrayutils"
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
		switchCmd.Run(cmd, args)
	},
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


