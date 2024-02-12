package form

import (
	"fmt"
	"log"
)

func AskName(firstName, lastName *string) {
	fmt.Println("What is your first name?")
	_, err := fmt.Scanf("%s", firstName)
	if err != nil {
		log.Fatal("Input Error")
	}

	fmt.Println("What is your last name?")
	_, err = fmt.Scanf("%s", lastName)
	if err != nil {
		log.Fatal("Input Error")
	}

}
