package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Frontend Journey: Day 1")

	// Let's do a "Speed Demon" test: Calculating a large sum
	start := time.Now()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	elapsed := time.Since(start)

	fmt.Printf("Calculation result: %d\n", sum)
	fmt.Printf("Time taken in Go WASM: %s\n", elapsed)
}
