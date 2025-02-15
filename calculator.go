package main

import "fmt"

func calculate(a int, b int) []float64 {
	sum := float64(a + b)
	difference := float64(a - b)
	product := float64(a * b)
	quotient := float64(a) / float64(b)
	results := []float64{sum, difference, product, quotient}
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
