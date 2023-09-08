#include <iostream>
using namespace std;

// สร้าง if เพื่อคำนวณเกรด
char GradeCalWithIf(int score)
{
    char grade;
    if (score >= 80)
    {
        grade = 'A';
    }
    else
    {
        grade = 'F';
    }
    return grade;
}

// ใช้ else-if เมื่อมีหลายเงื่อนไข
void GradeCalWithElseIf(int score)
{
    char grade;
    if (score >= 80)
    {
        grade = 'A';
        std::cout << grade << std::endl;
    }
    else if (score >= 70 && score < 80) // ใช้ else if เมื่อมีเงื่อนไขต่อไปเรื่อยๆ
    {
        grade = 'B';
        std::cout << grade << std::endl;
    }
    else if (score >= 60 && score < 70)
    {
        grade = 'C';
        std::cout << grade << std::endl;
    }
    else
    { // ถ้าเป็นเงื่อนไขสุดท้าย ใช้ else
        grade = 'F';
        std::cout << grade << std::endl;
    }
}

// variable = (condition) ? expressionTrue : expressionFalse;
// เขียนในรูปแบบ short statement ในกรณีที่เงื่อนไขไม่ซับซ้อน basic
void ShortHandIf(int status)
{
    enum Status
    {
        TURN_ON = 1,
        TURN_OFF = 0
    };
    string result = (status == 0) ? "device turn off!!\n" : "device turn on!!\n";
    cout << result;
}

// แสดงข้อความพร้อมเกรด โดยใช้ switch case
void ShowGrade(char grade)
{
    switch (grade)
    {
    case 'A':
        cout << "Your grade is A\n"
             << endl;
        break;
    case 'B':
        cout << "Your grade is B\n"
             << endl;
        break;
    case 'C':
        cout << "Your grade is C\n"
             << endl;
        break;
    case 'F':
        cout << "Your grade is F\n"
             << endl;
        break;
    default:
        cout << "Invalid grade\n"
             << endl;
        break;
    }
}
