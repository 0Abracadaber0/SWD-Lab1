package main

import (
	"Lab_1/form"
	"fmt"
)

func main() {
	var firstName, lastName string

	form.AskName(&firstName, &lastName)

	fmt.Printf("Hello, %s %s\n", firstName, lastName)
}
