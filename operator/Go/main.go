package main

import (
	"fmt"
	"goBasic/operators"
)

func main() {
	fmt.Println("******** Below is Compare Operator ********")
	operators.IntCompare()
	operators.DecCompare()
	operators.StrCompare()
	operators.CharCompare()

	fmt.Println("******** Below is Logic Operator ********")
	operators.LogicOr()
	operators.LogicAnd()
	operators.LogicNot()

}
