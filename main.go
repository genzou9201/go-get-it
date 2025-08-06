package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Go get it! You provided: %s\n", os.Args[1])
	} else {
		fmt.Println("Go get it! Welcome to the go-get-it application!")
		fmt.Println("Usage: go run main.go <your-message>")
	}
}
