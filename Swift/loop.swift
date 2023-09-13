// For-In Loops ใช้วนซ้ำค่าที่อยู่ใน arra
func forInLoop(){
    let number = [1,2,3,4,5]
    for num in number{
        print("for in loop count \(num) ")
    }
    print("\n")
}

// while จะทำงานในขณะที่เงื่อนไขเป็นจริง
func whileLoop() {
    var i = 1
    while i<5 { // กำหนดเงื่อนไขให้ทำซ้ำจนกว่า i>=5 (เป็นเท็จ)
        print("while loop count ",i)
        i += 1
    }
    print("\n")
}

// repeat while คล้ายกับ do while ในภาษา java
// ให้ทำซ้ำๆจนกว่าเงื่อนไขจะเป็นเท็จ
func repeatWhile(){
    var num = 1
    repeat{
        print("repeat count is ", num)
        num+=1
    }while num<=5
}

forInLoop()
whileLoop()
repeatWhile()
