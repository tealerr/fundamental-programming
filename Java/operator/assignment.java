package operator;

public class Assignment {

  public static void main(String[] args) {
    System.out.println("This Assignment operator class");
  }

  public static void primaryAssign() {
    int x;

    x = 10;
    System.out.println("x = " + x);

    x += 3;
    System.out.println("x +=3 is " + x);

    x = 10;
    x -= 3;
    System.out.println("x -=3 is " + x);

    x = 10;
    x *= 3;
    System.out.println("x *=3 is " + x);

    x = 10;
    x /= 3;
    System.out.println("x /=3 is " + x);
  }

  public static void otherAssign() {
    int x;

    x = 10;
    x &= 3;
    System.out.println("x &=3 is " + x);

    x = 10;
    x |= 3;
    System.out.println("x |=3 is " + x);

    x = 10;
    x ^= 3;
    System.out.println("x ^=3 is " + x);

    x = 10;
    x >>= 3;
    System.out.println("x >>=3 is " + x);

    x = 10;
    x <<= 3;
    System.out.println("x <<=3 is " + x);
  }
}
