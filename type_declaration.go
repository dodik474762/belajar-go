package main

import "fmt"

func main() {

	type NoKTP string
	ktpDodik := "1234423542542423"
	noKTP := NoKTP(ktpDodik)
	fmt.Println(noKTP)
}
