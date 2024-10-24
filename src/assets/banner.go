package assets

import (
	"fmt"
)

func PrintBanner() {
	fmt.Printf("%s\n", Green)
	fmt.Println(" __ __ _  _  _  _ ___ __ _   /^--^\\ ")
	fmt.Println("(_ |_ |V||_)| \\|_| | |_ |_) |      | ")
	fmt.Println("__)|__| ||  |_/| | | |__| \\  \\____/ ")
	fmt.Println("---- Beta v1.0")
	fmt.Println("")
}