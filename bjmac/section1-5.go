package bjmac

import "fmt"

type hotdog int
type bjmac int

var hd hotdog
var xBjMac bjmac

var testNum = 10
var xInt int
var yString string
var zBool bool

func LearnTheBasics() {
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

func NinjaExercise1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%s is %d years old, is that true? %t\n", y, x, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func NinjaExercise2() {
	fmt.Println(xInt, yString, zBool)
}

func NinjaExercise3() {
	xInt = 42
	yString = "James Bond"
	zBool = true

	s := fmt.Sprintf("%v\t%s\t%t", xInt, yString, zBool)
	fmt.Println(s)
}

func NinjaExercise4() {
	fmt.Println(xBjMac)
	fmt.Printf("%T\n", xBjMac)
	xBjMac = 42
	fmt.Println(xBjMac)
}
