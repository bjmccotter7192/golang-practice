package bjmac

import "fmt"

func ControlFlowFun() {
	fmt.Println("Working with control flows (loops and conditionals)")

	// FOR LOOP
	for i := 0; i <= 100; i++ {
		fmt.Printf("Counting is fun :%d\n", i)
	}

	// Nested Loop
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			fmt.Println("Printing", i, j)
		}
	}
}

func NinjaLevel_3_Exercise_1() {
	for x := 0; x <= 10000; x++ {
		fmt.Println(x)
	}
}

func NinjaLevel_3_Exercise_2() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		fmt.Printf("\t%#U\n", x)
		fmt.Printf("\t%#U\n", x)
		fmt.Printf("\t%#U\n", x)
	}
}

func NinjaLevel_3_Exercise_3() {
	birthYear := 1992
	x := 2022

	for x >= birthYear {
		fmt.Println(x)
		x--
	}
}

func NinjaLevel_3_Exercise_4() {
	birthYear := 1992
	x := 2022

	for {
		if x < birthYear {
			break
		}
		fmt.Println(x)
		x--
	}
}

func NinjaLevel_3_Exercise_5() {
	for x := 10; x <= 100; x++ {
		remainder := x % 4
		fmt.Println(remainder)
	}
}

func NinjaLevel_3_Exercise_6() {
	for x := 10; x <= 100; x++ {
		remainder := x % 4

		if remainder == 0 {
			fmt.Println("Remainder is 0")
		} else if remainder == 1 {
			fmt.Println("Remainder is 1")
		} else if remainder == 2 {
			fmt.Println("Remainder is 2")
		} else {
			fmt.Println("Remainder is 3")
		}
	}
}
