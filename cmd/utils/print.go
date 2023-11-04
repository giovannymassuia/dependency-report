package utils

import "fmt"

func PrintError(msg string) {
	// in red
	fmt.Println("\033[31mError:\033[0m", msg)
}
