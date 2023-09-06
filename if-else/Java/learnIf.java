public class learnIf {
    public static void main(String[] args){
        //Calculate grade
        char score1 = gradeCal(82);
        char score2 = gradeCal(73);
        char score3 = gradeCal(49);

        //Show grade
        showGrade(score1);
        showGrade(score2);
        showGrade(score3);

        //Check device status
        int turnOn = 1;
        int turnOff = 0;
        isStatus(turnOn);
        isStatus(turnOff);

    }

    //ใช้ if else เพื่อคำนวณคะแนนออกมาเป็นเกรด เช่น A, B, F
    public static char gradeCal(int score) {
        char grade;

        if (score>=80){
            grade = 'A';
        }else if(score>70&&score<80){
            grade = 'B';
        }else{
            grade = 'F';
        }
        return grade;
    }

    //variable = (condition) ? expressionTrue :  expressionFalse;
    // if ในรูปแบบ short statement
    public static void isStatus(int status){
        
        String result = (status == 0) ? "device turn off!!\n" : "device turn on!!\n";
        System.out.println(result);    }


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
}
     */

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
