package hello

import "fmt"

// Hello func only return 'Hello, World!'
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Ardi"))
}
