package handle

import (
	"fmt"
	"os/exec"
	"sempdater/src/assets"
)

func checkoutBranch() string {
	if tryCheckout("develop") {
		return "develop"
	}
	if tryCheckout("homolog") {
		return "homolog"
	}

	fmt.Printf("%s> Neither 'develop' nor 'homolog' branches exist. Staying on current branch.%s\n", assets.Red, assets.Nc)
	currentBranch := RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	return currentBranch
}

func tryCheckout(branch string) bool {
	fmt.Printf("%s> Attempting to checkout branch '%s'...%s\n", assets.Cyan, branch, assets.Nc)
	cmd := exec.Command("git", "checkout", branch)
	if err := cmd.Run(); err == nil {
		fmt.Printf("%s> Successfully switched to branch '%s'.%s\n", assets.Green, branch, assets.Nc)
		return true
	}
	fmt.Printf("%s> Branch '%s' does not exist or checkout failed.%s\n", assets.Red, branch, assets.Nc)
	return false
}
