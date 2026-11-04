package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func modifyValue(ptr *int) {
	// Dereferencing ptr changes the original variable.
	*ptr += 50
	fmt.Printf("Value inside modifyValue(): %d\n", *ptr)
}

func main() {
	num := 100

	fmt.Println("Part A: Referencing and Dereferencing")
	fmt.Printf("Original variable value: %d\n", num)
	fmt.Printf("Memory address using &: %p\n", &num)

	ptr := &num
	fmt.Printf("Pointer address: %p\n", ptr)
	fmt.Printf("Value accessed using *: %d\n", *ptr)

	fmt.Println("\nPart B: Pointer Function")
	fmt.Printf("Value before modification: %d\n", num)
	modifyValue(&num)
	fmt.Printf("Updated original variable value: %d\n", num)

	fmt.Println("\nPart C: Struct Using new()")
	student := new(Student)
	student.Name = "Ananya"
	student.Age = 20
	student.Marks = 87.5

	fmt.Printf("Student details after allocation: Name=%s, Age=%d, Marks=%.2f\n", student.Name, student.Age, student.Marks)

	student.Marks = 92.0
	fmt.Printf("Updated student details: Name=%s, Age=%d, Marks=%.2f\n", student.Name, student.Age, student.Marks)
}
