package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/adaviloper/gco/internal/str_utils"
	"github.com/spf13/viper"
)

func Run(args ...string) ([]byte, error) {
	checkoutCommand := exec.Command("git", args...)
	output, err := checkoutCommand.Output()

	if err != nil {
		return nil, err
	}

	return output, nil
}

func CheckoutBranch(branch string) {
	// fmt.Printf("Switching to [%s]", branch)
	_, err := Run("checkout", "-B", branch)

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
	fmt.Printf("ticket prefix: %s\n", ticketPrefix)
	branch := str_utils.Slugify(description)
	fmt.Printf("Branch: %s/%s-%s\n", prefix, ticketPrefix, branch)
	CheckoutBranch(fmt.Sprintf("%s%s%s-%s", prefix, viper.GetString("separator"), ticketPrefix, branch))
}

func TicketPrefix() string {
    // Known repository name → ticket prefix mappings.
    prefixes := map[string]string{
        "ultimate-tic-tac-toe": "UTTT",
        "redwood":              "RW",
    }

    // Allow tests or callers to override repo name detection.
    if override := os.Getenv("GCO_REPO_NAME"); override != "" {
        if p, ok := prefixes[override]; ok {
            return p
        }
        return ""
    }

    // Try to resolve repository top-level via git, then map base directory.
    if out, err := Run("rev-parse", "--show-toplevel"); err == nil {
        repoPath := strings.TrimSpace(string(out))
        if repoPath != "" {
            base := filepath.Base(repoPath)
            if p, ok := prefixes[base]; ok {
                return p
            }
        }
    }

    // Fallback: use current working directory name.
    if wd, err := os.Getwd(); err == nil {
        base := filepath.Base(wd)
        if p, ok := prefixes[base]; ok {
            return p
        }
    }

    return ""
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
	return viper.GetString(key), nil
}

func GetPrefixForRepo(category string) (string, error) {
	prefix, _ := GetCurrentRepoProperty(category)

	if prefix == "" {
		return category, nil
	}

	return prefix, nil
}

