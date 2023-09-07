package main

import "fmt"

func main() {
	ImplicitDeclare()
	ExplicitDeclare()
}

func ImplicitDeclare() {
	//Implicit declare
	myNum := 112
	myFloat := 1.12
	myChar := 'T'
	myName := "Teeramate"
	myStatus := true

	//Output
	fmt.Println("##################### Implicit declare #####################")
	fmt.Printf("myNum's type is %T and its value is %v\n", myNum, myNum)
	fmt.Printf("myFloat's type is %T and its value is %v\n", myFloat, myFloat)
	fmt.Printf("myChar's type is %T and its value is %v\n", myChar, myChar)
	fmt.Printf("myName's type is %T and its value is %v\n", myName, myName)
	fmt.Printf("myStatus's type is %T and its value is %v\n", myStatus, myStatus)

}

func ExplicitDeclare() {
	//Explicit declare

	var a uint = 500
	var b float32 = 1.12
	var c int = -230
	var myName byte = 'T'
	var Name string = "Teeramate"

	//Output
	fmt.Println("##################### Explicit declare #####################")
	fmt.Printf("a's type is %T and its value is %v\n", a, a)
	fmt.Printf("b's type is %T and its value is %v\n", b, b)
	fmt.Printf("c's type is %T and its value is %v\n", c, c)
	fmt.Printf("Name's type is %T and its value is %v\n", Name, Name)
	fmt.Printf("myName's type is %T and its value is %c\n", myName, myName)

}
