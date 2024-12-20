package handle

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sempdater/src/assets"
	"strconv"
	"strings"
	"time"
)

func SelectRepos(repoFile string) []string {
	allRepos, err := ReadLines(repoFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, repo := range allRepos {
		fmt.Printf("%d) %s\n", i+1, repo)
	}

	fmt.Printf("\n%s> Enter the number of the desired repositories you want to update (separated by spaces) or 'A' to select all:%s\n", assets.Cyan, assets.Nc)

	var selectedRepos string
	fmt.Scanln(&selectedRepos)

	var repos []string

	if selectedRepos == "A" {
		repos = allRepos
	} else {
		for _, i := range strings.Split(selectedRepos, " ") {
			index, err := strconv.Atoi(i)
			if err != nil || index < 1 || index > len(allRepos) {
				fmt.Printf("%s> Invalid selection: %s%s\n", assets.Red, i, assets.Nc)
				os.Exit(1)
			}
			repos = append(repos, allRepos[index-1])
		}
	}
	return repos
}

func updateRepositories(baseBranch, fileToDistribute, repoName string) {
	branchName := "sempdater-auto-update-" + time.Now().Format("20060102")
	fmt.Printf("%s> Creating the branch: %s...%s\n", assets.Cyan, branchName, assets.Nc)
	RunCommand("git", "checkout", "-b", branchName)

	if _, err := os.Stat(".github/workflows"); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(".github/workflows", os.ModePerm)
	}

	UpdateRemoteRepo(branchName, baseBranch, fileToDistribute, repoName)
}


func ProcessRepo(repo, fileToDistribute string) {
	fmt.Printf("\n%s> Cloning %s...%s\n", assets.Cyan, repo, assets.Nc)

	repoName := filepath.Base(strings.TrimSuffix(repo, ".git"))
	if _, err := os.Stat(repoName); errors.Is(err, os.ErrNotExist) {
		RunCommand("git", "clone", repo)
	}

	os.Chdir(repoName)

	fmt.Printf("%s> Accessing: %s...%s\n", assets.Cyan, repo, assets.Nc)

	baseBranch := checkoutBranch()
	updateRepositories(baseBranch, fileToDistribute, repoName)
}

func UpdateRemoteRepo(branchName, baseBranch, fileToDistribute, repoName string) {
	RunCommand("cp", "../"+fileToDistribute, ".github/workflows/")
	RunCommand("git", "add", ".")

	fmt.Printf("%s> Committing the changes...%s\n", assets.Cyan, assets.Nc)
	RunCommand("git", "commit", "-m", "Updating semgrep configuration")

	fmt.Printf("%s> Pushing to origin: %s...%s\n", assets.Cyan, branchName, assets.Nc)
	RunCommand("git", "push", "--set-upstream", "origin", branchName)

	fmt.Printf("%s> Creating a pull request...%s\n", assets.Cyan, assets.Nc)
	RunCommand("gh", "pr", "create", "--base", baseBranch, "--head", branchName, "--title", "Sempdater Automatic Pull Request " + time.Now().Format("20060102"), "--body-file", "../test-files/pull-request-template.md")
	
	fmt.Printf("%s> Repository successfully updated, check the GitHub remote repository!%s\n", assets.Green, assets.Nc)

	os.Chdir("..")

	fmt.Printf("%s> Removing repo folder...%s\n", assets.Cyan, assets.Nc)
	RunCommand("rm", "-rf", repoName)
}

func RunCommand(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s> Failed to run command: %s %v%s\n", assets.Red, name, arg, assets.Nc)
        fmt.Printf("%s> Output: %s%s\n", assets.Red, string(out), assets.Nc)
		os.Exit(1)
	}
	return strings.TrimSpace(string(out))
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
