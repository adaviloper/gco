package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/adaviloper/gco/internal/str_utils"
	"github.com/spf13/viper"
)

func Run(args ...string) ([]byte, error) {
	checkoutCommand := exec.Command("git", args...)
	output, err := checkoutCommand.Output()
	viper.GetBool("debug")

	if err != nil {
		return nil, err
	}

	return output, nil
}

func CheckoutBranch(branch string) {
	if viper.GetBool("debug") {
		fmt.Printf("Switching to [%s]", branch)
	}

	_, err := Run("checkout", "-B", branch)

	if err != nil {
		return
	}
}

func CheckoutBranchFromRemote(branch string) {
	if viper.GetBool("debug") {
		fmt.Printf("Switching to [%s]", branch)
	}

	_, err := Run("checkout", branch)

	if err != nil {
		return
	}
}

func CheckForUncommittedWork() bool {
	hasWork, err := Run("status", "--short")

	if err != nil {
		fmt.Println("An error occurred while trying to check the status.")
		return true
	}

	return len(hasWork) > 0
}

type TicketData struct {
	Type string
	ID int
	Description string
}

func GenerateBranchName(category string, description string) {
	prefix, err := GetPrefixForRepo(category)
	if err != nil {
	  return
	}
	ticketPrefix, _ := GetCurrentRepoProperty("ticket_prefix")
  if viper.GetBool("debug") {
		fmt.Printf("ticket prefix: %s\n", ticketPrefix)
  }
	branch := str_utils.Slugify(description)
  if viper.GetBool("debug") {
		fmt.Printf("Branch: %s/%s-%s\n", prefix, ticketPrefix, branch)
  }
  if ticketPrefix == "" {
		CheckoutBranch(fmt.Sprintf("%s%s%s", prefix, viper.GetString("separator"), branch))
  } else {
		CheckoutBranch(fmt.Sprintf("%s%s%s-%s", prefix, viper.GetString("separator"), ticketPrefix, branch))
  }
}

func GetCurrentRepoName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "default", nil
	}

	// Extract last directory name (repo name)
	parts := strings.Split(strings.TrimSpace(string(output)), "/")
	return parts[len(parts)-1], nil
}

func GetCurrentRepoProperty(prop string) (string, error) {
	repo, err := GetCurrentRepoName()
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("repositories.%s.%s", repo, prop)
	property := viper.GetString(key)
	if property == "" {
		key := fmt.Sprintf("repositories.%s.%s", "default", prop)
		return viper.GetString(key), nil
	}

	return property, nil
}

func GetPrefixForRepo(category string) (string, error) {
	prefix, _ := GetCurrentRepoProperty(category)

	if prefix == "" {
		return category, nil
	}

	return prefix, nil
}

