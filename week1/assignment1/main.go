package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Goland Bootcamp Week 1 Assignment")
    fmt.Println(message())
	fmt.Print("> ")

	for scanner.Scan() {
        res, err := eval(scanner.Text())
        if err != nil {
        	fmt.Println(err.Error())
        	fmt.Print("> ")
        	continue
		}
		fmt.Println(res)
        fmt.Print("> ")
    }
}
