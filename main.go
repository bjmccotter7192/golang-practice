package main

import (
	"fmt"
	"golang-practice/bjmac"
)

func main() {
	/*
		Section 1 - 5 Basics
		Ninja Level 1
	*/
	bjmac.LearnTheBasics() //Just playing around with the basics
	bjmac.NinjaExercise1() // First hands on example of Ninja/Jedi exercises
	bjmac.NinjaExercise2() // Second hands on example of Ninja/Jedi exercies
	bjmac.NinjaExercise3() // Third hands on example of Ninja/Jedi exercies
	bjmac.NinjaExercise4() // Forth hands on example of Ninja/Jedi exercies

	fmt.Println("================================================")

	/*
		Section 6 - 7 Programming Fundamentals
		Ninja Level 2
	*/
	bjmac.BoolTypePractice()
	bjmac.WorkingWithInts()
	bjmac.RuntimePlayground()
	bjmac.PlayingWithStrings()
	bjmac.NumeralSystemsPlayground()
	bjmac.BitShiftingFun()

	bjmac.NinjaLevel_2_Exercise_1()
	bjmac.NinjaLevel_2_Exercise_2()
	bjmac.NinjaLevel_2_Exercise_3()
	bjmac.NinjaLevel_2_Exercise_4()
	bjmac.NinjaLevel_2_Exercise_5()
	bjmac.NinjaLevel_2_Exercise_6()

	fmt.Println("================================================")

	/*
		Section 8 - 9 Control Flow
		Ninja Level 3
	*/
	bjmac.ControlFlowFun()
	bjmac.NinjaLevel_3_Exercise_1()
	bjmac.NinjaLevel_3_Exercise_2()
	bjmac.NinjaLevel_3_Exercise_3()
	bjmac.NinjaLevel_3_Exercise_4()
	bjmac.NinjaLevel_3_Exercise_5()
	bjmac.NinjaLevel_3_Exercise_6()
	bjmac.NinjaLevel_3_Exercise_7()

	fmt.Println("================================================")

	/*
		Section 10 - 11 Grouping Data
		Ninja Level 4
	*/
	bjmac.ArrayFun()
	bjmac.SliceFun()
	bjmac.SlicingASlice()
	bjmac.AppendingSlices()
	bjmac.DeletingFromSlice()
	bjmac.SliceUsingMake()
	bjmac.MultiSliceSlice()
	bjmac.MapFun()
	bjmac.AddToMap()
	bjmac.DeleteEntryFromMap()

	fmt.Println("================================================")

	/*
		Section 12 - 13 Structs
		Ninja Level 5
	*/
	bjmac.StructPlayground()
	bjmac.AnonymousStructs()
	bjmac.StructInStruct()
}
