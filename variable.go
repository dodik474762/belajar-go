package main

import "fmt"

func main() {
	var name string

	name = "dodik"
	fmt.Println(name)

	var names = "dodik rismawan affrudin"
	fmt.Println(names)

	namesNew := "dodik rismawan affrudin s.kom"
	fmt.Println(namesNew)

	var (
		firstName = "dodik"
		lastName  = "rismawan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	const nickname = "dodik"

	fmt.Println(nickname)
}
