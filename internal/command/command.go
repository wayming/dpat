package command

import "rsc.io/quote"

import "fmt"

func Pprint() {
	fmt.Println("command package")
	fmt.Println(quote.Hello())
}
