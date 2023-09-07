package operators

import "fmt"

func LogicAnd() {
	a := 2
	b := 3

	fmt.Println("##################### Logic AND #####################")
	if a >= 1 && b <= 2 {
		fmt.Println("Hello World")
	} else {
		fmt.Println("This false")
	}
}

func LogicOr() {
	a := 2
	b := 3

	fmt.Println("##################### Logic OR #####################")
	if a >= 1 || b <= 2 {
		fmt.Println("Hello World")
	} else {
		fmt.Println("This false")
	}
}

func LogicNot() {
	a := 2
	b := 3

	fmt.Println("##################### Logic NOT #####################")
	if a != 2 || b != 3 {
		fmt.Println("Hello World")
	} else {
		fmt.Println("This false")
	}
}
