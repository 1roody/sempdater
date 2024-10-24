package check

import (
	"errors"
	"fmt"
	"os"
	"sempdater/src/assets"
	"strings"
)

func FileToDistribute(arg string) string {
	if !strings.HasPrefix(arg, "-fD=") {
		Usage()
		os.Exit(1)
	}
	filePath := strings.Split(arg, "=")[1]
	fileDoesNotExist(filePath)
	return filePath
}

func RepositoriesList(arg string) string {
	if !strings.HasPrefix(arg, "-rL=") {
		Usage()
		os.Exit(1)
	}
	filePath := strings.Split(arg, "=")[1]
	fileDoesNotExist(filePath)
	return filePath
}

func fileDoesNotExist(filePath string) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s> The file %s hasn't been found. Please check if it exists.%s\n", assets.Red, filePath, assets.Nc)
		os.Exit(1)
	}
}

func Usage() {
	fmt.Printf("%sFor usage details: ./sempdater --help%s\n", assets.Yellow, assets.Nc)
	fmt.Println("")
	os.Exit(1)
}