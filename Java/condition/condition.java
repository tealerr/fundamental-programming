package condition;

public class Condition {

  public static void main(String[] args) {
    System.out.println("This condition class");
  }

  //ใช้ if else เพื่อคำนวณคะแนนออกมาเป็นเกรด เช่น A, B, F
  public static char gradeCal(int score) {
    char grade;

    if (score >= 80) {
      grade = 'A';
    } else if (score > 70 && score < 80) {
      grade = 'B';
    } else {
      grade = 'F';
    }
    return grade;
  }

  // if ในรูปแบบ short statement
  public static void isStatus(int status) {
    String result =
        (status == 0) ? "device turn off!!\n" : "device turn on!!\n";
    System.out.println(result);
  }

  //ใช้ switch case เพื่อหาเกรดที่ตรงตามเคส และแสดงผลออกมา
  public static void showGrade(char grade) {
    switch (grade) {
    case 'A':
      System.out.println("Your grade is A\n");
      break;
    case 'B':
      System.out.println("Your grade is B\n");
      break;
    case 'F':
      System.out.println("Your grade is F\n");
      break;
    default:
      System.out.println("Invalid grade\n");
      break;
    }
  }
}
