package main

import (
	"flag"
	"fmt"
)

func main() {
	var greeting string
	flag.StringVar(&greeting, "hello", "Hello, Slate", "Slate greeting message")
	flag.Parse()
	fmt.Println(greeting)
}
