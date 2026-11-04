# Experiment 3: Collection Management Using Arrays, Slices, and Maps

## Aim

To demonstrate collection management in Go using slices and maps.

## Requirements

- Go 1.26 or a compatible newer Go version
- A terminal or command prompt

## Slice Operations

The program creates a slice of student names and demonstrates:

1. **Display:** Prints the original slice.
2. **Add:** Uses `append` to add a new student.
3. **Remove:** Removes the element at a selected index by joining the slice before and after that index.
4. **Update:** Replaces the value at a selected index.

The remove and update operations validate their indexes before accessing the slice, preventing an out-of-range runtime error.

## Map Operations

The program creates a map of subject names and marks and demonstrates:

1. **Insert:** Adds a new subject and its marks.
2. **Delete:** Uses Go's built-in `delete` function to remove a subject.
3. **Lookup:** Uses the comma-ok pattern to check whether a subject exists before displaying its marks.
4. **Missing subject:** Displays a clear message when a requested subject is not present.

## How to Execute

From the `lab-02/experiment-03-collections` directory, run:

```text
go run main.go
```

To run the module tests and compile-check all packages, run:

```text
go test ./...
```

## Sample Output

Map entries may appear in a different order because Go map iteration order is not guaranteed.

```text
=== Part A: Slice Operations ===
Original students: [Utkarsh Rahul Aman Priya]
After adding Neha: [Utkarsh Rahul Aman Priya Neha]
After removing index 1: [Utkarsh Aman Priya Neha]
After updating index 1: [Utkarsh Kavya Priya Neha]

=== Part B: Map Operations ===
Original marks: map[Go:85 Java:90 Python:88]
After inserting C++: map[C++:82 Go:85 Java:90 Python:88]
After deleting Java: map[C++:82 Go:85 Python:88]
Lookup Go: 85 marks
Lookup Rust: subject not found
```