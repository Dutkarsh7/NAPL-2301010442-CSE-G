package mathutil

import "testing"

func TestStringFunctions(t *testing.T) {
	if got := Reverse("Hello World"); got != "dlroW olleH" {
		t.Errorf("Reverse() = %q, want %q", got, "dlroW olleH")
	}

	if got := CountVowels("Hello World"); got != 3 {
		t.Errorf("CountVowels() = %d, want %d", got, 3)
	}
}

func TestMathFunctions(t *testing.T) {
	if got := Factorial(5); got != 120 {
		t.Errorf("Factorial() = %d, want %d", got, 120)
	}

	if got := Power(2, 8); got != 256 {
		t.Errorf("Power() = %d, want %d", got, 256)
	}
}
