package main

import "fmt"

func main() {
	p := new(int)
	*p = 42
	fmt.Println("The value of p is:", *p)
}
