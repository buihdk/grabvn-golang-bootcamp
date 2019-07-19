package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func cleanSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func eval(text string) (result float64, err error) {
	args := strings.Split(cleanSpaces(text), " ")
	if len(args) < 3 {
		return 0, errors.New("Please provide arguments in the right format <1st num> <arithmetic> <2nd num>")
	}
	x, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return
	}
	y, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return
	}
	op := args[1]

	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	default:
		fmt.Println("Invalid arguments")
	}
	return
}
