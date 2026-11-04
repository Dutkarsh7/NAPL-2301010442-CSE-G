# Experiment 5: Pointer Manipulation - Referencing and Dereferencing

## Aim

To understand pointers in Go by demonstrating referencing, dereferencing, pointer parameters, modification of an original variable, and dynamic struct allocation using `new()`.

## Software Requirements

- Go 1.26 or later
- Visual Studio Code or any Go-compatible editor
- A terminal or command prompt

## Referencing and Dereferencing

The address operator `&` obtains the memory address of a variable. A pointer stores that address. In this experiment, `ptr := &num` stores the address of `num` in `ptr`.

The dereferencing operator `*` accesses the value stored at the address held by a pointer. Therefore, `*ptr` accesses the same value as `num`. The variable value and the memory address are different: the value is `100`, while the address is printed in hexadecimal and can differ between executions.

## Pointer Function Modification

The `modifyValue()` function accepts an integer pointer of type `*int`. It dereferences the pointer with `*ptr += 50`, so it changes the original variable rather than making a separate copy. The program prints the value before the function call, inside the function, and after the call.

## Struct Allocation Using `new()`

The statement `student := new(Student)` allocates a `Student` struct and returns a pointer to it. The fields are assigned through the pointer. Go automatically allows field access such as `student.Name`; writing `(*student).Name` is not required. The program then modifies the `Marks` field through the pointer.

## Project Structure

```text
lab-03/experiment-05-pointers/
|-- main.go
|-- go.mod
`-- README.md
```

## Steps to Execute

From the `lab-03/experiment-05-pointers` directory, run:

```text
go run main.go
```

To run the module tests:

```text
go test ./...
```

There are no test files in this experiment, but the command verifies that the module builds successfully.

## Sample Output

Memory addresses are expected to differ between executions.

```text
Part A: Referencing and Dereferencing
Original variable value: 100
Memory address using &: 0xc0000120d0
Pointer address: 0xc0000120d0
Value accessed using *: 100

Part B: Pointer Function
Value before modification: 100
Value inside modifyValue(): 150
Updated original variable value: 150

Part C: Struct Using new()
Student details after allocation: Name=Ananya, Age=20, Marks=87.50
Updated student details: Name=Ananya, Age=20, Marks=92.00
```
