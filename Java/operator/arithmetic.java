package operator;

public class Arithmetic {

  public static void main(String[] args) {
    System.out.println("This Arithmetic operator class");
  }

  public static void numberArithmetic() {
    int sum = 5 + 3;
    double total = 10.5 + 2.5;
    System.out.println("sum value is: " + sum);
    System.out.println("total value is: " + total);
  }

  public static void textArithmetic() {
    char b = 'a' + 1;

    // String
    String firstName = "John";
    String lastName = "Doe";
    String fullName = firstName + " " + lastName;

    // Output
    System.out.println("b value is: " + b);
    System.out.println("fullName is: " + fullName);
  }
}
