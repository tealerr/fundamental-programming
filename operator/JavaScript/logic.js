function NumberLogicAnd() {
    let a = 1
    let b = 2
    if (a > b && b < 5) {
        console.log("Hello world")
    } else {
        console.log("Try again")
    }
}

function NumberLogicOr() {
    let a = 1
    let b = 2
    if (a > b || b < 5) {
        console.log("Hello world")
    } else {
        console.log("Try again")
    }
}

function NumberLogicNot() {
    let a = 1
    let b = 2
    if (a != b && b != 2) {
        console.log("Hello world")
    } else {
        console.log("Try again")
    }
}

function StringLogicAnd() {
    let user = "teera"
    let password = "12345"

    if (user === "teera" && password === "12345") {
        return console.log("success")
    } else {
        return console.log("try again")
    }

}

function StringLogicOr() {
    let user = "Teera"
    let password = 12345

    if (user === "teera" || password === "12345") {
        return console.log("success")
    } else {
        return console.log("try again")
    }
}

function StringLogicNot() {
    let user = "teera"
    let password = "12345"
    let gender = 'M'

    if (user !== "Teera" || password !== 12345 || gender !== 'F') {
        return console.log("success")
    } else {
        return console.log("try again")
    }
}



console.log("######## Number ########")
NumberLogicAnd()
NumberLogicOr()
NumberLogicNot()
console.log("\n")

console.log("######## String ########")
StringLogicAnd()
StringLogicOr()
StringLogicNot()