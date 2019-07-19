package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)



func eval(text string) (result float64, err error) {
	args := strings.Split(cleanSpaces(text), " ")
	if len(args) != 3 {
		return 0, errors.New(message())
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
		fmt.Println(message())
	}
	return
}
