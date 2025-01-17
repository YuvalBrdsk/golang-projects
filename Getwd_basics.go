package main

import (
	"fmt"
	"os"
)

func main() {
	// Print a greeting
	fmt.Println("Ahlan, Yuval!")

	// Get and print the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Println("Current Working Directory:", workingDir)
}
