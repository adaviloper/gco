/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package ticketCreator

import (
	"strings"

	"github.com/adaviloper/gco/internal/git"
)

func Run(ticketType string, args []string) {
	git.GenerateBranchName(ticketType, strings.Join(args, "-"))
}

