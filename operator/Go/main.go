package main

import (
	"fmt"
	"goBasic/operators"
)

func main() {
	fmt.Println("******** Below is Assignment Operator ********")
	operators.Assignment()

	fmt.Println("############## Other Assignment Operator ###############")
	operators.OtherAssign()

	fmt.Println("******** Below is Arithmetic Operator ********")
	operators.NumberArithmetic()
	operators.StringArithmetic()
	operators.CharArithmetic()

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
