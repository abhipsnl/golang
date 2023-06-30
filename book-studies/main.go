package main

import "fmt"

func main() {

	var x int = 20
	x *= 2
	fmt.Printf(`Hello
	"World"
	`)
	test(x)
	byt()
}

// this is test function
func test(inp int) {
	fmt.Println(inp)
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)
}

func byt() {
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println(sum3, sum4)
}
