package operators

import "fmt"

func Assignment() {
	var x int

	x = 10
	fmt.Println("x = ", x)

	x = 10
	x += 3
	fmt.Println("x +=3 is ", x)

	x = 10
	x -= 3
	fmt.Println("x -=3 is ", x)

	x = 10
	x *= 3
	fmt.Println("x *= 3 is ", x)

	x = 10
	x /= 3
	fmt.Println("x /= 3 is ", x)

}

func OtherAssign() {
	var x int

	x = 10
	x &= 3
	fmt.Println("x &= 3 is ", x)

	x = 10
	x |= 3
	fmt.Println("x -=3 is ", x)

	x = 10
	x ^= 3
	fmt.Println("x ^= 3 is ", x)

	x = 10
	x >>= 3
	fmt.Println("x >>= 3 is ", x)

	x = 10
	x <<= 3
	fmt.Println("x <<= 3 is ", x)
}
