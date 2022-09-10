package bjmac

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func StructPlayground() {
	// Inner type gets promoted to outer type
	sa1 := secretAgent{
		person: person{
			first: "Something",
			last:  "Bond",
			age:   33,
		},
		ltk: true,
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	p2 := person{
		first: "Martha",
		last:  "Stewart",
		age:   65,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
}

func AnonymousStructs() {
	ap1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(ap1)

}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func StructInStruct() {
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(t.color)
	fmt.Println(t.doors)
	fmt.Println(t.fourWheel)

	fmt.Println(s)
	fmt.Println(s.color)
	fmt.Println(s.doors)
	fmt.Println(s.luxury)

}
