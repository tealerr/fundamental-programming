function NumberArithmetic() {
  let x = 10;
  let y = 5;

  console.log("******** Number ********");
  console.log("x + y =", x + y);
  console.log("x - y =", x - y);
  console.log("x * y =", x * y);
  console.log("x ** y =", x ** y);
  console.log("x / y =", x / y);
  console.log("x % y =", x % y);
  console.log("x++ =", x++);
  console.log("x-- =", x--);
  console.log("\n");
}

function DecimalArithmetic() {
  let x = 12.34;
  let y = 2.23;

  console.log("******** Decimal ********");
  console.log("x + y =", x + y);
  console.log("x - y =", x - y);
  console.log("x * y =", x * y);
  console.log("x ** y =", x ** y);
  console.log("x / y =", x / y);
  console.log("x % y =", x % y);
  console.log("x++ =", x++);
  console.log("x-- =", x--);
  console.log("\n");
}

function StringArithmetic() {
  let firstname = "Jack";
  let lastname = "Dawson";

  console.log("******** String ********");
  console.log("firstname + lastname is", firstname + " " + lastname);
  console.log("\n");
}

function MixArithmetic() {
  let number = 10;
  let numStr = "5";
  let str = "hello";

  console.log("******** Mix data type ********");
  console.log("10 + '5' is ", number + numStr); // type is String
  console.log("10 - 'hello' is ", number - str);
  console.log("10 * '5' is ",
              number * numStr); // when use -, *, /, %, ** type is Number
  console.log("10 / '5' is ", number / numStr);
  console.log("10 % '5' is ", number % numStr);
  console.log("10 ** '5' is ", number ** numStr);
  console.log("\n");
}

NumberArithmetic();
DecimalArithmetic();
StringArithmetic();
MixArithmetic();
