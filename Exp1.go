package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	intMin   = -1000000
	intMax   = 1000000
	floatMin = -1000000.0
	floatMax = 1000000.0
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Simple Calculator")
		fmt.Println("1. Integer calculations")
		fmt.Println("2. Floating-point calculations")
		fmt.Println("3. Exit")

		choice := readMenuChoice(reader)
		if choice == 3 {
			fmt.Println("Goodbye.")
			return
		}

		if choice == 1 {
			first := readIntInRange(reader, "Enter the first integer (-1000000 to 1000000): ")
			second := readIntInRange(reader, "Enter the second integer (-1000000 to 1000000): ")

			fmt.Println("Integer results:")
			fmt.Println("Addition:", first+second)
			fmt.Println("Subtraction:", first-second)
			fmt.Println("Multiplication:", first*second)
		} else if choice == 2 {
			first := readFloatInRange(reader, "Enter the first float (-1000000 to 1000000): ")
			second := readFloatInRange(reader, "Enter the second float (-1000000 to 1000000): ")

			fmt.Println("Floating-point results:")
			fmt.Println("Addition:", first+second)
			fmt.Println("Subtraction:", first-second)
			fmt.Println("Multiplication:", first*second)
		}

		fmt.Println()
	}
}

func readMenuChoice(reader *bufio.Reader) int {
	for {
		fmt.Print("Choose an option (1-3): ")
		text := readLine(reader)
		choice, err := strconv.Atoi(text)
		if err == nil && choice >= 1 && choice <= 3 {
			return choice
		}
		fmt.Println("Invalid choice. Enter 1, 2, or 3.")
	}
}

func readIntInRange(reader *bufio.Reader, prompt string) int64 {
	for {
		fmt.Print(prompt)
		text := readLine(reader)
		value, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Enter a numeric integer value.")
			continue
		}
		if value < intMin || value > intMax {
			fmt.Printf("Value out of range. Enter a value between %d and %d.\n", intMin, intMax)
			continue
		}
		return value
	}
}

func readFloatInRange(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		text := readLine(reader)
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Invalid input. Enter a numeric floating-point value.")
			continue
		}
		if value < floatMin || value > floatMax {
			fmt.Printf("Value out of range. Enter a value between %.0f and %.0f.\n", floatMin, floatMax)
			continue
		}
		return value
	}
}

func readLine(reader *bufio.Reader) string {
	text, err := reader.ReadString('\n')
	if err != nil && len(text) == 0 {
		return ""
	}
	return strings.TrimSpace(text)
}
