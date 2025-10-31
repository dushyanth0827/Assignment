package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func main() {
	result := add(3, 4)
	result1 := subtract(10, 20)
	fmt.Println("The sum is:", result)
	fmt.Println("The sum is:", result1)
}
