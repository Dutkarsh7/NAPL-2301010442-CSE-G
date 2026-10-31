package main

import "fmt"

// isValidIndex reports whether index can be used with the slice.
func isValidIndex(index int, length int) bool {
	return index >= 0 && index < length
}

func main() {
	// Part A: Slice operations.
	students := []string{"Utkarsh", "Rahul", "Aman", "Priya"}

	fmt.Println("=== Part A: Slice Operations ===")
	fmt.Println("Original students:", students)

	// Add a student with append.
	students = append(students, "Neha")
	fmt.Println("After adding Neha:", students)

	// Remove the student at index 1 after validating the index.
	removeIndex := 1
	if isValidIndex(removeIndex, len(students)) {
		students = append(students[:removeIndex], students[removeIndex+1:]...)
		fmt.Printf("After removing index %d: %v\n", removeIndex, students)
	} else {
		fmt.Printf("Cannot remove index %d: index is out of range\n", removeIndex)
	}

	// Update the student at index 1 after validating the index.
	updateIndex := 1
	if isValidIndex(updateIndex, len(students)) {
		students[updateIndex] = "Kavya"
		fmt.Printf("After updating index %d: %v\n", updateIndex, students)
	} else {
		fmt.Printf("Cannot update index %d: index is out of range\n", updateIndex)
	}

	// Part B: Map operations.
	marks := map[string]int{
		"Go":     85,
		"Java":   90,
		"Python": 88,
	}

	fmt.Println("\n=== Part B: Map Operations ===")
	fmt.Println("Original marks:", marks)

	// Insert a new subject and its marks.
	marks["C++"] = 82
	fmt.Println("After inserting C++:", marks)

	// Delete a subject with the built-in delete function.
	delete(marks, "Java")
	fmt.Println("After deleting Java:", marks)

	// Look up an existing subject using the comma-ok pattern.
	if score, exists := marks["Go"]; exists {
		fmt.Printf("Lookup Go: %d marks\n", score)
	} else {
		fmt.Println("Lookup Go: subject not found")
	}

	// Demonstrate the message shown for a missing subject.
	if score, exists := marks["Rust"]; exists {
		fmt.Printf("Lookup Rust: %d marks\n", score)
	} else {
		fmt.Println("Lookup Rust: subject not found")
	}
}
