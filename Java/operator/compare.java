package operator;

import java.util.Objects;

public class Compare {

  public static void main(String[] args) {
    System.out.println("This Compare operator class");
  }

  public static void intCompare() {
    // Integer Comparison
    int a = 1;
    int b = 2;

    // Output
    System.out.println("##################### Integer #####################");
    System.out.println("1 == 2 is: " + (a == b));
    System.out.println("1 != 2 is: " + (a != b));
    System.out.println("1 > 2 is: " + (a > b));
    System.out.println("1 < 2 is: " + (a < b));
    System.out.println("1 >= 2 is: " + (a >= b));
    System.out.println("1 <= 2 is: " + (a <= b));
  }

  public static void decCompare() {
    // Decimal Comparison
    double a = 10.2;
    float b = 22.25f;

    // Output
    System.out.println("##################### Decimal #####################");
    System.out.println("10.2 == 22.25 is: " + (a == b));
    System.out.println("10.2 != 22.25 is: " + (a != b));
    System.out.println("10.2 > 22.25 is: " + (a > b));
    System.out.println("10.2 < 22.25 is: " + (a < b));
    System.out.println("10.2 >= 22.25 is: " + (a >= b));
    System.out.println("10.2 <= 22.25 is: " + (a <= b));
  }

  public static void strCompare() {
    // String compare
    String a = "Hello";
    String b = "hello";

    // Output
    System.out.println("##################### String #####################");
    System.out.println(
        "a == b is: " +
        (Objects.equals(a, b))); // change == to use (Objects.equals)
    System.out.println(
        "a != b is: " +
        (!Objects.equals(
            a, b))); // change != to use !(Objects.equals), it's mean not equals
  }

  public static void charCompare() {
    // Character compare
    char name1 = 'A';
    char name2 = 'a';

    // Output
    System.out.println("##################### Character #####################");
    System.out.println("'A' == 'a' is: " + (name1 == name2));
    System.out.println("'A' != 'a' is: " + (name1 != name2));
  }
}
