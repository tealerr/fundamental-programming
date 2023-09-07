package condition

// สร้าง function เพื่อคำนวณเกรด
// ใช้ if, else if โดยที่เกรด A ต้อง 80 ขึ้นไป, เกรด B ได้ระหว่าง 70-79, ต่ำกว่านั้น F
func CalculateGrade(score int) rune {
	var grade rune

	if score >= 80 {
		grade = 'A'
	} else if score > 70 && score < 80 {
		grade = 'B'
	} else {
		grade = 'F'
	}
	return grade
}

// แสดงเกรด โดยใช้ switch case
// รับอักขระของเกรดเข้ามา และ return ข้อความออกไป
func ShowGrade(grade rune) string {
	switch grade {
	case 'A':
		return "Your grade is A"
	case 'B':
		return "Your grade is B"
	default:
		return "Your grade is F"
	}

}

// เขียน if-else ในรูปแบบ short-statement
// เป็น function ที่ใช้เพื่อเปิดปิดอุปกรณ์
func IsActive(num int) string {
	if num == 1 { // 1 แทนการเปิดเครื่อง
		return "Device is On" //ถ้าเงื่อนไขเป็จริง จะเข้าการทำงานใน block นี้
	}
	return "Device is Off" // ถ้าเป็น false จะทำงานตรงนี้
}
