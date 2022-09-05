package main

import (
	"fmt"
)

var testNum = 10

type hotdog int

var hd hotdog

func main() {
	learnTheBasics() //Just playing around with the basics

	ninjaExercise1() // First hand on example of Ninja/Jedi exercises
}

func learnTheBasics() {
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

	testNum = int(hd) // Converting variables just use T(val)
}

func ninjaExercise1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%s is %d years old, is that true? %t\n", y, x, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
