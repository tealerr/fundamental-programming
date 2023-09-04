package main

import "fmt"

func main() {
	fmt.Println("####### For loop #######")
	forLoop()

	fmt.Println("####### While loop #######")
	whileLoop()

	fmt.Println("####### Range in for loop #######")
	rangeLoop()
}

//This for loop
func forLoop() {
	for i := 1; i < 5; i++ {
		fmt.Println("This For loop ", i)
	}
	fmt.Println("")
}

//this while loop
func whileLoop() {
	num := 1

	for num < 5 {
		fmt.Println("This While loop ", num)
		num++
	}
	fmt.Println("")

}

//shouldn't run, it infinity! อย่าหารัน เป็นตัวอย่างเฉยๆ
func infLoop() {
	for {
		fmt.Println("Hello World!")
	}
}

//ใช้ range กับ for loop เพื่อเช็คสมาชิกทั้งหมดใน array และ map
func rangeLoop() {
	fmt.Println("******** For range Array ********")
	// สร้าง slice ของ string
	strings := []string{"A", "B", "C", "D"}

	fmt.Println("******** String slice ********")
	for index, value := range strings {
		fmt.Println("index ", index, "value ", value)
	}

	fmt.Println("")
	// สร้าง slice ของตัวเลข
	numbers := [5]int{1, 2, 3, 4, 5}

	// ใช้ range เพื่อวนลูปของ array number
	// index ใช้เก็บตำแหน่งของ index
	// value ใช้เก็บค่าประจำ index
	fmt.Println("******** Number array ********")
	for index, value := range numbers {
		fmt.Println("index ", index, "value ", value)
	}
	fmt.Println("")

	fmt.Println("******** For range Map ********")

	myMap := map[string]string{
		"JP": "Japan",
		"US": "United States",
		"TH": "Thailand",
	}
	for key, value := range myMap {
		fmt.Println("key = ", key, " value = ", value)
	}
}
