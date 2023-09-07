package condition

// สร้าง function เพื่อคำนวณเกรด
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

func IsActive(num int) string {
	if num == 1 {
		return "Device is On"
	}
	return "Device is Off"
}
