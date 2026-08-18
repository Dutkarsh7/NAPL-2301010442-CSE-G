package main

import "fmt"

func main() {
    intA := 12
    intB := 5

    floatA := 12.5
    floatB := 4.2

    fmt.Println("Integer calculations:")
    fmt.Println("Addition:", intA+intB)
    fmt.Println("Subtraction:", intA-intB)
    fmt.Println("Multiplication:", intA*intB)

    fmt.Println()
    fmt.Println("Floating-point calculations:")
    fmt.Println("Addition:", floatA+floatB)
    fmt.Println("Subtraction:", floatA-floatB)
    fmt.Println("Multiplication:", floatA*floatB)
}