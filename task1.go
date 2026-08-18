package main

import "fmt"

func main() {
	fmt.Println("Go environment is successfully verified!")
	fmt.Println("--- Hardcoded Calculator ---")
  fmt.Println("Hello! Dev sood")
  fmt.Println("2301730121")
  fmt.Println("Interger operations ")
	var intA int = 15
	var intB int = 4
	intAdd := intA + intB
	intSub := intA - intB
	intMul := intA * intB
	fmt.Printf("Integers: %d and %d\n", intA, intB)
	fmt.Printf("Addition: %d\n", intAdd)
	fmt.Printf("Subtraction: %d\n", intSub)
	fmt.Printf("Multiplication: %d\n\n", intMul)
  fmt.Printf("Float operations\n")
	var floatA float64 = 12.5
	var floatB float64 = 2.5
	floatAdd := floatA + floatB
	floatSub := floatA - floatB
	floatMul := floatA * floatB
	fmt.Printf("Floats: %.2f and %.2f\n", floatA, floatB)
	fmt.Printf("Addition: %.2f\n", floatAdd)
	fmt.Printf("Subtraction: %.2f\n", floatSub)
	fmt.Printf("Multiplication: %.2f\n", floatMul)
}
