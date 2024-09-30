package main

import "fmt"

func main() {
	const value = 10
	fmt.Println(value)

	const conversionValue = int32(value)
	fmt.Println(conversionValue)

	name := "dodik"
	e := name[0]
	eString := string(e)
	fmt.Println(eString)
}
