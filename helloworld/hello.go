package main

import (
	"fmt"
)


const (
	spanish = "Spanish"
	french = "French"

 	englishPrefix = "Hello "
 	spanishPrefix = "Hola "
 	frenchPrefix = "Bonjour "
)



func Hello(name, language string)  string {
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
	switch(language){
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix

	}
}

func main() {
	fmt.Println(Hello("World", ""))
}

