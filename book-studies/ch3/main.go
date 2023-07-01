package main

import "fmt"

func main() {
	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))
	//slice()
	//changedSlice()
	//sliceAppend()
	//confusedSlice()
	copySlice()
}

func slice() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func changedSlice() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}

func sliceAppend() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] //ab
	fmt.Println(cap(x), cap(y))
	y = append(y, "z")
	fmt.Println("x:", x) //abcd
	fmt.Println("y:", y) // abz
}

func confusedSlice() {
	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2]
	// z := x[2:]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, "i", "j", "k")
	// x = append(x, "x")
	// z = append(z, "y")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func copySlice() {
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[2:])
	fmt.Println(x, num)
}
