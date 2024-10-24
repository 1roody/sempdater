package check

import "os/exec"

func IsGitUserConfigured() bool {
	cmd := exec.Command("git", "config", "--global", "user.name")
	err := cmd.Run()
	return err == nil
}