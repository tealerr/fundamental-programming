// For-In Loops ใช้วนซ้ำค่าที่อยู่ใน array
func forInLoop() {
    let numbers = [1, 2, 3, 4, 5]
    for number in numbers {
        print("for-in loop count \(number)")
    }
    print("\n")
}

// while จะทำงานในขณะที่เงื่อนไขเป็นจริง
func whileLoop() {
    var counter = 1
    while counter < 5 { // Executes as long as counter is less than 5
        print("while loop count \(counter)")
        counter += 1
    }
    print("\n")
}

// repeat while คล้ายกับ do while ในภาษา java
// ให้ทำซ้ำๆจนกว่าเงื่อนไขจะเป็นเท็จ
func repeatWhile() {
    var number = 1
    repeat {
        print("repeat count is \(number)")
        number += 1
    } while number <= 5
}
