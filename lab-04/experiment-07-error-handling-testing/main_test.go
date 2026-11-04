package main

import (
	"math"
	"testing"
)

// TestDivide tests the Divide function for valid division and division by zero.
func TestDivide(t *testing.T) {
	t.Run("Valid division", func(t *testing.T) {
		a, b := 10.0, 2.0
		expected := 5.0
		result, err := Divide(a, b)

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if result != expected {
			t.Errorf("expected %.2f, got %.2f", expected, result)
		}
	})

	t.Run("Division by zero", func(t *testing.T) {
		a, b := 10.0, 0.0
		_, err := Divide(a, b)

		if err == nil {
			t.Fatal("expected an error for division by zero, but got nil")
		}
	})
}

// TestSquareRoot tests the SquareRoot function for positive numbers, zero, and negative numbers.
func TestSquareRoot(t *testing.T) {
	t.Run("Valid positive number", func(t *testing.T) {
		input := 16.0
		expected := 4.0
		result, err := SquareRoot(input)

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if math.Abs(result-expected) > 1e-9 {
			t.Errorf("expected %.2f, got %.2f", expected, result)
		}
	})

	t.Run("Square root of zero", func(t *testing.T) {
		input := 0.0
		expected := 0.0
		result, err := SquareRoot(input)

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if math.Abs(result-expected) > 1e-9 {
			t.Errorf("expected %.2f, got %.2f", expected, result)
		}
	})

	t.Run("Negative number error", func(t *testing.T) {
		input := -9.0
		_, err := SquareRoot(input)

		if err == nil {
			t.Fatal("expected an error for negative square root input, but got nil")
		}
	})
}

// TestValidateInput tests the ValidateInput function for valid strings and empty strings.
func TestValidateInput(t *testing.T) {
	t.Run("Valid non-empty string", func(t *testing.T) {
		input := "Hello, Go!"
		expected := "Hello, Go!"
		result, err := ValidateInput(input)

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})

	t.Run("Empty string error", func(t *testing.T) {
		input := ""
		_, err := ValidateInput(input)

		if err == nil {
			t.Fatal("expected an error for empty string input, but got nil")
		}
	})
}
