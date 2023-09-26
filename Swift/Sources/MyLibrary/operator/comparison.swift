func compareDoubles() {
    let doubleA: Double = 7.2
    let doubleB: Double = 3.8

    // Comparison operators with doubles
    print("######## Double Compare")
    print("\(doubleA) == \(doubleB) =", doubleA == doubleB)
    print("\(doubleA) != \(doubleB) =", doubleA != doubleB)
    print("\(doubleA) > \(doubleB) =", doubleA > doubleB)
    print("\(doubleA) < \(doubleB) =", doubleA < doubleB)
    print("\(doubleA) >= \(doubleB) =", doubleA >= doubleB)
    print("\(doubleA) <= \(doubleB) =", doubleA <= doubleB)
    print("\n")
}

func compareStrings() {
    let stringA = "Hello"
    let stringB = "hello"

    // Comparison operators with strings
    print("######## String Compare")
    print("\(stringA) == \(stringB) =", stringA == stringB)
    print("\(stringA) != \(stringB) =", stringA != stringB)
    print("\n")
}

func compareCharacters() {
    let charA: Character = "A"
    let charB: Character = "B"

    // Comparison operators with characters
    print("######## Characters Compare")
    print("\(charA) == \(charB) =", charA == charB)
    print("\(charA) != \(charB) =", charA != charB)
    print("\n")
}

func compareBooleans() {
    let boolA = true
    let boolB = false

    // Comparison operators with booleans
    print("######## Booleans Compare")
    print("\(boolA) == \(boolB) =", boolA == boolB)
    print("\(boolA) != \(boolB) =", boolA != boolB)
    print("\n")
}
func compareNumbers() {
    let intA = 5
    let intB = 10

    // Equality operator (==)
    print("######## Number Compare")
    print("\(intA) == \(intB) =", intA == intB)
    print("\(intA) != \(intB) =", intA != intB)
    print("\(intA) > \(intB) =", intA > intB)
    print("\(intA) < \(intB) =", intA < intB)
    print("\(intA) >= \(intB) =", intA >= intB)
    print("\(intA) <= \(intB) =", intA <= intB)
    print("\n")
}

func compareFloats() {
    let floatA: Float = 5.5
    let floatB: Float = 3.48

    // Comparison operators with floats
    print("######## Float Compare")
    print("\(floatA) == \(floatB) =", floatA == floatB)
    print("\(floatA) != \(floatB) =", floatA != floatB)
    print("\(floatA) > \(floatB) =", floatA > floatB)
    print("\(floatA) < \(floatB) =", floatA < floatB)
    print("\(floatA) >= \(floatB) =", floatA >= floatB)
    print("\(floatA) <= \(floatB) =", floatA <= floatB)
    print("\n")
}

func datatypeCompare() {
    compareNumbers()
    compareFloats()
    compareDoubles()
    compareStrings()
    compareCharacters()
    compareBooleans()
}
