package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	kirundi = "Kirundi"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Salut, "
	kirundiHelloPrefix = "Yambu, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case kirundi:
		prefix = kirundiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
func main() {
	fmt.Println(Hello("world", ""))
}
