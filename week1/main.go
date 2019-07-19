package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
        fmt.Println(eval(scanner.Text()))
        fmt.Print("> ")
    }
}