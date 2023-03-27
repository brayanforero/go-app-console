package users

import "strings"

func AddToList(user string, listOfuser *[]string) {

	*listOfuser = append(*listOfuser, strings.ToUpper(user))
}
