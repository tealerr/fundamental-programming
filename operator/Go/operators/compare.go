package operators

import "fmt"

func IntCompare() {
	a := 1
	b := 2

	fmt.Println("##################### Integer #####################")
	fmt.Println("a == b is: ", (a == b))
	fmt.Println("a != b is: ", (a != b))
	fmt.Println("a > b is: ", (a > b))
	fmt.Println("a < b is: ", (a < b))
	fmt.Println("a >= b is: ", (a >= b))
	fmt.Println("a <= b is: ", (a <= b))

}

func DecCompare() {
	var a = 10.2
	var b = 10.232345

	fmt.Println("##################### Decimal #####################")
	fmt.Printf("a == b is: %v\n", a == b)
	fmt.Printf("a != b is: %v\n", a != b)
	fmt.Printf("a > b is: %v\n", a > b)
	fmt.Printf("a < b is: %v\n", a < b)
	fmt.Printf("a >= b is: %v\n", a >= b)
	fmt.Printf("a <= b is: %v\n", a <= b)

}

func StrCompare() {
	a := "Hello"
	b := "hello"

	fmt.Println("#####################String#####################")
	fmt.Printf("a == b is: %v\n", a == b)
	fmt.Printf("a != b is: %v\n", a != b)

}

func CharCompare() {
	name1 := 'A'
	name2 := 'a'

	fmt.Println("#####################Character#####################")
	fmt.Printf("name1 == name2 is: %v\n", name1 == name2)
	fmt.Printf("name1 != name2 is: %v\n", name1 != name2)

}
