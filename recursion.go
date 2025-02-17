package main

import "fmt"

func printSquares(n int) {
	// base case
	if n == -6 {
		return
	}
	fmt.Printf("%d ", n*n)
	printSquares(n - 1)
}
func main() {
	printSquares(2)
}

// Input: n=2; Output: 4 1 0 1 4 9 16 25
// At each iteration, the printSquares function will print the square of n to the console and then call itself with n decremented by 1.
// When n is -6, the if statement will evaluate to true and the function will return, terminating the recursive calls.
