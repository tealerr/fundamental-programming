using namespace std;
#include <iostream>

void integerArithmetic();
void decimalArithmetic();
void charArithmetic();
void stringArithmetic();

int main()
{
  integerArithmetic();
  decimalArithmetic();
  charArithmetic();
  stringArithmetic();

  return 0;
}

void integerArithmetic()
{
  int x = 5;
  int y = 3;

  std::cout << "5 + 3 = " << x + y << std::endl;
  std::cout << "5 - 3 = " << x - y << std::endl;
  std::cout << "5 * 3 = " << x * y << std::endl;
  std::cout << "5 / 3 = " << x / y << std::endl;
  std::cout << "5 % 3 = " << x % y << std::endl;

  std::cout << "\n"
            << std::endl;
}

void decimalArithmetic()
{
  float x = 5.25;
  float y = 3.3;

  std::cout << "5.25 + 3.3 = " << x + y << std::endl;
  std::cout << "5.25 - 3.3 = " << x - y << std::endl;
  std::cout << "5.25 * 3.3 = " << x * y << std::endl;
  std::cout << "5.25 / 3.3 = " << x / y << std::endl;
  std::cout << "++x = " << ++x << std::endl;
  std::cout << "--x = " << --x << std::endl;

  std::cout << "\n"
            << std::endl;
}

void charArithmetic()
{
  char char1 = 'A';
  char char2 = 'B';

  std::cout << "char1 + char2 = " << int(char1 + char2) << std::endl;
  std::cout << "char1 - char2 = " << int(char1 - char2) << std::endl;
  std::cout << "char1 * char2 = " << int(char1 * char2) << std::endl;
  std::cout << "char1 / char2 = " << int(char1 / char2) << std::endl;

  std::cout << "\n"
            << std::endl;
}

void stringArithmetic()
{
  string firstname = "Teeramate";
  string lastname = "Kantima";

  std::cout << "My name is " << firstname + " " + lastname << std::endl;

  std::cout << "\n"
            << std::endl;
}
