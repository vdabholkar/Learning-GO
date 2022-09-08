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
	person := person{
		lastName:  "Foo",
		firstName: "Bar",
		contact: contactInfo{
			email:   "foo.bar@gmail.com",
			zipCode: 91101,
		},
	}

	fmt.Printf("%+v", person)
}
