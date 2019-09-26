package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const franceHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const france = "France"

// Hello func just give some output
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case france:
		prefix = franceHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
