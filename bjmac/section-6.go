package bjmac

import (
	"fmt"
	"runtime"
)

var x bool
var intX int
var float64Y float64

var testingVar rune

func BoolTypePractice() {
	fmt.Println("Working with booleans")

	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42

	fmt.Println(a == b) // False
	fmt.Println(a != b) // True
	fmt.Println(a >= b) // False
	fmt.Println(a <= b) // True
}

func WorkingWithInts() {
	fmt.Println("Ints, floats, numeric types")

	intX = 42
	float64Y = 42.34323

	fmt.Println(intX)
	fmt.Println(float64Y)

	fmt.Printf("%T\n", intX)
	fmt.Printf("%T\n", float64Y)
	fmt.Printf("%T\n", testingVar)
}

func RuntimePlayground() {
	fmt.Println("Using the runtime package")

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func PlayingWithStrings() {
	fmt.Println("Working with strings")

	s := "Hello, playgroud"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Strings can be converted to []byte / []uint8
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	//Unicode?
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Print("\n")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %x\n", i, v)
	}
}

const a = 42
const b = 42.78
const c = "James Bond"

const (
	d = iota
	e = iota
	f = iota
)

const (
	g = iota
	h
	j
)

func NumeralSystemsPlayground() {
	fmt.Println("Working with Constants")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(j)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T\n", h)
	fmt.Printf("%T\n", j)
}

// This is the Go idomatic way using iota
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func BitShiftingFun() {
	fmt.Println("Working with Bit Shifting")

	// This is easier to read and == the same as iota example
	// kb := 1024
	// mb := kb * 1024
	// gb := 1024 * mb

	x := 4
	fmt.Printf("%d\t\t%b\t", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
