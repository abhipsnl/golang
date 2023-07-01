package main

import (
	"fmt"
)

func main() {
	message := "Hi 👩 and 👨"
	rune := []rune(message)
	fmt.Println(string(rune[3]))

}
