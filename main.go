package main

import "fmt"
import "math"

func main() {
	fmt.Println("Building...")

	fmt.Println("Training...")

	fmt.Println("Done...")
}

func seq(x float64) float64 {
	return math.Pow(x, 3) + math.Pow(x, 2) - x - 1
}
