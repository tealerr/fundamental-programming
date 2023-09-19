function NumberCompare() {
  console.log("******** Number Compare ********");

  let x = 10;
  let y = 5.3;
  console.log("x == y is ", x == y);
  console.log("x === y is ", x === y);
  console.log("x != y is ", x != y);
  console.log("x !== y is ", x !== y);
  console.log("x > y is ", x > y);
  console.log("x < y is ", x < y);
  console.log("x >= y is ", x >= y);
  console.log("x <= y is ", x <= y);
  console.log("\n");
}

function StringCompare() {
  console.log("********String Compare ********");

  let text1 = "Hello";
  let text2 = "hello";
  console.log("Hello == hello is ", text1 == text2);
  console.log("Hello === hello is ", text1 === text2);
  console.log("Hello != hello is ", text1 != text2);
  console.log("Hello !== hello is ", text1 !== text2);
  console.log("Hello > hello is ", text1 > text2);
  console.log("Hello < hello is ", text1 < text2);
  console.log("Hello >= hello is ", text1 >= text2);
  console.log("Hello <= hello is ", text1 <= text2);
  console.log("\n");

  let numtText1 = "20";
  let numtText2 = "5";
  console.log("20 == 5 is ", numtText1 == numtText2);
  console.log("20 === 5 is ", numtText1 === numtText2);
  console.log("20 != 5 is ", numtText1 != numtText2);
  console.log("20 !== 5 is ", numtText1 !== numtText2);
  console.log("20 > 5 is ", numtText1 > numtText2);
  console.log("20 < 5 is ", numtText1 < numtText2);
  console.log("20 >= 5 is ", numtText1 >= numtText2);
  console.log("20 <= 5 is ", numtText1 <= numtText2);
  console.log("\n");

  let char1 = "A";
  let char2 = "b";
  console.log("A == b is ", char1 == char2);
  console.log("A === b is ", char1 === char2);
  console.log("A != b is ", char1 != char2);
  console.log("A !== b is ", char1 !== char2);
  console.log("A > b is ", char1 > char2);
  console.log("A < b is ", char1 < char2);
  console.log("A >= b is ", char1 >= char2);
  console.log("A <= b is ", char1 <= char2);
  console.log("\n");
}

function MixCompare() {
  let num = 1;
  let str = "1";
  console.log("num == string is ", num == str);
  console.log("num === string is ", num === str);
  console.log("num != string is ", num != str);
  console.log("num !== string is ", num !== str);
  console.log("num > string is ", num > str);
  console.log("num < string is ", num < str);
  console.log("num >= string is ", num >= str);
  console.log("num <= string is ", num <= str);
  console.log("\n");
}

NumberCompare();
StringCompare();
MixCompare();
