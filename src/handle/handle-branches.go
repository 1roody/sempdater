package handle

import (
	"fmt"
	"os/exec"
	"sempdater/src/assets"
)

func checkoutBranch() string {
	if branchExists("develop") {
		return switchBranch("develop")
	} else if branchExists("homolog") {
		return switchBranch("homolog")
	} else {
		fmt.Printf("%s> Neither 'develop' or 'homolog' branches exist. Using master/main instead...%s\n", assets.Red, assets.Nc)
		return RunCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	}
}

func branchExists(branch string) bool {
	cmd := exec.Command("git", "show-ref", "--quiet", "refs/heads/"+branch)
	return cmd.Run() == nil
}

func switchBranch(branch string) string {
	fmt.Printf("%s> Branch '%s' exists. Attempting to checkout...%s\n", assets.Cyan, branch, assets.Nc)
	RunCommand("git", "checkout", branch)
	fmt.Printf("%s> Switched to %s%s\n", assets.Green, branch, assets.Nc)
	return branch
}
