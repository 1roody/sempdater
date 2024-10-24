package assets

import (
	"fmt"
)

func CliHelper() {
	fmt.Printf("%s-fD= (The file you want to copy to the repo)%s", Yellow, Nc); fmt.Printf("%s REQUIRED%s\n", Red, Nc)
	fmt.Printf("%s-rL= (Your list of repositories to update)%s", Yellow, Nc); fmt.Printf("%s REQUIRED%s\n", Red, Nc)
}