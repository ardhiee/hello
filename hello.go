package hello

// Hello func just give some output
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}
