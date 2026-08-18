package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getIntInput(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readInput(reader, prompt)
		value, err := strconv.Atoi(inputStr)
		if err == nil {
			return value
		}
		fmt.Println("❌ Invalid input! Please enter a valid whole number (integer).")
	}
}

func getFloatInput(reader *bufio.Reader, prompt string) float64 {
	for {
		inputStr := readInput(reader, prompt)
		value, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return value
		}
		fmt.Println("❌ Invalid input! Please enter a valid decimal number (float).")
	}
}

func main() {
  fmt.Println("Hello! Dev sood")
  fmt.Println("2301730121")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== Simple Calculator Program ===")

	for {
		fmt.Println("\nSelect Data Type Scenario:")
		fmt.Println("1. Integer Operations")
		fmt.Println("2. Floating-Point Operations")
		fmt.Println("3. Exit")

		choice := readInput(reader, "Enter choice (1-3): ")

		if choice == "3" {
			fmt.Println("Exiting program. Goodbye!")
			break
		}

		if choice != "1" && choice != "2" {
			fmt.Println("❌ Invalid choice! Please select 1, 2, or 3.")
			continue
		}
		var op string
		for {
			op = readInput(reader, "Enter operation (+, -, *): ")
			if op == "+" || op == "-" || op == "*" {
				break
			}
			fmt.Println("❌ Invalid operation! Please choose from +, -, or *.")
		}
		if choice == "1" {
			fmt.Println("\n--- Integer Mode ---")
			num1 := getIntInput(reader, "Enter first integer: ")
			num2 := getIntInput(reader, "Enter second integer: ")

			var result int
			if op == "+" {
				result = num1 + num2
			} else if op == "-" {
				result = num1 - num2
			} else if op == "*" {
				result = num1 * num2
			}
			fmt.Printf("\n✅ Result: %d %s %d = %d\n", num1, op, num2, result)
		} else if choice == "2" {
			fmt.Println("\n--- Floating-Point Mode ---")
			num1 := getFloatInput(reader, "Enter first decimal: ")
			num2 := getFloatInput(reader, "Enter second decimal: ")

			var result float64
			if op == "+" {
				result = num1 + num2
			} else if op == "-" {
				result = num1 - num2
			} else if op == "*" {
				result = num1 * num2
			}
			fmt.Printf("\n✅ Result: %g %s %g = %g\n", num1, op, num2, result)
		}
	}
}
