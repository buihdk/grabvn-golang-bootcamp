package main

import (
    "fmt"
    "strconv"
    "strings"
)

func cleanSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func eval(text string) (result float64) {
    args := strings.Split(cleanSpaces(text), " ")
    x, _ := strconv.ParseFloat(args[0], 64)
    y, _ := strconv.ParseFloat(args[2], 64)
    op := args[1]

    switch op {
        case "+":
            result =  x + y
        case "-":
            result = x - y
        case "*":
            result = x * y
        case "/":
            result = x / y
        default:
            fmt.Println("Invalid arguments")
    }
    return result
}