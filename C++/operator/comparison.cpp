using namespace std;
#include <iostream>

void numberCompare();
void stringCompare();
void charCompare();
void decimalCompare();

int main() {
  numberCompare();
  stringCompare();
  charCompare();
  decimalCompare();
}

void numberCompare() {
  int x = 5;
  int y = 3;

  std::cout << "*** 0 is false; 1 is true;" << std::endl;
  std::cout << "5 == 3: " << (x == y) << std::endl;
  std::cout << "5 != 3: " << (x != y) << std::endl;
  std::cout << "5 > 3: " << (x > y) << std::endl;
  std::cout << "5 < 3: " << (x < y) << std::endl;
  std::cout << "5 >= 3: " << (x >= y) << std::endl;
  std::cout << "5 <= 3: " << (x <= y) << std::endl;

  std::cout << "\n" << std::endl;
}

void stringCompare() {
  string str1 = "Hello";
  string str2 = "hello";

  std::cout << "*** 0 is false; 1 is true;" << std::endl;
  std::cout << "Hello == hello: " << (str1 == str2) << std::endl;
  std::cout << "Hello != hello: " << (str1 != str2) << std::endl;
  std::cout << "Hello > hello: " << (str1 > str2) << std::endl;
  std::cout << "Hello < hello: " << (str1 < str2) << std::endl;
  std::cout << "Hello >= hello: " << (str1 >= str2) << std::endl;
  std::cout << "Hello <= hello: " << (str1 <= str2) << std::endl;

  std::cout << "\n" << std::endl;
}

void charCompare() {
  char char1 = 'a';
  char char2 = 'b';

  std::cout << "*** 0 is false; 1 is true;" << std::endl;
  std::cout << "char1 == char2: " << (char1 == char2) << std::endl;
  std::cout << "char1 != char2: " << (char1 != char2) << std::endl;
  std::cout << "char1 > char2: " << (char1 > char2) << std::endl;
  std::cout << "char1 < char2: " << (char1 < char2) << std::endl;
  std::cout << "char1 >= char2: " << (char1 >= char2) << std::endl;
  std::cout << "char1 <= char2: " << (char1 <= char2) << std::endl;

  std::cout << "\n" << std::endl;
}

void decimalCompare() {
  double x = 5.5;
  double y = 3.2;

  std::cout << "*** 0 is false; 1 is true;" << std::endl;
  std::cout << "5.5 == 3.2: " << (x == y) << std::endl;
  std::cout << "5.5 != 3.2: " << (x != y) << std::endl;
  std::cout << "5.5 > 3.2: " << (x > y) << std::endl;
  std::cout << "5.5 < 3.2: " << (x < y) << std::endl;
  std::cout << "5.5 >= 3.2: " << (x >= y) << std::endl;
  std::cout << "5.5 <= 3.2: " << (x <= y) << std::endl;

  std::cout << "\n" << std::endl;
}
