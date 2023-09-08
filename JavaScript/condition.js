//สร้าง function เพื่อคำนวณเกรด โดยใช้ if else-if และ else
function gradeCalculate(score) { // รับ parameter เป็นคะแนน(number)
    let grade;
    if (score >= 80) { //ใช้ if เงื่อนไขแรก
        grade = 'A';
    } else if (score >= 70 && score < 80) { // ใช้ else if ถ้ายังมีเงื่อนไขต่อไปเรื่อยๆ
        grade = 'B';
    } else { //ใช้ else เมื่อเป็นเงื่อนไขสุดท้าย
        grade = 'F';
    }
    return grade;
}

// function นี้ใช้โชว์เกรดออกมาเป็นข้อความ โดยใช้ switch case
function showGrade(grade) {
    let message;
    switch (grade) {
        case 'A':
            message = "Your grade is A";
            break;
        case 'B':
            message = "Your grade is B";
            break;
        case 'F':
            message = "Your grade is F"
            break;
        default:
            message = "Invalid grade";
            break;
    }

    console.log(message);
}

function main(){
    let studen1 = gradeCalculate(62); // ใส่คะแนนเพื่อคิดเกรด
    showGrade(studen1); //รับเกรดจาก score1 เพื่อแสดงผล
}

main(); // run code ใน function main