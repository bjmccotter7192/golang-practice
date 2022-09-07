package bjmac

import "fmt"

func NinjaLevel_2_Exercise_1() {
	x := 42
	fmt.Printf("%d, %b, %#x\n", x, x, x)
}

func NinjaLevel_2_Exercise_2() {
	a := 42
	b := 74
	eq2 := a == b
	lteq2 := a <= b
	gteq2 := a >= b
	neq2 := a != b
	lt := a < b
	gt := a > b

	fmt.Printf("%d is equal to %d? %t\n", a, b, eq2)
	fmt.Printf("%d is less than or equal to %d? %t\n", a, b, lteq2)
	fmt.Printf("%d is greater than or equal to %d? %t\n", a, b, gteq2)
	fmt.Printf("%d is not equal to %d? %t\n", a, b, neq2)
	fmt.Printf("%d is less than to %d? %t\n", a, b, lt)
	fmt.Printf("%d is greater than to %d? %t\n", a, b, gt)

}

const (
	INT_CONSTANT           = 42
	STRING_CONSTANT string = "James Bond"
)

func NinjaLevel_2_Exercise_3() {
	fmt.Printf("My number is %d and my name is %s\n", INT_CONSTANT, STRING_CONSTANT)
}

func NinjaLevel_2_Exercise_4() {
	x := 42
	fmt.Printf("%d, %b, %#x", x, x, x)
	fmt.Println()
	shiftedX := x << 1
	fmt.Printf("%d, %b, %#x", shiftedX, shiftedX, shiftedX)
	fmt.Println()
}

func NinjaLevel_2_Exercise_5() {
	rawString := `This is cool
	
	multi level string
	
	using literals`
	fmt.Println(rawString)
}

const (
	_ = iota + 2022
	year1
	year2
	year3
	year4
)

func NinjaLevel_2_Exercise_6() {
	fmt.Println(year1, year2, year3, year4)
}
