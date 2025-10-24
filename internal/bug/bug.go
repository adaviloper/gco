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
	"github.com/spf13/viper"
)

func Run(cmd *cobra.Command, args []string) {
	repos := viper.Get("repositories.gco")
	var prefix string
	if list, ok := repos.([]interface{}); ok {
		for _, item := range list {
			if m, ok := item.(map[string]interface{}); ok {
				if val, ok := m["bug"]; ok {
					prefix = fmt.Sprintf("%v", val)
					break
				}
			}
		}
	}

	fmt.Printf("Prefix: %s\n", prefix)
	branch := str_utils.Slugify(strings.Join(args, "-"))
	git.CheckoutBranch(fmt.Sprintf("%s/%s", prefix, branch))
}

