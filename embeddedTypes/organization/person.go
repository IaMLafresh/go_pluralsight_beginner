package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

func (th TwitterHandler) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Idenfifiable interface {
	ID() string
}

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	firstName string
	lastName  string
}

type Person struct {
	Name
	twitterHandler TwitterHandler
}

func (p Person) ID() string {
	return "12345"
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("Twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p Person) GetTwitterHandler() TwitterHandler {
	return p.twitterHandler
}
