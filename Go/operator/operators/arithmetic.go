package operators

import "fmt"

func NumberArithmetic() {
	a, b := 5, 3

	x, y := 5.5, 2.55

	fmt.Println("Integer Addition value is: ", a+b)
	fmt.Println("Float Addition value is: ", x+y)
	fmt.Println("Integer Subtraction value is: ", a-b)
	fmt.Println("Float Subtraction value is: ", x-y)
	fmt.Println("Integer Multiplication value is: ", a*b)
	fmt.Println("Float Multiplication value is: ", x*y)
	fmt.Println("Integer Division value is: ", a/b)
	fmt.Println("Float Division value is: ", x/y)
	fmt.Println("Integer Modulus value is: ", a%b)

}

func StringArithmetic() {
	firstName, lastName := "Teeramate", "Kantima"
	fullname := firstName + " " + lastName

	fmt.Println("My fullname  is: ", fullname)

}
