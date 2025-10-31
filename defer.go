package main

import "fmt"

func main() {
	defer fmt.Println("This is the first deferred statement.")
	defer fmt.Println("This is the second deferred statement.")
	fmt.Println("This is the main function.")
}
