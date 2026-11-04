package main

import (
	"example.cpm/app/mathutil"
	"fmt"
)

func main() {
	fmt.Println("=== String Manipulation Functions ===")

	// Reverse function
	testString := "Hello World"
	reversed := mathutil.Reverse(testString)
	fmt.Printf("Original: %s\n", testString)
	fmt.Printf("Reversed: %s\n\n", reversed)

	// CountVowels function
	vowelCount := mathutil.CountVowels(testString)
	fmt.Printf("String: %s\n", testString)
	fmt.Printf("Number of vowels: %d\n\n", vowelCount)

	fmt.Println("=== Math Utility Functions ===")

	// Factorial function
	result := mathutil.Factorial(5)
	fmt.Printf("Factorial(5) = %d\n", result)

	// Power function
	result = mathutil.Power(2, 8)
	fmt.Printf("Power(2, 8) = %d\n", result)
}
