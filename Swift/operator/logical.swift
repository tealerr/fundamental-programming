var myInt1 = 5
var myInt2 = 10

var myFloat1: Float = 5.5
var myFloat2: Float = 10.0

var myDouble1: Double = 7.2
var myDouble2: Double = 3.8

var myStr1 = "Hello"
var myStr2 = "World"

var myChar1: Character = "A"
var myChar2: Character = "B"

var myBool1 = true
var myBool2 = false

// Logical AND operator (&&)
print("Boolean Logic")
print("\(myBool1) && \(myBool2) =", myBool1 && myBool2)
print("\n")

// Logical OR operator (||)
print("\(myBool1) || \(myBool2) =", myBool1 || myBool2)
print("\n")

// Logical NOT operator (!)
print("!\(myBool1) =", !myBool1)
print("!\(myBool2) =", !myBool2)
print("\n")

// Logical operators with integers
print("Number Logic")
print("\(myInt1) > 0 && \(myInt2) < 20 =", myInt1 > 0 && myInt2 < 20)
print("\(myInt1) > 0 || \(myInt2) < 20 =", myInt1 > 0 || myInt2 < 20)
print("\n")

// Logical operators with floats
print("Float Logic")
print("\(myFloat1) > 0.0 && \(myFloat2) < 15.0 =", myFloat1 > 0.0 && myFloat2 < 15.0)
print("\(myFloat1) > 0.0 || \(myFloat2) < 15.0 =", myFloat1 > 0.0 || myFloat2 < 15.0)
print("\n")

// Logical operators with doubles
print("Double Logic")
print("\(myDouble1) > 0.0 && \(myDouble2) < 10.0 =", myDouble1 > 0.0 && myDouble2 < 10.0)
print("\(myDouble1) > 0.0 || \(myDouble2) < 10.0 =", myDouble1 > 0.0 || myDouble2 < 10.0)
print("\n")

// Logical operators with strings
print("String Logic")
print("\(myStr1) == \"Hello\" && \(myStr2) == \"World\" =", myStr1 == "hello" && myStr2 == "World")
print("\(myStr1) == \"Hello\" || \(myStr2) == \"Universe\" =", myStr1 == "Hello" || myStr2 == "Universe")
print("\n")

// Logical operators with characters
print("Character Logic")
print("\(myChar1) == \"C\" && \(myChar2) == \"B\" =", myChar1 == "C" && myChar2 == "B")
print("\(myChar1) == \"A\" || \(myChar2) == \"C\" =", myChar1 == "A" || myChar2 == "C")
print("\n")
