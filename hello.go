package hello

const englishHelloPrefix = "Hello, "

// Hello function just give simple output
func Hello(name string) string {

	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
