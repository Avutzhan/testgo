package main

import "fmt"

var golang, angular, vuejs bool
var react string = "react"

func main() {
	y := "test y"
	golang = true
	var x int = 10
	fmt.Printf("Hello World\n")
	fmt.Println(react)
	fmt.Println("Sum is ", add(10, 20))
	fmt.Println(x, golang, angular, vuejs, y)
}

func add(a int, b int) int {
	return a + b;
}
