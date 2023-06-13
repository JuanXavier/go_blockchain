package main

import (
	"fmt"
)
var log = fmt.Println

func main() {
    var num1, num2 float64
    var operator string

    for {
        log("First number: ")
        fmt.Scanln(&num1)
        log("Second number: ")
        fmt.Scanln(&num2)

        log("Select the operation to perform: ")
        log("1. Addition (+)")
        log("2. Subtraction (-)")
        log("3. Multiplication (*)")
        log("4. Division (/)")
        log("5. Exit")

        log("Option: ")
        fmt.Scanln(&operator)

        switch operator {
        case "1":
            result := add(num1, num2)
            log("Result:", result)
        case "2":
            result := sub(num1, num2)
            log("Result:", result)
        case "3":
            result := mul(num1, num2)
            log("Result:", result)
        case "4":
            if num2 != 0 {
                result := div(num1, num2)
                log("Result:", result)
            } else {
                log("Error: Cant divide by zero.")
            }
        case "5":
            log("Exiting...")
            return
        default:
            log("Invalid option. Try again")
        }
        log()
    }
}

func add(a float64, b float64) float64 {
    return a + b
}

func sub(a float64, b float64) float64 {
    return a - b
}

func mul(a float64, b float64) float64 {
    return a * b
}

func div(a float64, b float64) float64 {
    return a / b
}
