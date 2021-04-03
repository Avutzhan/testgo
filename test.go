package main

import "fmt"

var golang, angular, vuejs bool
var react string = "react"
var empty_balance [10] float32
var full_balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

type Vertex struct {
    X int
    Y int
}

func main() {
	y := "test y"
	golang = true
	var x int = 10
	fmt.Printf("Hello World\n")
	fmt.Println(react)
	fmt.Println("Sum is ", add(10, 20))
	fmt.Println(x, golang, angular, vuejs, y)
	fmt.Println("loop", addLoop())
	fmt.Println("empty balance", empty_balance)
	fmt.Println("full_balance", full_balance)

	var a [2]string
    a[0] = "Привет"
    a[1] = "Tproger"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    var s []int = primes[1:4]
    fmt.Println("s", s)

//     var numbers []int
//     numbers = make([]int,5,5)

    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}

func add(a int, b int) int {
	return a + b;
}

func addLoop() int {
    sum := 0
    for i := 0; i < 8; i++ {
        if i % 2 == 0 {
            sum += 1
        }

    }

    return sum
}
