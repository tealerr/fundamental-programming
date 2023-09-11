package operators

import "fmt"

func LogicAnd() {
	a := 2
	b := 3

	fmt.Println("##################### Logic AND #####################")
	if a >= 1 && b <= 2 {
		fmt.Println("True AND logic")
	} else {
		fmt.Println("False AND logic")
	}
}

func LogicOr() {
	a := 2
	b := 3

	fmt.Println("##################### Logic OR #####################")
	if a >= 1 || b <= 2 {
		fmt.Println("True OR logic")
	} else {
		fmt.Println("False OR logic")
	}
}

func LogicNot() {
	a := 2
	b := 3

	fmt.Println("##################### Logic NOT #####################")
	if a != 2 || b != 3 {
		fmt.Println("True NOT logic")
	} else {
		fmt.Println("False NOT logic")
	}
}
