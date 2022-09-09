package bjmac

import "fmt"

func ArrayFun() {
	// GO SAYS DONT USE ARRAYS, USE SLICES
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
}

func SliceFun() {
	// Composite Literals
	mySlice := []int{4, 5, 6, 7, 3} // Group together items of same type

	// Loop over slice
	// len = length
	for index, val := range mySlice {
		fmt.Println(index, val)
	}
}

func SlicingASlice() {
	x := []int{5, 6, 2, 3, 4}

	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	// Looping without range
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func AppendingSlices() {
	x := []int{4, 5, 6, 4, 3, 4, 5, 3}
	fmt.Println(x)
	x = append(x, 55, 44, 33, 22, 33)
	fmt.Println(x)

	// NOTE: ...T means unlimited amount of function arguments
	// T... places all elements from slice
	y := []int{234, 432, 453, 849}
	x = append(x, y...)
	fmt.Println(x)
}

func DeletingFromSlice() {
	x := []int{4, 5, 6, 4, 3, 4, 5, 3}

	fmt.Println(x)

	// Delete is just using append and slicing to remake a new slice
	// T... used because int[] != int so unpack all elements from sliced slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func SliceUsingMake() {
	// Make creates an array of a determined size
	// slices sit ontop of arrays
	// make([]T, length, capacity)
	x := make([]int, 10, 12)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 434

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// Setting the max to 12
	// Once it reaches over the capacity it will create a new slice
	// 2 times the size

	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

func MultiSliceSlice() {
	x := []string{"James", "Bond", "Martini"}
	y := []string{"Margret", "Tony", "Scotch"}
	fmt.Println(x)
	fmt.Println(y)

	multiSlice := [][]string{x, y}
	fmt.Println(multiSlice)

	for i, xs := range multiSlice {
		for j, v := range xs {
			fmt.Println(i, j, v)
		}
	}
}

func MapFun() {
	m := map[string]int{
		"James": 42,
		"Bondy": 33,
		"Helen": 24,
	}
	fmt.Println(m)
	fmt.Println(m["James"]) // Returns 42

	fmt.Println(m["Barnabas"]) // Returns 0 since it doesn't exist

	v, ok := m["Barnabas"] // This will check if the key exists
	fmt.Println(v)
	fmt.Println(ok)

	if j, ok := m["Bondy"]; ok {
		fmt.Println("THIS IS THE IF PRINT", j)
	}
}

func AddToMap() {
	m := map[string]int{
		"James": 42,
		"Bondy": 33,
		"Helen": 24,
	}
	fmt.Println(m)

	m["Beej"] = 55 // THIS WILL ADD ELEMENT TO MAP

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 43}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}

func DeleteEntryFromMap() {
	m := map[string]int{
		"James": 42,
		"Bondy": 33,
		"Helen": 24,
	}
	fmt.Println(m)

	delete(m, "James")

	fmt.Println(m)

	delete(m, "Some Random") // THIS WILL NOT ERROR OUT

	fmt.Println(m)

	if v, ok := m["Bondy"]; ok {
		fmt.Println("value:", v)
		delete(m, "Bondy")
	}

	fmt.Println(m)
}
