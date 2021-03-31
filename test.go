package main

import "fmt"

func main() {
	fmt.Printf("Hello World\n")
	fmt.Println("Sum is ", add(10, 20))
}

func add(a int, b int) int {
	return a + b;
}
