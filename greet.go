package greet

import "fmt"

func Greet(lang string) {
	if lang == "English" {
		fmt.Println("Hello")
	} else {
		fmt.Println("Bonjour")
	}
}
