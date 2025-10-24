package git

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
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
	fmt.Printf("Switching to [%s]", branch)
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

func Bug(id int, desc string) string {
	return GenerateBranchName(TicketData{Type: "bugfix", ID: id, Description: desc})
}

type TicketData struct {
	Type string
	ID int
	Description string
}

func GenerateBranchName(ticket TicketData) string {
    typePart := strings.ToLower(strings.TrimSpace(ticket.Type))
    if typePart == "" {
        typePart = "feature"
    }

    prefix := ticketPrefix()
    idPart := fmt.Sprintf("%d", ticket.ID)

    // slugify description: lowercase, replace non-alnum with '-', trim '-'
    desc := strings.ToLower(strings.TrimSpace(ticket.Description))
    // Replace any sequence of non [a-z0-9] with '-'
    var b strings.Builder
    prevDash := false
    for _, r := range desc {
        if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9' || r == '_') {
            b.WriteRune(r)
            prevDash = false
            continue
        }
        if !prevDash {
            b.WriteByte('-')
            prevDash = true
        }
    }
    slug := strings.Trim(b.String(), "-")

    base := idPart
    if prefix != "" {
        base = prefix + "-" + idPart
    }

    if slug != "" {
        return typePart + "/" + base + "-" + slug
    }
    return typePart + "/" + base
}

func ticketPrefix() string {
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


