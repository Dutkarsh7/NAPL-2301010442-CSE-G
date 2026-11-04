package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

var scanner = bufio.NewScanner(os.Stdin)

func readLine() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func (person *Person) ReadData() {
	fmt.Print("Enter name: ")
	person.Name = readLine()

	for {
		fmt.Print("Enter age: ")
		age, err := strconv.Atoi(readLine())
		if err == nil {
			person.Age = age
			break
		}
		fmt.Println("Please enter a valid whole number for age.")
	}

	fmt.Print("Enter job title: ")
	person.Job = readLine()

	for {
		fmt.Print("Enter salary: ")
		salary, err := strconv.ParseFloat(readLine(), 64)
		if err == nil {
			person.Salary = salary
			break
		}
		fmt.Println("Please enter a valid number for salary.")
	}
}

func (person Person) DisplayData() {
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Job: %s\n", person.Job)
	fmt.Printf("Salary: %.2f\n", person.Salary)
}

func main() {
	var firstPerson Person
	fmt.Println("\nEnter details for Person 1")
	firstPerson.ReadData()

	var secondPerson Person
	fmt.Println("\nEnter details for Person 2")
	secondPerson.ReadData()

	fmt.Println("\n===== Person 1 Details =====")
	firstPerson.DisplayData()

	fmt.Println("\n===== Person 2 Details =====")
	secondPerson.DisplayData()
}
