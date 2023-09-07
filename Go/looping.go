package main

import "fmt"

// func main() {
// 	fmt.Println("####### For loop #######")
// 	forLoop()

// 	fmt.Println("####### While loop #######")
// 	whileLoop()

// 	fmt.Println("####### Range in for loop #######")
// 	rangeLoop()

// 	fmt.Println("####### Use break and continue in loop #######")
// 	breakAndCon()
// }

// This for loop
func forLoop() {
	for i := 1; i < 5; i++ {
		fmt.Println("This For loop ", i)
	}
	fmt.Println("")
}

// this while loop
func whileLoop() {
	num := 1

	for num < 5 {
		fmt.Println("This While loop ", num)
		num++
	}
	fmt.Println("")

}

func breakAndCon() {
	//นับ 1 ถึง 10
	fmt.Println("Count 1 to 10 !")

	// การใช้ continue
	for i := 1; i <= 10; i++ {
		if i == 7 { // ให้เบรคเมื่อนับถึงครั้งที่ 7
			fmt.Println("I'm use break")
			break //ใช้เพื่อให้ออกจาก loop ทันทีโดยไม่สนใจเงื่อนไขของ loop
		}
		fmt.Println(i) // print ไปเรื่อยๆจนครบเงื่อนไข
	}
	fmt.Println("")

	// การใช้ break
	fmt.Println("Count 1 to 5 !")
	i := 1

	//ใช้ while loop in Go
	for i <= 5 {
		if i == 3 { //ข้ามการนับครั้งที่ 3
			fmt.Println("I'm using continue")
			i++
			continue //ใช้เพื่อข้ามการทำงานบางส่วนของลูป
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("")
}

// shouldn't run, it infinity! อย่าหารัน เป็นตัวอย่างเฉยๆ
func forEver() {
	for {
		fmt.Println("Hello World!")
	}
}

// ใช้ range กับ for loop เพื่อเช็คสมาชิกทั้งหมดใน array และ map
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
