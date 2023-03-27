package main

import (
	"fmt"

	"github.com/brayanforero/go-app-console/texts"
	"github.com/brayanforero/go-app-console/users"
)

func main() {
	lang := ""
	usersList := []string{}

	texts.PrintHeader()

	_, err := fmt.Scan(&lang)

	if err != nil {
		fmt.Println("Cannot read your information: ", err)
		return
	}

	userLang := texts.GetAllTranslatedText(lang)
	user := ""
	answer := "y"

	for answer == "y" {
		fmt.Println(userLang["header"] + ":")
		_, err = fmt.Scan(&user)

		if err != nil {
			fmt.Println("Cannot read your information: ", err)
			return
		}

		users.AddToList(user, &usersList)

		fmt.Print(userLang["alert"] + ": ")

		fmt.Scan(&answer)

		if err != nil {
			fmt.Println("Cannot read your information: ", err)
			break
		}
	}

	fmt.Println(userLang["body"] + ":")

	for index, u := range usersList {

		fmt.Printf("%d. %s", index+1, u)
	}
}
