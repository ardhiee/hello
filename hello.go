package hello

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello func only return 'Hello, World!'
func Hello(name, language string) string {

	if name == "" {
		name = "World!"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	if language == "French" {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Ardi", ""))
}
