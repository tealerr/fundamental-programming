func integerAssignment() {
    var myInt: Int

    // Integer assignment
    print("####### Integer assignment #######")
    myInt = 10
    print("int =", myInt)

    myInt = 10
    myInt += 5
    print("10 += 5 =", myInt)

    myInt = 10
    myInt -= 3
    print("10 -= 3 =", myInt)

    myInt = 10
    myInt *= 2
    print("10 *= 2 =", myInt)

    myInt = 10
    myInt /= 4
    print("10 /= 4 =", myInt)

    myInt = 10
    myInt %= 3
    print("10 %= 3 =", myInt)
    print("\n")
}

func floatingPointAssignment() {
    var myFloat: Float

    // Floating-point assignment
    print("####### Floating-point assignment #######")
    myFloat = 3.14
    print("float =", myFloat)

    myFloat = 3.14
    myFloat += 3.5
    print("3.14 += 3.5 =", myFloat)

    myFloat = 3.14
    myFloat -= 1.2
    print("3.14 -= 1.2 =", myFloat)

    myFloat = 3.14
    myFloat *= 2.2
    print("3.14 *= 2.2 =", myFloat)

    myFloat = 3.14
    myFloat /= 1.5
    print("3.14 /= 1.5 =", myFloat)
    print("\n")
}

func doubleAssignment() {
    var myDouble: Double

    // Double assignment
    print("####### Double assignment #######")
    myDouble = 5.25
    print("double =", myDouble)

    myDouble = 5.25
    myDouble += 3.5
    print("5.25 += 3.5 =", myDouble)

    myDouble = 5.25
    myDouble -= 1.2
    print("5.25 -= 1.2 =", myDouble)

    myDouble = 5.25
    myDouble *= 2.2
    print("5.25 *= 2.2 =", myDouble)

    myDouble = 5.25
    myDouble /= 1.5
    print("5.25 /= 1.5 =", myDouble)
    print("\n")
}

func stringAssignment() {
    var myStr: String

    // String assignment
    print("####### String assignment #######")
    myStr = "Hello, "
    myStr += "world!"
    print("my string is ")
    print("string1 + string2 is ", myStr)
    print("\n")
}

func characterAssignment() {
    var myChar: Character

    // Character assignment
    print("####### Character assignment #######")
    myChar = "T"
    print(myChar)
    print("\n")
}

func booleanAssignment() {
    var myBool: Bool

    // Boolean assignment
    print("####### Boolean assignment #######")
    myBool = true
    print(myBool)
    print("\n")
}

func assignment() {
    integerAssignment()
    floatingPointAssignment()
    doubleAssignment()
    stringAssignment()
    characterAssignment()
    booleanAssignment()
}
