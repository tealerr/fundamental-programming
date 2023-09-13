// สร้าง function เพื่อคำนวณเกรด โดยใช้ if และ else-if
func gradeCal(score: Int) -> Character {
    var grade: Character

    if score >= 80 {
        grade = "A" //if true, grade =80, false ไปเงื่อนไขต่อไป
    } else if score >= 70 && score < 80 {
        grade = "B"
    } else { //เงื่อนไขสุดท้าย
        grade = "F"
    }

    return grade // ส่งออกเกรดที่ผ่านการเข้าเงื่อนไขแล้ว
}

// แสดงเกรดที่ส่งมาจาก func  พร้อมข้อความ
func showGrade(grade: Character) -> String {
    var message: String

    switch grade {
    case "A":
        message = "Your grade is A"
    case "B":
        message = "Your grade is B"
    case "F":
        message = "Your grade is F"
    default:
        message = "Invalid grade"
    }

    return message
}

func main() {
    let score = gradeCal(score: 72)
    let message = showGrade(grade: score)
    print(message)
}

main()
