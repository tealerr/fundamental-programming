function NumberAssign() {
  const y = 3;
  let x = 10;

  console.log("Initial x =", x);
  console.log("Constant y =", y);
  console.log("******** Number Assignment ********");

  // JavaScript Assignment Operators
  x += y;
  console.log("x += y, x =", x);

  x = 10; // Reset x to its initial value
  x -= y;
  console.log("x -= y, x =", x);

  x = 10; // Reset x to its initial value
  x *= y;
  console.log("x *= y, x =", x);

  x = 10; // Reset x to its initial value
  x /= y;
  console.log("x /= y, x =", x);

  x = 10; // Reset x to its initial value
  x %= y;
  console.log("x %= y, x =", x);

  x = 10; // Reset x to its initial value
  x **= y;
  console.log("x **= y, x =", x);

  // Shift Assignment Operators
  x = 10; // Reset x to its initial value
  x <<= y;
  console.log("x <<= y, x =", x);

  x = 10; // Reset x to its initial value
  x >>= y;
  console.log("x >>= y, x =", x);

  x = 10; // Reset x to its initial value
  x >>>= y;
  console.log("x >>>= y, x =", x);

  // Bitwise Assignment Operators
  x = 10; // Reset x to its initial value
  x &= y;
  console.log("x &= y, x =", x);

  x = 10; // Reset x to its initial value
  x ^= y;
  console.log("x ^= y, x =", x);

  x = 10; // Reset x to its initial value
  x |= y;
  console.log("x |= y, x =", x);

  // Logical Assignment Operators
  x = 10; // Reset x to its initial value
  x &&= y;
  console.log("x &&= y, x =", x);

  x = 10; // Reset x to its initial value
  x ||= y;
  console.log("x ||= y, x =", x);

  x = 10; // Reset x to its initial value
  x ??= y;
  console.log("x ??= y, x =", x);

  console.log("\n");
}

function OtherAssign() {
  // This is normal assign value to variable by use "="
  let myName = "Teeramate";
  let myChar = "T";
  let myBool = true;

  console.log("******** Other type ********");
  console.log("My name is ", myName);
  console.log("My character is ", myChar);
  console.log("My boolean is ", myBool);
  console.log("\n");
}

// call function for run
NumberAssign();
OtherAssign();
