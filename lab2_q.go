package main

import (
	"fmt"
	"strings"
)

// Reverse takes a string and returns its characters in reverse order.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountVowels returns the number of vowels (a, e, i, o, u) in a string.
func CountVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// Factorial calculates the factorial of a non-negative integer.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Power returns the result of a base integer raised to an exponent power.
func Power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
} 
func main() {
	fmt.Println("=== Custom Go Package Demonstration ===")
	fmt.Println()
	testStr := "Dev sood"
	fmt.Printf("Original String:      \"%s\"\n", testStr)
	fmt.Printf("Reversed String:      \"%s\"\n", Reverse(testStr))
	fmt.Printf("Number of Vowels:     %d\n", CountVowels(testStr))
	fmt.Println()
	num := 5
	base, exp := 3, 4
	fmt.Printf("Factorial of %d:       %d\n", num, Factorial(num))
	fmt.Printf("%d raised to power %d:  %d\n", base, exp, Power(base, exp))
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("Name:                 Dev sood")
	fmt.Println("Roll No:              2301730121")
	fmt.Println("=======================================")
}
