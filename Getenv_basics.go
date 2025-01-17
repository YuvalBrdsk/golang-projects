package main

import (
	"fmt"
	"os"
)

func main() {
	// Get and print the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println("Current Directory:", currentDir)

	// Change to the home direcroty
	homeDir := os.Getenv("HOME") // Get the HOME environment variable.
	if homeDir == "" {
		fmt.Println("Error: HOME environment variable is not set")
		return
	}

	err = os.Chdir(homeDir)
	if err != nil {
		fmt.Println("Error changing to home directory:", err)
		return
	}

	// Get and print the new working directory
	newDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting new directory:", err)
		return
	}
	fmt.Println("New Directory (Home):", newDir)
}
