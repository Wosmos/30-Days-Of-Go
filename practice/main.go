package main

import (
	"fmt"
)

func main() {
    var num1, num2 float64
    var operator string

    // Input first number
    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)

    // Input operator
    fmt.Print("Enter operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    // Input second number
    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)

    // Perform calculation based on operator
    switch operator {
    case "+":
        fmt.Printf("Result: %.2f\n", num1+num2)
    case "-":
        fmt.Printf("Result: %.2f\n", num1-num2)
    case "*":
        fmt.Printf("Result: %.2f\n", num1*num2)
    case "/":
        if num2 != 0 {
            fmt.Printf("Result: %.2f\n", num1/num2)
        } else {
            fmt.Println("Error: Division by zero")
        }
    default:
        fmt.Println("Error: Invalid operator")
    }
}
