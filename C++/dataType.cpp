#include <iostream>
using namespace std;

void primitiveData() {
  // Integer ใช้เก็บข้อมูลจำนวนเต็ม มีขนาด min -2147483648 max 2147483647
  int score = 99;

  // Character ใช้เก็บอักขระ มีขนาด min -128 max 127
  char myChar = 'M';

  // Boolean ใช้เก็บค่า true, false
  bool isMe = true;

  // Floating Point ใช้เก็บเลขทศนิยม มีขนาด 32 bit เก็บได้ 6-7 ตำแหน่ง
  float myFloat = 12.7f;

  // Double Floating Point ใช้เก็บเลขทศนิยม มีขนาด 64 bit เก็บได้ 15 ตำแหน่ง
  double myDouble = 1.12123;

  // Wide Character ใช้เก็บอักขระ โดยเก็บได้มากกว่า char โดยมีขนาด min 2 bytes max 4
  // bytes
  wchar_t wideChar = L'ก';
}

// Valueless or Void ใช้เพื่อบอกว่า function นี้ไม่มีการส่งค่ากลับ
//  โดย function นี้ จะแสดงเฉพาะข้อความอย่างเดียว
void printMessage() { cout << "Hello from the void function!" << endl; }
