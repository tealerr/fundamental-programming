//for loop
function forLoop() {
  for (let i = 0; i < 5; i++) {
    console.log(i);
  }
  console.log("\n");
}

// The For In Loop ใช้ในการวน loop คุณสมบัติของวัตถุ
// for...in เป็นลูปที่ใช้ในการวนลูปผ่านคุณสมบัติ (properties) ของอ็อบเจกต์ (objects)
// โดยทั่วไป for...in ใช้กับอ็อบเจกต์เท่านั้น (object)
function forInLoop() {
  var myCar = {
    brand: "BMW",
    models: "i8",
    color: "black",
  };
  for (const key in myCar) {
    //สร้างตัวแปร key เพื่อเก็บชื่อของคุณสมบัติ (property name)
    console.log(key + ":" + myCar[key]); // เข้าถึงค่าของข้อมูล โดยใช้ key และ print ออกมา
  }
  console.log("\n");
}

// The For Of Loop
// for...of เป็นลูปที่ใช้ในการวนลูปผ่านค่าข้อมูล (values) ของอ็อบเจกต์ที่เป็น Iterable (ตัวอย่างเช่น อาร์เรย์, สตริง, แม็พ, Set)
// for...of ใช้กับ Iterable objects เท่านั้น (ตัวอย่างเช่น อาร์เรย์, สตริง, แม็พ, Set)
function forOfLoop() {
  var myColor = ["red", "green", "blue"];

  for (const color of myColor) {
    // สร้างตัวแปร color เพื่อเก็บค่าของสมาชิกแต่ละตัวใน array
    console.log("Color is " + color); // print ค่าจากสมาชิกแต่ละตัวออกมา
  }
  console.log("\n");
}

//The While Loop
function whileLoop() {
  console.log("while loop");

  var num = 0;
  while (num < 5) {
    console.log("Count is " + num);
    num++;
  }
  console.log("\n");
}

//The Do While Loop
function doWhileLoop() {
  console.log("do-while loop");

  var num = 0;
  do {
    console.log("Count is " + num); //ให้ทำการ print
    num++;
  } while (num < 5); //ในขณะที่ num<5

  console.log("\n");
}

function breakAndCon() {
  console.log("\n");
  console.log("Let's say hello 5 times\n");
  console.log("***** Use break *****");
  for (i = 1; i <= 5; i++) {
    if (i == 3) {
      console.log("I'm break ^<^");
      break; //ใช้ break เพื่อบังคับออกจาก loop ทันทีเมื่อเข้าเงื่อนไข
    }
    console.log("Hello " + i + " times");
  }
  console.log("\n");

  console.log("***** Use continue *****");
  for (i = 1; i <= 5; i++) {
    if (i == 3) {
      console.log("I'm continue step ^o^");
      continue; // ใช้ continue เพื่อข้ามการทำงาน เมื่อเข้าเงื่อนไข
    }
    console.log("Hello " + i + " times");
  }
  console.log("\n");
}

forLoop();
forInLoop();
forOfLoop();
whileLoop();
doWhileLoop();
breakAndCon();
