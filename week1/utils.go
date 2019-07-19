package main

import "strings"

func message() string {
	return "Please provide arguments in format: <1st num> <arithmetic> <2nd num>\nArithmetic operations: + - * /"
}

func cleanSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}