#include <iostream>
using namespace std;

char useIfElse(int score);
void useElseIf(int score);
void shortHandIf(int status);
void switchGrade(char grade);

int main()
{
    // calculate grade and return grade
    char student1 = useIfElse(87); // A
    char student2 = useIfElse(79); // F

    // check device status
    shortHandIf(0); // turn off device

    // show grade output
    switchGrade(student1);
    switchGrade(student2);
    return 0;
}

// สร้าง if เพื่อคำนวณเกรด
char useIfElse(int score)
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

// ใช้ else-if เมื่อมีหลายเงื่อน
void useElseIf(int score)
{
    char grade;
    if (score >= 80)
    {
        grade = 'A';
        std::cout << grade << std::endl;
    }
    else if (score >= 70 && score < 80) // ใช้ else if เมื่อมีเงื่อนไขต่อไปเรื่อยๆ
    {
        grade = 'B+';
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
void shortHandIf(int status)
{

    enum Status
    {
        TURN_ON = 1,
        TURN_OFF = 0
    };
    string result = (status == 0) ? "device turn off!!\n" : "device turn on!!\n";
    cout << result;
}

/*
switch(expression) {
  case x:
    // code block
    break;
  case y:
    // code block
    break;
  default:
    // code block
}*/

// แสดงข้อความพร้อมเกรด โดยใช้ switch case
void switchGrade(char grade)
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
