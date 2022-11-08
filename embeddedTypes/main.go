/*
This course talked about creating structs and interfaces in a separate class/package
It also discussed the importance of using pointer based objects for setting and receiving variables
*/
package main

import (
	"fmt"

	"github.com/go_pluralsight_beginner/embeddedTypes/organization"
)

func main() {
	p := organization.NewPerson("Mike", "LaBreche")
	err := p.SetTwitterHandler("@IAM")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting Twitter handler: %s\n", err.Error())
	}
	println(p.GetTwitterHandler())
	println(p.ID())
	println(p.FullName())
}
