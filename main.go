package main

import (
	"fmt"
	"golang-practice/bjmac"
)

func main() {
	/*
		This section covers Variables, values, and type
	*/
	bjmac.LearnTheBasics() //Just playing around with the basics
	bjmac.NinjaExercise1() // First hands on example of Ninja/Jedi exercises
	bjmac.NinjaExercise2() // Second hands on example of Ninja/Jedi exercies
	bjmac.NinjaExercise3() // Third hands on example of Ninja/Jedi exercies
	bjmac.NinjaExercise4() // Forth hands on example of Ninja/Jedi exercies

	fmt.Println("================================================")

	/*
		This section covers programming fundamentals
	*/
	bjmac.BoolTypePractice()
	bjmac.WorkingWithInts()
	bjmac.RuntimePlayground()
	bjmac.PlayingWithStrings()
	bjmac.NumeralSystemsPlayground()
	bjmac.BitShiftingFun()

	/*
		NINJA Level 2 Hands On Exercises
	*/
	bjmac.NinjaLevel2_Exercise1()
}
