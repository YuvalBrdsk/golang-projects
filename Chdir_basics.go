package main

import (
	"fmt"
	"os"
)

func main() {
	// Print initial working directory
	initialDir, _ := os.Getwd()
	fmt.Println("Initial Directory:", initialDir)

	// Change to a different directory
	err := os.Chdir("/tmp")
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Print new working directory
	newDir, _ := os.Getwd()
	fmt.Println("New Working Directory:", newDir)
}
