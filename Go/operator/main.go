package main

import (
	"basicLoop/operator/operators"
	"fmt"
)

func main() {
	fmt.Println("******** Below is Assignment Operator ********")
	operators.Assignment()

	fmt.Println("############## Other Assignment Operator ###############")
	operators.OtherAssign()

	fmt.Println("******** Below is Arithmetic Operator ********")
	operators.NumberArithmetic()
	operators.StringArithmetic()

	fmt.Println("******** Below is Compare Operator ********")
	operators.IntCompare()
	operators.DecCompare()
	operators.StrCompare()
	operators.CharCompare()

	fmt.Println("******** Below is Logic Operator ********")
	operators.LogicAnd()
	operators.LogicOr()
	operators.LogicNot()

}
