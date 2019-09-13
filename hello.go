package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello func only receive string and return string
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
