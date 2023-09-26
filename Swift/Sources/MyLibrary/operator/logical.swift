func logicBoolean() {
    let myBool1 = true
    let myBool2 = false

    print("\(myBool1) && \(myBool2) =", myBool1 && myBool2)
    print("\(myBool1) || \(myBool2) =", myBool1 || myBool2)
    print("!\(myBool1) =", !myBool1)
    print("!\(myBool2) =", !myBool2)
    print("\n")
}

func logicNumber() {
    let myInt1 = 5
    let myInt2 = 10

    print("\(myInt1) == \(myInt2) =", myInt1 == myInt2)
    print("\(myInt1) != \(myInt2) =", myInt1 != myInt2)
    print("\(myInt1) > \(myInt2) =", myInt1 > myInt2)
    print("\(myInt1) < \(myInt2) =", myInt1 < myInt2)
    print("\(myInt1) >= \(myInt2) =", myInt1 >= myInt2)
    print("\(myInt1) <= \(myInt2) =", myInt1 <= myInt2)
    print("\n")
}

func logicFloat() {
    let myFloat1: Float = 5.5
    let myFloat2: Float = 10.0

    print("\(myFloat1) == \(myFloat2) =", myFloat1 == myFloat2)
    print("\(myFloat1) != \(myFloat2) =", myFloat1 != myFloat2)
    print("\(myFloat1) > \(myFloat2) =", myFloat1 > myFloat2)
    print("\(myFloat1) < \(myFloat2) =", myFloat1 < myFloat2)
    print("\(myFloat1) >= \(myFloat2) =", myFloat1 >= myFloat2)
    print("\(myFloat1) <= \(myFloat2) =", myFloat1 <= myFloat2)
    print("\n")
}

func logicDouble() {
    let myDouble1: Double = 7.2
    let myDouble2: Double = 3.8

    print("\(myDouble1) == \(myDouble2) =", myDouble1 == myDouble2)
    print("\(myDouble1) != \(myDouble2) =", myDouble1 != myDouble2)
    print("\(myDouble1) > \(myDouble2) =", myDouble1 > myDouble2)
    print("\(myDouble1) < \(myDouble2) =", myDouble1 < myDouble2)
    print("\(myDouble1) >= \(myDouble2) =", myDouble1 >= myDouble2)
    print("\(myDouble1) <= \(myDouble2) =", myDouble1 <= myDouble2)
    print("\n")
}

func logicString() {
    let myStr1 = "Hello"
    let myStr2 = "World"

    print("\(myStr1) == \(myStr2) =", myStr1 == myStr2)
    print("\(myStr1) != \(myStr2) =", myStr1 != myStr2)
    print("\n")
}

func logicCharacter() {
    let myChar1: Character = "A"
    let myChar2: Character = "B"

    print("\(myChar1) == \(myChar2) =", myChar1 == myChar2)
    print("\(myChar1) != \(myChar2) =", myChar1 != myChar2)
    print("\n")
}

func datatypeLogic() {
    logicBoolean()
    logicNumber()
    logicFloat()
    logicDouble()
    logicString()
    logicCharacter()
}
