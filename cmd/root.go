/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/adaviloper/gco/internal/branch"
	"github.com/adaviloper/gco/internal/tickets/bug"
	"github.com/adaviloper/gco/internal/tickets/spike"
	"github.com/adaviloper/gco/internal/tickets/story"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gco",
	Short: "Convenient git branch switcher",
	Long: `Shell utility for quickly switching between git branches by pattern-matching:
	gco full/branch-name
	gco <some-substring>
	gco <ticket_number>`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initializeConfig(cmd)
	},
	// Args: cobra.NoArgs(),
	// Uncoment the following line if your bare application
	// has an action associated with it:
  RunE: func(cmd *cobra.Command, args []string) error {
  	fmt.Printf("%v\n", viper.Get("repositories"))
  	ticketTypes := []string{
  		"bug",
  		"story",
  		"spike",
  	}

		selectedType := ""
  	for _, ticketType := range ticketTypes {
  		val, _ := cmd.Flags().GetBool(ticketType)
  		if val == true {
  			selectedType = ticketType
  			break;
  		}
  	}
  	switch selectedType {
  	case "bug":
  		bug.Run(cmd, args)
  		return nil
  	case "spike":
  		spike.Run(cmd, args)
  		return nil
  	case "story":
  		story.Run(cmd, args)
  		return nil
  	}

    branch.Run(cmd, args)
  	return nil
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gco/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("remote", "r", false, "Include remote branches")

	rootCmd.Flags().BoolP("bug", "b", false, "Create a bug ticket")
	rootCmd.Flags().BoolP("story", "s", false, "Create a story ticket")
	rootCmd.Flags().BoolP("spike", "p", false, "Create a spike ticket")
}

func initializeConfig(cmd *cobra.Command) error {
	// 1. Set up Viper to use environment variables.
	viper.SetEnvPrefix("GCO")
	// Allow for nested keys in environment variables (e.g. `MYAPP_DATABASE_HOST`)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()
	viper.SetDefault(
		"repositories",
		map[string]map[string]string{
			"default": {
				"bug": "bug",
				"story": "feat",
				"spike": "spike",
				"epic": "epic",
				"task": "task",
			},
		},
		)
	viper.SetDefault(
		"separator",
		"/",
		)

	// 2. Handle the configuration file.
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for a config file in default locations.
		home, err := os.UserHomeDir()
		// Only panic if we can't get the home directory.
		cobra.CheckErr(err)

		// Search for a config file with the name "config" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + "/.config/gco")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// 3. Read the configuration file.
	// If a config file is found, read it in. We use a robust error check
	// to ignore "file not found" errors, but panic on any other error.
	if err := viper.ReadInConfig(); err != nil {
		// It's okay if the config file doesn't exist.
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	// 4. Bind Cobra flags to Viper.
	// This is the magic that makes the flag values available through Viper.
	// It binds the full flag set of the command passed in.
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	// This is an optional but useful step to debug your config.
	// fmt.Println("Configuration initialized. Using config file:", viper.ConfigFileUsed())
	return nil
}
