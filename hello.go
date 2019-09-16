package hello

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello func only return 'Hello, World!'
func Hello(name string) string {

	if name == "" {
		name = "World!"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Ardi"))
}
