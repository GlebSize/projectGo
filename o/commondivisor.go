package main

import (
	"fmt"
	"math"
)

func main() {
	a := 56
	b := 42

	gcd := math.GCD(a, b)
	fmt.Println(gcd)
}
