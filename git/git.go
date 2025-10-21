package git

import (
	"fmt"
	"os/exec"
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
