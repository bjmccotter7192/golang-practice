package main

import (
	"fmt"
)

var testNum = 10

type hotdog int

var hd hotdog

func main() {
	fmt.Println("Testing this out")

	myNum := 15
	myNum2 := 16

	total := myNum * myNum2

	fmt.Println(total)
	fmt.Println(testNum)

	fmt.Printf("Let's use the printf with this number: %d\n", testNum)

	fmt.Printf("%T\n", testNum) //This shows the type of a variable

	hd = 43
	fmt.Printf("%T\n", hd)
}
