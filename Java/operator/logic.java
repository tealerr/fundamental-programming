package operator;

public class Logic {

  public static void main(String[] args) {
    System.out.println("This Logic operator class");
  }

  public static void logicAnd() {
    // Logic AND
    int a = 2;
    int b = 3;

    if (a >= 1 && b <= 2) {
      System.out.println("Hello World!");
    } else {
      System.out.println("This false!");
    }
  }

  public static void logicOr() {
    // Logic AND
    int a = 2;
    int b = 3;

    if (a >= 1 || b <= 2) {
      System.out.println("Hello World");
    } else {
      System.out.println("This false");
    }
  }

  public static void logicNot() {
    // Logic AND
    int a = 2;
    int b = 3;

    if (a != 2 || b != 3) {
      System.out.println("Hello World");
    } else {
      System.out.println("This false");
    }
  }
}
