/*
This course talked about creating structs and interfaces in a separate class/package
It also discussed the importance of using pointer based objects for setting and receiving variables
*/
package main

import (
	"fmt"

	"github.com/go_pluralsight_beginner/interfaces_structs/organization"
)

func main() {
	p := organization.NewPerson("Mike", "LaBreche")
	err := p.SetTwitterHandler("@IAM")
	if err != nil {
		fmt.Printf("An error occurred setting Twitter handler: %s\n", err.Error())
	}
	println(p.GetTwitterHandler())
	println(p.ID())
	println(p.FullName())
}
