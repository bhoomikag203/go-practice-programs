 package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const kannadaHelloPrefix = "Namaskara "
const spanish = "Spanish"
const french = "French"
const kannada = "Kannada"

func Hello(name string, language string) string{
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}


func greetingPrefix(language string)(prefix string) {
	switch language {
	case kannada:
		prefix = kannadaHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish: 
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(Hello("Raju", ""))
}
