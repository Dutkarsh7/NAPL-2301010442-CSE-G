package mathutil

// CountVowels counts the number of vowels in a string.
func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, character := range s {
		for _, vowel := range vowels {
			if character == vowel {
				count++
				break
			}
		}
	}
	return count
}
