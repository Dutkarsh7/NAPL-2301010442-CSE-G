package main

import (
	"fmt"
	"math"
)

// Divide performs floating-point division of 'a' by 'b'.
// It returns an error if the divisor 'b' is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is undefined: cannot divide %.2f by 0", a)
	}
	return a / b, nil
}

// SquareRoot computes the square root of a non-negative float64 number.
// It returns an error if the input 'x' is negative.
func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("math error: square root of negative number %.2f is not defined in real numbers", x)
	}
	return math.Sqrt(x), nil
}

// ValidateInput verifies that the given input string is not empty.
// It returns an error if the input string is empty.
func ValidateInput(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("validation error: input string cannot be empty")
	}
	return input, nil
}

func main() {
	fmt.Println("================================================================")
	fmt.Println("  Lab 4 - Experiment 7: Error Handling & Unit Testing")
	fmt.Println("================================================================")

	// ---------------------------------------------------------
	// Scenario 1: Division by Zero
	// ---------------------------------------------------------
	fmt.Println("\n--- Scenario 1: Division Operation ---")

	// Valid input case
	val1, err1 := Divide(10.0, 2.0)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Valid Case   : Divide(10.0, 2.0) = %.2f\n", val1)
	}

	// Invalid input case (division by zero)
	val2, err2 := Divide(10.0, 0.0)
	if err2 != nil {
		fmt.Printf("Invalid Case : Divide(10.0, 0.0) -> Handled Error: %v\n", err2)
	} else {
		fmt.Printf("Result: %.2f\n", val2)
	}

	// ---------------------------------------------------------
	// Scenario 2: Square Root of a Negative Number
	// ---------------------------------------------------------
	fmt.Println("\n--- Scenario 2: Square Root Operation ---")

	// Valid positive input case
	sqrtVal1, sqrtErr1 := SquareRoot(25.0)
	if sqrtErr1 != nil {
		fmt.Printf("Error: %v\n", sqrtErr1)
	} else {
		fmt.Printf("Valid Case   : SquareRoot(25.0) = %.2f\n", sqrtVal1)
	}

	// Valid zero input case
	sqrtVal2, sqrtErr2 := SquareRoot(0.0)
	if sqrtErr2 != nil {
		fmt.Printf("Error: %v\n", sqrtErr2)
	} else {
		fmt.Printf("Valid Case   : SquareRoot(0.0)  = %.2f\n", sqrtVal2)
	}

	// Invalid negative input case
	sqrtVal3, sqrtErr3 := SquareRoot(-16.0)
	if sqrtErr3 != nil {
		fmt.Printf("Invalid Case : SquareRoot(-16.0) -> Handled Error: %v\n", sqrtErr3)
	} else {
		fmt.Printf("Result: %.2f\n", sqrtVal3)
	}

	// ---------------------------------------------------------
	// Scenario 3: Empty String Input Validation
	// ---------------------------------------------------------
	fmt.Println("\n--- Scenario 3: Input String Validation ---")

	// Valid non-empty input case
	str1, strErr1 := ValidateInput("Go Language Practical")
	if strErr1 != nil {
		fmt.Printf("Error: %v\n", strErr1)
	} else {
		fmt.Printf("Valid Case   : ValidateInput(\"Go Language Practical\") -> Validated: \"%s\"\n", str1)
	}

	// Invalid empty string input case
	str2, strErr2 := ValidateInput("")
	if strErr2 != nil {
		fmt.Printf("Invalid Case : ValidateInput(\"\") -> Handled Error: %v\n", strErr2)
	} else {
		fmt.Printf("Result: \"%s\"\n", str2)
	}

	fmt.Println("\n================================================================")
	fmt.Println("  All error scenarios handled gracefully without program crash!")
	fmt.Println("================================================================")
}
