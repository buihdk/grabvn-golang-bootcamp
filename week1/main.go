package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Format: <1st num> <arithmetic> <2nd num>")
	fmt.Print("> ")

	for scanner.Scan() {
        res, err := eval(scanner.Text())
        if err != nil {
        	fmt.Println(err.Error())
			fmt.Print("> ")
        	continue
		} else {
			fmt.Println(res)
		}
        fmt.Print("> ")
    }
}