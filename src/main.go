package main

import (
	"fmt"
	"os"
	"sempdater/src/assets"
	"sempdater/src/check"
	"sempdater/src/handle"
)

func main() {
	assets.PrintBanner()

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		assets.CliHelper()
		os.Exit(0)
	}
	
	if len(os.Args) < 3 {
		check.Usage()
	}

	if !check.IsGitUserConfigured() {
		fmt.Printf("%s> You must be logged in to do this operation.%s\n", assets.Red, assets.Nc)
		os.Exit(1)
	}

	fileToDistribute := check.FileToDistribute(os.Args[1])
	repoFile := check.RepositoriesList(os.Args[2])

	repos := handle.SelectRepos(repoFile)

	if len(repos) == 0 {
		fmt.Printf("%s> No repositories selected :(%s\n", assets.Red, assets.Nc)
		os.Exit(1)
	}

	for _, repo := range repos {
		handle.ProcessRepo(repo, fileToDistribute)
	}

	fmt.Printf("%s> Operation completed successfully!%s\n", assets.Green, assets.Nc)
}