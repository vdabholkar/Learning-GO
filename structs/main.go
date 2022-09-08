package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	lastName  string
	firstName string
	contact   contactInfo
}

func main() {
	foo := person{
		lastName:  "Foo",
		firstName: "Bar",
		contact: contactInfo{
			email:   "foo.bar@gmail.com",
			zipCode: 91101,
		},
	}

	foo.updateLastName("Molly")
	foo.print()
}

func (p *person) updateLastName(ln string) {
	(*p).lastName = ln
}

func (p *person) print() {
	fmt.Printf("%+v", *p)
}
